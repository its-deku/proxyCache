package httphandler

import (
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
)

type proxy struct {
	URL    string
	muLock sync.Mutex
	cache  map[string]*http.Response
}

func Init(port string, forwardUrl string) {
	pry := proxy{
		URL:   forwardUrl,
		cache: map[string]*http.Response{},
	}
	http.HandleFunc("/", pry.Handle)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (pry *proxy) Handle(w http.ResponseWriter, req *http.Request) {
	newUrl := pry.URL + strings.TrimLeft(req.URL.Path, "/")

	// check if the url exists in the cache
	if _, exists := pry.cache[newUrl]; exists {
		pry.cache[newUrl].Header.Set("X-Cache", "HIT")
		io.WriteString(w, pry.cache[newUrl].Header["X-Cache"][0]+"\t")
		io.WriteString(w, pry.cache[newUrl].Status)
	} else {
		res, err := http.Get(newUrl)
		if err != nil {
			panic(err)
		}
		res.Header.Add("X-Cache", "MISS")
		io.WriteString(w, res.Header["X-Cache"][0]+"\t")
		io.WriteString(w, res.Status)

		// cache the response (add concurrency safety)
		pry.muLock.Lock()
		pry.cache[newUrl] = res
		pry.muLock.Unlock()
	}
}

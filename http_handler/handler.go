package httphandler

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type proxy struct {
	URL string
}

func Init(port string, forwardUrl string) {
	pry := proxy{
		URL: forwardUrl,
	}
	http.HandleFunc("/", pry.Handle)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func (pry *proxy) Handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("any path accepted, url %v", pry.URL+req.URL.Path))
}

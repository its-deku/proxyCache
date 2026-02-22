package httphandler

import (
	"fmt"
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, fmt.Sprintf("any path acceoted, url %v", req.URL.Path))
}

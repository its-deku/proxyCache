package main

import (
	"log"
	"net/http"

	httphandler "example.com/v2/proxy_cache/http_handler"
)

func main() {
	http.HandleFunc("/", httphandler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

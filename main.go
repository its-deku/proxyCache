package main

import (
	"flag"
	"fmt"

	httphandler "example.com/v2/proxy_cache/http_handler"
)

func main() {
	port := flag.String("port", "8080", "port on which the server runs")
	url := flag.String("origin", "", "URL for the intended website for which the caching should be enabled")
	flag.Parse()

	fmt.Printf("Proxy server started on port %s\n", *port)

	httphandler.Init(*port, *url)
}

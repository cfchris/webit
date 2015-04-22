package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port string
var path string
var cors bool

func main() {
	flag.StringVar(&port, "p", "3333", "Port for server to run on - defaults to 3333")
	flag.StringVar(&path, "d", "./", "Server directory - defaults to ./")
	flag.BoolVar(&cors, "c", false, "Add cors support")
	flag.Parse()

	handler := http.FileServer(http.Dir(path))

	if cors {
		handler = corsHandler(handler)
	}

	fmt.Printf("Starting web server on port %s from %s\n", port, path)
	http.ListenAndServe(":"+port, handler)
}

func corsHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

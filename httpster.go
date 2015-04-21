package main

import (
	"flag"
	"fmt"
	"net/http"
)

var port string
var path string

func main() {
	flag.StringVar(&port, "p", "3333", "Port for server to run on - defaults to 3333")
	flag.StringVar(&path, "d", "./", "Server directory - defaults to ./.")
	flag.Parse()

	fmt.Printf("Starting HTTPster on port %s from %s\n", port, path)
	http.ListenAndServe(":"+port, http.FileServer(http.Dir(path)))
}

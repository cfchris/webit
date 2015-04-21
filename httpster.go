package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

var port string
var path string

func main() {
	workdir, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln("Error getting current working directory:", err.Error())
	}

	flag.StringVar(&port, "p", "3333", "Port for server to run on - defaults to 3333")
	flag.StringVar(&path, "d", workdir, "Server directory - defaults to ./.")
	flag.Parse()

	httpDir := http.Dir(path)

	fmt.Printf("Starting HTTPster on port %s from %s\n", port, httpDir)
	http.ListenAndServe(":"+port, http.FileServer(httpDir))
}

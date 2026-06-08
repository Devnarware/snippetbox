package main

import (
	"flag"
	"log"
	"net/http"
)

func main(){

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// stripPrefix() -> it is used to remove the prefix from the request URL before passing it to the file server, it is used to serve static files from the "./ui/static" directory when the request URL starts with "/static/"

	// .handle -> 

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	// create all the handler function

	log.Printf("Server is starting at the port: %s", *addr)

	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err)
}
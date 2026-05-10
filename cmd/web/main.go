package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// stripPrefix() -> it is used to remove the prefix from the request URL before passing it to the file server, it is used to serve static files from the "./ui/static" directory when the request URL starts with "/static/"

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	// create all the handler function

	fmt.Printf("Starting the server at port :4000")

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err) ;
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POSt /snippet/create", snippetCreatePost)
	// create all the handler function

	fmt.Printf("Starting the server at port :4000")

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err) ;
	}
}
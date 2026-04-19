package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)


	log.Print("starting the server on :4000")

	http.ListenAndServe(":4000", mux)
}

func home(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("hello from snippetbox"))

	if err != nil {
		http.Error(w, "idk", http.StatusNotAcceptable)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	
	w.Header().Add("server", "GO")
	w.Write([]byte("Hello from snippet box"))

}
// this is our Home function, which will handle the "/" route


func snippetView(w http.ResponseWriter, r *http.Request){
	
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
// this is the snippetView Function, which will hanlde the get request for "/snippet/view/{id}"


func snippetCreate(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("Display a form for creating a new snippet..."))
}
// this is the snippetCreate function, which will handle the get request for "/snippet/create"


func snippetCreatePost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new Snippet..."))
}
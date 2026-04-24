package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request){
	// ts = template set
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	// tempplate.ParseFiles() -> it is used to read the content of template file and returns template set and err if any

	if err != nil {
		log.Print(err.Error())
		// it is used to print the error message in the console, not for client
		http.Error(w,"Internal server error", http.StatusInternalServerError)
		//it used to send the error to the client
		// it is a light weight function which returns text respone with the status code and the msg u want to send to the client
		return
	}

	err = ts.Execute(w, nil) 
	// executing the template set and write it as a response


	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return 
	}

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
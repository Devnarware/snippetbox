package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

func home(w http.ResponseWriter, r *http.Request){

	w.Header().Add("server", "Go")


	// now adding all the neccessary htlm templates here
	files := []string{

				"./ui/html/base.html",
				"./ui/html/partials/nav.html",
				"./ui/html/pages/home.html",

			}

	// ts = template set
	// (...) this ois the spread operator just like for each, it says every value in the slice
	ts, err := template.ParseFiles(files...)
	// tempplate.ParseFiles() -> it is used to read the content of template file and returns template set and err if any

	if err != nil {
		log.Print(err.Error())
		// it is used to print the error message in the console, not for client
		http.Error(w,"Internal server error", http.StatusInternalServerError)
		//it used to send the error to the client
		// it is a light weight function which returns text respone with the status code and the msg u want to send to the client
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	// executing the template set and write it as a response
	// use ececuteTemplate() instead of execute() because we have multiple templates


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
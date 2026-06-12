package main

import (
	"fmt"
	"net/http"
	"strconv"
	"html/template"
)

func (app *application)home(w http.ResponseWriter, r *http.Request){
	// have to recieve the application struct as a parameter because we have to use the custom logger in this function, and the custom logger is a field of the application struct, so we have to use app.logger to access the custom logger in this function

	w.Header().Add("server", "Go")


	// now adding all the neccessary htlm templates here
	files := []string{

				"./ui/html/base.html",
				"./ui/html/partials/nav.html",
				"./ui/html/pages/home.html",

			}

	// ts = template set
	// (...) this is the spread operator just like for each, it says every value in the slice
	ts, err := template.ParseFiles(files...)
	// template.ParseFiles() -> it is used to read the content of template file and returns template set and err if any

	if err != nil {
		app.serverError(w, r, err)
		// it is used to print the error message in the console, not for developer
		// here we are using the custom looger we created

		http.Error(w,"Internal server error", http.StatusInternalServerError)
		//it used to send the error to the client
		// it is a light weight function which returns text respone with the status code and the msg u want to send to the client
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	// executing the template set and write it as a response
	// use ececuteTemplate() instead of execute() because we have multiple templates


	if err != nil {
		app.serverError(w, r, err)
		return 
	}

}
// this is our Home function, which will handle the "/" route


func (app *application)snippetView(w http.ResponseWriter, r *http.Request){
	
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	// Fprintf -> it is used to write the formatted string to the response writer, it is similar to Sprintf but it writes the formatted string to the response writer instead of returning it as a string
}
// this is the snippetView Function, which will hanlde the get request for "/snippet/view/{id}"


func (app *application)snippetCreate(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("Display a form for creating a new snippet..."))
}
// this is the snippetCreate function, which will handle the get request for "/snippet/create"


func (app *application)snippetCreatePost(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new Snippet..."))
}
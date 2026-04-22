package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	// instializing mux as a new serve mux
	mux.HandleFunc("GET /{$}", home)
	// handlefunc -> tell the server that which function will execute when user visit a particular url, in this case the urll is "/"

	// {$} -> it is used to stop the restricting the sub tree or traling slash, url will only work if the exact url would match

	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// 
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	log.Print("Starting the server on th port :4000")
	// log.Print -> it will print the message on the localhost page

	err := http.ListenAndServe(":4000", mux)
	// ListenAndServe -> it will start the server on the port 4000 and it will use mux as a handler to handle the incoming request, if we didn't use mux and create default server then it will we send nil as a handler and it will use the default server to handle the incoming request

	// by using mux or any other handler we can create multiple routes and handle them with different functions at different ports or urls

	// it holds the programm here if we haven't any error and keep the server running unitl manually stop

	if err != nil {
		log.Fatal(err)
	}

	// log.Fatal -> it will print the error on the terminal and kills the programm immediately

}

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello from snippetbox"))
	// what WRITE did -> it write the response in the response body as byte slice

	w.Header().Add("Server", "GO")
	// Header().Add() -> is used to add a custom header the response

}

// w is a response writer which is used to write or basically send the resposne to the client
// r is a request by the clent, basically it is used to read what the user wants to do, it contains all the information about the request like method, url, headers, body etc

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))

	//strconv -> it is a package in go which is used to convert the string to other data types like int, float, etc
	// Atoi -> it is used to convert the string to integer
	// PathValue -> it is used to get the value of the path parameter

	if err != nil || id < 1 {
		http.NotFound(w,r)
		return
	}
	msg := fmt.Sprintf("Display a specific snippet...%d", id)
	// Sprintf -> it is used to format the string and return the formatted string
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}
// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated) 
	// writeHeader -> used to update the status code of the response
	// http.StatusCreated -> we are updating the status code to 201 which means that the resource has been created successfully
	w.Write([]byte("Save a new Snippet...."))
}

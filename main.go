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
	mux.HandleFunc("/{$}", home)
	// handlefunc -> tell the server that which function will execute when user visit a particular url, in this case the urll is "/"

	// {$} -> it is used to stop the restricting the sub tree or traling slash, url will only work if the exact url would match

	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)


	log.Print("Starting the server on th port :4000")
	// log.Print -> it will print the message on the localhost page

	err := http.ListenAndServe(":4000", mux)
	// ListenAndServe -> it will start the server on the port 4000 and it will use mux as a handler to handle the incoming request, if we didn't use mux and create default server then it will we send nil as a handler and it will use the default server to handle the incoming request

	// by using mux or any other handler we can create multiple routes and handle them with different functions at different ports or urls

	if err != nil {
		log.Fatal(err)
	}

	// log.Fatal -> it will print the error message on the localhost page and it will stop the server if there is any error while starting the server
}

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello from snippetbox"))
	// what WRITE did -> it write the response in the response body as byte slice

}

// w is a response writer which is used to write or basically send the resposne to the client
// r is a request by the clent, basically it is used to read what the user wants to do, it contains all the information about the request like method, url, headers, body etc

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi("id")

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return 
	}
	msg := fmt.Sprintf("Display a specific snippet...%d", id)
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

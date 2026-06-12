package main

import "net/http"

func (app *application) routes() *http.ServeMux{


	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// stripPrefix() -> it is used to remove the prefix from the request URL before passing it to the file server, it is used to serve static files from the "./ui/static" directory when the request URL starts with "/static/"

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// create all the handler function
	// we have to use app.func because we have to use the custom logger in the handler function, and the custom logger is a field of the application struct, so we have to use app.func to access the custom logger in the handler function

	return mux
}
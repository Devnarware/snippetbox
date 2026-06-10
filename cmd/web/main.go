package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct{
	logger *slog.Logger	
}

func main(){

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	// stripPrefix() -> it is used to remove the prefix from the request URL before passing it to the file server, it is used to serve static files from the "./ui/static" directory when the request URL starts with "/static/"

	// .handle -> 

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// create all the handler function

	logger.Info("Starting the server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
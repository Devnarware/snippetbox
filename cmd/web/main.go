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
// it is the application struct help us to create custom logger and use it in all the handler function, it is a good practice to create a custom logger for your application, so that you can easily log the error and other information in a structured way, and also you can easily change the logging format and destination in one place, instead of changing it in all the handler function

// why it is a good practice to create a custom logger for your application?
// 1. it helps us to log the error and other information in a structured way, which makes it easier to read and understand the logs
// 2. it helps us to easily change the logging format and destination in one place, instead of changing it in all the handler function
// 3. it helps us to easily add additional fields to the logs, such as request method, request URI, etc, which can be very useful for debugging and monitoring the application

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

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	// create all the handler function
	// we have to use app.func because we have to use the custom logger in the handler function, and the custom logger is a field of the application struct, so we have to use app.func to access the custom logger in the handler function

	logger.Info("Starting the server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)

	logger.Error(err.Error())
	os.Exit(1)
}
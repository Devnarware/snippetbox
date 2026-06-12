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

	mux := app.routes()
	// calling the route function to get the server mux so that we can serve it

	logger.Info("Starting the server", "addr", *addr)

	err := http.ListenAndServe(*addr, mux)
	// we can directly call the route function here without saving it in a variable

	logger.Error(err.Error())
	os.Exit(1)
}
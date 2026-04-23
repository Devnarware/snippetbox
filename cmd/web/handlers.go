package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request){
	
	w.Header().Add("server", "GO")
	w.Write([]byte("Hello from snippet box"))

}
// this is our Home function, which will handle the "/" route


func snippetView(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("Hello from snippet box"))
}



func snippetCreate(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("Hello from snippet box"))
}



func snippetCreatePost(w http.ResponseWriter, r *http.Request){
	
	w.Write([]byte("Hello from snippet box"))
}
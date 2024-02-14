package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/healthz", health)
	mux.HandleFunc("/version", version)
	mux.HandleFunc("/v1/book/view", bookView)
	mux.HandleFunc("/v1/book/create", bookCreate)
	mux.HandleFunc("/v1/user/view", userView)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters: The TCP network address to listen on (in this case ":8080")
	// and the servemux we just created.
	// If http.ListenAndServe() returns an error, we use the log.Fatal() function
	// to log the error message and exit.
	// Note that any error returned by http.ListenAndServe is always non-nil.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

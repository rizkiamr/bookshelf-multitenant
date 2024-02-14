package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Bookshelf" as the response body
func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches "/".
	// If it doesn't, use the http.NotFound() function
	// to send a 404 response to the client.
	// Importantly, we then return from the handler.
	// If we don't return the handler, the handler would keep executing
	// and also write the "Hello from Bookshelf" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Bookshelf"))
}

// Add a bookView handler function
func bookView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display books"))
}

// Add a userView handler function
func userView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display users"))
}

func main() {
	// Use the http.NewServeMux function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/v1/books", bookView)
	mux.HandleFunc("/v1/users", userView)

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

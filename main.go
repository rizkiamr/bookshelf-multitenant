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
	w.Write([]byte("Display book"))
}

// Add a bookCreate handler function
func bookCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST method or not.
	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header
		// to the response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", "POST")

		// If it's not, use the w.WriteHeader() method to send a 405 status code
		// and the w.Write() method to write a "Method Not Allowed" response body.
		// We then return from the function so that the subsequent code is not executed.
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Create a new book"))
}

// Add a userView handler function
func userView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display user"))
}

// health handler function for health check
func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"name\": \"bookshelf-api\", \"healthy\": true}"))
}

// version handler function for app version
func version(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"name\": \"bookshelf-api\", \"version\": \"0.0.1\"}"))
}

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

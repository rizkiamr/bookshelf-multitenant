package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the value of the id parameter from the query string and
	// try to convert it to an integer using strconv.Atoi() function.
	// If it can't be converted to an integer, or the value is less than 1,
	// we return 404 page not found response
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

// Add a bookCreate handler function
func bookCreate(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using POST method or not.
	if r.Method != "POST" {
		// Use the Header().Set() method to add an 'Allow: POST' header
		// to the response header map. The first parameter is the header name, and
		// the second parameter is the header value.
		w.Header().Set("Allow", http.MethodPost)

		// Use the http.Error() function to send a 405 status code and
		// "Method Not Allowed" string as the response body.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
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
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"name": "bookshelf-api", "healthy": true}`))
}

// version handler function for app version
func version(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(`{"name": "bookshelf-api", "version": "0.0.1"}`))
}

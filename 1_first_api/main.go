package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Go API server!")
}

// Handler for the /api route
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, API world!")
}

func main() {
	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api", apiHandler)

	// Start the server
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

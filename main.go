package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Define the handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from a Platform Engineer!")
	})

	// Create a new server and specify the timeouts
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,   // Timeout for reading the request
		WriteTimeout: 10 * time.Second,  // Timeout for writing the response
		IdleTimeout:  15 * time.Second,  // Timeout for idle connections
	}

	// Start the server
	fmt.Println("Server starting on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
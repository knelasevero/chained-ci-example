package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from a Platform Engineer!")
}

func main() {
    server := &http.Server{
        Addr:         ":8080",
        Handler:      http.HandlerFunc(handler),
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  15 * time.Second,
    }

    fmt.Println("Server starting on port 8080...")
    if err := server.ListenAndServe(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
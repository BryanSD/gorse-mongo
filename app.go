package main

import (
    "net/http"

    "github.com/bryansd/gorse-mongo/handlers"
)


func main() {
    // http.HandleFunc("/", handler) Will handle all requests?
    http.HandleFunc("/health", handlers.HealthGet)
    http.ListenAndServe(":8080", nil)
}

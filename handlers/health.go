package handlers

import (
    "fmt"
    "net/http"
)


func HealthGet(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

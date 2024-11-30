package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from Dockerized Go! I an finally working")
    })

    log.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

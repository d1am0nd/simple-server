package main

import (
    "os"
    "fmt"
    "net/http"
)

func main() {
    port := "3000"
    // First argument is port
    if len(os.Args) > 1 {
        port = os.Args[1]
        fmt.Println("port", os.Args[1])
    }

    http.HandleFunc("/", newHandler)

    fmt.Println("Starting simple server on port", port)

    http.ListenAndServe(":" + port, nil)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
    dir, _ := os.Getwd()

    path := r.URL.Path;

    http.ServeFile(w, r, dir + path)
}

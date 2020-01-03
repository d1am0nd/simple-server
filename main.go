package main

import (
    "os"
    "fmt"
    "net/http"
)

func main() {
    port := "3000"
    serveRoot := false

    // First argument is port
    if len(os.Args) > 1 {
        port = os.Args[1]
        fmt.Println("port", os.Args[1])
    }

    // Second argument allows for not appending path
    // but always serving from the root
    // Useful for developing SPAs with routers
    if len(os.Args) > 2 {
        serveRoot = os.Args[2] == "true" || os.Args[2] == "1"
    }

    http.HandleFunc("/", createHandler(serveRoot))

    fmt.Println("Starting simple server on port", port)

    http.ListenAndServe(":" + port, nil)
}

func createHandler(serveRoot bool) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        dir, _ := os.Getwd()

        // Check if file exists
        _, err := os.Stat(dir + r.URL.Path)
        if serveRoot == false || ! os.IsNotExist(err) {
            dir += r.URL.Path
        }

        http.ServeFile(w, r, dir)
    }
}

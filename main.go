package main

import (
    "fmt"
    "net/http"

    "github.com/julienschmidt/httprouter"
)

func main() {
    fmt.Print("Serve on port: ")
    port := ""
    fmt.Scanln(&port)
    if len(port) == 0 {
        port = "3000"
    }

    fmt.Println("Serve from folder: ")
    from := ""
    fmt.Scanln(&from)

    router := NewRouter(from)

    fmt.Println("Serving on port ", port)
    http.ListenAndServe("localhost:" + port, router)
}

func NewRouter(from string) *httprouter.Router {
    r := httprouter.New()
    fmt.Println("Serving folder ", from)
    r.ServeFiles("/*filepath", http.Dir(from))
    return r
}

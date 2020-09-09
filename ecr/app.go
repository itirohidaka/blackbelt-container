package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "<h1>Hello, BlackBelt!</h1> \n %s \n", r.URL.Path[1:])
}

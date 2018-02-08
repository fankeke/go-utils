package main

import (
    "fmt"
    "net/http"
    "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello, ok")
}
func handler_hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func main() {
    if len(os.Args) !=2 {
        fmt.Println("Usage: http <port>")
        os.Exit(1)
    }
    http.HandleFunc("/", handler)
    http.HandleFunc("/hello", handler_hello)
    http.ListenAndServe("127.0.0.1:"+os.Args[1], nil)
}
       

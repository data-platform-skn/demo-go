package main

import (
    "net/http" //package for http based web programs
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Fprintf(w, "Hello Docker world!")
}

func main() {
    http.HandleFunc("/", handler) // redirect all urls to the handler function
    http.ListenAndServe(":9999", nil) // listen for connections at port 9999 on the local machine
}


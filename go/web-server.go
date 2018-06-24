package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth\n")
}

package main

import (
	"fmt"
	"html"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
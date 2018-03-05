package main

import (
    "fmt"
    "html"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func testHandler(w http.ResponseWriter, r *http.Request, response http.Response) {
    fmt.Fprintf(w, "これはテスト〜, %q", html.EscapeString(r.URL.Path))
}

func main() {
    http.HandleFunc("/hello", handler)
    http.ListenAndServe(":8080", nil)
    response, err := http.Get("http://toyokeizai.net/list/feed/rss")
}


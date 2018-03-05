package main

import (
    "fmt"
    "html"
    "net/http"
    "io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    response, err := http.Get("http://toyokeizai.net/list/feed/rss")
    if err != nil {
        return
    }
    bodyBytes, err := ioutil.ReadAll(response.Body)
    fmt.Fprintf(w, "これはテスト〜, %q", string(bodyBytes))
}

func main() {
    http.HandleFunc("/hello", handler)
    http.HandleFunc("/test", testHandler)
    http.ListenAndServe(":8080", nil)
}


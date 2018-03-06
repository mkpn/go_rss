// 参考: https://qiita.com/ytkhs/items/948f516ec82c82eaa882
package main

import (
    "net/http"

)
func main() {
    http.HandleFunc("/test", testHandler)
    http.ListenAndServe(":8080", nil)
}

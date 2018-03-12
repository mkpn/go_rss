// 参考: https://qiita.com/ytkhs/items/948f516ec82c82eaa882
package main

import (
	"net/http"
	"github.com/mkpn/go_rss/controller"
)

func main() {
	http.HandleFunc("/test", controller.RssHandler)
	http.HandleFunc("/rss_json", controller.RssAsJsonHandler)
	http.HandleFunc("/rx", controller.RxHandler)
	http.ListenAndServe(":8080", nil)
}

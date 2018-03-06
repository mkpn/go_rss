package controller

import (
	"net/http"
	"fmt"
	"time"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	xml, err := GetFeedArticles("http://b.hatena.ne.jp/hotentry/it.rss")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, articles := range xml.Articles {
		// timeParseの規格の選びやすさ良き〜〜〜
		datetime, _ := time.Parse(time.RFC3339, articles.Date)

		fmt.Fprintf(w, "%v\n", datetime.Format("2006/01/02 15:04:05"))
		fmt.Fprintf(w, "%s\n", articles.Title)
		fmt.Fprintf(w, "%v\n", articles.Link)
		fmt.Fprintln(w)
	}
}

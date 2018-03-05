// 参考: https://qiita.com/ytkhs/items/948f516ec82c82eaa882
package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"time"
)

type XML struct {
	Articles []struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		// パース前とパース後で全然違うフォーマットになってるのｲﾐﾌﾜﾛﾀ
		// あと元のデータpubDateってなってるのにdateじゃないとアクセスできなかったのもｲﾐﾌ
		Date        string `xml:"date"`
	} `xml:"item"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	xmlFeed := getXmlFeed("http://b.hatena.ne.jp/hotentry/it.rss")

	result := XML{}
	err := xml.Unmarshal([]byte(xmlFeed), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, articles := range result.Articles {
		// timeParseの規格の選びやすさ良き〜〜〜
		datetime, _ := time.Parse(time.RFC3339, articles.Date)

		fmt.Fprintf(w, "%v\n", datetime.Format("2006/01/02 15:04:05"))
		fmt.Fprintf(w, "%s\n", articles.Title)
		fmt.Fprintf(w, "%v\n", articles.Link)
		fmt.Fprintln(w)
	}
}

func getXmlFeed(url string) string { //返り値の構文ミス補完してくれなかった。
	response, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(bodyBytes)
}

func main() {
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}

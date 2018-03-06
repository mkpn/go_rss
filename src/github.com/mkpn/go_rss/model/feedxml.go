package model

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type Xml struct {
	Articles []struct { //これなんで入れ子(内部クラス？)にする必要あるんだ？
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		// パース前とパース後で全然違うフォーマットになってるのｲﾐﾌﾜﾛﾀ
		// あと元のデータpubDateってなってるのにdateじゃないとアクセスできなかったのもｲﾐﾌ
		Date string `xml:"date"`
	} `xml:"item"`
}

func GetFeedArticles(url string) (Xml, error) { //返り値の構文ミス補完してくれなかった。
	response, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	feedXmlStr := string(bodyBytes)
	feedXml := Xml{}
	err := xml.Unmarshal([]byte(feedXmlStr), &feedXml)
	return feedXml, err
}

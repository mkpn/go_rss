package model

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"github.com/mkpn/go_rss/entity"
	"github.com/mkpn/go_rss/data"
)

func GetArticleList(url string) (entity.FeedXml, error) { //返り値の構文ミス補完してくれなかった。
	response, _ := http.Get(url)
	bodyBytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	feedXmlStr := string(bodyBytes)
	feedXml := entity.FeedXml{}
	err := xml.Unmarshal([]byte(feedXmlStr), &feedXml)
	//return feedXml.Articles, err ←本当はこうやって書きたいんだけどどうすれば・・・

	data.Insert()
	return feedXml, err
}

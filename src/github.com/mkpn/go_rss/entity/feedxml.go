package entity


type FeedXml struct {
	Articles []struct { //これなんで入れ子(内部クラス？)にする必要あるんだ？
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		// パース前とパース後で全然違うフォーマットになってるのｲﾐﾌﾜﾛﾀ
		// あと元のデータpubDateってなってるのにdateじゃないとアクセスできなかったのもｲﾐﾌ
		Date string `xml:"date"`
	} `xml:"item"`
}

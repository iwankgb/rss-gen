package rss

type Channel struct {
	XMLName       string  `xml:"rss"`
	Version       float32 `xml:"attr,version"`
	Title         string  `xml:"channel>title"`
	Link          string  `xml:"channel>link"`
	Description   string  `xml:"channel>description"`
	Language      string  `xml:"channel>language"`
	LastBuildDate string  `xml:"channel>lastBuildDate"`
	Items         []Item  `xml:"channel>item"`
}

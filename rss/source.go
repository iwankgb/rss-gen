package rss

type Source struct {
	Url   string `xml:"url,attr"`
	Value string `xml:",chardata"`
}

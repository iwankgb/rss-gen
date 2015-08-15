package rss

type Channel struct {
	XMLName       string `xml:"rss"`
	Version       string `xml:"version,attr"`
	Title         string `xml:"channel>title"`
	Link          string `xml:"channel>link"`
	Description   string `xml:"channel>description"`
	Language      string `xml:"channel>language"`
	LastBuildDate string `xml:"channel>lastBuildDate"`
	Generator     string `xml:"channel>generator"`
	Items         []Item `xml:"channel>item"`
}

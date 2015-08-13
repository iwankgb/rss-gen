package rss

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Source      string `xml:"source"`
	PubDate     string `xml:"pubDate"`
}

package rss

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Source      Source `xml:"source"`
	Guid        Guid   `xml:"guid"`
	PubDate     string `xml:"pubDate"`
}

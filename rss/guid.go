package rss

type Guid struct {
	Value  string `xml:",chardata"`
	IsLink bool   `xml:"isPermaLink,attr"`
}

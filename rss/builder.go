// Rss builder funtions
package rss

import y "github.com/iwankgb/rss-gen/yaml"
import t "time"
import s "crypto/sha512"
import "fmt"

func NewBuilder(knownDates map[string]*string, itemLimit *int) *rssBuilder {
	return &rssBuilder{
		knownDates,
		itemLimit,
		Channel{},
	}
}

type rssBuilder struct {
	knownDates map[string]*string
	itemLimit  *int
	rss        Channel
}

// Builds RSS object from yaml representation
func (builder *rssBuilder) BuildRssFromYaml(yaml *y.Channel) {
	builder.buildChannel(yaml)
	builder.buildItems(yaml)
}

// Builds main channel object
func (builder *rssBuilder) buildChannel(yaml *y.Channel) {
	builder.rss.Version = "2.0"
	builder.rss.Description = yaml.Description
	builder.rss.Language = yaml.Language
	builder.rss.LastBuildDate = t.Now().Format(t.RFC822)
	builder.rss.Link = yaml.Link
	builder.rss.Title = yaml.Title
	builder.rss.Generator = "https://github.com/iwankgb/rss-gen"
}

// Builds array of items for RSS
func (builder *rssBuilder) buildItems(yaml *y.Channel) {
	for i := 0; i < len(yaml.Items) && i < *builder.itemLimit; i++ {
		item := Item{}
		item.Description = yaml.Items[i].Description
		item.Link = yaml.Items[i].Link
		item.Source.Value = yaml.Items[i].Source.Name
		item.Source.Url = yaml.Items[i].Source.Url
		item.Title = yaml.Items[i].Title
		item.Guid.IsLink = false
		item.Guid.Value = fmt.Sprintf("%x", s.Sum512([]byte(item.Link)))
		existingDate, dateExists := builder.knownDates[item.Link]
		if dateExists {
			item.PubDate = *existingDate
		} else {
			item.PubDate = t.Now().Format(t.RFC822)
		}

		builder.rss.Items = append(builder.rss.Items, item)
	}
}

func (builder *rssBuilder) GetRss() *Channel {
	return &builder.rss
}

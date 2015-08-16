// Rss builder funtions
package rss

import y "github.com/iwankgb/rss-gen/yaml"
import t "time"
import s "crypto/sha512"
import "fmt"

func NewBuilder(knownDates DateResolver, itemLimit *int) *rssBuilder {
	return &rssBuilder{
		knownDates,
		itemLimit,
		Channel{},
	}
}

type rssBuilder struct {
	knownDates DateResolver
	itemLimit  *int
	rss        Channel
}

func (builder *rssBuilder) BuildRssFromYaml(yaml *y.Channel) {
	builder.buildChannelElement(yaml)
	builder.buildItemElements(yaml)
}

func (builder *rssBuilder) buildChannelElement(yaml *y.Channel) {
	builder.rss.Version = "2.0"
	builder.rss.Description = yaml.Description
	builder.rss.Language = yaml.Language
	builder.rss.LastBuildDate = t.Now().Format(t.RFC822)
	builder.rss.Link = yaml.Link
	builder.rss.Title = yaml.Title
	builder.rss.Generator = "https://github.com/iwankgb/rss-gen"
}

func (builder *rssBuilder) buildItemElements(yaml *y.Channel) {
	realLength := len(yaml.Items)
	for i := realLength - 1; i >= realLength-*builder.itemLimit && i >= 0; i-- {
		item := Item{}
		item.Description = yaml.Items[i].Description
		item.Link = yaml.Items[i].Link
		item.Source.Value = yaml.Items[i].Source.Name
		item.Source.Url = yaml.Items[i].Source.Url
		item.Title = yaml.Items[i].Title
		item.Guid.IsLink = false
		item.Guid.Value = fmt.Sprintf("%x", s.Sum512([]byte(item.Link)))
		item.PubDate = builder.knownDates.GetDate(&item)
		builder.rss.Items = append(builder.rss.Items, item)
	}
}

func (builder *rssBuilder) GetRss() *Channel {
	return &builder.rss
}

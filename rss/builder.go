// Rss builder funtions
package rss

import y "github.com/iwankgb/rss-gen/yaml"
import t "time"

// Builds RSS object from yaml representation
func BuildRssFromYaml(yaml y.Channel, itemLimit int) (rss Channel) {
	buildChannel(&rss, &yaml)
	buildItems(&rss, &yaml, itemLimit)
	return rss
}

// Builds main channel object
func buildChannel(rss *Channel, yaml *y.Channel) {
	rss.Version = 2.0
	rss.Description = yaml.Description
	rss.Language = yaml.Language
	rss.LastBuildDate = t.Now().Format(t.RFC822)
	rss.Link = yaml.Link
	rss.Title = yaml.Title
}

// Builds array of items for RSS
func buildItems(rss *Channel, yaml *y.Channel, itemLimit int) {
	for i := 0; i < len(yaml.Items) && i < itemLimit; i++ {
		item := Item{}
		item.Description = yaml.Items[i].Description
		item.Link = yaml.Items[i].Link
		item.PubDate = t.Now().Format(t.RFC822)
		item.Source = yaml.Items[i].Source
		item.Title = yaml.Items[i].Title
		rss.Items = append(rss.Items, item)
	}
}

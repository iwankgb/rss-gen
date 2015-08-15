// Rss builder funtions
package rss

import y "github.com/iwankgb/rss-gen/yaml"
import t "time"
import s "crypto/sha512"

import "fmt"

// Builds RSS object from yaml representation
func BuildRssFromYaml(yaml y.Channel, itemLimit int, dates map[string]*string) (rss Channel) {
	buildChannel(&rss, &yaml)
	buildItems(&rss, &yaml, itemLimit, dates)
	return rss
}

// Builds main channel object
func buildChannel(rss *Channel, yaml *y.Channel) {
	rss.Version = "2.0"
	rss.Description = yaml.Description
	rss.Language = yaml.Language
	rss.LastBuildDate = t.Now().Format(t.RFC822)
	rss.Link = yaml.Link
	rss.Title = yaml.Title
	rss.Generator = "https://github.com/iwankgb/rss-gen"
}

// Builds array of items for RSS
func buildItems(rss *Channel, yaml *y.Channel, itemLimit int, dates map[string]*string) {
	//	fmt.Println(len(yaml.Items))
	for i := 0; i < len(yaml.Items) && i < itemLimit; i++ {
		item := Item{}
		item.Description = yaml.Items[i].Description
		item.Link = yaml.Items[i].Link
		existingDate, dateExists := dates[item.Link]
		if dateExists {
			item.PubDate = *existingDate
		} else {
			item.PubDate = t.Now().Format(t.RFC822)
		}
		item.Source.Value = yaml.Items[i].Source.Name
		item.Source.Url = yaml.Items[i].Source.Url
		item.Title = yaml.Items[i].Title
		item.Guid.IsLink = false
		item.Guid.Value = fmt.Sprintf("%x", s.Sum512([]byte(item.Link)))
		//		fmt.Printf("%x\n", item.Guid.Value)
		//		fmt.Printf("%+v", item)
		rss.Items = append(rss.Items, item)
	}
}

package rss

import (
	"github.com/iwankgb/rss-gen/yaml"
	"testing"
	"time"
)

func TestBuildingRss(t *testing.T) {
	yamlObject := prepareYamlObject()
	rss := createBuilderAndBuild(yamlObject)
	assertChannel(rss, yamlObject, t)
	assertItems(rss, t)
}

func prepareYamlObject() *yaml.Channel {
	yamlObject := new(yaml.Channel)
	yamlObject.Description = "This is Channel description"
	yamlObject.Language = "en"
	yamlObject.Link = "http://example.com"
	yamlObject.Title = "The Title"
	firstItem := yaml.Item{}
	firstItem.Description = "First item description"
	firstItem.Link = "http://example.com/1"
	firstItem.Source.Url = "http://example.com/some/rss/feed"
	firstItem.Source.Name = "A source name"
	firstItem.Title = "What a wonderful title!"
	secondItem := yaml.Item{}
	secondItem.Description = "Second item description"
	secondItem.Link = "http://example.com/2"
	secondItem.Source.Url = "http://example.com/yet/another/feed"
	secondItem.Source.Name = "Some random source name"
	secondItem.Title = "అందమైన టైటిల్"
	thirdItem := yaml.Item{}

	yamlObject.Items = append(yamlObject.Items, firstItem, secondItem, thirdItem)
	return yamlObject
}

func createBuilderAndBuild(yamlObject *yaml.Channel) *Channel {
	dates := new(knownDatesMock)
	limit := 2
	rssBuilder := NewBuilder(dates, &limit)
	rssBuilder.BuildRssFromYaml(yamlObject)
	return rssBuilder.GetRss()
}

func assertChannel(rss *Channel, yamlObject *yaml.Channel, t *testing.T) {
	if rss.Description != yamlObject.Description {
		t.Errorf("Wrong description: found %v instead of %v", rss.Description, yamlObject.Description)
	}
	if rss.Version != "2.0" {
		t.Error("Version 2.0 needs to be specified")
	}
	if rss.Language != yamlObject.Language {
		t.Errorf("Wrong language: found %v instead of %v", rss.Language, yamlObject.Language)
	}
	if !isCorrectDate(rss.LastBuildDate) {
		t.Errorf("Wrong build date: found %v", rss.LastBuildDate)
	}
	if rss.Link != yamlObject.Link {
		//		t.Errorf("Wrong link: found %v instead of %v", rss.Link, yamlObject.Link)
	}
	if rss.Title != yamlObject.Title {
		t.Errorf("Wrong title: found %v instead of %v", rss.Title, yamlObject.Title)
	}
	if rss.Generator != "https://github.com/iwankgb/rss-gen" {
		t.Errorf("Wrong generator: found %v instead of https://github.com/iwankgb/rss-gen", rss.Generator)
	}
}

func isCorrectDate(date string) bool {
	rssUpdateTime, _ := time.Parse(time.RFC822, date)
	return rssUpdateTime.Unix()-time.Now().Unix() <= 5
}

func assertItems(rss *Channel, t *testing.T) {
	if len(rss.Items) != 2 {
		t.Errorf("Wrong items count: found %v instead of 1", len(rss.Items))
	} else {
		if rss.Items[0].Description != "First item description" {
			t.Errorf("Wrong item description: found %v instead of First item description\n", rss.Items[0].Description)
		}
		if rss.Items[0].Link != "http://example.com/1" {
			t.Errorf("Wrong item link: found %v instead of http://example.com/1\n", rss.Items[0].Link)
		}
		if !isCorrectDate(rss.Items[0].PubDate) {
			t.Errorf("Wrong publication date found: %v\n", rss.Items[0].PubDate)
		}
		if rss.Items[0].Source.Url != "http://example.com/some/rss/feed" {
			t.Errorf("Wrong source URL: found %v instead of http://example.com/some/rss/feed\n", rss.Items[0].Source.Url)
		}
		if rss.Items[0].Source.Value != "A source name" {
			t.Errorf("Wrong source name: found %v instead of A source name\n", rss.Items[1].Source.Value)
		}
		if rss.Items[0].Title != "What a wonderful title!" {
			t.Errorf("Wrong title: found %v instead of What a wonderful title!\n", rss.Items[0].Title)
		}
		if rss.Items[0].Guid.IsLink != false {
			t.Error("Guid is not a link")
		}
		if rss.Items[0].Guid.Value != "c683bad1870807ad44a9b0413f9b4ac9764e1a0c00ce1eadb4e3e50dd09a2372e921a6ff3e7d2ec7c6ed5023d25900b07f8a7c226979444895acd2d2eab7b981" {
			t.Errorf("Wrong guid value: found %v instead of c683bad1870807ad44a9b0413f9b4ac9764e1a0c00ce1eadb4e3e50dd09a2372e921a6ff3e7d2ec7c6ed5023d25900b07f8a7c226979444895acd2d2eab7b981\n", rss.Items[0].Guid.Value)
		}
		if rss.Items[1].Description != "Second item description" {
			t.Errorf("Wrong item description: found %v instead of First item description\n", rss.Items[1].Description)
		}
		if rss.Items[1].Link != "http://example.com/2" {
			t.Errorf("Wrong item link: found %v instead of http://example.com/2\n", rss.Items[1].Link)
		}
		if rss.Items[1].PubDate != "very fake date" {
			t.Errorf("Wrong publication date found: %v instead of very fake date\n", rss.Items[1].PubDate)
		}
		if rss.Items[1].Source.Url != "http://example.com/yet/another/feed" {
			t.Errorf("Wrong source URL: found %v instead of http://example.com/yet/another/feed\n", rss.Items[1].Source)
		}
		if rss.Items[1].Source.Value != "Some random source name" {
			t.Errorf("Wrong source name: found %v instead of Some random source name\n", rss.Items[1].Source.Value)
		}
		if rss.Items[1].Title != "అందమైన టైటిల్" {
			t.Errorf("Wrong title: found %v instead of అందమైన టైటిల్\n", rss.Items[1].Title)
		}
		if rss.Items[1].Guid.IsLink != false {
			t.Errorf("Guid is not a link")
		}
		if rss.Items[1].Guid.Value != "859c263e81032ee1b9cc8225727124bebfa8de52ceb61ef8b8a82355ae2dbecb678d416aecf154970e57a80d81446b48269b9773b6bedecde1fb45c3259f66ce" {
			t.Errorf("Wrong guid value: found %v instead of 859c263e81032ee1b9cc8225727124bebfa8de52ceb61ef8b8a82355ae2dbecb678d416aecf154970e57a80d81446b48269b9773b6bedecde1fb45c3259f66ce\n", rss.Items[1].Guid.Value)
		}
	}
}

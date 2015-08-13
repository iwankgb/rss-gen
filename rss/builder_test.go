package rss

import (
	"github.com/iwankgb/rss-gen/yaml"
	"testing"
	"time"
)

func TestBuilding(t *testing.T) {
	yamlObject := prepareYamlObject()
	rss := BuildRssFromYaml(yamlObject, 1)
	assertChannel(&rss, &yamlObject, t)
	assertItems(&rss, t)
}

// Prepares input structure
func prepareYamlObject() (yamlObject yaml.Channel) {
	yamlObject.Description = "This is Channel description"
	yamlObject.Language = "en"
	yamlObject.Link = "http://example.com"
	yamlObject.Title = "The Title"
	firstItem := yaml.Item{}
	firstItem.Description = "First item description"
	firstItem.Link = "http://example.com/1"
	firstItem.Source = "http://example.com/some/rss/feed"
	firstItem.Title = "What a wonderful title!"
	secondItem := yaml.Item{}
	secondItem.Description = "Second item description"
	secondItem.Link = "http://example.com/2"
	secondItem.Source = "http://example.com/yet/another/feed"
	secondItem.Title = "అందమైన టైటిల్"
	yamlObject.Items = append(yamlObject.Items, firstItem, secondItem)
	return yamlObject
}

// Asserts structure of channel
func assertChannel(rss *Channel, yamlObject *yaml.Channel, t *testing.T) {
	if rss.Description != yamlObject.Description {
		t.Errorf("Wrong description: found %v instead of %v", rss.Description, yamlObject.Description)
	}
	if rss.Version != 2.0 {
		t.Error("Version 2.0 needs to be specified")
	}
	if rss.Language != yamlObject.Language {
		t.Errorf("Wrong language: found %v instead of %v", rss.Language, yamlObject.Language)
	}
	if !isCorrectDate(rss.LastBuildDate) {
		t.Errorf("Wrong build date: found %v", rss.LastBuildDate)
	}
	if rss.Link != yamlObject.Link {
		t.Errorf("Wrong link: found %v instead of %v", rss.Link, yamlObject.Link)
	}
	if rss.Title != yamlObject.Title {
		t.Errorf("Wrong title: found %v instead of %v", rss.Title, yamlObject.Title)
	}
}

// Dirty way of validation the date
func isCorrectDate(date string) bool {
	rssUpdateTime, _ := time.Parse(time.RFC822, date)
	return rssUpdateTime.Unix()-time.Now().Unix() <= 5
}

// Asserts structure if channel items
func assertItems(rss *Channel, t *testing.T) {
	if len(rss.Items) != 1 {
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
		if rss.Items[0].Source != "http://example.com/some/rss/feed" {
			t.Errorf("Wrong source: found %v instead of http://example.com/some/rss/feed\n", rss.Items[0].Source)
		}
		if rss.Items[0].Title != "What a wonderful title!" {
			t.Errorf("Wrong title: found %v instead of అందమైన టైటిల్What a wonderful title!\n", rss.Items[0].Title)
		}
	}
}

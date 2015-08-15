package rss

import (
	"github.com/iwankgb/rss-gen/yaml"
	"testing"
	"time"
)

func TestBuilding(t *testing.T) {
	yamlObject := prepareYamlObject()
	dates := make(map[string]*string)
	fakeDate := "very fake date"
	dates["http://example.com/2"] = &fakeDate
	rss := BuildRssFromYaml(yamlObject, 2, dates)
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
	thirdItem := yaml.Item{}

	yamlObject.Items = append(yamlObject.Items, firstItem, secondItem, thirdItem)
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
		if rss.Items[0].Source != "http://example.com/some/rss/feed" {
			t.Errorf("Wrong source: found %v instead of http://example.com/some/rss/feed\n", rss.Items[0].Source)
		}
		if rss.Items[0].Title != "What a wonderful title!" {
			t.Errorf("Wrong title: found %v instead of What a wonderful title!\n", rss.Items[0].Title)
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
		if rss.Items[1].Source != "http://example.com/yet/another/feed" {
			t.Errorf("Wrong source: found %v instead of http://example.com/yet/another/feed\n", rss.Items[1].Source)
		}
		if rss.Items[1].Title != "అందమైన టైటిల్" {
			t.Errorf("Wrong title: found %v instead of అందమైన టైటిల్\n", rss.Items[1].Title)
		}
	}
}

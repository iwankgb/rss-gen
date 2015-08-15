package rss

import (
	"testing"
)

func TestDatesDictionary(t *testing.T) {
	rssObject := new(Channel)
	firstItem := new(Item)
	firstItem.Link = "http://example.com/1"
	firstItem.PubDate = "first date"
	secondItem := new(Item)
	secondItem.Link = "http://example.com/2"
	secondItem.PubDate = "second date"
	rssObject.Items = append(rssObject.Items, *firstItem, *secondItem)

	dates := NewDates(rssObject)
	firstDate, firstDateError := dates["http://example.com/1"]
	if firstDateError != true {
		t.Error("No date for http://example.com/1")
	}
	if *firstDate != "first date" {
		t.Errorf("Wrong date for http://example.com/1; expected first date got %v", firstDate)
	}
	secondDate, secondDateError := dates["http://example.com/2"]
	if secondDateError != true {
		t.Error("No date for http://example.com/1")
	}
	if *secondDate != "second date" {
		t.Errorf("Wrong date for http://example.com/1; expected second date got %v", firstDate)
	}
}

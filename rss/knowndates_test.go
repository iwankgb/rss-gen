package rss

import "testing"

func TestKnownDate(t *testing.T) {
	item := new(Item)
	item.Link = "known url"
	knownDates := prepareKnownDates()
	date := knownDates.GetDate(item)
	if date != "known date" {
		t.Errorf("Invalid date: %s", date)
	}
}

func prepareKnownDates() *knownDates {
	datesMap := make(map[string]*string, 1)
	date := "known date"
	datesMap["known url"] = &date
	return &knownDates{datesMap}
}

func TestUknownDate(t *testing.T) {
	item := new(Item)
	item.Link = "unknown url"
	item.PubDate = "a brand new date"
	knownDates := prepareKnownDates()
	date := knownDates.GetDate(item)
	if !isCorrectDate(date) {
		t.Errorf("Invalid date: %s", date)
	}
}

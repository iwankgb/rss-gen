package rss

import t "time"

func NewKnownDates(dates map[string]*string) *knownDates {
	return &knownDates{dates}
}

type knownDates struct {
	dates map[string]*string
}

func (kd *knownDates) GetDate(item *Item) (correctDate string) {
	datePointer, dateExists := kd.dates[item.Link]
	if dateExists {
		correctDate = *datePointer
	} else {
		correctDate = t.Now().Format(t.RFC822)
	}
	return correctDate
}

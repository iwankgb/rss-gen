package rss

func NewDates(rss *Channel) *knownDates {
	dates := make(map[string]*string)
	for i := 0; i < len(rss.Items); i++ {
		dates[rss.Items[i].Link] = &rss.Items[i].PubDate
	}
	return NewKnownDates(dates)
}

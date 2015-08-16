package rss

type DateResolver interface {
	GetDate(item *Item) string
}

func NewDates(rss *Channel) DateResolver {
	dates := make(map[string]*string)
	for i := 0; i < len(rss.Items); i++ {
		dates[rss.Items[i].Link] = &rss.Items[i].PubDate
	}
	return NewKnownDates(dates)
}

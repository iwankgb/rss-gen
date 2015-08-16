package rss

type knownDatesMock struct{}

func (*knownDatesMock) GetDate(*Item) string {
	return "very fake date"
}

package site

import "time"

type (
	Options func(options *options)

	options struct {
		url *string
	}
)

func WithUrl(url string) Options {
	return func(options *options) {
		options.url = &url
	}
}

type Top struct {
	SpiderTime time.Time `json:"spider_time"`
	Site       string    `json:"site"`
	Ranking    int32     `json:"ranking"`
	Title      string    `json:"title"`
	URL        string    `json:"url"`
	Extra      string    `json:"extra"`
	Original   string    `json:"original"`
}

func (t Top) StringSlice() []string {
	return []string{
		t.SpiderTime.Format(time.DateTime),
		t.Site,
		t.Title,
		t.URL,
		t.Extra,
		t.Original,
	}
}

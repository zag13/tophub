package endpoint

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
	Endpoint   string    `json:"endpoint"`
	Ranking    int64     `json:"ranking"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	Extra      string    `json:"extra"`
	Original   string    `json:"original"`
}

func (t Top) StringSlice() []string {
	return []string{
		t.SpiderTime.Format("2006-01-02 15:04:05"),
		t.Endpoint,
		t.Title,
		t.Url,
		t.Extra,
		t.Original,
	}
}

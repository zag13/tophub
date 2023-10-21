package site

import (
	"github.com/spf13/cast"
	"time"
)

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
	Rank       int32     `json:"rank"`
	Title      string    `json:"title"`
	Url        string    `json:"url"`
	Extra      string    `json:"extra"`
}

func (t Top) StringSlice() []string {
	return []string{
		t.SpiderTime.Format(time.DateTime),
		t.Site,
		cast.ToString(t.Rank),
		t.Title,
		t.Url,
		t.Extra,
	}
}

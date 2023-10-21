package site

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"time"
)

const (
	ZHIHU_SITE   = "zhihu"
	ZHIHU_URL    = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	ZHIHU_PREFIX = "https://www.zhihu.com/question/"
)

func ZhiHu(opts ...Options) (tops []Top, err error) {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	var url string
	if options.url != nil && *options.url != "" {
		url = *options.url
	} else {
		url = ZHIHU_URL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	j, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}

	spiderTime := time.Now()
	data := j.Get("data")
	for i := range data.MustArray() {
		datum := data.GetIndex(i)

		extra, err := json.Marshal(map[string]string{
			"image": datum.Get("children").GetIndex(0).Get("thumbnail").MustString(),
		})
		if err != nil {
			return nil, err
		}

		tops = append(tops, Top{
			SpiderTime:  spiderTime,
			Site:        ZHIHU_SITE,
			Rank:        cast.ToInt32(i + 1),
			Title:       datum.Get("target").Get("title").MustString(),
			Url:         ZHIHU_PREFIX + cast.ToString(datum.Get("target").Get("id").MustInt64()),
			Description: datum.Get("target").Get("excerpt").MustString(),
			Extra:       cast.ToString(extra),
		})
	}

	return tops, nil
}

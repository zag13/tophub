package endpoint

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"time"
)

const (
	ZHIHU_ENDPOINT   = "zhihu"
	ZHIHU_URL        = "https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true"
	ZHIHU_URL_PREFIX = "https://www.zhihu.com/question/"
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

	data := j.Get("data")
	for i := range data.MustArray() {
		datum := data.GetIndex(i)

		extra, err := json.Marshal(map[string]string{
			"description": datum.Get("target").Get("excerpt").MustString(),
			"image":       datum.Get("children").GetIndex(0).Get("thumbnail").MustString(),
		})
		if err != nil {
			return nil, err
		}
		original, err := json.Marshal(datum)
		if err != nil {
			return nil, err
		}

		tops = append(tops, Top{
			SpiderTime: time.Now(),
			Endpoint:   ZHIHU_ENDPOINT,
			Ranking:    cast.ToInt64(i + 1),
			Title:      datum.Get("target").Get("title").MustString(),
			Url:        ZHIHU_URL_PREFIX + datum.Get("target").Get("id").MustString(),
			Extra:      cast.ToString(extra),
			Original:   cast.ToString(original),
		})
	}

	return tops, nil
}

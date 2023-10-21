package site

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

const (
	WEIBO_SITE   = "weibo"
	WEIBO_URL    = "https://s.weibo.com/top/summary"
	WEIBO_PREFIX = "https://s.weibo.com"
	// WEIBO_URL2    = "https://m.weibo.cn/api/container/getIndex?containerid=106003type%3D25%26t%3D3%26disable_hot%3D1%26filter_type%3Drealtimehot"
)

func WeiBo(opts ...Options) (tops []Top, err error) {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	var url string
	if options.url != nil && *options.url != "" {
		url = *options.url
	} else {
		url = WEIBO_URL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Set("cookie", "SUB=_2AkMV0R1pdcPxrAVTn_EQyWvgZY5H-jymBHSfAn7uJhIyOhhu7nECqSVutBF-XDXcz_XKfP8ConVYk8o_32dN56GH; SUBP=0033WrSXqPxfM72wWs9jqgMF555t")

	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	spiderTime := time.Now()
	doc.Find("#pl_top_realtimehot table tbody tr").Each(func(i int, selection *goquery.Selection) {
		rank := cast.ToInt32(selection.Find(".td-01").Text())
		if rank == 0 {
			return
		}

		u, _ := selection.Find(".td-02 a").Attr("href")
		tops = append(tops, Top{
			SpiderTime: spiderTime,
			Site:       WEIBO_SITE,
			Rank:       rank,
			Title:      selection.Find(".td-02 a").Text(),
			Url:        WEIBO_PREFIX + u,
			Extra:      "{}",
		})
	})
	return
}

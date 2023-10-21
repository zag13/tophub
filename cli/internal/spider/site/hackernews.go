package site

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cast"
	"net/http"
	"strings"
	"time"
)

const (
	HACKERNEWS_SITE   = "hacker-news"
	HACKERNEWS_URL    = "https://news.ycombinator.com/front"
	HACKERNEWS_PREFIX = "https://news.ycombinator.com/"
)

func HackerNews(opts ...Options) (tops []Top, err error) {
	var options options
	for _, opt := range opts {
		opt(&options)
	}

	var url string
	if options.url != nil && *options.url != "" {
		url = *options.url
	} else {
		url = HACKERNEWS_URL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")

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
	doc.Find("#hnmain tbody tr:nth-child(3) td table tbody tr[class='athing']").Each(func(i int, selection *goquery.Selection) {
		tops = append(tops, Top{
			SpiderTime: spiderTime,
			Site:       HACKERNEWS_SITE,
			Rank:       cast.ToInt32(strings.Trim(selection.Find("span[class='rank']").Text(), ".")),
			Title:      selection.Find("span[class='titleline']").Text(),
			Url:        HACKERNEWS_PREFIX + selection.Next().Find("span[class='subline'] a:nth-child(1)").AttrOr("href", ""),
			Extra:      "{}",
		})
	})
	return tops, nil
}

package site

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const (
	V2EX_SITE = "v2ex"
	V2EX_URL  = "https://www.v2ex.com/?tab=hot"
)

func V2EX(ctx context.Context) (tops []Top, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", V2EX_URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(".item_title").Each(func(i int, selection *goquery.Selection) {
		_, _ = selection.Find("a").Attr("href")
		_ = selection.Find("a").Text()
	})
	return
}

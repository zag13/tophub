package site

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const (
	GITHUB_SITE = "github"
	GITHUB_URL  = "https://github.com/trending"
)

func Github(ctx context.Context) (tops []Top, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", GITHUB_URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("TODO").Each(func(i int, selection *goquery.Selection) {

	})
	return
}

package storage

import (
	"context"
	"github.com/zag13/tophub/cli/bootstrap"
	"github.com/zag13/tophub/cli/dal"
	"github.com/zag13/tophub/cli/dal/model"
	"github.com/zag13/tophub/cli/internal/spider/site"
	"time"
)

func DB(tops []site.Top) error {
	bootstrap.InitDatabase()

	news := make([]*model.News, len(tops))
	spiderTime := time.Now()
	for i, top := range tops {
		news[i] = &model.News{
			SpiderTime: spiderTime,
			Site:       top.Site,
			Ranking:    top.Ranking,
			Title:      top.Title,
			URL:        top.URL,
		}
	}

	return dal.Q.News.WithContext(context.Background()).CreateInBatches(news, 20)
}

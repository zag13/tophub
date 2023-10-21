package storage

import (
	"context"
	"github.com/zag13/tophub/cli/bootstrap"
	"github.com/zag13/tophub/cli/dal"
	"github.com/zag13/tophub/cli/dal/model"
	"github.com/zag13/tophub/cli/internal/spider/site"
	"github.com/zag13/tophub/cli/utils/stringz"
)

func DB(tops []site.Top) error {
	bootstrap.InitDatabase()

	topsInsert := make([]*model.Top, len(tops))
	for i, top := range tops {
		topsInsert[i] = &model.Top{
			SpiderTime:  top.SpiderTime,
			Site:        top.Site,
			Rank:        top.Rank,
			Title:       stringz.TruncateString(top.Title, 128),
			URL:         top.Url,
			Description: stringz.TruncateString(top.Description, 512),
			Extra:       top.Extra,
		}
	}

	return dal.Q.Top.WithContext(context.Background()).CreateInBatches(topsInsert, 20)
}

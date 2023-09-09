package dal

import (
	"context"
	"github.com/zag13/tophub/server/dal/model"
)

var _ NewsModel = (*customNewsModel)(nil)

type (
	NewsModel interface {
		FindOne(ctx context.Context, id int64) (*model.News, error)
		FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.News, error)
	}

	customNewsModel struct {
		Q *Query
	}
)

func NewNewsModel(q *Query) NewsModel {
	return &customNewsModel{Q: q}
}

func (m *customNewsModel) FindOne(ctx context.Context, id int64) (*model.News, error) {
	news := m.Q.News

	return m.Q.News.WithContext(ctx).Where(news.ID.Eq(id)).Take()
}

func (m *customNewsModel) FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.News, error) {
	do := m.Q.News.WithContext(ctx)
	news := m.Q.News

	qry = PreprocessForQuery(qry)
	if val, ok := qry["Id"].(int64); ok && val != 0 {
		do.Where(news.ID.Eq(val))
	} else {
		if val, ok := qry["Site"].(string); ok && val != "" {
			do.Where(news.Site.Eq(val))
		}
	}

	offset, limit := GetOffsetAndLimit(page, pageSize)
	return do.Offset(offset).Limit(limit).Find()
}

package dal

import (
	"context"
	"github.com/zag13/tophub/server/dal/model"
	"gorm.io/gen"
)

var _ NewsModel = (*customNewsModel)(nil)

type (
	NewsModel interface {
		FindOne(ctx context.Context, id int64) (*model.News, error)
		FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.News, error)
		FindLatest(ctx context.Context, qry map[string]any) ([]*model.News, error)
	}

	customNewsModel struct {
		Q *Query
	}
)

func NewNewsModel(q *Query) NewsModel {
	return &customNewsModel{Q: q}
}

func (m *customNewsModel) FindOne(ctx context.Context, id int64) (*model.News, error) {
	return m.Q.News.WithContext(ctx).FindOne(id)
}

func (m *customNewsModel) FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.News, error) {
	news := m.Q.News
	var conds []gen.Condition

	qry = PreprocessForQuery(qry)
	if val, ok := qry["Id"].(int64); ok && val != 0 {
		conds = append(conds, news.ID.Eq(val))
	} else {
		if val, ok := qry["Site"].(string); ok && val != "" {
			conds = append(conds, news.Site.Eq(val))
		}
	}

	offset, limit := GetOffsetAndLimit(page, pageSize)

	return m.Q.News.WithContext(ctx).Offset(offset).Limit(limit).Find()
}

func (m *customNewsModel) FindLatest(ctx context.Context, qry map[string]any) ([]*model.News, error) {
	news := m.Q.News
	var conds []gen.Condition

	qry = PreprocessForQuery(qry)
	if val, ok := qry["Id"].(int64); ok && val != 0 {
		conds = append(conds, news.ID.Eq(val))
	} else {
		if val, ok := qry["Site"].(string); ok && val != "" {
			conds = append(conds, news.Site.Eq(val))
		}
	}

	conds = append(conds, news.WithContext(ctx).Columns(news.SpiderTime).Eq(news.WithContext(ctx).Select(news.SpiderTime.Max())))

	return m.Q.News.WithContext(ctx).Where(conds...).Find()
}

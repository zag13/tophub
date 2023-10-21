package dal

import (
	"context"
	"github.com/zag13/tophub/server/dal/model"
	"gorm.io/gen"
)

var _ TopsModel = (*customTopsModel)(nil)

type (
	TopsModel interface {
		FindOne(ctx context.Context, id int64) (*model.Top, error)
		FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.Top, error)
		FindLatest(ctx context.Context, qry map[string]any) ([]*model.Top, error)
	}

	customTopsModel struct {
		Q *Query
	}
)

func NewTopsModel(q *Query) TopsModel {
	return &customTopsModel{Q: q}
}

func (m *customTopsModel) FindOne(ctx context.Context, id int64) (*model.Top, error) {
	return m.Q.Top.WithContext(ctx).FindOne(id)
}

func (m *customTopsModel) FindPage(ctx context.Context, qry map[string]any, page, pageSize int) ([]*model.Top, error) {
	news := m.Q.Top
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

	return m.Q.Top.WithContext(ctx).Where(conds...).Offset(offset).Limit(limit).Find()
}

func (m *customTopsModel) FindLatest(ctx context.Context, qry map[string]any) ([]*model.Top, error) {
	news := m.Q.Top
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

	return m.Q.Top.WithContext(ctx).Where(conds...).Find()
}

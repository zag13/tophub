package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"tophub/app/topdata/dal/model"
	"tophub/app/topdata/service/internal/biz"
	"tophub/pkg/util/pagination"
)

var _ biz.TopdataRepo = (*topdataRepo)(nil)

type topdataRepo struct {
	data *Data
	log  *log.Helper
}

func NewTopdataRepo(data *Data, logger log.Logger) biz.TopdataRepo {
	return &topdataRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/order")),
	}
}

func (r *topdataRepo) List(ctx context.Context, pageNum, pageSize int64) ([]*biz.Topdata, error) {
	var td []*model.Topdatum
	res := r.data.db.WithContext(ctx).
		Limit(int(pageSize)).
		Offset(int(pagination.GetPageOffset(pageNum, pageSize))).
		Find(&td)
	if res.Error != nil {
		return nil, res.Error
	}

	rv := make([]*biz.Topdata, 0)
	for _, v := range td {
		rv = append(rv, &biz.Topdata{
			ID:         v.ID,
			Tab:        v.Tab,
			Position:   v.Position,
			Title:      v.Title,
			URL:        v.URL,
			Extra:      v.Extra,
			SpiderTime: v.SpiderTime,
		})
	}
	return rv, nil
}

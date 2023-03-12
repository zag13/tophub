package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"tophub/app/topdata/service/internal/biz"

	v1 "tophub/api/topdata/service/v1"
)

type TopdataService struct {
	v1.UnimplementedTopdataServer

	tduc *biz.TopdataUseCase
	log  *log.Helper
}

func NewTopDataService(tduc *biz.TopdataUseCase, logger log.Logger) *TopdataService {
	return &TopdataService{
		tduc: tduc,
		log:  log.NewHelper(log.With(logger, "module", "topdata-service/service")),
	}
}

func (s *TopdataService) List(ctx context.Context, req *v1.ListRequest) (*v1.ListReply, error) {
	rv, err := s.tduc.List(ctx, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	rs := make([]*v1.ListReply_List, 0)
	for _, v := range rv {
		rs = append(rs, &v1.ListReply_List{
			Position: 0,
			Title:    v.Title,
			Url:      "v.Url",
			Extra:    nil,
		})
	}

	return &v1.ListReply{
		List: rs,
	}, nil
}

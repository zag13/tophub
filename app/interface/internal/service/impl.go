package service

import (
	"context"
	"time"

	v1 "tophub/api/interface/v1"
	"tophub/app/interface/internal/data"
)

// Data implements data.GreeterServer.
func (s *InterfaceService) Data(ctx context.Context, req *v1.DataRequest) (*v1.DataResponse, error) {
	var dataInDB []data.Data

	if err := s.db.Where("tab = ? AND spider_time >= ? AND spider_time <= ?", req.Tab, time.Now().Add(-10*time.Minute), time.Now().Unix()).Find(&dataInDB).Error; err != nil {
		return nil, err
	}

	var list []*v1.DataResponse_Data

	for _, datum := range dataInDB {
		list = append(list, &v1.DataResponse_Data{
			Position:    uint32(datum.Position),
			Title:       datum.Title,
			Url:         datum.Url,
			Description: datum.Description,
			Image:       datum.Image,
			Extra:       datum.Extra,
		})
	}

	return &v1.DataResponse{List: list}, nil
}

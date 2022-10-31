package service

import (
	"context"
	"encoding/json"

	v1 "tophub/api/interface/v1"
	taskv1 "tophub/api/task/v1"
)

// Data implements data.GreeterServer.
func (s *InterfaceService) Data(ctx context.Context, req *v1.DataRequest) (*v1.DataResponse, error) {
	resp, err := s.tc.Data(ctx, &taskv1.DataRequest{Tab: req.Tab})
	if err != nil {
		return nil, err
	}

	var list []*v1.DataResponse_Data
	for _, data := range resp.List {
		var extra map[string]string
		if err := json.Unmarshal([]byte(data.Extra), &extra); err != nil {
			return nil, err
		}
		list = append(list, &v1.DataResponse_Data{
			Position: data.Position,
			Title:    data.Title,
			Url:      data.Url,
			Extra:    extra,
		})
	}

	return &v1.DataResponse{List: list}, nil
}

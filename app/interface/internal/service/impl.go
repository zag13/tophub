package service

import (
	"context"

	v1 "tophub/api/interface/v1"
)

// Data implements data.GreeterServer.
func (s *InterfaceService) Data(ctx context.Context, req *v1.DataRequest) (*v1.DataResponse, error) {
	return &v1.DataResponse{}, nil
}

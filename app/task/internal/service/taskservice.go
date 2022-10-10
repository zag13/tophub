package service

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "tophub/api/task/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTaskService)

type TaskService struct {
	pb.UnimplementedTaskServer

	log *klog.Helper
}

func NewTaskService(logger klog.Logger) (*TaskService, func(), error) {
	cleanup := func() {
		klog.NewHelper(logger).Info("closing the data resources")
	}
	return &TaskService{
		log: klog.NewHelper(klog.With(logger, "module", "task/service")),
	}, cleanup, nil
}

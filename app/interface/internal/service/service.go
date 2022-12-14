package service

import (
	"context"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	v1 "tophub/api/interface/v1"
	taskv1 "tophub/api/task/v1"
	"tophub/app/interface/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewInterfaceService, NewDiscovery, NewRegistrar, NewTaskClient)

// InterfaceService is an Interface service.
type InterfaceService struct {
	v1.UnimplementedInterfaceServer

	tc  taskv1.TaskClient
	log *klog.Helper
}

// NewInterfaceService new an interface service.
func NewInterfaceService(logger klog.Logger, tc taskv1.TaskClient) (*InterfaceService, func(), error) {
	cleanup := func() {
		klog.NewHelper(logger).Info("closing the data resources")
	}
	return &InterfaceService{
		tc:  tc,
		log: klog.NewHelper(klog.With(logger, "module", "interface/service")),
	}, cleanup, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := clientv3.Config{}
	c.Endpoints = conf.Etcd.Endpoints
	cli, err := clientv3.New(c)
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := clientv3.Config{}
	c.Endpoints = conf.Etcd.Endpoints
	cli, err := clientv3.New(c)
	if err != nil {
		panic(err)
	}
	r := etcd.New(cli)
	return r
}

func NewTaskClient(r registry.Discovery, tp *tracesdk.TracerProvider) taskv1.TaskClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///tophub.task"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return taskv1.NewTaskClient(conn)
}

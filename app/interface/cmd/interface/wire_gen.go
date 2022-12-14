// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"tophub/app/interface/internal/conf"
	"tophub/app/interface/internal/server"
	"tophub/app/interface/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, registry *conf.Registry, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	discovery := service.NewDiscovery(registry)
	taskClient := service.NewTaskClient(discovery, tracerProvider)
	interfaceService, cleanup, err := service.NewInterfaceService(logger, taskClient)
	if err != nil {
		return nil, nil, err
	}
	grpcServer := server.NewGRPCServer(confServer, interfaceService, logger)
	httpServer := server.NewHTTPServer(confServer, interfaceService, logger)
	registrar := service.NewRegistrar(registry)
	app := newApp(logger, grpcServer, httpServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}

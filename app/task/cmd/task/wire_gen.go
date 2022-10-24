// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
	"tophub/app/task/internal/conf"
	"tophub/app/task/internal/server"
	"tophub/app/task/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, registry *conf.Registry, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	taskService, cleanup, err := service.NewTaskService(logger)
	if err != nil {
		return nil, nil, err
	}
	grpcServer := server.NewGRPCServer(confServer, taskService, tracerProvider, logger)
	registrar := service.NewRegistrar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}

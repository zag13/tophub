//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"tophub/app/topdata/cronjob/internal/conf"
	"tophub/app/topdata/cronjob/internal/data"
	"tophub/app/topdata/cronjob/internal/server"
	"tophub/app/topdata/cronjob/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}

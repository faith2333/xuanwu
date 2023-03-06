//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github/faith2333/xuanwu/internal/biz"
	"github/faith2333/xuanwu/internal/conf"
	"github/faith2333/xuanwu/internal/data/base"
	"github/faith2333/xuanwu/internal/server"
	"github/faith2333/xuanwu/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, base.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	application2 "github.com/faith2333/xuanwu/internal/biz/application"
	user2 "github.com/faith2333/xuanwu/internal/biz/user"
	"github.com/faith2333/xuanwu/internal/conf"
	"github.com/faith2333/xuanwu/internal/data/application"
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/faith2333/xuanwu/internal/data/user"
	"github.com/faith2333/xuanwu/internal/server"
	application3 "github.com/faith2333/xuanwu/internal/service/application"
	user3 "github.com/faith2333/xuanwu/internal/service/user"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	baseData, err := base.NewData(data, logger)
	if err != nil {
		return nil, nil, err
	}
	iRepoUser := user.NewRepoUser(baseData)
	userInterface, err := user2.NewUserFactory(iRepoUser, confServer)
	if err != nil {
		return nil, nil, err
	}
	serviceUser := user3.NewServiceUser(userInterface)
	iAppRepo := application.NewAppRepo(baseData)
	biz := application2.NewBiz(iAppRepo)
	appSvc := application3.NewAppSvc(biz)
	grpcServer := server.NewGRPCServer(confServer, serviceUser, appSvc, logger)
	httpServer := server.NewHTTPServer(confServer, serviceUser, appSvc, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
	}, nil
}

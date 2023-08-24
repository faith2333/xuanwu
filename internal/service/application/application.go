package application

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/application/v1"
	bizApp "github.com/faith2333/xuanwu/internal/biz/application"
	"github.com/faith2333/xuanwu/internal/service/base"
)

type AppSvc struct {
	pb.UnimplementedApplicationSvcServer
	base.Base
	biz *bizApp.Biz
}

func NewAppSvc(biz *bizApp.Biz) *AppSvc {
	return &AppSvc{
		biz: biz,
	}
}

func (s *AppSvc) CreateApplication(ctx context.Context, req *pb.CreateAppRequest) (*pb.Application, error) {

}

func (s *AppSvc) ListApplications(ctx context.Context, req *pb.ListAppRequest) (*pb.ListAppResponse, error) {

}

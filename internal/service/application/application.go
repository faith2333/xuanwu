package application

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/application/v1"
	bizApp "github.com/faith2333/xuanwu/internal/biz/application"

	basePB "github.com/faith2333/xuanwu/api/base/v1"
	"github.com/faith2333/xuanwu/internal/biz/application/types"
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
	bizCreateReq := &bizApp.CreateAppReq{
		Name:     req.Name,
		AppType:  types.AppType(req.AppType),
		Category: req.Category,
		Labels:   req.Labels,
	}

	err := s.PBStructUnmarshal(req.DevelopmentInfo, &bizCreateReq.DevelopmentInfo)
	if err != nil {
		return nil, err
	}

	err = s.PBStructUnmarshal(req.TestInfo, &bizCreateReq.TestInfo)
	if err != nil {
		return nil, err
	}

	err = s.PBStructUnmarshal(req.NotificationInfos, &bizCreateReq.NotificationInfos)
	if err != nil {
		return nil, err
	}

	bizResp, err := s.biz.CreateApp(ctx, bizCreateReq)
	if err != nil {
		return nil, err
	}

	pbApp, err := s.bizAppToPBApp(bizResp)
	if err != nil {
		return nil, err
	}

	return pbApp, nil
}

func (s *AppSvc) ListApplications(ctx context.Context, req *pb.ListAppRequest) (*pb.ListAppResponse, error) {
	bizListReq := &bizApp.ListAppReq{}
	err := s.TransformJSON(req, &bizListReq)
	if err != nil {
		return nil, err
	}

	bizResp, err := s.biz.ListApps(ctx, bizListReq)
	if err != nil {
		return nil, err
	}

	pbResp := &pb.ListAppResponse{
		PageInfo: &basePB.PageInfo{
			Total:     bizResp.PageInfo.Total,
			PageIndex: bizResp.PageInfo.PageIndex,
			PageSize:  bizResp.PageInfo.PageSize,
		},
	}

	for _, bApp := range bizResp.Data {
		pbApp, err := s.bizAppToPBApp(bApp)
		if err != nil {
			return nil, err
		}
		pbResp.Data = append(pbResp.Data, pbApp)
	}

	return pbResp, nil
}

func (s *AppSvc) bizAppToPBApp(bApp *bizApp.Application) (*pb.Application, error) {
	pbApp := &pb.Application{
		Id:       bApp.ID,
		Name:     bApp.Name,
		Code:     bApp.Code,
		AppType:  bApp.AppType.String(),
		Category: bApp.Category,
		Labels:   bApp.Labels,
	}

	var err error
	pbApp.DevelopmentInfo, err = s.NewPBStruct(bApp.DevelopmentInfo)
	if err != nil {
		return nil, err
	}

	pbApp.TestInfo, err = s.NewPBStruct(bApp.TestInfo)
	if err != nil {
		return nil, err
	}

	pbApp.NotificationInfos, err = s.NewPBStruct(bApp.NotificationInfos)
	if err != nil {
		return nil, err
	}

	return pbApp, nil
}

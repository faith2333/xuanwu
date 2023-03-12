package service

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/cicd/v1"
	"github.com/faith2333/xuanwu/internal/biz"
	bizCICD "github.com/faith2333/xuanwu/internal/biz/cicd"
)

type CICDService struct {
	pb.UnimplementedCICDServer

	biz *biz.Biz
}

func NewCICDService(biz *biz.Biz) *CICDService {
	return &CICDService{
		biz: biz,
	}
}

func (s *CICDService) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.Template, error) {
	resp, err := s.biz.CICD.TemplateRepo.CreateTemplate(ctx, &bizCICD.CreateTemplate{CreateTemplateRequest: req})
	if err != nil {
		return nil, err
	}
	return resp.Template, nil
}
func (s *CICDService) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateReply, error) {
	return &pb.UpdateTemplateReply{}, nil
}
func (s *CICDService) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateReply, error) {
	return &pb.DeleteTemplateReply{}, nil
}
func (s *CICDService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.Template, error) {
	return &pb.Template{}, nil
}
func (s *CICDService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	return &pb.ListTemplateReply{}, nil
}

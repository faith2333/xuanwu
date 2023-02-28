package service

import (
	"context"

	pb "github/faith2333/xuanwu/api/cicd/v1"
)

type CICDService struct {
	pb.UnimplementedCICDServer
}

func NewCICDService() *CICDService {
	return &CICDService{}
}

func (s *CICDService) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	return &pb.CreateTemplateReply{}, nil
}
func (s *CICDService) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateReply, error) {
	return &pb.UpdateTemplateReply{}, nil
}
func (s *CICDService) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateReply, error) {
	return &pb.DeleteTemplateReply{}, nil
}
func (s *CICDService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateReply, error) {
	return &pb.GetTemplateReply{}, nil
}
func (s *CICDService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	return &pb.ListTemplateReply{}, nil
}

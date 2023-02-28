package service

import (
	"context"

	pb "xuanwu/api/cicd/v1"
)

type TemplateService struct {
	pb.UnimplementedTemplateServer
}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (s *TemplateService) CreateTemplate(ctx context.Context, req *pb.CreateTemplateRequest) (*pb.CreateTemplateReply, error) {
	return &pb.CreateTemplateReply{}, nil
}
func (s *TemplateService) UpdateTemplate(ctx context.Context, req *pb.UpdateTemplateRequest) (*pb.UpdateTemplateReply, error) {
	return &pb.UpdateTemplateReply{}, nil
}
func (s *TemplateService) DeleteTemplate(ctx context.Context, req *pb.DeleteTemplateRequest) (*pb.DeleteTemplateReply, error) {
	return &pb.DeleteTemplateReply{}, nil
}
func (s *TemplateService) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateReply, error) {
	return &pb.GetTemplateReply{}, nil
}
func (s *TemplateService) ListTemplate(ctx context.Context, req *pb.ListTemplateRequest) (*pb.ListTemplateReply, error) {
	return &pb.ListTemplateReply{}, nil
}

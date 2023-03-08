package cicd

import (
	"context"
	pb "github/faith2333/xuanwu/api/cicd/v1"
)

type TemplateRepo interface {
	CreateTemplate(ctx context.Context, request *CreateTemplate) (*Template, error)
}

type CreateTemplate struct {
	*pb.CreateTemplateRequest
}

type Template struct {
	*pb.Template
}

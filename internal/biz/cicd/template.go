package cicd

import pb "github/faith2333/xuanwu/api/cicd/v1"

type TemplateRepo interface {
	CreateTemplate(request *CreateTemplate) (Template, error)
}

type CreateTemplate struct {
	pb.CreateTemplateRequest
}

type Template struct {
	pb.Template
}

package organization

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/organization/v1"
)

type IRepoOrganization interface {
	Create(ctx context.Context, req *CreateOrgReq) (*Organization, error)
	List(ctx context.Context, req *ListOrgReq) (*ListOrgReply, error)
}

type Organization struct {
	pb.Organization
}

type CreateOrgReq struct {
	pb.CreateOrgRequest
}

func (biz *Biz) CreateOrg(ctx context.Context, req *CreateOrgReq) (*Organization, error) {
	return biz.repo.Create(ctx, req)
}

type ListOrgReq struct {
	pb.ListOrgsRequest
}

type ListOrgReply struct {
	pb.ListOrgsResponse
}

func (biz *Biz) ListOrgs(ctx context.Context, req *ListOrgReq) (*ListOrgReply, error) {
	if req.PageIndex <= 0 {
		req.PageIndex = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	return biz.repo.List(ctx, req)
}

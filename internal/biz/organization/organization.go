package organization

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/organization/v1"
)

type IRepoOrganization interface {
	Create(ctx context.Context, req *CreateOrgReq) (*Organization, error)
	List(ctx context.Context, req *ListOrgReq) (*ListOrgReply, error)
	GetByCode(ctx context.Context, code string) (*Organization, error)
	Update(ctx context.Context, org *Organization) error
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

func (biz *Biz) UpdateOrg(ctx context.Context, req *CreateOrgReq) (*Organization, error) {
	org, err := biz.repo.GetByCode(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	if req.Name != org.Name {
		org.Name = req.Name
	}

	if req.Enabled != org.Enabled {
		org.Enabled = req.Enabled
	}

	if req.Desc != org.Desc {
		org.Desc = req.Desc
	}

	return org, biz.repo.Update(ctx, org)
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

type ChangeOrgStatusReq struct {
	Code   string `json:"code"`
	Enable bool   `json:"enable"`
}

func (biz *Biz) ChangeOrgStatus(ctx context.Context, req *ChangeOrgStatusReq) error {
	org, err := biz.repo.GetByCode(ctx, req.Code)
	if err != nil {
		return err
	}

	if org.Enabled == req.Enable {
		return nil
	}

	org.Enabled = req.Enable

	return biz.repo.Update(ctx, org)
}

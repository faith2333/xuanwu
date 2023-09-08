package organization

import (
	"context"
	pb "github.com/faith2333/xuanwu/api/organization/v1"
	bizORG "github.com/faith2333/xuanwu/internal/biz/organization"
	"github.com/faith2333/xuanwu/internal/service/base"
)

type OrgSvc struct {
	base.Base
	pb.UnimplementedOrganizationServiceServer
	biz *bizORG.Biz
}

func NewOrgSvc(biz *bizORG.Biz) *OrgSvc {
	return &OrgSvc{
		biz: biz,
	}
}

func (s *OrgSvc) CreateOrganization(ctx context.Context, req *pb.CreateOrgRequest) (*pb.Organization, error) {
	bizReq := &bizORG.CreateOrgReq{}
	err := s.TransformJSON(req, &bizReq)
	if err != nil {
		return nil, err
	}

	bizResp, err := s.biz.CreateOrg(ctx, bizReq)
	if err != nil {
		return nil, err
	}

	pbResp := &pb.Organization{}
	err = s.TransformJSON(bizResp, &pbResp)
	if err != nil {
		return nil, err
	}

	return pbResp, nil
}

func (s *OrgSvc) ListOrganizations(ctx context.Context, req *pb.ListOrgsRequest) (*pb.ListOrgsResponse, error) {
	bizReq := &bizORG.ListOrgReq{}
	err := s.TransformJSON(req, &bizReq)
	if err != nil {
		return nil, err
	}

	bizResp, err := s.biz.ListOrgs(ctx, bizReq)
	if err != nil {
		return nil, err
	}

	pbResp := &pb.ListOrgsResponse{}
	err = s.TransformJSON(bizResp, &pbResp)
	if err != nil {
		return nil, err
	}

	return pbResp, nil
}

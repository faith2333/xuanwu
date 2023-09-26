package organization

import (
	"context"
	basePB "github.com/faith2333/xuanwu/api/base/v1"
	pbORG "github.com/faith2333/xuanwu/api/organization/v1"
	bizORG "github.com/faith2333/xuanwu/internal/biz/organization"
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/faith2333/xuanwu/pkg/xerrors"
	"github.com/pkg/errors"
	"sync"
)

type Organization struct {
	ID      int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"type:varchar(64);not null"`
	Code    string `json:"code" gorm:"type:varchar(64);not null;uniqueIndex:CODE_DELETED"`
	Desc    string `json:"desc" gorm:"type:varchar(400)"`
	Enabled bool   `json:"enabled"`
	base.Model
}

const (
	TablePrefix = "org_"
)

func (Organization) TableName() string {
	return TablePrefix + "orgs"
}

var orgRunOnce = &sync.Once{}

type OrgRepo struct {
	base.RepoBase
	data *base.Data
}

func NewOrgRepo(data *base.Data) bizORG.IRepoOrganization {
	orgRunOnce.Do(func() {
		err := data.DB(context.Background()).AutoMigrate(&Organization{})
		if err != nil {
			panic(err)
		}
	})

	return &OrgRepo{
		data: data,
	}
}

func (repo *OrgRepo) Create(ctx context.Context, req *bizORG.CreateOrgReq) (*bizORG.Organization, error) {
	dbOrg := &Organization{}
	err := repo.Transform(req, &dbOrg)
	if err != nil {
		return nil, err
	}

	err = repo.data.DB(ctx).Create(dbOrg).Error
	if err != nil {
		return nil, errors.Errorf("create org failed: %v", err)
	}

	bizResp := &bizORG.Organization{}
	err = repo.Transform(dbOrg, &bizResp)
	if err != nil {
		return nil, err
	}

	return bizResp, nil
}

func (repo *OrgRepo) GetByCode(ctx context.Context, code string) (*bizORG.Organization, error) {
	org, err := repo.getByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	bizOrg := &bizORG.Organization{}
	err = repo.Transform(org, &bizOrg)
	if err != nil {
		return nil, err
	}

	return bizOrg, nil
}

func (repo *OrgRepo) Update(ctx context.Context, org *bizORG.Organization) error {
	dbOrg := &Organization{}
	err := repo.Transform(org, &dbOrg)
	if err != nil {
		return err
	}

	err = repo.data.DB(ctx).Model(&Organization{}).Save(&dbOrg).Error
	if err != nil {
		return errors.Errorf("update failed: %v", err)
	}

	return nil
}

func (repo *OrgRepo) getByCode(ctx context.Context, code string) (*Organization, error) {
	orgs := make([]*Organization, 0)
	err := repo.data.DB(ctx).Model(&Organization{}).Where("code = ? and deleted = 0", code).Find(&orgs).Error
	if err != nil {
		return nil, errors.Errorf("query failed: %v", err)
	}

	if len(orgs) == 0 {
		return nil, xerrors.ErrNotFound
	}

	if len(orgs) > 1 {
		return nil, xerrors.ErrMultipleValues
	}

	return orgs[0], nil
}

func (repo *OrgRepo) List(ctx context.Context, req *bizORG.ListOrgReq) (*bizORG.ListOrgReply, error) {
	offset := (req.PageIndex - 1) * req.PageSize
	dbOrgs := make([]*Organization, 0)

	query := repo.data.DB(ctx).Model(&Organization{}).Where("deleted = 0")
	if req.Id != 0 {
		query = query.Where("id = ?", req.Id)
	}

	if req.Code != "" {
		query = query.Where("code = ?", req.Code)
	}

	if req.Name != "" {
		query = query.Where("name like ?", "%"+req.Name+"%")
	}

	if req.Enabled {
		query = query.Where("enabled = 1")
	}

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, errors.Errorf("count failed: %v", err)
	}

	err = query.Offset(int(offset)).Limit(int(req.PageSize)).Find(&dbOrgs).Error
	if err != nil {
		return nil, errors.Errorf("list failed: %v", err)
	}

	pbORGs := make([]*pbORG.Organization, 0)
	err = repo.Transform(dbOrgs, &pbORGs)
	if err != nil {
		return nil, err
	}

	resp := &bizORG.ListOrgReply{
		ListOrgsResponse: pbORG.ListOrgsResponse{
			Data: pbORGs,
			PageInfo: &basePB.PageInfo{
				Total:     total,
				PageSize:  req.PageSize,
				PageIndex: req.PageIndex,
			},
		},
	}

	return resp, nil
}

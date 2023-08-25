package application

import (
	"context"
	"fmt"
	bizApp "github.com/faith2333/xuanwu/internal/biz/application"
	"github.com/faith2333/xuanwu/internal/biz/application/types"
	bizBase "github.com/faith2333/xuanwu/internal/biz/base"
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/faith2333/xuanwu/pkg/xerrors"
	"github.com/pkg/errors"
	"sync"
)

const (
	TablePrefix = "app_"
)

type Application struct {
	ID                 int64                     `json:"id" gorm:"primaryKey;autoIncrement"`
	Code               string                    `json:"code" gorm:"type:varchar(64);uniqueIndex:CODE_DELETED"`
	Name               string                    `json:"name" gorm:"type:varchar(64)"`
	AppType            types.AppType             `json:"appType" gorm:"type:varchar(64)"`
	Category           string                    `json:"category" gorm:"type:varchar(64)"`
	Labels             base.TypeSlice            `json:"labels" gorm:"type:json"`
	Language           types.DevelopmentLanguage `json:"language" gorm:"type:varchar(64)"`
	GitAddress         string                    `json:"gitAddress" gorm:"type:varchar(64)"`
	DevelopmentManager string                    `json:"developmentManager" gorm:"type:varchar(64)"`
	Developers         base.TypeSlice            `json:"developers" gorm:"type:json"`
	TestManager        string                    `json:"testManager" gorm:"type:varchar(64)"`
	NotificationInfos  types.NotificationInfos   `json:"notificationInfos" gorm:"type:json"`
	base.Model
}

var appRunOnce = &sync.Once{}

func (Application) TableName() string {
	return TablePrefix + "apps"
}

type AppRepo struct {
	base.RepoBase
	data *base.Data
}

func NewAppRepo(data *base.Data) bizApp.IAppRepo {
	appRunOnce.Do(func() {
		err := data.DB(context.Background()).AutoMigrate(&Application{})
		if err != nil {
			panic(err)
		}
	})
	return &AppRepo{
		data: data,
	}
}

func (repo *AppRepo) Create(ctx context.Context, app *bizApp.Application) (*bizApp.Application, error) {
	dbApp := repo.bizToDbApp(app)
	err := repo.data.DB(ctx).Create(dbApp).Error
	if err != nil {
		return nil, errors.Wrap(err, "create app failed")
	}

	return repo.dbToBizApp(dbApp), nil
}

func (repo *AppRepo) List(ctx context.Context, req *bizApp.ListAppReq) (*bizApp.ListAppReply, error) {
	offset := (req.PageIndex - 1) * req.PageSize
	dbApps := make([]*Application, 0)

	query := repo.data.DB(ctx).Model(&Application{}).Where("deleted = 0")
	if req.ID != 0 {
		query = query.Where("id = ?", req.ID)
	}
	if req.Code != "" {
		query = query.Where("code like ", fmt.Sprintf("%%%s%%", req.Code))
	}
	if req.Name != "" {
		query = query.Where("name like ?", fmt.Sprintf("%%%s%%", req.Name))
	}
	if req.AppType != "" {
		query = query.Where("app_type = ?", req.AppType)
	}
	if len(req.Labels) != 0 {
		query = query.Where("labels @> ?", req.Labels)
	}
	if req.DevelopmentLanguage != "" {
		query = query.Where("language = ?", req.DevelopmentLanguage)
	}
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, errors.Wrap(err, "count failed")
	}

	err = query.Offset(int(offset)).Limit(int(req.PageSize)).Find(&dbApps).Error
	if err != nil {
		return nil, errors.Wrap(err, "list failed")
	}

	bizApps := make([]*bizApp.Application, 0)
	for _, dbApp := range dbApps {
		bizApps = append(bizApps, repo.dbToBizApp(dbApp))
	}

	return &bizApp.ListAppReply{
		Data: bizApps,
		PageInfo: bizBase.PageInfo{
			Total:     total,
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
		},
	}, nil
}

func (repo *AppRepo) GetByCode(ctx context.Context, code string) (*bizApp.Application, error) {
	dbApp, err := repo.getByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return repo.dbToBizApp(dbApp), nil
}

func (repo *AppRepo) getByCode(ctx context.Context, code string) (*Application, error) {
	dbApps := make([]*Application, 0)
	err := repo.data.DB(ctx).Where("code = ? and deleted = 0", code).Find(&dbApps).Error
	if err != nil {
		return nil, errors.Errorf("get app by code failed, code: %s", code)
	}

	if len(dbApps) == 0 {
		return nil, xerrors.ErrNotFound
	}

	if len(dbApps) > 1 {
		return nil, xerrors.ErrMultipleValues
	}

	return dbApps[0], nil
}

func (repo *AppRepo) dbToBizApp(dbApp *Application) *bizApp.Application {
	return &bizApp.Application{
		ID:       dbApp.ID,
		Code:     dbApp.Code,
		Name:     dbApp.Name,
		AppType:  dbApp.AppType,
		Category: dbApp.Category,
		Labels:   dbApp.Labels.ToSlice(),
		DevelopmentInfo: bizApp.DevelopmentInfo{
			Language:           dbApp.Language,
			GitAddress:         dbApp.GitAddress,
			DevelopmentManager: dbApp.DevelopmentManager,
			Developers:         dbApp.Developers.ToSlice(),
		},
		TestInfo: bizApp.TestInfo{
			TestManager: dbApp.TestManager,
		},
		NotificationInfos: dbApp.NotificationInfos.ToSlice(),
	}
}

func (repo *AppRepo) bizToDbApp(bizApp *bizApp.Application) *Application {
	return &Application{
		ID:                 bizApp.ID,
		Code:               bizApp.Code,
		Name:               bizApp.Name,
		AppType:            bizApp.AppType,
		Category:           bizApp.Category,
		Labels:             base.NewTypeSlice(bizApp.Labels),
		Language:           bizApp.DevelopmentInfo.Language,
		GitAddress:         bizApp.DevelopmentInfo.GitAddress,
		DevelopmentManager: bizApp.DevelopmentInfo.DevelopmentManager,
		Developers:         base.NewTypeSlice(bizApp.DevelopmentInfo.Developers),
		TestManager:        bizApp.TestInfo.TestManager,
		NotificationInfos:  bizApp.NotificationInfos,
	}
}

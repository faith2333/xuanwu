package application

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/application/types"
	"github.com/faith2333/xuanwu/internal/biz/base"
)

type IAppRepo interface {
	Create(ctx context.Context, app *Application) (*Application, error)
	List(ctx context.Context, req *ListAppReq) (*ListAppReply, error)
	GetByCode(ctx context.Context, code string) (*Application, error)
	DeleteByCode(ctx context.Context, code string) error
}

type Application struct {
	Code    string        `json:"code"`
	Name    string        `json:"name"`
	AppType types.AppType `json:"appType"`
	// the organization of the application, it is used for the organization.proto of the application.
	Organization      string                    `json:"organization"`
	Labels            []string                  `json:"labels"`
	Desc              string                    `json:"desc"`
	DevelopmentInfo   DevelopmentInfo           `json:"developmentInfo"`
	TestInfo          TestInfo                  `json:"testInfo"`
	NotificationInfos []*types.NotificationInfo `json:"notificationInfos"`
	base.Model
}

type DevelopmentInfo struct {
	Language   types.DevelopmentLanguage `json:"language"`
	GitAddress string                    `json:"gitAddress"`
	// it will use for the permissions lately, include deploy, alert notify and so on.
	DevelopmentManager string   `json:"developmentManager"`
	Developers         []string `json:"developers"`
}

type TestInfo struct {
	TestManager string `json:"testManager"`
}

type CreateAppReq struct {
	Name              string                    `json:"name"`
	Code              string                    `json:"code"`
	AppType           types.AppType             `json:"appType"`
	Labels            []string                  `json:"labels"`
	Organization      string                    `json:"Organization"`
	DevelopmentInfo   DevelopmentInfo           `json:"developmentInfo"`
	TestInfo          TestInfo                  `json:"testInfo"`
	NotificationInfos []*types.NotificationInfo `json:"notificationInfos"`
	Desc              string                    `json:"desc"`
}

func (biz *Biz) CreateApp(ctx context.Context, req *CreateAppReq) (*Application, error) {
	if !req.AppType.IsSupported() {
		return nil, types.ErrAppTypeNotSupported
	}

	if !req.DevelopmentInfo.Language.IsSupported() {
		return nil, types.ErrDevelopmentLanguageNotSupported
	}

	app := &Application{
		Name:              req.Name,
		Code:              req.Code,
		AppType:           req.AppType,
		Labels:            req.Labels,
		Organization:      req.Organization,
		DevelopmentInfo:   req.DevelopmentInfo,
		TestInfo:          req.TestInfo,
		NotificationInfos: req.NotificationInfos,
		Desc:              req.Desc,
	}

	return biz.appRepo.Create(ctx, app)
}

type ListAppReq struct {
	ID                  int64    `json:"id"`
	Code                string   `json:"code"`
	Name                string   `json:"name"`
	AppType             string   `json:"appType"`
	Labels              []string `json:"labels"`
	DevelopmentLanguage string   `json:"developmentLanguage"`
	Category            string   `json:"category"`
	PageIndex           int64    `json:"pageIndex"`
	PageSize            int64    `json:"pageSize"`
}

type ListAppReply struct {
	Data     []*Application `json:"data"`
	PageInfo base.PageInfo  `json:"pageInfo"`
}

func (biz *Biz) ListApps(ctx context.Context, req *ListAppReq) (*ListAppReply, error) {
	if req.PageIndex <= 0 {
		req.PageIndex = 1
	}

	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	return biz.appRepo.List(ctx, req)
}

func (biz *Biz) DeleteApp(ctx context.Context, code string) error {
	return biz.appRepo.DeleteByCode(ctx, code)
}

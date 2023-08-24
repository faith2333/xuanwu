package application

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/application/types"
)

type IAppRepo interface {
	Create(ctx context.Context, app *Application) (*Application, error)
	List(ctx context.Context, req *ListAppReq) (*ListAppReply, error)
	GetByCode(ctx context.Context, code string) (*Application, error)
}

type Application struct {
	ID               int64            `json:"id"`
	Code             string           `json:"code"`
	Name             string           `json:"name"`
	AppType          types.AppType    `json:"appType"`
	DevelopmentInfo  DevelopmentInfo  `json:"developmentInfo"`
	TestInfo         TestInfo         `json:"testInfo"`
	NotificationInfo NotificationInfo `json:"NotificationInfo"`
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

type NotificationInfo struct {
	AlertNotificationInfo innerNotificationInfo `json:"alertNotificationInfo"`
	BuildNotificationInfo innerNotificationInfo `json:"buildNotificationInfo"`
}

type innerNotificationInfo struct {
	NotifyUsers []string `json:"notifyUsers"`
	// the id of the notify config in notification center.
	NotifyID string `json:"notifyId"`
}

type ListAppReq struct {
}

type ListAppReply struct {
}

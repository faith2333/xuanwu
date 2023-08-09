package application

import "github.com/faith2333/xuanwu/internal/biz/application/types"

type IAppRepo interface {
}

type Application struct {
	ID              int64           `json:"id"`
	Code            string          `json:"code"`
	Name            string          `json:"name"`
	AppType         types.AppType   `json:"appType"`
	DevelopmentInfo DevelopmentInfo `json:"developmentInfo"`
	TestInfo        TestInfo        `json:"testInfo"`
	NotifyInfo      NotifyInfo      `json:"notifyInfo"`
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

type NotifyInfo struct {
}

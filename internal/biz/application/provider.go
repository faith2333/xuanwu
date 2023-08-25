package application

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewBiz)

type Biz struct {
	appRepo IAppRepo
}

func NewBiz(appRepo IAppRepo) *Biz {
	return &Biz{
		appRepo: appRepo,
	}
}

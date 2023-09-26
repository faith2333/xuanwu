package organization

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewBiz)

type Biz struct {
	repo IRepoOrganization
}

func NewBiz(repo IRepoOrganization) *Biz {
	return &Biz{repo: repo}
}

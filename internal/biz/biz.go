package biz

import (
	"github.com/google/wire"
	cicdRepo "github/faith2333/xuanwu/internal/biz/cicd"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz)

type Biz struct {
	CICD *cicdRepo.Repo
}

func NewBiz(cicd *cicdRepo.Repo) *Biz {
	return &Biz{
		CICD: cicd,
	}
}

package biz

import (
	cicdRepo "github.com/faith2333/xuanwu/internal/biz/cicd"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewBiz, cicdRepo.NewRepo)

type Biz struct {
	CICD *cicdRepo.Repo
}

func NewBiz(cicd *cicdRepo.Repo) *Biz {
	return &Biz{
		CICD: cicd,
	}
}

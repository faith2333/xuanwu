package data

import (
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/faith2333/xuanwu/internal/data/cicd"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(base.ProviderSet, cicd.ProviderSet)

package biz

import (
	"github.com/faith2333/xuanwu/internal/biz/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.ProviderSet,
)

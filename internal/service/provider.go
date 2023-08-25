package service

import (
	"github.com/faith2333/xuanwu/internal/service/application"
	"github.com/faith2333/xuanwu/internal/service/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.ProviderSet,
	application.ProviderSet,
)

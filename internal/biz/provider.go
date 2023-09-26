package biz

import (
	"github.com/faith2333/xuanwu/internal/biz/application"
	"github.com/faith2333/xuanwu/internal/biz/organization"
	"github.com/faith2333/xuanwu/internal/biz/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	user.ProviderSet,
	application.ProviderSet,
	organization.ProviderSet,
)

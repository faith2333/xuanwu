package data

import (
	"github.com/faith2333/xuanwu/internal/data/application"
	"github.com/faith2333/xuanwu/internal/data/base"
	"github.com/faith2333/xuanwu/internal/data/organization"
	"github.com/faith2333/xuanwu/internal/data/user"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	base.ProviderSet,
	user.ProviderSet,
	application.ProviderSet,
	organization.ProviderSet,
)

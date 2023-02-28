package service

import (
	"github.com/google/wire"
	cicd "github/faith2333/xuanwu/internal/service/cicd"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(cicd.NewCICDService)

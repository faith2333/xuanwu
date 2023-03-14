package service

import (
	cicd "github.com/faith2333/xuanwu/internal/service/cicd"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(cicd.NewCICDService)

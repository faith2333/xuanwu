package organization

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewOrgRepo,
)

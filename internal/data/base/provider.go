package base

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewData,
	NewTransaction,
)

package user

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserFactory)

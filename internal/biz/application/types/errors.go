package types

import "github.com/pkg/errors"

var (
	ErrAppTypeNotSupported             = errors.New("app type not support")
	ErrDevelopmentLanguageNotSupported = errors.New("development language not support")
)

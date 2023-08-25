package xerrors

import "github.com/pkg/errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrMultipleValues = errors.New("multiple values")
)

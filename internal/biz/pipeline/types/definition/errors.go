package definition

import "github.com/pkg/errors"

var (
	ErrNameIsEmpty  = errors.New("name is empty")
	ErrMultipleName = errors.New("multiple name in stage")

	ErrRepeatFromVariableTypeNotSupported = errors.New("the repeat is true but repeatFromVariable in global variables was not SLICE or MAP in stage")
	ErrRepeatFromVariableNotSpecified     = errors.New("the repeat is true but repeatFromVariable was not specified in stage")
	ErrRepeatFromVariableNotFound         = errors.New("the repeat is true but repeatFromVariable was not found in stage")
)

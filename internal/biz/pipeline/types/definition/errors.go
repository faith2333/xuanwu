package definition

import "github.com/pkg/errors"

var (
	ErrNameIsEmpty      = errors.New("name is empty")
	ErrTypeIsEmpty      = errors.New("type is empty")
	ErrTypeIsNotSupport = errors.New("type is not supported")
	ErrMultipleName     = errors.New("multiple name in stage")

	ErrRepeatFromVariableTypeNotSupported = errors.New("the repeat is true but repeatFromVariable in global variables was not SLICE or MAP in stage")
	ErrRepeatFromVariableNotSpecified     = errors.New("the repeat is true but repeatFromVariable was not specified in stage")
	ErrRepeatFromVariableNotFound         = errors.New("the repeat is true but repeatFromVariable was not found in stage")

	ErrVariableNotSupported = errors.New("the variable type has not been supported")
)

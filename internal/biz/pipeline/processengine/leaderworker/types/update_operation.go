package types

import "strings"

type UpdateOperation struct {
	OperationType OperationType `json:"operationType"`
	Data          interface{}   `json:"data"`
}

const (
	OperationTypeCancel        OperationType = "CANCEL"
	OperationTypeContinue      OperationType = "CONTINUE"
	OperationTypeUpdateContext OperationType = "UPDATE_CONTEXT"
)

type OperationType string

func (ot OperationType) String() string {
	return string(ot)
}

func (ot OperationType) Upper() OperationType {
	return OperationType(strings.ToUpper(ot.String()))
}

func (ot OperationType) IsCancel() bool {
	return ot == OperationTypeCancel
}

func (ot OperationType) IsContinue() bool {
	return ot == OperationTypeContinue
}

func (ot OperationType) IsUpdateContext() bool {
	return ot == OperationTypeUpdateContext
}

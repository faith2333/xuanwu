package definition

import "strings"

type Variable struct {
	Key          string       `yaml:"key"`
	Type         VariableType `yaml:"type"`
	DefaultValue interface{}  `json:"defaultValue"`
	Required     bool         `json:"required"`
}

const (
	VariableTypeInt    VariableType = "INT"
	VariableTypeString VariableType = "STRING"
	VariableTypeSlice  VariableType = "SLICE"
	VariableTypeBool   VariableType = "BOOL"
	VariableTypeMap    VariableType = "MAP"
)

var AllVariableTypes = []VariableType{
	VariableTypeInt,
	VariableTypeString,
	VariableTypeSlice,
	VariableTypeBool,
	VariableTypeMap,
}

type VariableType string

func (vt VariableType) String() string {
	return string(vt)
}

func (vt VariableType) Upper() VariableType {
	return VariableType(strings.ToUpper(vt.String()))
}

func (vt VariableType) IsSupported() bool {
	for _, allVT := range AllVariableTypes {
		if vt.Upper() == allVT {
			return true
		}
	}
	return false
}

func (vt VariableType) IsInt() bool {
	return vt.Upper() == VariableTypeInt
}

func (vt VariableType) IsString() bool {
	return vt.Upper() == VariableTypeString
}

func (vt VariableType) IsBool() bool {
	return vt.Upper() == VariableTypeBool
}

func (vt VariableType) IsSlice() bool {
	return vt.Upper() == VariableTypeSlice
}

func (vt VariableType) IsMap() bool {
	return vt.Upper() == VariableTypeMap
}

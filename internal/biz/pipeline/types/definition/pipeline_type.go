package definition

import "strings"

const (
	PipelineTypeCICD PipelineType = "CICD"
)

var AllPipelineTypes = []PipelineType{
	PipelineTypeCICD,
}

type PipelineType string

func (t PipelineType) String() string {
	return string(t)
}

func (t PipelineType) Upper() PipelineType {
	return PipelineType(strings.ToUpper(t.String()))
}

func (t PipelineType) IsCICD() bool {
	return t.Upper() == PipelineTypeCICD
}

func (t PipelineType) IsSupported() bool {
	for _, tmpT := range AllPipelineTypes {
		if t.Upper() == tmpT {
			return true
		}
	}
	return false
}

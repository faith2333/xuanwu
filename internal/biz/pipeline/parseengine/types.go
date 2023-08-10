package parseengine

import "github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"

type innerVariable struct {
	Key   string                  `json:"key"`
	Value interface{}             `json:"value"`
	Type  definition.VariableType `json:"type"`
}

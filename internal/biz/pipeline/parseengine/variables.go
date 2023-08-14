package parseengine

import (
	"fmt"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
)

const (
	VariableLeftSymbol  = "${"
	VariableRightSymbol = "}$"
)

type innerVariable struct {
	Key   string                  `json:"key"`
	Value interface{}             `json:"value"`
	Type  definition.VariableType `json:"type"`
}

// key => ${key}$
func (eng *defaultEngine) makeVariableDefinitionWithKey(key string) string {
	return fmt.Sprintf("%s%s%s", VariableLeftSymbol, key, VariableRightSymbol)
}

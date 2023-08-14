package parseengine

import (
	"encoding/json"
	"fmt"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

func (eng *defaultEngine) Repeat(stage *definition.Stage, variables map[string]*innerVariable) (repeatedStages []*definition.Stage,
	newVariables map[string]*innerVariable, err error) {
	repeatVariable, ok := variables[stage.RepeatFromVariable]
	if !ok {
		return nil, nil, errors.Errorf("repeat variable %q is not exist", stage.RepeatFromVariable)
	}

	repeatVarMap := make(map[string]string)
	switch repeatVariable.Type {
	case definition.VariableTypeMap:
		rVar, ok := repeatVariable.Value.(map[string]string)
		if !ok {
			return nil, nil, errors.Errorf("repeat variable element request string type, but got: %s", reflect.TypeOf(repeatVariable.Value))
		}

		for k, v := range rVar {
			repeatVarMap[k] = v
		}
	case definition.VariableTypeSlice:
		rVar, ok := repeatVariable.Value.([]string)
		if !ok {
			return nil, nil, errors.Errorf("repeat variable element request string type, but got: %s", reflect.TypeOf(repeatVariable.Value))
		}

		for i := range rVar {
			repeatVarMap[fmt.Sprintf("%d", i)] = rVar[i]
		}
	default:
		return nil, nil, errors.Errorf("repeat variable type has not been supported")
	}

	// replace the origin variable key with generated repeated key
	stageBytes, err := json.Marshal(stage)
	if err != nil {
		return nil, nil, errors.Errorf("marshal stage as json failed: %v", err)
	}

	for k, v := range repeatVarMap {
		newVariableKey := eng.makeRepeatVariable(repeatVariable.Key, k)

		newStageStr := strings.Replace(string(stageBytes), eng.makeVariableDefinitionWithKey(repeatVariable.Key),
			eng.makeVariableDefinitionWithKey(newVariableKey), -1)

		newStage := &definition.Stage{}
		err = json.Unmarshal([]byte(newStageStr), &newStage)
		if err != nil {
			return nil, nil, errors.Errorf("unmarshal new stage failed: %v", err)
		}

		newStage.Name = eng.makeRepeatStageName(stage.Name, k)

		repeatedStages = append(repeatedStages, newStage)
		newVariables[newVariableKey] = &innerVariable{
			Key:   newVariableKey,
			Type:  definition.VariableTypeString,
			Value: v,
		}
	}

	for k, v := range variables {
		if k == stage.RepeatFromVariable {
			continue
		}

		newVariables[k] = v
	}

	return
}

func (eng *defaultEngine) makeRepeatVariable(origin, addonKey string) string {
	return fmt.Sprintf("%s-%s", origin, addonKey)
}

func (eng *defaultEngine) makeRepeatStageName(origin, addonKey string) string {
	return fmt.Sprintf("%s-%s", origin, addonKey)
}

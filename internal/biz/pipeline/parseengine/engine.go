package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type defaultEngine struct{}

func (eng *defaultEngine) ParseAndGenerate(ctx context.Context, pipeline *definition.Pipeline, variables map[string]interface{}) (*runtime.Instance, []*runtime.Node, error) {
	err := pipeline.Validate()
	if err != nil {
		return nil, nil, errors.Errorf("validate pipeline definition failed: %v", err)
	}

	innerVariables, err := eng.validateVariables(pipeline, variables)
	if err != nil {
		return nil, nil, errors.Errorf("validate variables failed: %v", err)
	}

	inst, err := eng.generateInstance(ctx, pipeline, innerVariables)
	if err != nil {
		return nil, nil, errors.Errorf("generate instance failed: %v", err)
	}

	nodes, err := eng.generateNodes(ctx, pipeline, inst.Code, innerVariables)
	if err != nil {
		return nil, nil, errors.Errorf("generate nodes failed: %v", err)
	}

	return inst, nodes, nil
}

func (eng *defaultEngine) generateInstance(ctx context.Context, pipeline *definition.Pipeline, variables map[string]*innerVariable) (*runtime.Instance, error) {
	inst := &runtime.Instance{
		Code:            uuid.New().String(),
		Name:            pipeline.Name,
		DefinitionCode:  pipeline.Code,
		GlobalVariables: make(map[string]interface{}),
		Status:          runtime.StatusInitialization,
	}

	for _, v := range variables {
		inst.GlobalVariables[v.Key] = v.Value
	}

	return inst, nil
}

func (eng *defaultEngine) validateVariables(pipeline *definition.Pipeline, variables map[string]interface{}) (map[string]*innerVariable, error) {
	innerVariables := make(map[string]*innerVariable)
	for _, vDef := range pipeline.GlobalVariables {
		v, ok := variables[vDef.Key]
		if !ok {
			if vDef.Required {
				return nil, errors.Errorf("variable %q is required, but not specified when instantiated", vDef.Key)
			}

			// use default value
			if vDef.DefaultValue != nil {
				innerVariables[vDef.Key] = &innerVariable{
					Key:   vDef.Key,
					Value: vDef.DefaultValue,
					Type:  vDef.Type,
				}
			}
		} else {
			vType := definition.GetVariableType(v)
			if vType != vDef.Type {
				return nil, errors.Errorf("variable %q type error, request %q but got %q", vDef.Key, vType, vDef.Type)
			}
			innerVariables[vDef.Key] = &innerVariable{
				Key:   vDef.Key,
				Value: v,
				Type:  vType,
			}
		}
	}
	return innerVariables, nil
}

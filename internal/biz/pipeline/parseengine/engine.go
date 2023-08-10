package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/base"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type defaultEngine struct {
	instRepo IInstanceRepo
	nodeRepo INodeRepo
	tx       base.Transaction
}

func (de *defaultEngine) ParseAndGenerate(ctx context.Context, pipeline *definition.Pipeline, variables map[string]interface{}) (*runtime.Instance, []*runtime.Node, error) {
	err := pipeline.Validate()
	if err != nil {
		return nil, nil, errors.Errorf("validate pipeline definition failed: %v", err)
	}

	innerVariables, err := de.validateVariables(pipeline, variables)
	if err != nil {
		return nil, nil, errors.Errorf("validate variables failed: %v", err)
	}

	inst, err := de.generateInstance(ctx, pipeline, innerVariables)
	if err != nil {
		return nil, nil, errors.Errorf("generate instance failed: %v", err)
	}

	nodes, err := de.generateNodes(ctx, pipeline, inst.Code, innerVariables)
	if err != nil {
		return nil, nil, errors.Errorf("generate nodes failed: %v", err)
	}

	return inst, nodes, nil
}

func (de *defaultEngine) generateInstance(ctx context.Context, pipeline *definition.Pipeline, variables map[string]*innerVariable) (*runtime.Instance, error) {
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

func (de *defaultEngine) Save(ctx context.Context, instance *runtime.Instance, nodes []*runtime.Node) error {
	err := de.tx.ExecTx(ctx, func(ctx context.Context) error {
		_, err := de.instRepo.Create(ctx, instance)
		if err != nil {
			return err
		}

		for _, node := range nodes {
			_, err = de.nodeRepo.Create(ctx, node)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (de *defaultEngine) validateVariables(pipeline *definition.Pipeline, variables map[string]interface{}) (map[string]*innerVariable, error) {
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

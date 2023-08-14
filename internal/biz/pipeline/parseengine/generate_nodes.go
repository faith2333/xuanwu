package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (eng *defaultEngine) generateNodes(ctx context.Context, pipeline *definition.Pipeline, instCode string, variables map[string]*innerVariable) ([]*runtime.Node, error) {
	nodes := make([]*runtime.Node, 0)

	// generate nodes by stage definition
	var preStagesName []string
	for _, stages := range pipeline.Stages {
		var curStepNodes []*runtime.Node
		var err error
		curStepNodes, preStagesName, err = eng.genNodeByStageWithParams(stages, instCode, preStagesName, variables)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, curStepNodes...)
	}
	return nodes, nil
}

func (eng *defaultEngine) genNodeByStageWithParams(stages []*definition.Stage, instCode string, preStagesName []string, variables map[string]*innerVariable) ([]*runtime.Node, []string, error) {
	curStagesName := make([]string, 0)
	nodes := make([]*runtime.Node, 0)

	for _, stage := range stages {
		// repeat the stages if stage repeat is specified
		if stage.Repeat {
			repeatedNodes, repeatedStagesName, err := eng.repeatStageAndGenNodes(stage, instCode, preStagesName, variables)
			if err != nil {
				return nil, nil, err
			}

			nodes = append(nodes, repeatedNodes...)
			curStagesName = append(curStagesName, repeatedStagesName...)
			continue
		}

		node := &runtime.Node{
			Name:         stage.Name,
			Code:         uuid.New().String(),
			InstanceCode: instCode,
			ExecutorName: stage.ExecutorName,
			ExecuteInfo:  stage.ExecuteInfo,
			RunAfter:     preStagesName,
		}

		nodes = append(nodes, node)
		if stage.NextStages == nil {
			curStagesName = append(curStagesName, stage.Name)
			continue
		}

		for _, tmpStages := range stage.NextStages {
			nextNodes, nextLastNodeNames, err := eng.genNodeByStageWithParams(tmpStages, instCode, []string{stage.Name}, variables)
			if err != nil {
				return nil, nil, err
			}

			nodes = append(nodes, nextNodes...)
			curStagesName = append(curStagesName, nextLastNodeNames...)
		}
	}

	return nodes, curStagesName, nil
}

func (eng *defaultEngine) repeatStageAndGenNodes(stage *definition.Stage, instCode string, preStagesName []string, variables map[string]*innerVariable) ([]*runtime.Node, []string, error) {
	repeatedStages, newVariables, err := eng.Repeat(stage, variables)
	if err != nil {
		return nil, nil, errors.Errorf("repeat stage failed: %v", err)
	}
	repeatedNodes, lastStagesName, err := eng.genNodeByStageWithParams(repeatedStages, instCode, preStagesName, newVariables)
	if err != nil {
		return nil, nil, err
	}

}

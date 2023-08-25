package processengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker"
	lwTypes "github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
)

type Engine struct {
	lw leaderworker.Interface
}

func NewEngine(lw leaderworker.Interface) *Engine {
	return &Engine{
		lw: lw,
	}
}

func (e *Engine) Start(ctx context.Context) error {
	//todo start the engine
	return nil
}

func (e *Engine) Stop(ctx context.Context) error {
	//todo stop the engine
	return nil
}

func (e *Engine) StartPipeline(ctx context.Context, pipeline *definition.Pipeline) error {
	//todo distribute start a pipeline
	return nil
}

func (e *Engine) UpdatePipeline(ctx context.Context, pipelineID string, operation *lwTypes.UpdateOperation) error {
	//todo distribute update pipeline
	return nil
}

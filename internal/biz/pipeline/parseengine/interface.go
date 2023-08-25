package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
)

type Interface interface {
	// ParseAndGenerate parse pipeline definition and generate instance and nodes
	ParseAndGenerate(ctx context.Context, pipeline *definition.Pipeline, variables map[string]interface{}) (*runtime.Instance, []*runtime.Node, error)
}

func New() Interface {
	return &defaultEngine{}
}

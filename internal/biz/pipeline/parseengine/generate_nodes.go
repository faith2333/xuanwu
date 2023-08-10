package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
)

func (de *defaultEngine) generateNodes(ctx context.Context, pipeline *definition.Pipeline, instCode string, variables map[string]*innerVariable) ([]*runtime.Node, error) {

}

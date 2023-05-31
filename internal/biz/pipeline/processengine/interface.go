package processengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types"
)

type Interface interface {
	// Start implement the Server start method in kratos transport
	Start(ctx context.Context) error
	// Stop implement the Server stop method in kratos transport
	Stop(ctx context.Context) error

	// StartPipeline start schedule and execute pipeline
	StartPipeline(ctx context.Context, pipeline *types.Pipeline) error
	// StopPipeline stop the pipeline
	StopPipeline(ctx context.Context, pipeline *types.Pipeline) error
}

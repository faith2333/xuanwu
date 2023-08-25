package processengine

import (
	"context"
	lwTypes "github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
)

type Interface interface {
	// Start implement the Server start method in kratos transport
	Start(ctx context.Context) error
	// Stop implement the Server stop method in kratos transport
	Stop(ctx context.Context) error

	// StartPipeline start schedule and execute pipeline
	StartPipeline(ctx context.Context, pipeline *definition.Pipeline) error
	// UpdatePipeline update the pipeline with specified operation
	UpdatePipeline(ctx context.Context, pipelineID string, operation *lwTypes.UpdateOperation) error
}

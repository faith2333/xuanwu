package parseengine

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/base"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/definition"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/types/runtime"
)

type Interface interface {
	// ParseAndGenerate parse pipeline definition and generate instance and nodes
	ParseAndGenerate(ctx context.Context, pipeline *definition.Pipeline, variables map[string]interface{}) (*runtime.Instance, []*runtime.Node, error)
	// Save deposit the instance and nodes to persistent storage
	Save(ctx context.Context, instance *runtime.Instance, nodes []*runtime.Node) error
}

// IInstanceRepo the database operation interface for pipeline instance
type IInstanceRepo interface {
	Create(ctx context.Context, inst *runtime.Instance) (*runtime.Instance, error)
	UpdateStatus(ctx context.Context, instCode string, status runtime.Status) error
	GetByCode(ctx context.Context, instCode string) (*runtime.Instance, error)
	GetRunningInstByPipelineCode(ctx context.Context, defCode string) (*runtime.Instance, error)
}

// INodeRepo the database operation interface for pipeline node
type INodeRepo interface {
	Create(ctx context.Context, node *runtime.Node) (*runtime.Node, error)
	UpdateStatus(ctx context.Context, nodeCode string, status runtime.Status) error
	ListByInstCode(ctx context.Context, instCode string) ([]*runtime.Node, error)
}

func New(instRepo IInstanceRepo, nodeRepo INodeRepo, tx base.Transaction) Interface {
	return &defaultEngine{
		instRepo: instRepo,
		nodeRepo: nodeRepo,
		tx:       tx,
	}
}

package leaderworker

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Interface interface {
	ILeader
	IWorker
	IGeneral
}

// ILeader the leader is responsible for worker add or remove and assign task.
type ILeader interface {
	// RunLeader run leader instances
	RunLeader(ctx context.Context) error
	// OnLeader register listener which webhook function will be called before and after leader execute.
	OnLeader(l types.Listener)
	// InitWorkersList init the list which contains all worker information.
	InitWorkersList(ctx context.Context) error
	// OnWorkerAdd register webhook function which will be called when the instance run as leader and
	// add worker to its worker list
	OnWorkerAdd(webhookFunc types.WebhookFunc)
	// OnWorkerRemove register webhook function which will be called when instance run as leader and
	// remove worker from its worker list
	OnWorkerRemove(webhookFunc types.WebhookFunc)
	// AssignLogicTaskToWorker assign task to worker
	AssignLogicTaskToWorker(ctx context.Context, workerID types.WorkerID, task types.LogicTask) error
}

// IWorker the worker is responsible for task handler.
type IWorker interface {
	// RunWorker run a worker instance
	RunWorker(ctx context.Context) error
	// HandleTask handle the task which assigned by worker
	HandleTask(ctx context.Context, task types.LogicTask) error
}

type IGeneral interface {
	// ListWorkers list all worker
	ListWorkers(ctx context.Context) ([]types.WorkerID, error)
	// WatchPrefix watch events occurred by etcd key with specified prefix, and do the passed handler function.
	WatchPrefix(ctx context.Context, prefix string, putHandler, deleteHandler func(context.Context, *clientv3.Event))
	// UpdateLogicTask update the logic task, it can stop a running task, continue a stopped task or update a task context.
	UpdateLogicTask(ctx context.Context, logicTask types.LogicTask, operation *types.UpdateOperation) error
}

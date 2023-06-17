package leaderworker

import (
	"context"
	"github.com/buraksezer/consistent"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/ratelimitqueue"
	"github.com/faith2333/xuanwu/pkg/election"
	clientV3 "go.etcd.io/etcd/client/v3"
)

type forLeaderUse struct {
	etcdClient *clientV3.Client

	election       election.Interface
	rateLimitQueue ratelimitqueue.Interface

	// the consistent map was used by leader to assign tasks to the workers
	workerConsistent *consistent.Consistent

	allWorkers     map[types.WorkerID]struct{}
	workerTasksMap map[types.WorkerID]map[types.LogicTaskID]types.LogicTask
	taskWorkerMap  map[types.LogicTaskID]types.WorkerID

	listeners              []types.Listener
	webhooksOnWorkerAdd    []types.WebhookFunc
	webhooksOnWorkerRemove []types.WebhookFunc
}

type forLeaderUseConfig struct {
}

func (lw *defaultLeaderWorker) RunLeader(ctx context.Context) error {
	//todo
	return nil
}

func (lw *defaultLeaderWorker) OnLeader(l types.Listener) {
	lw.lock.Lock()
	defer lw.lock.Unlock()

	lw.listeners = append(lw.listeners, l)
}

func (lw *defaultLeaderWorker) OnWorkerAdd(fn types.WebhookFunc) {
	lw.lock.Lock()
	defer lw.lock.Unlock()

	lw.webhooksOnWorkerAdd = append(lw.webhooksOnWorkerAdd, fn)
}

func (lw *defaultLeaderWorker) OnWorkerRemove(fn types.WebhookFunc) {
	lw.lock.Lock()
	defer lw.lock.Unlock()

	lw.webhooksOnWorkerRemove = append(lw.webhooksOnWorkerRemove, fn)
}

func (lw *defaultLeaderWorker) InitWorkersList(ctx context.Context) error {
	//todo
	return nil
}

func (lw *defaultLeaderWorker) AssignLogicTaskToWorker(ctx context.Context, workerID types.WorkerID, logicTask types.LogicTask) error {
	//todo
	return nil
}

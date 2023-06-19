package leaderworker

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	"github.com/faith2333/xuanwu/pkg/safego"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type Config struct {
	leaderElectionKey  string `default:"/election/leader"`
	forLeaderUseConfig forLeaderUseConfig
	forWorkerUseConfig forWorkerUseConfig
}

type defaultLeaderWorker struct {
	cfg  *Config
	lock *sync.RWMutex
	log  *log.Helper
	forLeaderUse
	forWorkerUse
}

func (lw *defaultLeaderWorker) ListWorkers(ctx context.Context) ([]types.WorkerID, error) {
	//todo
	return nil, nil
}

func (lw *defaultLeaderWorker) WatchPrefix(ctx context.Context, prefix string, putHandler, deleteHandler func(ctx context.Context, event *clientV3.Event)) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// the function will be blocked here, if something error happened in watchPrefix function,
			// it will auto retry until ctx done received
			lw.watchPrefix(ctx, prefix, putHandler, deleteHandler)
		}
	}
}

func (lw *defaultLeaderWorker) watchPrefix(ctx context.Context, prefix string, putHandler, deleteHandler func(ctx context.Context, event *clientV3.Event)) {
	wCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	watchChan := lw.etcdClient.Watch(wCtx, prefix, clientV3.WithPrefix())

	for {
		select {
		case <-ctx.Done():
			lw.log.Warnf("context done received, watch for etcd events with prefix %s stopped", prefix)
			return
		case watchResp, ok := <-watchChan:
			if !ok {
				lw.log.Debugf("watch chan closed, (auto  retry)")
				continue
			}
			if watchResp.Err() != nil {
				lw.log.Errorf("watch response error(auto retry) %v", watchResp.Err())
				continue
			}

			for _, event := range watchResp.Events {
				lw.log.Debugf("event %#v for prefix %s received", prefix, event)
				if event.Kv == nil {
					continue
				}
				switch event.Type {
				case mvccpb.PUT:
					safego.Go(func() {
						putHandler(wCtx, event)
					})
				case mvccpb.DELETE:
					safego.Go(func() {
						deleteHandler(wCtx, event)
					})
				}
			}
		}
	}
}

func (lw *defaultLeaderWorker) UpdateLogicTask(ctx context.Context, logicTask types.LogicTask, operation *types.UpdateOperation) error {
	//todo
	return nil
}

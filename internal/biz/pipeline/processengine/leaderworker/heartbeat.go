package leaderworker

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/etcdkey"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
)

func (lw *defaultLeaderWorker) HeartBeatForWorker(ctx context.Context, workerID types.WorkerID) {

}

func (lw *defaultLeaderWorker) probeOnce(ctx context.Context, workerID types.WorkerID) error {
	resp, err := lw.etcdClient.Get(ctx, etcdkey.MakeWorkerHeartBeatKeyByWorkerID(workerID))
	if err != nil {
		return err
	}

}

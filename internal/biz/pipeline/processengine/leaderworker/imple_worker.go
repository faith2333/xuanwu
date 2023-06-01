package leaderworker

import (
	"context"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"
	"github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/reconciler"
)

type forWorkerUse struct {
	reconciler reconciler.Interface
}

type forWorkerUseConfig struct {
}

func (lw *defaultLeaderWorker) RunWorker(ctx context.Context) error {
	//todo
	return nil
}

func (lw *defaultLeaderWorker) HandleTask(ctx context.Context, task types.LogicTask) error {
	//todo
	return nil
}

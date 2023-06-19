package etcdkey

import "github.com/faith2333/xuanwu/internal/biz/pipeline/processengine/leaderworker/types"

const (
	LeaderWorkerPrefix = "/xuanwu/lw"
)

// eg: /xuanwu/lw/heartbeat/worker
func makeWorkerHeartBeatPrefix() string {
	return LeaderWorkerPrefix + "/" + "heartbeat" + "/" + "worker"
}

// MakeWorkerHeartBeatKeyByWorkerID eg: /xuanwu/lw/heartbeat/worker/worker-123
func MakeWorkerHeartBeatKeyByWorkerID(workerID types.WorkerID) string {
	return makeWorkerHeartBeatPrefix() + "/" + workerID.String()
}

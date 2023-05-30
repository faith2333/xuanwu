package election

import (
	"context"
	"github.com/faith2333/xuanwu/pkg/safego"
	"github.com/faith2333/xuanwu/pkg/xforloop"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"sync"
)

type EtcdElection struct {
	log        *log.Helper
	lock       *sync.RWMutex
	etcdClient *clientv3.Client
	isLeader   bool
	isReady    bool
}

func NewEtcdElection(etcdClient *clientv3.Client) *EtcdElection {
	return &EtcdElection{
		lock:       &sync.RWMutex{},
		etcdClient: etcdClient,
		isLeader:   false,
		isReady:    false,
	}
}

func (e *EtcdElection) Champaign(ctx context.Context, prefix, name string) {
	safego.Go(func() {
		xforloop.LoopWithContextAndRetryBackoffLimit(ctx, func(ctx context.Context) bool {
			// clear status before do anything
			e.Resign(ctx)

			session, err := concurrency.NewSession(e.etcdClient)
			if err != nil {
				e.log.Errorf("new session failed (auto retry): %v", err)
				return true
			}

			election := concurrency.NewElection(session, prefix)

			// the election is ready
			e.lock.Lock()
			e.isReady = true
			e.lock.Unlock()

			if err = election.Campaign(ctx, name); err != nil {
				e.log.Errorf("election campaign failed (auto retry): %v", err)
				return true
			}

			// the election is success
			e.lock.Lock()
			e.isLeader = true
			e.lock.Unlock()

			select {
			case <-ctx.Done():
				e.log.Errorf("context done received, election finished without retry")
				// we should resign from etcd when context done
				err = election.Resign(ctx)
				if err != nil {
					e.log.Errorf("node %s resign failed: %v", name, err)
				}
				// clear status here
				e.Resign(ctx)
				return false
			case <-session.Done():
				e.log.Errorf("session done received, retry election")
				return true
			}
		})
	})
}

func (e *EtcdElection) Resign(ctx context.Context) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.isReady = false
	e.isLeader = false
}

func (e *EtcdElection) IsReady() bool {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.isReady
}

func (e *EtcdElection) IsLeader() bool {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.isLeader
}

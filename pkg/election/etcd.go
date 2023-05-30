package election

import (
	"context"
	"github.com/faith2333/xuanwu/pkg/safego"
	"github.com/faith2333/xuanwu/pkg/xforloop"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
)

type EtcdElection struct {
	log        *log.Helper
	lock       *sync.RWMutex
	etcdClient *clientv3.Client
	session    *concurrency.Session
	election   *concurrency.Election
	isLeader   bool
	isReady    bool
}

func NewEtcdElection(etcdClient *clientv3.Client, logger log.Logger) *EtcdElection {
	return &EtcdElection{
		log:        log.NewHelper(logger),
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
				return false
			}

			e.lock.Lock()
			e.session = session
			e.election = concurrency.NewElection(session, prefix)
			e.isReady = true
			e.lock.Unlock()

			if err = e.election.Campaign(ctx, name); err != nil {
				e.log.Errorf("election campaign failed (auto retry): %v", err)
				return false
			}
			e.log.Infof("leader election success for %s", name)

			// the election is success
			e.lock.Lock()
			e.isLeader = true
			e.lock.Unlock()

			select {
			case <-ctx.Done():
				e.log.Errorf("context done received, election finished without retry")

				// we should resign from etcd and clear status when received context done event
				e.Resign(ctx)
				return true
			case <-session.Done():
				e.log.Errorf("session done received, retry election")
				return false
			}
		})
	})
}

// Resign do resign from etcd and clear status
func (e *EtcdElection) Resign(ctx context.Context) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.isReady = false
	e.isLeader = false

	if e.election != nil {
		_ = e.election.Resign(ctx)
		e.election = nil
	}

	if e.session != nil {
		_ = e.session.Close()
		e.session = nil
	}
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

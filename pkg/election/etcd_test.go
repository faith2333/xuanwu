package election

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.etcd.io/etcd/client/v3"
	"os"
	"testing"
	"time"
)

var etcdElection1 *EtcdElection
var etcdElection2 *EtcdElection

func TestMain(m *testing.M) {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}

	etcdElection1 = NewEtcdElection(etcdClient, log.NewStdLogger(os.Stdout))
	etcdElection2 = NewEtcdElection(etcdClient, log.NewStdLogger(os.Stdout))
	os.Exit(m.Run())
}

func TestEtcdElection_Champaign(t *testing.T) {
	etcdElection1.Champaign(context.Background(), "/election/leader", "node01")
	etcdElection2.Champaign(context.Background(), "/election/leader", "node02")

	count := 0
	for {
		if count > 20 {
			break
		}
		count++

		if !etcdElection1.IsReady() && !etcdElection2.IsReady() {
			fmt.Println("node01 and node02 both not ready")
			time.Sleep(time.Second * 3)
			continue
		}

		if etcdElection1.IsLeader() {
			fmt.Println("node01 is leader")
			etcdElection1.Resign(context.TODO())
		} else if etcdElection2.IsLeader() {
			fmt.Println("node02 is leader")
			etcdElection2.Resign(context.TODO())
		}
		time.Sleep(time.Second * 5)
	}
}

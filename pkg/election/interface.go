package election

import "context"

type Interface interface {
	Champaign(ctx context.Context, prefix, name string)
	Resign(ctx context.Context)
	IsLeader() bool
	IsReady() bool
}

package base

import "context"

type Transaction interface {
	ExecTx(ctx context.Context, fn func(ctx context.Context) error) error
}

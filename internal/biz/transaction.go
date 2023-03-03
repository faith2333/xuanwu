package biz

import "context"

type Transaction interface {
	// ExecTx Execute anonymous method with transaction object which found in context or just created.
	ExecTx(context.Context, func(ctx context.Context) error) error
}

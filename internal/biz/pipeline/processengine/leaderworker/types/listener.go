package types

import "context"

type Listener interface {
	BeforeExec(ctx context.Context)
	AfterExec(ctx context.Context)
}

type DefaultListener struct {
	beforeExecFunc WebhookFunc
	afterExecFunc  WebhookFunc
}

func NewDefaultListener(beforeExecFun, afterExecFunc WebhookFunc) *DefaultListener {
	return &DefaultListener{
		beforeExecFunc: beforeExecFun,
		afterExecFunc:  afterExecFunc,
	}
}

func (dl *DefaultListener) BeforeExec(ctx context.Context) {
	if dl.beforeExecFunc == nil {
		return
	}
	dl.beforeExecFunc(ctx)
}

func (dl *DefaultListener) AfterExec(ctx context.Context) {
	if dl.afterExecFunc == nil {
		return
	}
	dl.afterExecFunc(ctx)
}

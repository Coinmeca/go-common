package commoncontext

import (
	"context"
	"github.com/coinmeca/go-common/commondefine"
	"time"
)

type Context interface {
	context.Context
	ChainId(chainName string) string
	UserId() string
}

type commonContext struct {
	context.Context
	userId string
}

func (c *commonContext) ChainId(chainName string) string {
	return commondefine.ChainIdMap[chainName]
}

func (c *commonContext) UserId() string {
	return c.userId
}

func NewCommonContext(ctx context.Context, userId string) Context {
	return &commonContext{
		Context: ctx,
		userId:  userId,
	}
}

func WithCancel(parent Context) (Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(parent)
	return NewCommonContext(ctx, parent.UserId()), cancel
}

func WithTimeout(parent Context, timeout time.Duration) (Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(parent, timeout)
	return NewCommonContext(ctx, parent.UserId()), cancel
}

func WithValue(parent Context, key, val interface{}) Context {
	ctx := context.WithValue(parent, key, val)
	return NewCommonContext(ctx, parent.UserId())
}

func Background() Context {
	return NewCommonContext(context.Background(), "")
}

func TODO() Context {
	return NewCommonContext(context.TODO(), "")
}

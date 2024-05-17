package commoncontext

import (
	"context"
	"github.com/coinmeca/go-common/commondefine"
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

func NewCommonContext(ctx context.Context, userId string, chainIdMap map[string]string) Context {
	return &commonContext{
		Context: ctx,
		userId:  userId,
	}
}

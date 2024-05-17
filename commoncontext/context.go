package commoncontext

import "context"

type Context interface {
	context.Context
	ChainId() string
	UserId() string
}

type commonContext struct {
	context.Context
	chainId string
	userId  string
}

func (c *commonContext) ChainId() string {
	return c.chainId
}

func (c *commonContext) UserId() string {
	return c.userId
}

func NewCommonContext(ctx context.Context, chainId, userId string) Context {
	return &commonContext{
		Context: ctx,
		chainId: chainId,
		userId:  userId,
	}
}

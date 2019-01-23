package client

import "github.com/yijizhichang/wechat-sdk/mp/core"

type Client struct {
	*core.Context
}

func NewClient(context *core.Context) *Client {
	clt := new(Client)
	clt.Context = context
	return clt
}

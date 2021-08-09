//企业微信实例
package client

import "github.com/yijizhichang/wechat-sdk/work/core"

type Client struct {
	*core.Context
}

func NewQyClient(context *core.Context) *Client {
	clt := new(Client)
	clt.Context = context
	return clt
}

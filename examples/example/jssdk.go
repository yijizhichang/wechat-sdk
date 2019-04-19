package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

func GetJSSign() {
	jssdk := wxconf.WechatClient.GetJSSDK()
	tickect, err := jssdk.GetTicket()
	if err != nil {
		fmt.Println("获取ticket error", err)
		return
	}

	sign := jssdk.MakeSign(tickect, "http://localhost")
	fmt.Println("签名参数", *sign)
}

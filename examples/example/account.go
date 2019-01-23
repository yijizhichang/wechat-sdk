package example

import (
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"fmt"
)

func AccountManage()  {
	account := wxconf.WechatClient.GetAccount()

	//创建临时二维码 scene_id
	//re, err := account.CreateQrCodeSceneId(false,300,1001)
	//fmt.Println("创建二维码：", re,  "Err:", err)

	//创建永久二维码 scene_id
	//re, err := account.CreateQrCodeSceneId(true,1001,0)
	//fmt.Println("创建二维码：", re,  "Err:", err)

	//创建临时二维码 scene_str
	//re, err := account.CreateQrCodeSceneStr(false,"test_001",300)
	//fmt.Println("创建二维码：", re,  "Err:", err)

	//创建永久二维码 scene_str
	re, err := account.CreateQrCodeSceneStr(true,"test_004",0)
	fmt.Println("创建二维码：", re,  "Err:", err)

	//通过ticket换取二维码
	qrcodeUrl := account.GetQrCodeUrl(re.Ticket)
	fmt.Println("Url：", qrcodeUrl)

	//长链接转短链接
	re2,err := account.ShortUrl(qrcodeUrl)
	fmt.Println("长链接转短链接：", re2,err)

}


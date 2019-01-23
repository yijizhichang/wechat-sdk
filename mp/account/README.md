## 账号管理

- [账号管理](#账号管理)
	- [生成带参数的二维码](#生成带参数的二维码)
	    - 创建临时二维码 scene_id
        - 创建永久二维码 scene_id
        - 创建临时二维码 scene_str
        - 创建永久二维码 scene_str
	- [获取二维码](#获取二维码)
		- 通过ticket换取二维码
	- [短链接](#短链接)
		- 长链接转短链接



## 账号管理

```go
为了满足用户渠道推广分析和用户帐号绑定等场景的需要，公众平台提供了生成带参数二维码的接口。使用该接口可以获得多个带不同场景值的二维码，用户扫描后，公众号可以接收到事件推送。
```

具体参数请参考微信文档：[账号管理](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1443433542)


## 二维码管理

##### 生成带参数的二维码
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	account := wxconf.WechatClient.GetAccount()


	//创建临时二维码 scene_id
	re, err := account.CreateQrCodeSceneId(false,300,1001)


	//创建永久二维码 scene_id
	//re, err := account.CreateQrCodeSceneId(true,1001,0)


	//创建临时二维码 scene_str
	//re, err := account.CreateQrCodeSceneStr(false,"test_001",300)


	//创建永久二维码 scene_str
	re, err := account.CreateQrCodeSceneStr(true,"test_004",0)

```


##### 获取二维码

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	account := wxconf.WechatClient.GetAccount()

	//获取二维码地址 参数ticket为创建二维时返回的ticket
	qrcodeUrl := account.GetQrCodeUrl(re.Ticket)

```


##### 长链接转短链接

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	account := wxconf.WechatClient.GetAccount()

	//长链接转短链接
	re,err := account.ShortUrl(longUrl)

```

详细Demo：[examples/example/account.go](examples/example/account.go)

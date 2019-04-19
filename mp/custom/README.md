## 客服管理

- [客服管理](#客服管理)
	- [客服设置](#客服设置)
	    - 增加客服账号
        - 修改客服账号
        - 删除客服账号
        - 设置客服头像
	- [客服消息](#客服消息)
		- 发送文本消息
		- 发送图片消息
		- 发送语音消息
		- 发送视频消息
		- 发送音乐消息
		- 发送图文消息
		- 发送卡片消息
		- 发送小程序卡片
	- [会话控制](#会话控制)
	    - 创建会话
	    - 关闭会话
	    - 获取客户会话状态
	    - 获取客服会话列表
	    - 获取未接入会话列表
	- [聊天记录](#聊天记录)

## 客服管理

```go
当用户和公众号产生特定动作的交互时（具体动作列表请见下方说明），微信将会把消息数据推送给开发者，
开发者可以在一段时间内（目前修改为48小时）调用客服接口，通过POST一个JSON数据包来发送消息给普通用户。
此接口主要用于客服等有人工消息处理环节的功能，方便开发者为用户提供更加优质的服务。


```

具体参数请参考微信文档：[客服消息](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140547)
[新版客服功能](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1458044813)

## 客服设置
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	kf := wc.GetCustom()

	//增加客服账号  账号，昵称
    re,err := kf.AddKfAccount("zhangsan@test", "小宜")

    //修改客服账号
    re,err := kf.UpdateKfAccount("zhangsan@test", "小宜2")

    //删除客服账号
    re,err := kf.DelKfAccount("zhangsan@test2")

    //获取客服列表
    re,err := kf.GetKfList()

    //获取在线客服列表
    re,err := kf.GetKfOnlineList

    //邀请绑定客服帐号  参数 客服账号，微信号
    re, err := kf.InviteWorker("zhangsan33@test","ppwqer-102")


    //设置客服头像  账号，用FORM表单方式上传一个多媒体文件
    re,err := kf.SetHeadImgURL("zhangsan@test","./util/upload/test.jpg")

```

## 客服消息
```go

    //文本消息
    text := custom.NewText("客服文本消息发送测试")  //文本

	text.ToUser = "oEYzpw3sceZnppwqerybyzNdWL3Ic"   //toUser
	text.MsgType = "text"                           //类型
	text.CustomService.KfAccount = "zhangsan@test"  //指定客服账号

	re,err := kf.SendMsgByKf(text)


	//图片消息
	img := custom.NewImage("123124435245")  //mediald
	img.ToUser = "abcd1234abcd1234abcd1234"
	img.MsgType = "image"
	img.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(img)
	fmt.Println(re,err)


	//语音
	voice := custom.NewVoice("mediald")
	voice.ToUser = "abcd1234abcd1234abcd1234"
	voice.MsgType = "voice"
	voice.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(voice)


	//视频
	video := custom.NewVideo("mediald", "thumbMediald", "title","description")
	video.ToUser = "abcd1234abcd1234abcd1234"
	video.MsgType = "video"
	video.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(video)


	//音乐
	music := custom.NewMusic("title", "des	cription", "musicURL", "HQMusicUrl", "thumbMediald")
	music.ToUser = "abcd1234abcd1234abcd1234"
	music.MsgType = "video"
	music.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(music)


	//图文  发送图文消息（点击跳转到外链） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
	ar := custom.NewArticle("图文消息", "我是一条图文消息", "https://www.baidu.com/img/bd_logo1.png", "")
	var newsList []*custom.Article
	newsList = append(newsList, ar)
	ars := custom.NewNews(newsList)
	ars.ToUser = "abcd1234abcd1234abcd1234"
	ars.MsgType = "news"
	ars.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(ars)


	//mpnews  发送图文消息（点击跳转到图文消息页面） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
	mpnews := custom.NewMpNews("mediald")
	mpnews.ToUser = "abcd1234abcd1234abcd1234"
	mpnews.MsgType = "video"
	mpnews.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(mpnews)


	//卡券
	card := custom.NewCard("cardid")
	card.ToUser = "abcd1234abcd1234abcd1234"
	card.MsgType = "video"
	card.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(card)


	//小程序卡片
	minipg := custom.NewMiniProgramPage("title","appid","pagepath","thumbmediald")
	minipg.ToUser = "abcd1234abcd1234abcd1234"
	minipg.MsgType = "video"
	minipg.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(minipg)


```

## 会话控制

```go
	//config配置文件省略...
	wc := wechat.NewWechat(config)
	kf := wc.GetCustom()

	//创建会话
	re, err := kf.CreateKfSession("zhangsan@test","abcd1234abcd1234abcd1234")

	//关闭会话
	re, err := kf.CloseKfSession("zhangsan@test","abcd1234abcd1234abcd1234")

	//获取客户会话状态
	re, err := kf.GetKfSession("abcd1234abcd1234abcd1234")

	//获取客服会话列表
	re, err := kf.GetKfSessionList("zhangsan@test")

	//获取未接入会话列表
	re, err := kf.GetWaitCaseList()
```

## 聊天记录

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	kf := wc.GetCustom()

    //获取聊天记录 方法返回对应结构体
    参数: 开始时间，结束时间，信息id, 数量
    re, err := kf.GetMsgList("2018-11-16 00:10:00","2018-11-16 23:10:00",1,1000)

    //说明：
    //时间格式 2018-11-16 00:10:00，每次查询时段不能超过24小时
    //信息id,第一次从1开始，如果一次没获取完，下次输入接口返回的msgid，为下次起始id
    //数量：每次获取条数，最多10000条


    //返回字段说明
    {
      "recordlist"   : [
         {
            "openid"   :  "oDF3iY9WMaswOPWjCIp_f3Bnpljk" ,
            "opercode"   : 2002,
            "text"   :  " 您好，客服test1为您服务。" ,
            "time"   : 1400563710,
            "worker"   :  "test1@test"
         },
         {
            "openid"   :  "oDF3iY9WMaswOPWjCIp_f3Bnpljk" ,
            "opercode"   : 2003,
            "text"   :  "你好，有什么事情？" ,
            "time"   : 1400563731,
            "worker"   :  "test1@test"
         }
      ],
      "number":2,
      "msgid":20165267
   }

```

详细Demo：[examples/example/custom.go](../../examples/example/custom.go)

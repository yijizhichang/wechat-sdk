## 群发管理

- [群发管理](#群发管理)
	- [群发消息](#群发消息)
	    - 根据标签进行群发
        - 根据OpenID列表群发
        - 预览接口
        - 删除群发
        - 查询状态
        - 控制群发速度


## 群发管理

```go
在公众平台网站上，为订阅号提供了每天一条的群发权限，为服务号提供每月（自然月）4条的群发权限。而对于某些具备开发能力的公众号运营者，可以通过高级群发接口，实现更灵活的群发能力。

注意事项：
1、对于认证订阅号，群发接口每天可成功调用1次，此次群发可选择发送给全部用户或某个标签；
2、对于认证服务号虽然开发者使用高级群发接口的每日调用限制为100次，但是用户每月只能接收4条，无论在公众平台网站上，还是使用接口群发，用户每月只能接收4条群发消息，多于4条的群发将对该用户发送失败；
3、开发者可以使用预览接口校对消息样式和排版，通过预览接口可发送编辑好的消息给指定用户校验效果；
4、群发过程中，微信后台会自动进行图文消息原创校验，请提前设置好相关参数(send_ignore等)；
5、开发者可以主动设置 clientmsgid 来避免重复推送。
6、群发接口每分钟限制请求60次，超过限制的请求会被拒绝。
7、图文消息正文中插入自己帐号和其他公众号已群发文章链接的能力。

```

具体参数请参考微信文档：[群发管理](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1481187827_i0l21)


## 群发消息

##### 根据标签进行群发
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	ms := wxconf.WechatClient.GetMass()


	//文本
 	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithTextOption("hello tag1"),
	)


	//图文
	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithMpnewsOption("asdfafasdfsdfasf",0),
	)


	//语音/音频
	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithVoiceOption("asdfafasdfsdfasf"),
	)

	//图片
	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithImageOption("asdfafasdfsdfasf"),
	)

	//视频
	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithMpvideoOption("asdfafasdfsdfasf"),
	)

	//视频
	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithWxcardOption("asdfafasdfsdfasf"),
	)


```

##### 根据OpenID列表群发
```go

    res, err := ms.MassSend(
		mass.WithTouserOption([]string{"oxcLI1WmMa4JVsf_75oDQ17qKqGg", "oxcLI1WmMa4JVsf_75oDQ17qKqGg"}),
		mass.WithTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)


```

##### 预览接口

```go
    //预览接口【订阅号与服务号认证后均可用】
	//开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用。

	//指定openid
	res, err := ms.MassPreview(
		mass.WithPreviewTouserOption("oxcLI1WmMa4JVsf_75oDQ17qKqGg"),
		mass.WithPreviewTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)

	//指定微信号
	res, err := ms.MassPreview(
		mass.WithPreviewTowxnameOption("wxname"),
		mass.WithPreviewTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)
```

##### 删除群发

```go

    err := ms.MassDel(30124,2)

    参数	    是否必须	说明
msg_id	         是	发送出去的消息ID
article_idx	    否	要删除的文章在图文消息中的位置，第一篇编号为1，该字段不填或填0会删除全部文章

```

##### 查询群发消息发送状态

```go

    res,err := ms.MassGet("301234")

    参数	        是否必须	说明
    msg_id	         是	    发送出去的消息ID

```


##### 查询群发消息发送状态

```go

    res,err := ms.MassSpeedSet(2)

    //参数说明

    参数	是否必须	说明
    speed	是	群发速度的级别  群发速度的级别，是一个0到4的整数，数字越大表示群发速度越慢。

    speed 与 realspeed 的关系如下：

    speed	realspeed
    0	    80w/分钟
    1	    60w/分钟
    2	    45w/分钟
    3	    30w/分钟
    4	    10w/分钟

```

详细Demo：[examples/example/custom.go](examples/example/mass.go)

## 菜单管理

- [菜单管理](#菜单管理)
	- [自定义菜单](#自定义菜单)
	    - 创建菜单
        - 获取菜单
        - 删除菜单
	- [个性化菜单](#个性化菜单)
		- 创建个性化菜单
		- 删除个性化菜单
		- 个性化菜单测试
	- [菜单配置](#菜单配置)
		- 获取自定义菜单配置



## 菜单管理

```go
自定义菜单能够帮助公众号丰富界面，让用户更好更快地理解公众号的功能。

注意：
1、自定义菜单最多包括3个一级菜单，每个一级菜单最多包含5个二级菜单。
2、一级菜单最多4个汉字，二级菜单最多7个汉字，多出来的部分将会以“...”代替。
3、创建自定义菜单后，菜单的刷新策略是，在用户进入公众号会话页或公众号profile页时，如果发现上一次拉取菜单的请求在5分钟以前，就会拉取一下菜单，如果菜单有更新，就会刷新客户端的菜单。测试时可以尝试取消关注公众账号后再次关注，则可以看到创建后的效果。

```

具体参数请参考微信文档：[菜单管理](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421141013)


## 自定义菜单

##### 创建菜单
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

	//二级菜单列表
	subMenuList1 := menu.SetButton(
		menu.WithClickButton("赞我们一下","V1001_GOOD"),    //不同的菜单类型，调用不用的menu.WithXXXButton()方法
		menu.WithViewButton("搜一下","http://www.soso.com/"),
		menu.WithLocationSelectButton("上报位置","wz2039_fdei"),
		menu.WithMiniprogramButton("跳转小程序","http://mp.weixin.qq.com","wx286b93c14bbf93aa","pages/lunar/index"),
	)
	subMenuList2 := menu.SetButton(
		menu.WithScanCodeWaitMsgButton("扫码带提示","rselfmenu_0_0"),
		menu.WithScanCodePushButton("扫码推事件","rselfmenu_0_1"),
	)

	//一级菜单列表
	parentMenu1 := menu.SetButton(
		menu.WithSubButton("菜单1",subMenuList1),
		menu.WithSubButton("菜单2",subMenuList2),
		menu.WithPicPhotoOrAlbumButton("菜单3","22222"),
	)


	//创建菜单
	res, err := mu.SetMenu(parentMenu1...)

```

##### 获取菜单
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

    #获取菜单
    res, err := mu.GetMenu()

```

##### 删除菜单
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

    #删除菜单 包含个性化菜单
    err := mu.DeleteMenu()

```

## 个性化菜单

##### 新增个性化菜单
```go

    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

	//二级菜单列表
	subMenuList1 := menu.SetButton(
		menu.WithClickButton("我是个性菜单","V1001_GOODdfd"),
		menu.WithLocationSelectButton("报告位置","wz2039_fdeidffd"),
	)
	//一级菜单列表
	parentMenu1 := menu.SetButton(
		menu.WithClickButton("今日歌曲","vmufddf"),
		menu.WithSubButton("菜单1",subMenuList1),
	)

	//匹配规则
	matchRule := menu.MatchRule{
		GroupID:"2",
	}

	//创建个性菜单
	res,err := mu.AddConditional(parentMenu1, &matchRule)

```

##### 删除个性化菜单

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

    //删除个性化菜单，参数menuid为菜单id
    err := mu.DeleteConditional(420025855)

```


##### 测试个性化菜单匹配结果

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

	#测试个性化菜单匹配结果  参数：user_id可以是粉丝的OpenID，也可以是粉丝的微信号
    res, err := mu.MenuTryMatch("pisdf0733")

```

## 菜单配置
##### 获取自定义菜单配置

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mu := wxconf.WechatClient.GetMenu()

	#获取自定义菜单配置
    res, err := mu.GetSelfMenuInfo()

```



详细Demo：[examples/example/menu.go](../../examples/example/menu.go)

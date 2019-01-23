## 用户管理

- [用户管理](#用户管理)
	- [用户标签管理](#用户标签管理)
	    - 创建标签
        - 获取公众号已创建的标签
        - 编辑标签
        - 删除标签
        - 获取标签下粉丝列表
        - 批量为用户打标签
        - 批量为用户取消标签
        - 获取用户身上的标签列表
	- [用户备注管理](#用户备注管理)
		- 设置用户备注名
    - [获取用户基本信息](#获取用户基本信息)
		- 获取用户基本信息(UnionID机制)
		- 批量获取用户基本信息
	- [获取用户列表](#获取用户列表)
		- 获取用户列表
	- [黑名单管理](#黑名单管理)
		- 获取公众号的黑名单列表
		- 拉黑用户
		- 取消拉黑用户

## 用户管理

```go
开发者可以使用用户标签管理的相关接口，实现对公众号的标签进行创建、查询、修改、删除等操作，也可以对用户进行打标签、取消标签等操作。


```

具体参数请参考微信文档：[用户管理](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140837)


## 用户标签管理

##### 用户标签
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	uTag := wx.GetUser()

	#创建标签
	re, err :=uTag.CreateTag("山西")
	fmt.Println("创建标签：", re, err)

	#获取公众号已创建的标签
	re, err := uTag.GetTag()
	fmt.Println("获取公众号已创建的标签：", re, err)

	#编辑标签
	err :=uTag.UpdateTag("山西2",100)
	fmt.Println("编辑标签：",  err)

	#删除标签
	err :=uTag.DeleteTag(100)
	fmt.Println("删除标签：",  err)

	#获取标签下粉丝列表
	re, err := uTag.GetTagsUser(2, "")
	fmt.Println("删除标签：",re,  err


	#批量为用户打标签
	err := uTag.BatchTagging(101, [] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232","abcd1234abcd1234abcd1234"})
	fmt.Println("批量为用户打标签：",  err)

	#批量为用户取消标签
	err := uTag.BatchUntagging(101, [] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232"})
	fmt.Println("批量为用户取消标签：",  err)

	#获取用户身上的标签列表
	re, err := uTag.Getidlist("abcd1234abcd1234abcd1234")
	fmt.Println("获取用户身上的标签列表：",  re, err)

```


## 用户备注管理

##### 设置用户备注名
```go

    //config配置文件省略...
	wc := wechat.NewWechat(config)
	uTag := wx.GetUser()

	//设置用户备注名
	re, err :=uTag.UpdateRemark("abcd1234abcd1234abcd1234","张三丰")
	fmt.Println("设置用户备注名：", re, err)

```

## 获取用户基本信息

##### 获取用户信息

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	uTag := wx.GetUser()

	//获取用户基本信息  返回对应结构体
	re, err :=uTag.GetUserInfo("abcd1234abcd1234abcd1234","zh_CN")
	fmt.Println("获取用户基本信息（包括UnionID机制）：", re, err)

	//字段描述
	{
        "subscribe": 1,
        "openid": "abcd1234abcd1234abcd1234",
        "nickname": "Band",
        "sex": 1,
        "language": "zh_CN",
        "city": "广州",
        "province": "广东",
        "country": "中国",
        "headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
        "subscribe_time": 1382694957,
        "unionid": " o6_bmasdasdsad6_2sgVt7hXXXXX"
        "remark": "",
        "groupid": 0,
        "tagid_list":[128,2],
        "subscribe_scene": "ADD_SCENE_QR_CODE",
        "qr_scene": 98765,
        "qr_scene_str": ""
    }

	//批量获取用户基本信息
	openidList := [] user.OpenIDs{
			{
				"abcd1234abcd1234abcd1231",
				"zh_CN",
			},
			{
				"abcd1234abcd1234abcd1232",
				"zh_CN",
			},
			{
				"abcd1234abcd1234abcd1234",
				"zh_CN",
			},
		}
	re, err :=uTag.GetUserInfoList(openidList)
	fmt.Println("批量获取用户基本信息：", re, err)


```

## 获取用户列表

##### 用户列表

```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	uTag := wx.GetUser()

	//获取用户列表
	re, err :=uTag.GetUserList("")
	fmt.Println("获取用户列表：", re, err)

```

## 黑名单管理
##### 黑名单
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	uTag := wx.GetUser()

    //获取公众号的黑名单列表
	re, err :=uTag.GetBlacklist("")
	fmt.Println("获取黑名单列表：", re, err)

	//拉黑用户
	//err :=uTag.BatchBlacklist([] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232","abcd1234abcd1234abcd1234"})
	//fmt.Println("拉黑用户：", err)

	//取消拉黑用户
	//err :=uTag.BatchUnblacklist([] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232","abcd1234abcd1234abcd1234"})
	//fmt.Println("取消拉黑用户：", err)
```

详细Demo：[examples/example/user.go](examples/example/user.go)

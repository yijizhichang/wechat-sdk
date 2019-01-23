## 微信网页授权

- [微信网页授权](#微信网页授权)
	- [网页授权流程](#网页授权流程)
	    - 第一步：用户同意授权，获取code
        - 第二步：通过code换取网页授权access_token
        - 第三步：刷新access_token（如果需要）
        - 第四步：拉取用户信息(需scope为 snsapi_userinfo)
        - 附：检验授权凭证（access_token）是否有效



## 微信网页授权

```go
如果用户在微信客户端中访问第三方网页，公众号可以通过微信网页授权机制，来获取用户基本信息，进而实现业务逻辑。


```

具体参数请参考微信文档：[微信网页授权](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842)

## 网页授权流程

```go

    //config配置文件省略...
	wc := wechat.NewWechat(config)
	oauth := wc.GetOauth()

    //第一步：用户同意授权，获取code  引导关注者打开如下页面
	url,err := oauth.GetRedirectURL(redirectUri, scope, state)

    //如果用户同意授权，页面将跳转至 redirect_uri/?code=CODE&state=STATE。

    //第二步：通过code换取网页授权access_token
    access_token, err := oauth.GetGrantAccessToken(code)
    
    //第三步：刷新access_token（如果需要）
    access_token, err := oauth.RefreshAccessToken(refresh_token)

    //第四步：拉取用户信息(需scope为 snsapi_userinfo)
    re, err := oauth.GetUserInfo(accessToken, openId, lang)   #zh_CN 简体，zh_TW 繁体，en 英语

    //附：检验授权凭证（access_token）是否有效
    b, err := oauth.CheckAccessToken(accessToken, openId)

```
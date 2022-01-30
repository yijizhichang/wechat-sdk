## 身份验证管理

- [身份验证](#身份验证)
    - [网页授权登录](#网页授权登录)
        - 构造网页授权链接
        - 获取访问用户身份
    - [扫码授权登录](#扫码授权登录)
        - 构造扫码登录链接
        - 获取访问用户身份


## 身份验证

```go
企业微信提供了OAuth的授权登录方式，可以让从企业微信终端打开的网页获取成员的身份信息，从而免去登录的环节。
企业应用中的URL链接（包括自定义菜单或者消息中的链接），均可通过OAuth2.0验证接口来获取成员的UserId身份信息。

```

具体参数请参考微信文档：[身份验证](https://developer.work.weixin.qq.com/document/path/91335)

## 网页授权登录
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetOauth2()
    
    //构造网页授权链接
    re,err := cli.GetQyOauth2AuthorizeURL(redirectUri, state)
    //获取访问用户身份
    re,err := cli.GetUserInfoByCode(accessToken, code)
    //构造独立窗口登录二维码
    re,err := cli.GetQySsoQrConnect(agentId, redirectUri, state)
    //获取企业的jsapi_ticket
    re,err := cli.GetJsapiTicket(token)
    //获取应用的jsapi_ticket
	re,err := cli.GetAgentJsapiTicket(token, agentId)
    //code2Session
    re,err := cli.GetCode2Session(token, jsCode, grantType)
```

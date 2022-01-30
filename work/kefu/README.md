## 微信客服

- [微信客服管理](#微信客服管理)
    - [客服账号管理](#客服账号管理)
        - 添加客服帐号
        - 删除客服帐号
        - 修改客服帐号
        - 获取客服帐号列表
        - 获取客服帐号链接
    - [接待人员管理](#接待人员管理)
      - 添加接待人员
      - 删除接待人员
      - 获取接待人员列表
    - [会话分配与消息收发](#会话分配与消息收发)
        - 分配客服会话
        - 接收消息和事件
        - 发送消息
        - 发送欢迎语等事件响应消息
        
    - [升级服务配置](#升级服务配置)
        - 
    - [其他基础信息获取](#其他基础信息获取)
        - 获取客户基础信息
    
## 微信客服

```go
什么是微信客服

微信客服由腾讯微信团队为企业打造，用于满足企业的客服需求，帮助企业做好客户服务。企业可以在微信内、外各个场景中接入微信客服，用户可以发起咨询，企业可以进行回复。
企业可在微信客服官网使用企业微信扫码开通微信客服，开通后即可使用。

企业在使用微信客服时，有两种选择。可选择在微信客服管理后台处独立使用微信客服，或选择由企业微信接管微信客服帐号和收发消息。选择由企业微信接管微信客服时，可实现：

可将企业员工配置为微信客服的接待人员，在企业微信里接收和回复用户在微信内、外发起的咨询消息
可在企业微信里通过API来管理微信客服帐号、分配客服会话和收发客服消息等。
可使用客服工具栏、「升级服务」等工具

```

具体参数请参考微信文档：[微信客服](https://developer.work.weixin.qq.com/document/path/94638)

## 客服账号管理
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetKefuAccount()

    //添加客服帐号
    re,err := cli.CreateKfAccount(token, req)
    //删除客服帐号
    re,err := cli.DelKfAccount(token, req)
    //修改客服帐号
    re,err := cli.UpdateKfAccount(token, req)
    //获取客服帐号列表
    re,err := cli.GetKfAccountList(token, req)
    //获取客服帐号链接
    re,err := cli.GetKfContactWay(token, req)

```

## 接待人员管理
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetKefuServicer()
    
    //添加接待人员
    re,err := cli.CreateKfServicer(token, req)
    //删除接待人员
    re,err := cli.DelKfServicer(token, req)
    //获取接待人员列表
    re,err := cli.GetKfServicerList(token, req)

```

## 会话分配与消息收发

```go
	//config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetKefuConverse()

    //获取会话状态
    re,err := cli.GetKfConverseState(token, req)
    //变更会话状态
    re,err := cli.UpdateKfConverseState(token, req)
    //读取消息
    re,err := cli.SyncKfConverseMsg(token, req)
    //发送消息
    re,err := cli.SendKfConverseMsg(token, req)
    
```

## 升级服务配置

```go
   //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetKefu()

    //获取配置的专员与客户群
    re,err := cli.GetKfCustomerUpgradeServiceConfig(token, req)
    //升级专员服务
    re,err := cli.UpgradeKfCustomerService(token, req)
    //为客户取消推荐
    re,err := cli.CancelKfCustomerService(token, req)
```

## 其他基础信息获取

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetKefu()
    
    //获取客户基础信息
    re,err := cli.GetKfCustomerList(token, req)
```

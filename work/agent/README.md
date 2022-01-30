## 企微应用管理

- [应用管理](#应用管理)
    - [应用设置](#应用设置)
        - 获取指定的应用详情
        - 获取access_token对应的应用列表
        - 设置应用
    - [自定义菜单](#自定义菜单)
        - 创建菜单
        - 获取菜单
        - 删除菜单
    - [设置工作台自定义展示](#设置工作台自定义展示)
        - 设置应用在工作台展示的模版
        - 获取应用在工作台展示的模版
        - 设置应用在用户工作台展示的数据


## 应用管理

```go
企业微信的应用具备以下能力：
收发消息
以普通会话的形式存在，可推送消息、接收消息，也可以设置自定义菜单。

通过“消息推送”API，推送应用消息给企业成员
开启“接收消息”，员工进入应用、发送消息、操作菜单等动作会以事件的方式转发给企业的应用后台
自定义8种个性化菜单

应用主页
应用主页可以配置为一个H5网页或者小程序，配置以后就会出现在工作台，点击以后直接进入H5网页或者小程序。

Oauth2用户身份识别。
JS-SDK调用原生客户端Native的能力。

授权登录
企业已有的Web网页、移动APP，可以使用企业微信的帐号登录。

扫一扫授权登录Web网页
嵌入登录SDK，一键登录移动APP
```

具体参数请参考微信文档：[应用管理](https://developer.work.weixin.qq.com/document/path/90226)

## 应用设置
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetAgent()

    //获取指定的应用详情
    re,err := cli.GetQyAgentView(token,req)
    //获取access_token对应的应用列表
    re,err := cli.GetQyAgentList(token,req)
    //设置应用
    re,err := cli.SetQyAgent(token,req)
```

## 自定义菜单
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetAgentMenu()

    //创建菜单
    re,err := cli.CreateQyMenu(token,req)
    //获取菜单
    re,err := cli.GetQyMenu(token,req)
    //删除菜单
    re,err := cli.DelQyMenu(token,req)

```

## 设置工作台自定义展示
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetAgentWorkbench()

    //设置应用在工作台展示的模版
    re,err := cli.SetWorkbenchTemplate(token,req)
    //获取应用在工作台展示的模版
    re,err := cli.GetWorkbenchTemplate(token,req)
    //设置应用在用户工作台展示的数据
    re,err := cli.SetWorkbenchData(token,req)

```






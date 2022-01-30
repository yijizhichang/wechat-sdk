## 消息推送管理

- [消息推送](#消息推送)
    - [应用消息](#应用消息)
        - 发送应用消息
        - 更新模版卡片消息
        - 撤回应用消息
        - 查询应用消息发送统计
    - [群聊会话](#群聊会话)
        - 创建群聊会话
        - 修改群聊会话
        - 获取群聊会话
        - 群聊会话应用推送消息
    - [互联企业消息推送](#互联企业消息推送)
        - 发送应用消息


## 消息推送

```go
企业微信开放了消息发送接口，企业可以使用这些接口让自定义应用与企业微信后台或用户间进行双向通信。
```

具体参数请参考微信文档：[应用消息](https://developer.work.weixin.qq.com/document/path/90235)

## 应用消息
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetMessage()

    //发送应用消息
    re,err := cli.SendTextMessage(token,req)
    re,err := cli.SendImageMessage(token,req)
    re,err := cli.SendVoiceMessage(token,req)
    re,err := cli.SendVideoMessage(token,req)
    re,err := cli.SendFileMessage(token,req)
    re,err := cli.SendTextCardMessage(token,req)
    re,err := cli.SendNewsMessage(token,req)
    re,err := cli.SendMpNewsMessage(token,req)
    re,err := cli.SendMarkdownMessage(token,req)
    re,err := cli.SendMiniprogramMessage(token,req)
    re,err := cli.SendTemplateCardTextMessage(token,req)
    re,err := cli.SendTemplateCardNewsMessage(token,req)
    re,err := cli.SendTemplateCardButtonMessage(token,req)
    re,err := cli.SendTemplateCardVoteMessage(token,req)
    re,err := cli.SendTemplateCardMultipleMessage(token,req)
    //更新模版卡片消息
    re,err := cli.UpdateTemplateCardButton(token,req)
    re,err := cli.UpdateTemplateCardTextMessage(token,req)
    re,err := cli.UpdateTemplateCardNewsMessage(token,req)
    re,err := cli.UpdateTemplateCardButtonInteractionMessage(token,req)
    re,err := cli.UpdateTemplateCardVoteMessage(token,req)
    re,err := cli.UpdateTemplateCardMultipleMessage(token,req)
    //撤回应用消息
    re,err := cli.QyRecallMessage(token,req)
    //查询应用消息发送统计
    re,err := cli.GetQyMessageStatistics(token,req)

```

## 群聊会话
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetMessageGroup()
    
    //创建群聊会话
    re,err := cli.CreateQyAppChat(token,req)
    //修改群聊会话
    re,err := cli.UpdateQyAppChat(token,req)
    //获取群聊会话
    re,err := cli.GetQyAppChat(token,req)
    //群聊会话应用推送消息
    re,err := cli.SendTextQyAppChat(token,req)

```

## 互联企业消息推送
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetMessageLinkedCorp()

    //发送应用消息
    re,err := cli.SendTextQyLinkedCorp(token,req)

```






## 客户联系

- [客户联系管理](#客户联系管理)
    - [企业服务人员管理](#企业服务人员管理)
        - 获取配置了客户联系功能的成员列表
        - 客户联系「联系我」管理
    - [客户管理](#客户管理)
      - 获取客户列表
      - 获取客户详情
      - 批量获取客户详情
      - 修改客户备注信息
      - 客户联系规则组管理
    - [客户标签管理](#客户标签管理)
        - 获取企业标签库
        - 添加企业客户标签
        - 编辑企业客户标签
        - 删除企业客户标签
        - 获取指定规则组下的企业客户标签
        - 为指定规则组创建企业客户标签
        - 编辑指定规则组下的企业客户标签
        - 删除指定规则组下的企业客户标签
        - 编辑客户企业标签
    - [在职继承](#在职继承)
        - 分配在职成员的客户
        - 查询客户接替状态
    - [离职继承](#离职继承)
        - 获取待分配的离职成员列表
        - 分配离职成员的客户
        - 查询客户接替状态
        - 分配离职成员的客户群
    - [客户群管理](#客户群管理)
        - 获取客户群列表
        - 获取客户群详情
        - 客户群opengid转换
    - [联系我与客户入群方式](#联系我与客户入群方式)
        - 客户联系「联系我」管理
        - 客户群「加入群聊」管理
    - [客户朋友圈](#客户朋友圈)
        - 企业发表内容到客户的朋友圈
        - 获取客户朋友圈全部的发表记录
        - 客户朋友圈规则组管理
    - [消息推送](#消息推送)
        - 创建企业群发
        - 获取企业的全部群发记录
        - 发送新客户欢迎语
        - 入群欢迎语素材管理
    - [统计管理](#统计管理)  
        - 获取「联系客户统计」数据
        - 获取「群聊数据统计」数据
    - [变更回调](#变更回调)
        - 事件格式

## 客户联系

```go
企业内的员工可以添加外部企业的联系人进行工作沟通，外部联系人分为企业微信联系人和微信联系人两种类型。
配置了客户联系功能的成员所添加的外部联系人为企业客户。

```

具体参数请参考微信文档：[客户联系](https://developer.work.weixin.qq.com/document/path/92109)

## 企业服务人员管理
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerFollow()

    //获取配置了客户联系功能的成员列表
    re,err := cli.GetCustomerFollowUserList(token)
    //配置客户联系「联系我」方式
    re,err := cli.CreateCustomerContactWay(token, req)
    //获取企业已配置的「联系我」方式
    re,err := cli.GetCustomerContactWay(token, req)
    //更新企业已配置的「联系我」方式
    re,err := cli.UpdateCustomerContactWay(token, req)
    //删除企业已配置的「联系我」方式
    re,err := cli.DelCustomerContactWay(token, req)
    //结束临时会话
    re,err := cli.CloseCustomerTempChat(token, req)

```

## 客户管理
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomer()
        
    //获取客户列表
    re,err := cli.GetQyCustomerList(token, req)
    //获取客户详情
    re,err := cli.GetQyCustomerView(token, req)
    //批量获取客户详情
    re,err := cli.GetQyCustomerViewBatch(token, req)
    //修改客户备注信息
    re,err := cli.UpdateQyCustomerRemark(token, req)
    //获取规则组列表
    re,err := cli.GetQyCustomerStrategyList(token, req)
    //获取规则组详情
    re,err := cli.GetQyCustomerStrategyView(token, req)
    //获取规则组管理范围
    re,err := cli.GetQyCustomerStrategyRange(token, req)
    //创建新的规则组
    re,err := cli.CreateQyCustomerStrategy(token, req)
    //编辑规则组及其管理范围
    re,err := cli.UpdateQyCustomerStrategy(token, req)
    //删除规则组
    re,err := cli.DelQyCustomerStrategy(token, req)

```

## 客户标签管理

```go
	//config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerTag()

    //获取企业标签库
    re,err := cli.GetCustomerTagList(token, req)
    //添加企业客户标签
    re,err := cli.CreateCustomerTag(token, req)
    //编辑企业客户标签
    re,err := cli.UpdateCustomerTag(token, req)
    //删除企业客户标签
    re,err := cli.DelCustomerTag(token, req)
    //编辑客户企业标签
    re,err := cli.MarkCustomerTag(token, req)
    //获取指定规则组下的企业客户标签
    re,err := cli.GetCustomerStrategyTagList(token, req)
    //为指定规则组创建企业客户标签
    re,err := cli.CreateCustomerStrategyTag(token, req)
    //编辑指定规则组下的企业客户标签
    re,err := cli.UpdateCustomerStrategyTag(token, req)
    //删除指定规则组下的企业客户标签
    re,err := cli.DelCustomerStrategyTag(token, req)
```

## 在职继承

```go
   //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerOnTransfer()

    //分配在职成员的客户
    re,err := cli.CreateCustomerOnTransfer(token, req)
    //查询客户接替状态
    re,err := cli.GetCustomerOnTransfer(token, req)
```

## 离职继承

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerOffTransfer()
    
    //获取待分配的离职成员列表
    re,err := cli.GetCustomerOffUnassignedList(token, req)
    //分配离职成员的客户
    re,err := cli.CreateCustomerOffTransfer(token, req)
    //查询客户接替状态
    re,err := cli.GetCustomerOffTransferResult(token, req)
    //分配离职成员的客户群
    re,err := cli.CreateCustomerOffGroupChatTransfer(token, req)
```

## 联系我与客户入群方式

```go

```

## 客户朋友圈

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerMoment()

    //获取企业全部的发表列表
    re,err := cli.GetQyMomentList(token, req)
    //获取客户朋友圈企业发表的列表
    re,err := cli.GetQyMomentTask(token, req)
    //获取客户朋友圈发表时选择的可见范围
    re,err := cli.GetQyMomentCustomerList(token, req)
    //获取客户朋友圈发表后的可见客户列表
    re,err := cli.GetQyMomentSendResult(token, req)
    //获取客户朋友圈的互动数据
    re,err := cli.GetQyMomentComments(token, req)
    //获取规则组列表
    re,err := cli.GetQyMomentStrategyList(token, req)
    //获取规则组详情
    re,err := cli.GetQyMomentStrategyView(token, req)
    //获取规则组管理范围
    re,err := cli.GetQyMomentStrategyRange(token, req)
    //创建新的规则组
    re,err := cli.CreateQyMomentStrategy(token, req)
    //编辑规则组及其管理范围
    re,err := cli.UpdateQyMomentStrategy(token, req)
    //删除规则组
    re,err := cli.DelQyMomentStrategy(token, req)
```

## 消息推送

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerMsg()
    
    //创建企业群发
    re,err := cli.CreateCustomerMsgTemplate(token, req)
    //获取群发记录列表
    re,err := cli.GetCustomerGroupMsgList(token, req)
    //获取群发成员发送任务列表
    re,err := cli.GetCustomerGroupMsgTask(token, req)
    //获取企业群发成员执行结果
    re,err := cli.GetQyGroupMsgSendResult(token, req)
    //发送新客户欢迎语
    re,err := cli.SendWelcomeMsg(token, req)
    //添加入群欢迎语素材
    re,err := cli.CreateGroupWelcomeTemplate(token, req)
    //编辑入群欢迎语素材
    re,err := cli.UpdateGroupWelcomeTemplate(token, req)
    //获取入群欢迎语素材
    re,err := cli.GetGroupWelcomeTemplate(token, req)
    //删除入群欢迎语素材
    re,err := cli.DelGroupWelcomeTemplate(token, req)
```

## 统计管理

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCustomerData()

    //获取「联系客户统计」数据
    re,err := cli.GetCustomerBehaviorData(token, req)
    //获取「群聊数据统计」数据 按群主聚合的方式
    re,err := cli.GetCustomerGroupChatData(token, req)
    //获取「群聊数据统计」数据 按自然日聚合的方式
    re,err := cli.GetCustomerGroupChatByDayData(token, req)

```

## 变更回调

```go
    具体参考微信消息回调
```
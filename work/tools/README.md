## 效率工具管理

- [效率工具](#效率工具)
    - [日程](#日程)
        - 创建日历
        - 更新日历
        - 获取日历详情
        - 删除日历
        - 创建日程
        - 更新日程
        - 获取日程详情
        - 取消日程
        - 获取日历下的日程列表
        - 新增日程参与者
        - 删除日程参与者
    - [直播](#直播)
        - 创建预约直播
        - 修改预约直播
        - 取消预约直播
        - 删除直播回放
        - 获取微信观看直播凭证
        - 获取成员直播ID列表
        - 获取直播详情
        - 获取直播观看明细
        - 获取跳转小程序商城的直播观众信息

## 效率工具

```go
日程：企业通过日程相关接口，可以很方便的将企业已有系统的会议、日程或提醒，同步到企业微信日历本
直播：企业和开发者通过直播接口，可以便捷地创建直播、管理直播、获取观看入口及查询直播间明细和统计信息。

```

具体参数请参考微信文档：[效率工具](https://developer.work.weixin.qq.com/document/path/93624)

## 日程
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetCalendar()

    //创建日历
    re,err := cli.CreateCalendar(token,req)
    //更新日历
    re,err := cli.UpdateCalendar(token,req)
    //获取日历详情
    re,err := cli.GetCalendar(token,req)
    //删除日历
    re,err := cli.DelCalendar(token,req)
    //创建日程
    re,err := cli.CreateSchedule(token,req)
    //更新日程
    re,err := cli.UpdateSchedule(token,req)
    //获取日程详情
    re,err := cli.GetSchedule(token,req)
    //取消日程
    re,err := cli.DelSchedule(token,req)
    //获取日历下的日程列表
    re,err := cli.ListSchedule(token,req)
    //新增日程参与者
    re,err := cli.AddAttendeesSchedule(token,req)
    //删除日程参与者
    re,err := cli.DelAttendeesSchedule(token,req)
```

## 直播
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetLiving()
    
    //创建预约直播
    re,err := cli.CreateLiving(token,req)
    //修改预约直播
    re,err := cli.UpdateLiving(token,req)
    //取消预约直播
    re,err := cli.CancelLiving(token,req)
    //删除直播回放
    re,err := cli.DeleteReplayLiving(token,req)
    //获取微信观看直播凭证
    re,err := cli.GetLivingCode(token,req)
    //获取成员直播ID列表
    re,err := cli.GetUserAllLivingid(token,req)
    //获取直播详情
    re,err := cli.GetLivingInfo(token,req)
    //获取直播观看明细
    re,err := cli.GetWatchStat(token,req)
    //获取跳转小程序商城的直播观众信息
    re,err := cli.GetLivingShareInfo(token,req)

```

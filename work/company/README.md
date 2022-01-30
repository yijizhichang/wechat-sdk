## 通讯录管理

- [通讯录管理](#通讯录管理)
    - [成员管理](#成员管理)
        - 创建成员
        - 读取成员
        - 更新成员
        - 删除成员
        - 批量删除成员
        - 获取部门成员
        - 获取部门成员详情
        - userid与openid互换
        - 二次验证
        - 邀请成员
        - 获取加入企业二维码
        - 获取企业活跃成员数
        - 手机号获取userid
    - [部门管理](#部门管理)
        - 创建部门
        - 更新部门
        - 删除部门
        - 获取部门列表
        - 获取子部门ID列表
        - 获取单个部门详情
    - [标签管理](#标签管理)
        - 创建标签
        - 更新标签名字
        - 删除标签
        - 获取标签成员
        - 增加标签成员
        - 删除标签成员
        - 获取标签列表
    - [异步批量接口](#异步批量接口)
        - 增量更新成员
        - 全量覆盖成员
        - 全量覆盖部门
        - 获取异步任务结果
    - [通讯录回调通知](#通讯录回调通知)
        - 成员变更通知
        - 部门变更通知
        - 标签变更通知
        - 异步任务完成通知
    - [互联企业](#互联企业)
        - 获取应用的可见范围
        - 获取互联企业成员详细信息
        - 获取互联企业部门成员
        - 获取互联企业部门成员详情
        - 获取互联企业部门列表
    - [异步导出接口](#异步导出接口)
        - 导出成员
        - 导出成员详情
        - 导出部门
        - 导出标签成员
        - 获取导出结果

## 通讯录管理

```go
通讯录同步相关接口，可以对部门、成员、标签等通讯录信息进行查询、添加、修改、删除等操作。
使用通讯录管理接口，原则上需要使用 通讯录管理secret，也可以使用 应用secret。
但是使用应用secret只能进行“查询”、“邀请”等非写操作，而且只能操作应用可见范围内的通讯录。

```

具体参数请参考微信文档：[通讯录管理](https://developer.work.weixin.qq.com/document/path/90193)

## 成员管理
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
	du := qw.GetDepartmentUser()

    //创建成员
    re,err := du.CreateUser(token, req)
    //读取成员
    re,err := du.GetUser(token, userid)
    //更新成员
    re,err := du.UpdateUser(token, req)
    //删除成员
    re,err := du.DelUser(token, userid)
    //批量删除成员
    re,err := du.DelUserBatch(token, req)
    //获取部门成员
	re,err := du.GetDepartmentSimpleUserList(token, departmentId, fetchChild)
    //获取部门成员详情
    re,err := du.GetDepartmentUserList(token, departmentId, fetchChild)
    //userid与openid互换
    re,err := du.ConvertToOpenid(token, req)
    //二次验证
	re,err := du.AuthSucc(token, req)
    //邀请成员
    re,err := du.BatchInvite(token, req)
    //获取加入企业二维码
    re,err := du.GetJoinCorpQrcode(token, sizeType)
    //获取企业活跃成员数
    re,err := du.GetActiveStat(token, req)
    //手机号获取userid
    re,err := du.GetUseridByMobile(token, req)

```

## 部门管理
```go

    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    dp := qw.GetDepartment()

	//创建部门
	re,err := dp.CreateDepartment(token, req)
	//更新部门
	re,err := dp.UpdateDepartment(token, req)
	//删除部门
	re,err := dp.DelDepartment(token, req)
	//获取部门列表
	re,err := dp.GetDepartmentList(token, req)
	//获取子部门ID列表
	re,err := dp.GetDepartmentSimpleList(token, req)
	//获取单个部门详情
	re,err := dp.GetDepartment(token, req)

```

## 标签管理

```go
	//config配置文件省略...
    qw := wechat.NewQyWechat(config)
    ctag := qw.GetCompanyTag()

    //创建标签
    re,err := ctag.CreateCompanyTag(token, req)
    //更新标签名字
    re,err := ctag.UpdateCompanyTag(token, req)
    //删除标签
    re,err := ctag.DelCompanyTag(token, req)
    //获取标签成员
    re,err := ctag.GetCompanyTagUser(token, req)
    //增加标签成员
    re,err := ctag.CreateCompanyTagUser(token, req)
    //删除标签成员
	re,err := ctag.DelCompanyTagUser(token, req)
    //获取标签列表
    re,err := ctag.GetCompanyTag(token)
```

## 异步批量接口

```go
   //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cb := qw.GetCompanyBatch()
    
    //增量更新成员
    re,err := cb.CreateBatchSyncUser(token, req)
    //全量覆盖成员
    re,err := cb.CreateBatchReplaceUser(token, req)
    //全量覆盖部门
    re,err := cb.CreateBatchReplaceParty(token, req)
    //获取异步任务结果
    re,err := cb.GetBatchResult(token, req)
    
```

## 通讯录回调通知

```go
   具体参考消息回调通知模块
```

## 互联企业

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    clc := qw.GetCompanyLinkedCorp()
    //获取应用的可见范围
    re,err := clc.GetLinkedCorpPermList(token, req)
    //获取互联企业成员详细信息
    re,err := clc.GetLinkedCorpUserInfo(token, req)
    //获取互联企业部门成员
    re,err := clc.GetLinkedCorpUserSimpleList(token, req)
    //获取互联企业部门成员详情
    re,err := clc.GetLinkedCorpUserList(token, req)
    //获取互联企业部门列表
    re,err := clc.GetLinkedCorpDepartmentList(token, req)
```

## 异步导出接口

```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    ce := qw.GetCompanyExport()

    //导出成员
    re,err := ce.QyExportSimpleUser(token, req)
    //导出成员详情
    re,err := ce.QyExportUser(token, req)
    //导出部门
    re,err := ce.QyExportDepartment(token, req)
    //导出标签成员
    re,err := ce.QyExportTagUser(token, req)
    //获取导出结果
    re,err := ce.QyExportResult(token, req)
```
        


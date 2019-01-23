package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/mp/user"
)

//用户标签管理
func UserTagManage() {

	uTag := wxconf.WechatClient.GetUser()

	//创建标签
	/*	re, err :=uTag.CreateTag("山西")
		fmt.Println("创建标签：", re, err)*/

	//获取公众号已创建的标签
	re, err := uTag.GetTag()
	fmt.Println("获取公众号已创建的标签：", re, err)

	//编辑标签
	//err :=uTag.UpdateTag("山西2",100)
	//fmt.Println("编辑标签：",  err)

	//删除标签
	//err :=uTag.DeleteTag(100)
	//fmt.Println("删除标签：",  err)

	//获取标签下粉丝列表
	//re, err := uTag.GetTagsUser(2, "")
	//fmt.Println("删除标签：",re,  err

	//批量为用户打标签
	//err := uTag.BatchTagging(101, [] string {"oEYzpw65XwZxFTHrsgo1RsUJGpsw","oEYzpw6f0t3OAHevr_EQpJ802msQ","abcd1234abcd1234abcd1234"})
	//fmt.Println("批量为用户打标签：",  err)

	//批量为用户取消标签
	//err := uTag.BatchUntagging(101, [] string {"oEYzpw65XwZxFTHrsgo1RsUJGpsw","oEYzpw6f0t3OAHevr_EQpJ802msQ"})
	//fmt.Println("批量为用户取消标签：",  err)

	//获取用户身上的标签列表
	//re, err := uTag.Getidlist("abcd1234abcd1234abcd1234")
	//fmt.Println("获取用户身上的标签列表：",  re, err)
}

//用户备注管理
func UserRemark() {
	uTag := wxconf.WechatClient.GetUser()

	//设置用户备注名
	re, err := uTag.UpdateRemark("abcd1234abcd1234abcd1234", "张三丰")
	fmt.Println("设置用户备注名：", re, err)
}

//获取用户基本信息（包括UnionID机制）
func UserInfo() {
	uTag := wxconf.WechatClient.GetUser()

	//获取用户基本信息
	//re, err :=uTag.GetUserInfo("abcd1234abcd1234abcd1234","zh_CN")
	//fmt.Println("获取用户基本信息（包括UnionID机制）：", re, err)

	//批量获取用户基本信息
	openidList := []user.OpenIDs{
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
	re, err := uTag.GetUserInfoList(openidList)
	fmt.Println("批量获取用户基本信息：", re, err)
}

//获取用户列表
func UserList() {
	uTag := wxconf.WechatClient.GetUser()

	//获取用户列表
	re, err := uTag.GetUserList("")
	fmt.Println("获取用户列表：", re, err)
}

//黑名单管理
func UserBlack() {
	uTag := wxconf.WechatClient.GetUser()

	//获取公众号的黑名单列表
	re, err := uTag.GetBlacklist("")
	fmt.Println("获取用户列表：", re, err)

	//拉黑用户
	//err :=uTag.BatchBlacklist([] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232","abcd1234abcd1234abcd1234"})
	//fmt.Println("拉黑用户：", err)

	//取消拉黑用户
	//err :=uTag.BatchUnblacklist([] string {"abcd1234abcd1234abcd1231","abcd1234abcd1234abcd1232","abcd1234abcd1234abcd1234"})
	//fmt.Println("取消拉黑用户：", err)
}

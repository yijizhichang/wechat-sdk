//客户联系-客户群管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCustomerGroupListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token=%s"  //获取客户群列表
	GetCustomerGroupViewURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token=%s"  //获取客户群详情
	GetOpengidToChatidURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/opengid_to_chatid?access_token=%s"  //客户群opengid转换
)

//CustomerGroup 客户管理
type CustomerGroup struct {
	*core.Context
}

//NewCustomerGroup 实例化
func NewCustomerGroup(context *core.Context) *CustomerGroup {
	cg := new(CustomerGroup)
	cg.Context = context
	return cg
}

//获取客户群列表
type CusGroupReq struct {
	StatusFilter int32 `json:"status_filter"`
	OwnerFilter struct{
		UseridList []string `json:"userid_list"`
	} `json:"owner_filter"`
	Cursor string `json:"cursor"`
	Limit int32 `json:"limit"`
}
type cusGroupList struct {
	util.WxError
	GroupChatList []groupChatItem `json:"group_chat_list"`
	NextCursor string `json:"next_cursor"`
}
type groupChatItem struct {
	ChatId string `json:"chat_id"`
	Status int32 `json:"status"`
}
func (cg *CustomerGroup) GetCustomerGroupList(accessToken string, req CusGroupReq)(result *cusGroupList, err error){
	qyUrl := fmt.Sprintf(GetCustomerGroupListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cg.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户群详情
type CusGroupViewReq struct {
	ChatId string `json:"chat_id"`
	NeedName int32 `json:"need_name"`
}
type cusGroupView struct {
	util.WxError
	GroupChat struct{
		ChatId string `json:"chat_id"`
		Name string `json:"name"`
		Owner string `json:"owner"`
		CreateTime int64 `json:"create_time"`
		Notice string `json:"notice"`
	} `json:"group_chat"`
	MemberList []memberItem `json:"member_list"`
	AdminList []adminItem `json:"admin_list"`
}
type memberItem struct {
	Userid string `json:"userid"`
	Type int32 `json:"type"`
	JoinTime int64 `json:"join_time"`
	JoinScene int32 `json:"join_scene"`
	Invitor struct{
		Userid string `json:"userid"`
	} `json:"invitor"`
	GroupNickname string `json:"group_nickname"`
	Name string `json:"name"`
	Unionid string `json:"unionid"`
}
type adminItem struct {
	Userid string `json:"userid"`
}
func (cg *CustomerGroup) GetCustomerGroupView(accessToken string, req CusGroupViewReq)(result *cusGroupView, err error){
	qyUrl := fmt.Sprintf(GetCustomerGroupViewURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cg.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//客户群opengid转换
type CusGroupOpengidReq struct {
	Opengid string `json:"opengid"`  //小程序在微信获取到的群ID
}
type cusGroupChatId struct {
	util.WxError
	ChatId string `json:"chat_id"`  //客户群ID，可以用来调用获取客户群详情
}
func (cg *CustomerGroup) GetCustomerGroupChatId(accessToken string, req CusGroupOpengidReq)(result *cusGroupChatId, err error){
	qyUrl := fmt.Sprintf(GetOpengidToChatidURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cg.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupChatId error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

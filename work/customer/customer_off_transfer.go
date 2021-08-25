//客户联系-离职继承
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCustomerOffUnassignedListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list?access_token=%s"  //获取待分配的离职成员列表
	CreateCustomerOffTransferURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_customer?access_token=%s"  //分配离职成员的客户
	GetCustomerOffTransferResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/resigned/transfer_result?access_token=%s"  //查询客户接替状态
	CreateCustomerOffGroupChatTransferURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/transfer?access_token=%s"  //分配离职成员的客户群

)

//CustomerOffTransfer 消息推送
type CustomerOffTransfer struct {
	*core.Context
}

//NewCustomerMsg 实例化
func NewCustomerOffTransfer(context *core.Context) *CustomerOffTransfer {
	cft := new(CustomerOffTransfer)
	cft.Context = context
	return cft
}
//获取待分配的离职成员列表
type CustomerOffUnassignedReq struct {
	PageId   int32    `json:"page_id"`
	Cursor   string   `json:"cursor"`
	PageSize int32    `json:"page_size"`
}
type customerOffUnassignedList struct {
	util.WxError
	Info    []struct {
		HandoverUserid string `json:"handover_userid"`
		ExternalUserid string `json:"external_userid"`
		DimissionTime  int32  `json:"dimission_time"`
	} `json:"info"`
	IsLast     bool   `json:"is_last"`
	NextCursor string `json:"next_cursor"`
}
func (cft *CustomerOffTransfer) GetCustomerOffUnassignedList(accessToken string, req CustomerOffUnassignedReq)(result *customerOffUnassignedList, err error){
	qyUrl := fmt.Sprintf(GetCustomerOffUnassignedListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cft.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerOffUnassignedList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//分配离职成员的客户
type CreateCustomerOffTransferReq struct {
	HandoverUserid string   `json:"handover_userid"`
	TakeoverUserid string   `json:"takeover_userid"`
	ExternalUserid []string `json:"external_userid"`
}
type createCustomerOffTransferRep struct {
	util.WxError
	Customer []struct {
		ExternalUserid string `json:"external_userid"`
		Errcode        int32    `json:"errcode"`
	} `json:"customer"`
}
func (cft *CustomerOffTransfer) CreateCustomerOffTransfer(accessToken string, req CreateCustomerOffTransferReq)(result *createCustomerOffTransferRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerOffTransferURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cft.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerOffTransfer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//查询客户接替状态
type CustomerOffTransferResultReq struct {
	HandoverUserid string `json:"handover_userid"`
	TakeoverUserid string `json:"takeover_userid"`
	Cursor         string `json:"cursor"`
}
type customerOffTransferResultRep struct {
	util.WxError
	Customer []struct {
		ExternalUserid string `json:"external_userid"`
		Status         int32    `json:"status"`
		TakeoverTime   int64    `json:"takeover_time"`
	} `json:"customer"`
	NextCursor string `json:"next_cursor"`
}
func (cft *CustomerOffTransfer) GetCustomerOffTransferResult(accessToken string, req CustomerOffTransferResultReq)(result *customerOffTransferResultRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerOffTransferResultURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cft.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerOffTransferResult error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//分配离职成员的客户群
type CreateCustomerOffGroupChatTransferReq struct {
	ChatIdList []string `json:"chat_id_list"`
	NewOwner   string   `json:"new_owner"`
}
type createCustomerOffGroupChatTransferRep struct {
	util.WxError
	FailedChatList []struct {
		ChatId  string `json:"chat_id"`
		Errcode int32    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	} `json:"failed_chat_list"`
}
func (cft *CustomerOffTransfer) CreateCustomerOffGroupChatTransfer(accessToken string, req CreateCustomerOffGroupChatTransferReq)(result *createCustomerOffGroupChatTransferRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerOffGroupChatTransferURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cft.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerOffGroupChatTransfer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

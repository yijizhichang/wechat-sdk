//客户联系-在职继承
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateCustomerOnTransferURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_customer?access_token=%s"  //分配在职成员的客户
	GetCustomerOnTransferURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/transfer_result?access_token=%s"  //查询客户接替状态
)

//CustomerOnTransfer 消息推送
type CustomerOnTransfer struct {
	*core.Context
}

//NewCustomerMsg 实例化
func NewCustomerOnTransfer(context *core.Context) *CustomerOnTransfer {
	cot := new(CustomerOnTransfer)
	cot.Context = context
	return cot
}

//分配在职成员的客户
type CreateCustomerOnTransferReq struct {
	HandoverUserid     string   `json:"handover_userid"`
	TakeoverUserid     string   `json:"takeover_userid"`
	ExternalUserid     []string `json:"external_userid"`
	TransferSuccessMsg string   `json:"transfer_success_msg"`
}
type createCustomerOnTransferRep struct {
	util.WxError
	Customer []struct {
		ExternalUserid string `json:"external_userid"`
		Errcode        int    `json:"errcode"`
	} `json:"customer"`
}
func (cot *CustomerOnTransfer) CreateCustomerOnTransfer(accessToken string, req CreateCustomerOnTransferReq)(result *createCustomerOnTransferRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerOnTransferURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cot.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerOnTransfer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//查询客户接替状态
type CustomerOnTransferReq struct {
	HandoverUserid string `json:"handover_userid"`
	TakeoverUserid string `json:"takeover_userid"`
	Cursor         string `json:"cursor"`
}
type customerOnTransferRep struct {
	util.WxError
	Customer []struct {
		ExternalUserid string `json:"external_userid"`
		Status         int    `json:"status"`
		TakeoverTime   int    `json:"takeover_time"`
	} `json:"customer"`
	NextCursor string `json:"next_cursor"`
}
func (cot *CustomerOnTransfer) GetCustomerOnTransfer(accessToken string, req CustomerOnTransferReq)(result *customerOnTransferRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerOnTransferURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cot.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerOnTransfer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

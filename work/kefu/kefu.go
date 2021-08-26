//微信客服-客户基本信息获取，升级服务配置
package kefu

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetKfCustomerListURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/batchget?access_token=%s"  //客户基本信息获取
	GetKfCustomerUpgradeServiceConfigURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/get_upgrade_service_config?access_token=%s"  //获取配置的专员与客户群
	UpgradeKfCustomerServiceURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/upgrade_service?access_token=%s"  //升级专员服务
	CancelKfCustomerServiceURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/customer/cancel_upgrade_service?access_token=%s"  //为客户取消推荐

)

//Kefu
type Kefu struct {
	*core.Context
}

//NewKefu 实例化
func NewKefu(context *core.Context) *Kefu {
	kf := new(Kefu)
	kf.Context = context
	return kf
}

//客户基本信息获取
type KfCustomerListReq struct {
	ExternalUseridList []string `json:"external_userid_list"`
}
type KfCustomerListRep struct {
	util.WxError
	CustomerList []struct {
		ExternalUserid string `json:"external_userid"`
		Nickname       string `json:"nickname"`
		Avatar         string `json:"avatar"`
		Gender         int32  `json:"gender"`
		Unionid        string `json:"unionid"`
	} `json:"customer_list"`
	InvalidExternalUserid []string `json:"invalid_external_userid"`
}
func (kf *KefuAccount) GetKfCustomerList(accessToken string, req KfCustomerListReq)(result *KfCustomerListRep, err error){
	qyUrl := fmt.Sprintf(GetKfCustomerListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfCustomerList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取配置的专员与客户群
type KfCustomerUpgradeServiceConfig struct {
	util.WxError
	MemberRange struct {
		UseridList       []string `json:"userid_list"`
		DepartmentIdList []int    `json:"department_id_list"`
	} `json:"member_range"`
	GroupchatRange struct {
		ChatIdList []string `json:"chat_id_list"`
	} `json:"groupchat_range"`
}
func (kf *KefuAccount) GetKfCustomerUpgradeServiceConfig(accessToken string)(result *KfCustomerUpgradeServiceConfig, err error){
	qyUrl := fmt.Sprintf(GetKfCustomerUpgradeServiceConfigURL, accessToken)

	response, err := util.HTTPGet(qyUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfCustomerUpgradeServiceConfig error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//升级专员服务
type UpgradeKfCustomerServiceReq struct {
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	Type           int32  `json:"type"`
	Member         struct {
		Userid  string `json:"userid"`
		Wording string `json:"wording"`
	} `json:"member,omitempty"`
	Groupchat struct {
		ChatId  string `json:"chat_id"`
		Wording string `json:"wording"`
	} `json:"groupchat,omitempty"`
}
func (kf *KefuAccount) UpgradeKfCustomerService(accessToken string, req UpgradeKfCustomerServiceReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpgradeKfCustomerServiceURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpgradeKfCustomerService error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//为客户取消推荐
type CancelKfCustomerServiceReq struct {
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
}
func (kf *KefuAccount) CancelKfCustomerService(accessToken string, req CancelKfCustomerServiceReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(CancelKfCustomerServiceURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CancelKfCustomerService error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

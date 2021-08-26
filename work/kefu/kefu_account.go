//微信客服-客服账号管理
package kefu

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateKfAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/add?access_token=%s"  //添加客服帐号
	DelKfAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/del?access_token=%s"  //删除客服帐号
	UpdateKfAccountURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/update?access_token=%s"  //修改客服帐号
	GetKfAccountListURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/account/list?access_token=%s"  //获取客服帐号列表
	GetKfContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/add_contact_way?access_token=%s"  //获取客服帐号链接
)

//KefuAccount
type KefuAccount struct {
	*core.Context
}

//NewKefuAccount 实例化
func NewKefuAccount(context *core.Context) *KefuAccount {
	kf := new(KefuAccount)
	kf.Context = context
	return kf
}

//添加客服帐号
type CreateKfAccountReq struct {
	Name    string `json:"name"`
	MediaId string `json:"media_id"`
}
type CreateKfAccountRep struct {
	util.WxError
	OpenKfid string `json:"open_kfid"`
}
func (kf *KefuAccount) CreateKfAccount(accessToken string, req CreateKfAccountReq)(result *CreateKfAccountRep, err error){
	qyUrl := fmt.Sprintf(CreateKfAccountURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateKfAccount error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除客服帐号
type DelKfAccountReq struct {
	OpenKfid string `json:"open_kfid"`
}
func (kf *KefuAccount) DelKfAccount(accessToken string, req DelKfAccountReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelKfAccountURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelKfAccount error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改客服帐号
type UpdateKfAccountReq struct {
	OpenKfid string `json:"open_kfid"`
	Name     string `json:"name"`
	MediaId  string `json:"media_id"`
}
func (kf *KefuAccount) UpdateKfAccount(accessToken string, req UpdateKfAccountReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateKfAccountURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateKfAccount error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客服帐号列表
type KfAccountList struct {
	util.WxError
	AccountList []struct {
		OpenKfid string `json:"open_kfid"`
		Name     string `json:"name"`
		Avatar   string `json:"avatar"`
	} `json:"account_list"`
}
func (kf *KefuAccount) GetKfAccountList(accessToken string)(result *KfAccountList, err error){
	qyUrl := fmt.Sprintf(GetKfAccountListURL, accessToken)

	response, err := util.HTTPGet(qyUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfAccountList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客服帐号链接
type KfContactWayReq struct {
	OpenKfid string `json:"open_kfid"`
	Scene    string `json:"scene"`
}
type KfContactWayRep struct {
	util.WxError
	Url     string `json:"url"`
}
func (kf *KefuAccount) GetKfContactWay(accessToken string, req KfContactWayReq)(result *KfContactWayRep, err error){
	qyUrl := fmt.Sprintf(GetKfContactWayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfContactWay error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


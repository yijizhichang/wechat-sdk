//微信客服-接待人员管理
package kefu

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateKfServicerURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/add?access_token=%s"  //添加接待人员
	DelKfServicerURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/del?access_token=%s"  //删除接待人员
	GetKfServicerListURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/servicer/list?access_token=%s&open_kfid=%s"  //获取接待人员列表
)

//KefuAccount
type KefuServicer struct {
	*core.Context
}

//KefuServicer 实例化
func NewKefuServicer(context *core.Context) *KefuServicer {
	kf := new(KefuServicer)
	kf.Context = context
	return kf
}

//添加接待人员
type CreateKfServicerReq struct {
	OpenKfid   string   `json:"open_kfid"`
	UseridList []string `json:"userid_list"`
}
type createKfServicerRep struct {
	util.WxError
	ResultList []struct {
		Userid  string `json:"userid"`
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	} `json:"result_list"`
}
func (kf *KefuServicer) CreateKfServicer(accessToken string, req CreateKfServicerReq)(result *createKfServicerRep, err error){
	qyUrl := fmt.Sprintf(CreateKfServicerURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateKfServicer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除接待人员
type DelKfServicerReq struct {
	OpenKfid   string   `json:"open_kfid"`
	UseridList []string `json:"userid_list"`
}
type delKfServicerRep struct {
	util.WxError
	ResultList []struct {
		Userid  string `json:"userid"`
		Errcode int32    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	} `json:"result_list"`
}
func (kf *KefuServicer) DelKfServicer(accessToken string, req DelKfServicerReq)(result *delKfServicerRep, err error){
	qyUrl := fmt.Sprintf(DelKfServicerURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelKfServicer error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取接待人员列表
type kfServicerList struct {
	util.WxError
	ServicerList []struct {
		Userid string `json:"userid"`
		Status int32  `json:"status"`
	} `json:"servicer_list"`
}
func (kf *KefuServicer) GetKfServicerList(accessToken string)(result *kfServicerList, err error){
	qyUrl := fmt.Sprintf(GetKfServicerListURL, accessToken)

	response, err := util.HTTPGet(qyUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfServicerList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


//会话控制
package custom

import (
	"github.com/yijizhichang/wechat-sdk/util"
	"fmt"
	"encoding/json"
)

const (
	CreateKfSessionURL         	= "https://api.weixin.qq.com/customservice/kfsession/create?access_token=%s"         					//创建会话
	CloseKfSessionURL      		= "https://api.weixin.qq.com/customservice/kfsession/close?access_token=%s"             				//关闭会话
	GetKfSessionURL         	= "https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=%s&openid=%s"         		//获取客户会话状态
	GetKfSessionListURL 		= "https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=%s&kf_account=%s" 		//获取客服会话列表
	GetKfWaitCaseListURL    	= "https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=%s"     					//获取未接入会话列表
)

//创建会话
type createKfSession struct {
	KfAccount 	string		`json:"kf_account"`
	Openid 		string 		`json:"openid"`
}

func (kf *Custom) CreateKfSession(kfAccount, openid string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(CreateKfSessionURL, accessToken)

	postData := new(createKfSession)
	postData.KfAccount = kfAccount
	postData.Openid = openid

	response, err := util.PostJSON(wxUrl, postData, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("CreateKfSession error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("创建会话错误", err)
	}
	return
}


//关闭会话
type closeKfSession struct {
	KfAccount 	string		`json:"kf_account"`
	Openid 		string 		`json:"openid"`
}

func (kf *Custom) CloseKfSession(kfAccount, openid string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(CloseKfSessionURL, accessToken)

	postData := new(closeKfSession)
	postData.KfAccount = kfAccount
	postData.Openid = openid

	response, err := util.PostJSON(wxUrl, postData, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("CloseKfSession error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("关闭会话错误", err)
	}
	return
}


//获取客户会话状态
type kfSessionResult struct{
	util.WxError
	KfAccount		string		`json:"kf_account"`
	CreateTime		int64		`json:"createtime"`
}
func (kf *Custom) GetKfSession(openid string) (result *kfSessionResult, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetKfSessionURL, accessToken, openid)

	response, err := util.HTTPGet(wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetKfSession error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("获取客户会话状态错误", err)
	}
	return
}


//获取客服会话列表
type kfSessionListResult struct{
	util.WxError
	SessionList		[] *sessionOpenid	`json:"sessionlist"`
}

type sessionOpenid struct{
	Openid     		string		`json:"openid"`
	CreateTime 		int64		`json:"createtime"`
}
func (kf *Custom) GetKfSessionList(kfAccount string) (result *kfSessionListResult, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetKfSessionListURL, accessToken, kfAccount)

	response, err := util.HTTPGet(wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetKfSessionList error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("获取客服会话列表错误", err)
	}
	return
}


//获取未接入会话列表
type waitCaseResult struct{
	util.WxError
	count  			int64		`json:"count"`
	WaitCaseList	[] *waitCaseOpenid	`json:"waitcaselist"`
}

type waitCaseOpenid struct{
	Openid     		string		`json:"openid"`
	LlatestTime 	int64		`json:"latest_time"`
}
func (kf *Custom) GetWaitCaseList() (result *waitCaseResult, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetKfWaitCaseListURL, accessToken)

	response, err := util.HTTPGet(wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetWaitCaseList error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("获取未接入会话列表错误", err)
	}
	return
}

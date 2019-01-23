//客服管理包
package custom

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	AddKfAccountURL    = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=%s"                         //添加客服账号
	UpdateKfAccountURL = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=%s"                      //修改客服账号
	DelKfAccountURL    = "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=%s&kf_account=%s"           //删除客服账号
	SetHeadImgURL      = "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=%s&kf_account=%s" //设置客服头像
	GetKfListURL       = "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=%s"                     //获取所有客服账号
	GetOnlineKfListURL = "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=%s"               //获取在线客服账号
	SendMsgByKfURL     = "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s"                         //客服接口-发消息
	TypingByKfURL      = "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=%s"                       //客服输入状态
	InviteWorker       = "https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=%s"                //邀请绑定客服帐号
)

//Custom 客服管理
type Custom struct {
	*core.Context
}

//NewCustom 实例化
func NewCustom(context *core.Context) *Custom {
	kf := new(Custom)
	kf.Context = context
	return kf
}

//添加客服账号
type addKfAccount struct {
	KfAccount string `json:"kf_account"`
	NickName  string `json:"nickname"`
	//Password 	string 		`json:"password"`   //新版不需要PWD
}

func (kf *Custom) AddKfAccount(kfAccount, nickName string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(AddKfAccountURL, accessToken)

	postData := new(addKfAccount)
	postData.KfAccount = kfAccount
	postData.NickName = nickName
	//postData.Password = password

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
		err = fmt.Errorf("AddKfAccount error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("添加客服账号错误", err)
	}
	return
}

//修改客服账号
type updateKfAccount struct {
	KfAccount string `json:"kf_account"`
	NickName  string `json:"nickname"`
	//Password 	string 		`json:"password"`
}

func (kf *Custom) UpdateKfAccount(kfAccount, nickName string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(UpdateKfAccountURL, accessToken)

	postData := new(updateKfAccount)
	postData.KfAccount = kfAccount
	postData.NickName = nickName
	//postData.Password = password

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
		err = fmt.Errorf("UpdateKfAccount error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("修改客服账号错误", err)
	}
	return
}

//删除客服账号
func (kf *Custom) DelKfAccount(kfAccount string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(DelKfAccountURL, accessToken, kfAccount)

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
		err = fmt.Errorf("DelKfAccount error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("删除客服账号错误", err)
	}
	return
}

//邀请绑定客服帐号
type inviteWorker struct {
	KfAccount string `json:"kf_account"`
	InviteWx  string `json:"invite_wx"`
}

func (kf *Custom) InviteWorker(kfAccount, inviteWx string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(InviteWorker, accessToken)

	postData := new(inviteWorker)
	postData.KfAccount = kfAccount
	postData.InviteWx = inviteWx

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
		err = fmt.Errorf("InviteWorker error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		kf.WXLog.Error("邀请绑定客服帐号错误", err)
	}
	return
}

//设置客服帐号的头像
func (kf *Custom) SetHeadImgURL(kfAccount, fileName string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(SetHeadImgURL, accessToken, kfAccount)

	response, err := util.PostFile("media", fileName, wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("SetHeadImgURL error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("设置客服帐号的头像错误", err)
	}
	return
}

//获取所有客服账号
type kfList struct {
	util.WxError
	KfList []*kfInfo `json:"kf_list"`
}

type kfInfo struct {
	KfAccount        string `json:"kf_account"`
	KfNickName       string `json:"kf_nick"`
	KfId             int64  `json:"kf_id"`
	KfHeadImgUrl     string `json:"kf_headimgurl"`
	KfWx             string `json:"kf_wx"`
	InviteWx         string `json:"invite_wx"`
	InviteExpireTime int64  `json:"invite_expire_time"`
	InviteStatus     string `json:"invite_status"`
}

func (kf *Custom) GetKfList() (result *kfList, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetKfListURL, accessToken)

	response, err := util.HTTPGet(wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	kf.WXLog.Debug("获取所有客服账号列表", string(response))

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetKfListURL error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("获取所有客服账号错误", err)
	}
	return
}

//获取在线客服账号
type kfOnlineList struct {
	util.WxError
	KfOnlineList []*kfOnlineInfo `json:"kf_online_list"`
}

type kfOnlineInfo struct {
	KfAccount    string `json:"kf_account"`
	KfId         int64  `json:"kf_id"`
	Status       int64  `json:"status"`
	AcceptedCase int64  `json:"accepted_case"`
}

func (kf *Custom) GetKfOnlineList() (result *kfOnlineList, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetOnlineKfListURL, accessToken)

	response, err := util.HTTPGet(wxUrl, kf.ProxyUrl)
	if err != nil {
		return
	}

	kf.WXLog.Debug("获取在线客服账号列表", string(response))

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetKfOnlineList error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("获取在线客服账号错误", err)
	}
	return
}

//发送客服消息
func (kf *Custom) SendMsgByKf(obj interface{}) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(SendMsgByKfURL, accessToken)

	response, err := util.PostJSON(wxUrl, obj, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("SendMsgByKf error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("发送客服消息错误", err)
	}
	return
}

//客服输入状态
type typingByKf struct {
	Touser  string `json:"touser"`
	Command string `json:"command"`
}

func (kf *Custom) TypingByKf(toUser string) (result util.WxError, err error) {
	accessToken, err := kf.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(TypingByKfURL, accessToken)

	postData := new(typingByKf)
	postData.Touser = toUser
	postData.Command = "Typing"

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
		err = fmt.Errorf("TypingByKf error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		kf.WXLog.Error("客服输入状态错误", err)
	}
	return
}

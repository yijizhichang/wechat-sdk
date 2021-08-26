//身份验证
package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetQyOauth2AuthorizeURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s#wechat_redirect"  //发送应用消息
	GetUserInfoByCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"  //获取访问用户身份
	GetQySsoQrConnectURL = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%s&redirect_uri=%s&state=%s"  //构造独立窗口登录二维码
)

type Oauth2 struct {
	*core.Context
}

func NewOauth2(context *core.Context) *Oauth2 {
	o := new(Oauth2)
	o.Context = context
	return o
}

//构造网页授权链接
//授权后重定向的回调链接地址，请使用urlencode对链接进行处理
func (o *Oauth2) GetQyOauth2AuthorizeURL(redirectUri string, state string)(url string){
	url = fmt.Sprintf(GetQyOauth2AuthorizeURL, o.CorpID, redirectUri, state)
	return
}

//获取访问用户身份
type UserInfoByCode struct {
	util.WxError
	UserId         string `json:"UserId"`
	DeviceId       string `json:"DeviceId"`
	OpenId         string `json:"OpenId"`
	ExternalUserid string `json:"external_userid"`
}
func (o *Oauth2) GetUserInfoByCode(accessToken string, code string)(result *UserInfoByCode, err error){
	qyUrl := fmt.Sprintf(GetUserInfoByCodeURL, accessToken, code)

	response, err := util.HTTPGet(qyUrl, o.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfoByCode error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//构造独立窗口登录二维码
func (o *Oauth2) GetQySsoQrConnect(agentId, redirectUri string, state string)(url string){
	url = fmt.Sprintf(GetQySsoQrConnectURL, o.CorpID, agentId, redirectUri, state)
	return
}



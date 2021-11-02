//身份验证
package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
	"time"
)

const (
	GetQyOauth2AuthorizeURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s#wechat_redirect"  //发送应用消息
	GetUserInfoByCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"  //获取访问用户身份
	GetQySsoQrConnectURL = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%s&redirect_uri=%s&state=%s"  //构造独立窗口登录二维码
	GetJsapiTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=%s"  //获取企业的jsapi_ticket
	GetAgentJsapiTicketURL = "https://qyapi.weixin.qq.com/cgi-bin/ticket/get?access_token=%s&type=agent_config"  //获取应用的jsapi_ticket
	QwJsapiTicketCachePrefix = "wechat_qy_jsapi_ticket_"
	QwAgentJsapiTicketCachePrefix = "wechat_qy_agent_jsapi_ticket_"
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

type repJsapiTicket struct {
	util.WxError
	Ticket string `json:"ticket"`
	ExpiresIn int64 `json:"expires_in"`
}

//获取企业的jsapi_ticket
func (o *Oauth2) GetJsapiTicket(token string)(ticket string, err error){
	//优先从缓存中获取
	key := QwJsapiTicketCachePrefix + o.CorpID
	val, err := o.Cache.Get(key)

	if err != nil {
		if o.Debug {
			fmt.Println("qy cache jsapi_ticket o.Cache.Get Err: %+v",err)
		}
	}else{
		if o.Debug {
			fmt.Println("qy cache jsapi_ticket test by token: %s, val: %+v",token, val)
		}
	}

	if val != "" {
		if o.Debug {
			fmt.Println("qy jsapi_ticket ticket from cache: %s",val)
		}
		ticket = val
		if ticket != "" {
			return
		}
	}

	//从微信服务器获取
	qyUrl := fmt.Sprintf(GetJsapiTicketURL, token)
	response, err := util.HTTPGet(qyUrl, o.ProxyUrl)
	if err != nil {
		return
	}

	var result repJsapiTicket

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetJsapiTicket error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}

	//accessToken存入缓存
	expires := result.ExpiresIn - 1500
	ticket = result.Ticket

	err = o.Cache.Set(key, ticket, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}

	return
}


//获取应用的jsapi_ticket
func (o *Oauth2) GetAgentJsapiTicket(token string, agentId string)(ticket string, err error){
	//优先从缓存中获取
	key := QwAgentJsapiTicketCachePrefix + o.CorpID + agentId
	val, err := o.Cache.Get(key)

	if err != nil {
		if o.Debug {
			fmt.Println("qy cache agent_jsapi_ticket o.Cache.Get Err: %+v",err)
		}
	}else{
		if o.Debug {
			fmt.Println("qy cache agent_jsapi_ticket test by token: %s, val: %+v",token, val)
		}
	}

	if val != "" {
		if o.Debug {
			fmt.Println("qy agent_jsapi_ticket ticket from cache: %s",val)
		}
		ticket = val
		if ticket != "" {
			return
		}
	}

	//从微信服务器获取
	qyUrl := fmt.Sprintf(GetAgentJsapiTicketURL, token)
	response, err := util.HTTPGet(qyUrl, o.ProxyUrl)
	if err != nil {
		return
	}

	var result repJsapiTicket

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAgentJsapiTicket error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
		return
	}

	//accessToken存入缓存
	expires := result.ExpiresIn - 1500
	ticket = result.Ticket

	err = o.Cache.Set(key, ticket, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}

	return
}



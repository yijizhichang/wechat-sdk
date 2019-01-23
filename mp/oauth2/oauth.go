package oauth2

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
	"net/http"
	"net/url"
)

const (
	redirectOauthUrl      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect" //授权回调
	accessTokenUrl        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"                        //通过code换取网页授权access_token
	refreshAccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"                             //刷新access_token
	userInfoUrl           = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"                                                       	//拉取用户信息
	checkAccessTokenUrl   = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"                                                                      //检验授权凭证
)

//保存用户授权信息
type Oauth struct {
	*core.Context
}

//获取用户授权access_token的返回结果
type GrantAccessToken struct {
	util.WxError
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

//用户授权获取到用户信息
type UserInfo struct {
	util.WxError

	OpenId     string   `json:"openid"`
	NickName   string   `json:"nickname"`
	Sex        int32    `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

//实例化授权信息
func NewOauth(context *core.Context) (auth *Oauth) {
	auth = new(Oauth)
	auth.Context = context
	return
}

//获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectUri, scope, state string) (string, error) {
	urlStr := url.QueryEscape(redirectUri)
	return fmt.Sprintf(redirectOauthUrl, oauth.AppID, urlStr, scope, state), nil
}

//跳转到网页授权
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectUri, scope, state string) error {
	location, err := oauth.GetRedirectURL(redirectUri, scope, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}

//通过网页授权的code 换取access_token，网页授权接口调用凭证(此access_token与基础支持的access_token不同)
func (oauth *Oauth) GetGrantAccessToken(code string) (result GrantAccessToken, err error) {
	wxUrl := fmt.Sprintf(accessTokenUrl, oauth.AppID, oauth.AppSecret, code)
	var response []byte
	response, err = util.HTTPGet(wxUrl, oauth.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetGrantAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//刷新access_token
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result GrantAccessToken, err error) {
	wxUrl := fmt.Sprintf(refreshAccessTokenUrl, oauth.AppID, refreshToken)
	var response []byte
	response, err = util.HTTPGet(wxUrl, oauth.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("RefreshAccessToken error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

//检验access_token是否有效
func (oauth *Oauth) CheckAccessToken(accessToken, openId string) (b bool, err error) {
	wxUrl := fmt.Sprintf(checkAccessTokenUrl, accessToken, openId)
	var response []byte
	response, err = util.HTTPGet(wxUrl, oauth.ProxyUrl)
	if err != nil {
		return
	}
	var result util.WxError
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		b = false
		return
	}
	b = true
	return
}

//如果scope为 snsapi_userinfo 则可以通过此方法获取到用户基本信息
func (oauth *Oauth) GetUserInfo(accessToken, openId, lang string) (result UserInfo, err error) {
	wxUrl := fmt.Sprintf(userInfoUrl, accessToken, openId, lang)
	var response []byte
	response, err = util.HTTPGet(wxUrl, oauth.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}

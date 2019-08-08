package core

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"time"
)

const (
	AccessTokenURL              = "https://api.weixin.qq.com/cgi-bin/token"
	AccessTokenCachePrefix      = "wechat_mp_access_token_appid_"
	AccessTokenThirdCachePrefix = "wechat_mp_access_token_third_appid_"
)

type AccessTokenInfo struct {
	util.WxError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (ctx *Context) GetAccessToken() (accessToken string, err error) {
	ctx.accessTokenLock.Lock()
	defer ctx.accessTokenLock.Unlock()

	//如果共享其它应用的accessToken,则优先取共享accessToken，共享的accessToken的更新获取机制不在当前项目维护
	if ctx.ThirdAccessToken {
		thirdKey := AccessTokenThirdCachePrefix + ctx.AppID
		thirdVal, _ := ctx.Cache.Get(thirdKey)

		if thirdVal != "" {
			AccessTokenThird := thirdVal
			accessToken = AccessTokenThird
			return
		}
		return
	}

	//优先从缓存中获取
	key := AccessTokenCachePrefix + ctx.AppID
	val, err := ctx.Cache.Get(key)

	fmt.Println("cache accessToken test err:", err)

	if val != "" {
		fmt.Println("accessToken from cache")
		accessToken = val
		if accessToken != "" {
			return
		}

	}

	//从微信服务器获取
	accessToken, err = ctx.GetAccessTokenFromServer()
	if err != nil {
		err = fmt.Errorf("GetAccessTokenFromMpServer error : errormsg=%v", err)
		return
	}

	return
}

//从微信服务器获取token
func (ctx *Context) GetAccessTokenFromServer() (accessToken string, err error) {
	wxUrl := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenURL, ctx.AppID, ctx.AppSecret)
	var response []byte
	response, err = util.HTTPGet(wxUrl, ctx.ProxyUrl)
	if err != nil {
		return
	}

	var result AccessTokenInfo
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.ErrMsg != "" {
		err = fmt.Errorf("GetAccessTokenFromMpServer error : errcode=%v , errormsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	//accessToken存入缓存
	key := AccessTokenCachePrefix + ctx.AppID
	expires := result.ExpiresIn - 1500
	accessToken = result.AccessToken

	err = ctx.Cache.Set(key, accessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	return
}

//设置共享的accessToken
func (ctx *Context) SetThirdAccessToken(thirdAccessToken string, expires int) (wxError util.WxError, err error) {
	thirdKey := AccessTokenThirdCachePrefix + ctx.AppID
	err = ctx.Cache.Set(thirdKey, thirdAccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	return
}

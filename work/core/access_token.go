//企微获取access_token
package core

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"time"
)

const (
	QyAccessTokenURL              = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	QyAccessTokenCachePrefix      = "wechat_qy_access_token_appid_"
	QyAccessTokenThirdCachePrefix = "wechat_qy_access_token_third_appid_"
)

type AccessTokenInfo struct {
	util.WxError
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func (ctx *Context) GetAccessToken(corpSecret string) (accessToken string, err error) {
	ctx.accessTokenLock.Lock()
	defer ctx.accessTokenLock.Unlock()

	//如果共享其它应用的accessToken,则优先取共享accessToken，共享的accessToken的更新获取机制不在当前项目维护
	if ctx.ThirdAccessToken {
		thirdKey := QyAccessTokenThirdCachePrefix + ctx.CorpID + corpSecret  //企微下每个应用有独立的secret，每个应用的access_token应该分开来处理
		thirdVal, _ := ctx.Cache.Get(thirdKey)

		if thirdVal != "" {
			AccessTokenThird := thirdVal
			accessToken = AccessTokenThird
			return
		}
		return
	}

	//优先从缓存中获取
	key := QyAccessTokenCachePrefix + ctx.CorpID + corpSecret
	val, err := ctx.Cache.Get(key)

	if err != nil {
		if ctx.Debug {
			fmt.Println("qy cache accessToken ctx.Cache.Get Err: %+v",err)
		}
	}else{
		if ctx.Debug {
			fmt.Println("qy cache accessToken test corpSecret: %s, val: %+v",corpSecret, val)
		}
	}

	if val != "" {
		if ctx.Debug {
			fmt.Println("qy accessToken from cache: %s",val)
		}
		accessToken = val
		if accessToken != "" {
			return
		}
	}

	//从微信服务器获取
	accessToken, err = ctx.GetAccessTokenFromServer(corpSecret)
	if err != nil {
		err = fmt.Errorf("GetAccessTokenFromMpServer error : errormsg=%v", err)
		return
	}
	if ctx.Debug {
		fmt.Println("qy accessToken from QyWxServer api: %s",accessToken)
	}
	return
}

//从微信服务器获取token
func (ctx *Context) GetAccessTokenFromServer(corpSecret string) (accessToken string, err error) {
	wxUrl := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", QyAccessTokenURL, ctx.CorpID, corpSecret)
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

	//公众号和企业微返回结构不同，所以此处判断也不同
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetAccessTokenFromQyServer error : errcode=%v , errormsg=%v", result.ErrCode, result.ErrMsg)
		return
	}

	//accessToken存入缓存
	key := QyAccessTokenCachePrefix + ctx.CorpID + corpSecret
	expires := result.ExpiresIn - 1500
	accessToken = result.AccessToken

	err = ctx.Cache.Set(key, accessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	return
}

//设置共享的accessToken
func (ctx *Context) SetThirdAccessToken(corpSecret, thirdAccessToken string, expires int) (wxError util.WxError, err error) {
	thirdKey := QyAccessTokenThirdCachePrefix + ctx.CorpID + corpSecret
	err = ctx.Cache.Set(thirdKey, thirdAccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	return
}


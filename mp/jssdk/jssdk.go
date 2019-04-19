package jssdk

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/util/cache"
	"strconv"
	"time"
)

const (
	JSAPITicketKeyCachePrefix = "wechat_mp_jsapi_ticket_"
	TicketURL                 = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="
)

func NewJSSDK(context *core.Context) *JSAPISDK {
	jssdk := new(JSAPISDK)
	jssdk.Context = context
	return jssdk
}

type JSAPISDK struct {
	*core.Context
}

func (j *JSAPISDK) GetTicket() (ticket string, err error) {
	key := JSAPITicketKeyCachePrefix + j.Context.AppID
	val, err := j.Context.Cache.Get(key)
	if err != nil && err.Error() != cache.FILENIL {
		return
	}
	if val != nil {
		ticket = val.(string)
		if ticket != "" {
			j.Context.WXLog.Debug("从缓存中获取ticket", ticket, "缓存方式", j.Context.CacheModel)
			return
		}
	}

	// 从微信服务器获取
	var apiTicket JSAPITicket
	apiTicket, err = j.GetTicketFromServer()
	if err != nil {
		err = fmt.Errorf("GetTicket error : errormsg=%v", err)
		j.Context.WXLog.Debug("从微信服务器获取ticket失败", err)
		return
	}
	err = j.Context.Cache.Set(key, apiTicket.Ticket, time.Duration(apiTicket.ExpiresIn)*time.Second-300)
	if err != nil {
		return
	}
	return apiTicket.Ticket, err
}

func (j *JSAPISDK) GetTicketFromServer() (ticket JSAPITicket, err error) {
	var accessToken string
	accessToken, err = j.GetAccessToken()
	if err != nil && err.Error() != cache.FILENIL {
		return
	}
	uri := TicketURL + accessToken
	var response []byte
	response, err = util.HTTPGet(uri, j.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &ticket)
	if err != nil {
		return
	}
	if ticket.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", ticket.ErrCode, ticket.ErrMsg)
		return
	}
	return
}

func (j *JSAPISDK) MakeSign(ticket, url string) (sign *Signs) {
	sign = &Signs{
		JsapiTicket: ticket,
		Noncestr:    util.RandomString(16),
		Timestamp:   strconv.Itoa(int(time.Now().Unix())),
		URL:         url,
	}
	sign.Signature = fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", sign.JsapiTicket, sign.Noncestr, sign.Timestamp, url)
	return
}

type JSAPITicket struct {
	util.WxError
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type Signs struct {
	JsapiTicket string
	Noncestr    string
	Timestamp   string
	URL         string
	Signature   string
}

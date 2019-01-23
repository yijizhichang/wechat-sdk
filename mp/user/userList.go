package user

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
)

// openID为空是从头开始拉取
func (u *User) GetUserList(nextOpenID string) (userInfo *UserList, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	if nextOpenID != "" {
		nextOpenID = commonNextOpenID + nextOpenID
	}
	uri := userGetURL + commonToken + accessToken + nextOpenID
	var response []byte
	response, err = util.HTTPGet(uri, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

// 获取公众号的黑名单列表
func (u *User) GetBlacklist(openID string) (userInfo *UserList, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	uri := blacklistURL + commonToken + accessToken
	var response []byte
	reqBegin := beginOpenID{BeginOpenID: openID}
	response, err = util.PostJSON(uri, reqBegin, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &userInfo)
	if err != nil {
		return
	}
	if userInfo.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", userInfo.ErrCode, userInfo.ErrMsg)
		return
	}
	return
}

// 拉黑用户,一次只能拉黑20个用户
func (u *User) BatchBlacklist(openIDs []string) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	uri := batchBlacklistURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	openIDList := openidList{OpenidList: openIDs}
	response, err = util.PostJSON(uri, openIDList, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &wxError)
	if err != nil {
		return
	}
	if wxError.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", wxError.ErrCode, wxError.ErrMsg)
		return
	}
	return
}

// 取消拉黑用户
func (u *User) BatchUnblacklist(openIDs []string) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	uri := batchUnblacklistURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	response, err = util.PostJSON(uri, openidList{OpenidList: openIDs}, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &wxError)
	if err != nil {
		return
	}
	if wxError.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", wxError.ErrCode, wxError.ErrMsg)
		return
	}
	return
}

type UserList struct {
	util.WxError
	Total int `json:"total"`
	Count int `json:"count"`
	Data struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type beginOpenID struct {
	BeginOpenID string `json:"begin_openid"`
}

type openidList struct {
	OpenidList []string `json:"openid_list"`
}

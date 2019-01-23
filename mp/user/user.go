package user

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	userFlag    = "用户管理"
	DefaultLang = "zh_CN"
	baseURL     = "https://api.weixin.qq.com/cgi-bin"
	userBaseURL = baseURL + "/user"
	tagsBaseURL = baseURL + "/tags"

	commonToken         = "?access_token="                          // 调用接口凭证
	commonOpenid        = "&openid="                                // 普通用户的标识，对当前公众号唯一
	commonLang          = "&lang="                                  // 返回国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语
	commonNextOpenID    = "&next_openid="                           // 拉取列表的最后一个用户的OPENID
	userInfoURL         = userBaseURL + "/info"                     // 获取用户基本信息(UnionID机制)
	userInfoBatchURL    = userInfoURL + "/batchget"                 // 批量获取用户基本信息
	updateremarkURL     = userInfoURL + "/updateremark"             // 设置用户备注名
	userGetURL          = userBaseURL + "/get"                      // 获取用户列表
	blacklistURL        = tagsBaseURL + "/members/getblacklist"     // 获取公众号的黑名单列表
	batchBlacklistURL   = tagsBaseURL + "/members/batchblacklist"   // 拉黑用户
	batchUnblacklistURL = tagsBaseURL + "/members/batchunblacklist" // 取消拉黑用户
	createTagsURL       = tagsBaseURL + "/create"
	getTagsURL          = tagsBaseURL + "/get"
	updateTagsURL       = tagsBaseURL + "/update"
	deleteTagsURL       = tagsBaseURL + "/delete"
	getTagsUserURL      = baseURL + "/user/tag/get"
	batchTagURL         = tagsBaseURL + "/members/batchtagging"
	batchUnTagURL       = tagsBaseURL + "/members/batchuntagging"
	getIDListTagsURL    = tagsBaseURL + "/getidlist"
)

type User struct {
	*core.Context
}

func NewUser(context *core.Context) *User {
	user := new(User)
	user.Context = context
	return user
}

// 获取用户基本信息（包括UnionID机制）
func (u *User) GetUserInfo(openID, lang string) (userInfo *ResUserInfo, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}

	if lang == "" {
		lang = DefaultLang
	}

	uri := userInfoURL + commonToken + accessToken + commonOpenid + openID + commonLang + lang
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

// 批量获取用户基本信息
func (u *User) GetUserInfoList(openIDs []OpenIDs) (userList *ResUserInfoList, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := userInfoBatchURL + commonToken + accessToken

	var response []byte
	response, err = util.PostJSON(uri, reqUserInfoList{UserList: openIDs}, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &userList)
	if err != nil {
		return
	}
	if userList.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", userList.ErrCode, userList.ErrMsg)
		return
	}
	return
}

// 设置用户备注名
func (u *User) UpdateRemark(openID, remark string) (userList *ResUserInfoList, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := updateremarkURL + commonToken + accessToken

	var response []byte
	response, err = util.PostJSON(uri, remarks{OpenID: openID, Remark: remark}, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &userList)
	if err != nil {
		return
	}
	if userList.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", userList.ErrCode, userList.ErrMsg)
		return
	}
	return
}

// 请求
type reqUserInfoList struct {
	UserList []OpenIDs `json:"user_list"`
}
type OpenIDs struct {
	OpenID string `json:"openid"`
	Lang   string `json:"lang,omitempty"`
}
type remarks struct {
	OpenID string `json:"openid"`
	Remark string `json:"remark"`
}

// 响应
type ResUserInfo struct {
	util.WxError
	Subscribe      int    `json:"subscribe,omitempty"`
	Openid         string `json:"openid,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	Sex            int    `json:"sex,omitempty"`
	Language       string `json:"language,omitempty"`
	City           string `json:"city,omitempty"`
	Province       string `json:"province,omitempty"`
	Country        string `json:"country,omitempty"`
	Headimgurl     string `json:"headimgurl,omitempty"`
	SubscribeTime  int    `json:"subscribe_time,omitempty"`
	Unionid        string `json:"unionid,omitempty"`
	Remark         string `json:"remark,omitempty"`
	Groupid        int    `json:"groupid,omitempty"`
	TagidList      []int  `json:"tagid_list,omitempty"`
	SubscribeScene string `json:"subscribe_scene,omitempty"`
	QrScene        int    `json:"qr_scene,omitempty"`
	QrSceneStr     string `json:"qr_scene_str,omitempty"`
}

type ResUserInfoList struct {
	util.WxError
	UserInfoList []struct {
		Subscribe      int    `json:"subscribe,omitempty"`
		Openid         string `json:"openid,omitempty"`
		Nickname       string `json:"nickname,omitempty"`
		Sex            int    `json:"sex,omitempty"`
		Language       string `json:"language,omitempty"`
		City           string `json:"city,omitempty"`
		Province       string `json:"province,omitempty"`
		Country        string `json:"country,omitempty"`
		Headimgurl     string `json:"headimgurl,omitempty"`
		SubscribeTime  int    `json:"subscribe_time,omitempty"`
		Unionid        string `json:"unionid,omitempty"`
		Remark         string `json:"remark,omitempty"`
		Groupid        int    `json:"groupid,omitempty"`
		TagidList      []int  `json:"tagid_list,omitempty"`
		SubscribeScene string `json:"subscribe_scene,omitempty"`
		QrScene        int    `json:"qr_scene,omitempty"`
		QrSceneStr     string `json:"qr_scene_str,omitempty"`
	} `json:"user_info_list"`
}

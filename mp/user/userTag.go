package user

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
)

// 创建标签
func (u *User) CreateTag(tagName string) (tag *Tag, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := createTagsURL + commonToken + accessToken
	var response []byte
	response, err = util.PostJSON(uri, Tag{Tag: userTag{Name: tagName}}, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tag)
	if err != nil {
		return
	}
	if tag.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", tag.ErrCode, tag.ErrMsg)
		return
	}
	return
}

// 获取公众号已创建的标签
func (u *User) GetTag() (tags *Tags, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := getTagsURL + commonToken + accessToken
	var response []byte
	response, err = util.HTTPGet(uri, u.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &tags)
	if err != nil {
		return
	}
	if tags.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", tags.ErrCode, tags.ErrMsg)
		return
	}
	return
}

// 编辑标签
func (u *User) UpdateTag(tagName string, tagID int) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := updateTagsURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	response, err = util.PostJSON(uri, Tag{Tag: userTag{Name: tagName, ID: tagID}}, u.ProxyUrl)
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

// 删除标签
func (u *User) DeleteTag(tagID int) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := deleteTagsURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	response, err = util.PostJSON(uri, Tag{Tag: userTag{ID: tagID}}, u.ProxyUrl)
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

// 获取标签下粉丝列表
func (u *User) GetTagsUser(tagID int, nextOpenID string) (resTagsUser *ResTagsUser, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := getTagsUserURL + commonToken + accessToken
	var response []byte
	response, err = util.PostJSON(uri, reqTagsUser{Tagid: tagID, NextOpenID: nextOpenID}, u.ProxyUrl)
	err = json.Unmarshal(response, &resTagsUser)
	if err != nil {
		return
	}
	if resTagsUser.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", resTagsUser.ErrCode, resTagsUser.ErrMsg)
		return
	}
	return
}

// 批量为用户打标签,标签功能目前支持公众号为用户打上最多20个标签,每次传入的openid列表个数不能超过50个
func (u *User) BatchTagging(tagID int, OpenidList []string) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := batchTagURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	response, err = util.PostJSON(uri, reqOpenidList{Tagid: tagID, OpenidList: OpenidList}, u.ProxyUrl)
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

// 批量为用户取消标签,每次传入的openid列表个数不能超过50个
func (u *User) BatchUntagging(tagID int, OpenidList []string) (err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := batchUnTagURL + commonToken + accessToken
	var (
		response []byte
		wxError  util.WxError
	)
	response, err = util.PostJSON(uri, reqOpenidList{Tagid: tagID, OpenidList: OpenidList}, u.ProxyUrl)
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

// 获取用户身上的标签列表
func (u *User) Getidlist(OpenID string) (resTagidList *ResTagidList, err error) {
	var accessToken string
	accessToken, err = u.GetAccessToken()
	if err != nil {
		return
	}
	uri := getIDListTagsURL + commonToken + accessToken
	var response []byte
	response, err = util.PostJSON(uri, openID{OpenID: OpenID}, u.ProxyUrl)
	err = json.Unmarshal(response, &resTagidList)
	if err != nil {
		return
	}
	if resTagidList.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", resTagidList.ErrCode, resTagidList.ErrMsg)
		return
	}
	return
}

type Tag struct {
	util.WxError
	Tag userTag `json:"tag"`
}

type Tags struct {
	util.WxError
	Tags []userTag `json:"tags"`
}

type userTag struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Count int    `json:"count,omitempty"`
}

// 请求
type reqTagsUser struct {
	Tagid      int    `json:"tagid"`
	NextOpenID string `json:"next_openid,omitempty"`
}

type reqOpenidList struct {
	OpenidList []string `json:"openid_list"`
	Tagid      int      `json:"tagid"`
}

type openID struct {
	OpenID string `json:"openid"`
}

// 响应
type ResTagsUser struct {
	util.WxError
	Count int `json:"count"`
	Data struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`
}

type ResTagidList struct {
	util.WxError
	TagidList []int `json:"tagid_list"`
}

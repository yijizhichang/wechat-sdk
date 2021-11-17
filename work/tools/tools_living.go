/**
 * @Time: 2021/11/17 10:34 上午
 * @Author: soupzhb@gmail.com
 * @File: tools_live.go
 * @Software: GoLand
 */

package tools

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	QyLivingCreateURL = "https://qyapi.weixin.qq.com/cgi-bin/living/create?access_token=%s"  //创建预约直播
	QyLivingUpdateURL = "https://qyapi.weixin.qq.com/cgi-bin/living/modify?access_token=%s" //修改预约直播
	QyLivingCancelURL = "https://qyapi.weixin.qq.com/cgi-bin/living/cancel?access_token=%s" //取消预约直播
	QyLivingDeleteReplayURL = "https://qyapi.weixin.qq.com/cgi-bin/living/delete_replay_data?access_token=%s" //删除直播回放
	QyLivingGetLivingCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_code?access_token=%s" //获取微信观看直播凭证
	QyLivingGetUserAllLivingIdURL = "https://qyapi.weixin.qq.com/cgi-bin/living/get_user_all_livingid?access_token=%s" //获取成员直播ID列表
	QyLivingGetLivingInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_info?access_token=%s&livingid=%s" //获取直播详情
	QyLivingGetWatchStatURL = "https://qyapi.weixin.qq.com/cgi-bin/living/get_watch_stat?access_token=%s" //获取直播观看明细
	QyLivingGetLivingShareInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/living/get_living_share_info?access_token=%s" //获取跳转小程序商城的直播观众信息
)

//Message 消息推送
type Living struct {
	*core.Context
}

//NewMessager 实例化
func NewLiving(context *core.Context) *Living {
	l := new(Living)
	l.Context = context
	return l
}

//创建预约直播
type CreateLivingReq struct {
	AnchorUserid         string `json:"anchor_userid"`
	Theme                string `json:"theme"`
	LivingStart          int64  `json:"living_start"`
	LivingDuration       int64  `json:"living_duration"`
	Description          string `json:"description"`
	Type                 int64  `json:"type"`
	Agentid              int64  `json:"agentid"`
	RemindTime           int64  `json:"remind_time"`
	ActivityCoverMediaid string `json:"activity_cover_mediaid"`
	ActivityShareMediaid string `json:"activity_share_mediaid"`
	ActivityDetail       struct {
		Description string   `json:"description"`
		ImageList   []string `json:"image_list"`
	} `json:"activity_detail"`
}
type CreateLivingRep struct {
	util.WxError
	Livingid string `json:"livingid"`
}
func (l *Living) CreateLiving(accessToken string, req CreateLivingReq)(result *CreateLivingRep, err error){
	qyUrl := fmt.Sprintf(QyLivingCreateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateLiving error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改预约直播
type UpdateLivingReq struct {
	Livingid       string `json:"livingid"`
	Theme          string `json:"theme"`
	LivingStart    int64  `json:"living_start"`
	LivingDuration int64  `json:"living_duration"`
	Description    string `json:"description"`
	Type           int64  `json:"type"`
	RemindTime     int64  `json:"remind_time"`
}
func (l *Living) UpdateLiving(accessToken string, req UpdateLivingReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(QyLivingUpdateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateLiving error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//取消预约直播
type CancelLivingReq struct {
	Livingid string `json:"livingid"`
}
func (l *Living) CancelLiving(accessToken string, req CancelLivingReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(QyLivingCancelURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CancelLiving error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除直播回放
type DeleteReplayLivingReq struct {
	Livingid string `json:"livingid"`
}
func (l *Living) DeleteReplayLiving(accessToken string, req DeleteReplayLivingReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(QyLivingDeleteReplayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DeleteReplayLiving error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取微信观看直播凭证
type LivingCodeReq struct {
	Livingid string `json:"livingid"`
	Openid   string `json:"openid"`
}
type LivingCodeRep struct {
	util.WxError
	LivingCode string `json:"living_code"`
}
func (l *Living) GetLivingCode(accessToken string, req LivingCodeReq)(result *LivingCodeRep, err error){
	qyUrl := fmt.Sprintf(QyLivingGetLivingCodeURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLivingCode error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取成员直播ID列表
type UserAllLivingidReq struct {
	Userid string `json:"userid"`
	Cursor string `json:"cursor"`
	Limit  int64  `json:"limit"`
}
type UserAllLivingidRep struct {
	util.WxError
	NextCursor   string   `json:"next_cursor"`
	LivingidList []string `json:"livingid_list"`
}
func (l *Living) GetUserAllLivingid(accessToken string, req UserAllLivingidReq)(result *UserAllLivingidRep, err error){
	qyUrl := fmt.Sprintf(QyLivingGetUserAllLivingIdURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserAllLivingid error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取直播详情
type LivingInfoRep struct {
	util.WxError
	LivingInfo struct {
		Theme                 string   `json:"theme"`
		LivingStart           int64    `json:"living_start"`
		LivingDuration        int64    `json:"living_duration"`
		Status                int64    `json:"status "`
		ReserveStart          int64    `json:"reserve_start"`
		ReserveLivingDuration int64    `json:"reserve_living_duration"`
		Description           string   `json:"description"`
		AnchorUserid          string   `json:"anchor_userid"`
		MainDepartment        int64    `json:"main_department"`
		ViewerNum             int64    `json:"viewer_num"`
		CommentNum            int64    `json:"comment_num"`
		MicNum                int64    `json:"mic_num"`
		OpenReplay            int64    `json:"open_replay"`
		ReplayStatus          int64    `json:"replay_status"`
		Type                  int64    `json:"type"`
		PushStreamUrl         string   `json:"push_stream_url"`
		OnlineCount           int64    `json:"online_count"`
		SubscribeCount        int64    `json:"subscribe_count"`
	} `json:"living_info"`
}
func (l *Living) GetLivingInfo(accessToken string, livingid string)(result *UserAllLivingidRep, err error){
	qyUrl := fmt.Sprintf(QyLivingGetLivingInfoURL, accessToken, livingid)

	response, err := util.HTTPGet(qyUrl, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLivingInfo error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取直播观看明细
type WatchStatReq struct {
	Livingid string `json:"livingid"`
	NextKey  string `json:"next_key"`
}
type WatchStatRep struct {
	util.WxError
	Ending   int64  `json:"ending"`
	NextKey  string `json:"next_key"`
	StatInfo struct {
		Users []struct {
			Userid    string `json:"userid"`
			WatchTime int64  `json:"watch_time"`
			IsComment int64  `json:"is_comment"`
			IsMic     int64  `json:"is_mic"`
		} `json:"users"`
		ExternalUsers []struct {
			ExternalUserid string `json:"external_userid"`
			Type           int64  `json:"type"`
			Name           string `json:"name"`
			WatchTime      int64  `json:"watch_time"`
			IsComment      int64  `json:"is_comment"`
			IsMic          int64  `json:"is_mic"`
		} `json:"external_users"`
	} `json:"stat_info"`
}
func (l *Living) GetWatchStat(accessToken string, req WatchStatReq)(result *WatchStatRep, err error){
	qyUrl := fmt.Sprintf(QyLivingGetWatchStatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetWatchStat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取跳转小程序商城的直播观众信息
type LivingShareInfoReq struct {
	WwShareCode string `json:"ww_share_code"`
}
type LivingShareInfoRep struct {
	util.WxError
	Livingid              string `json:"livingid"`
	ViewerUserid          string `json:"viewer_userid"`
	ViewerExternalUserid  string `json:"viewer_external_userid"`
	InvitorUserid         string `json:"invitor_userid"`
	InvitorExternalUserid string `json:"invitor_external_userid"`
}
func (l *Living) GetLivingShareInfo(accessToken string, req LivingShareInfoReq)(result *LivingShareInfoRep, err error){
	qyUrl := fmt.Sprintf(QyLivingGetLivingShareInfoURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, l.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLivingShareInfo error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

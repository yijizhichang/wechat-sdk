//微信客服-会话分配与消息收发
package kefu

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetKfConverseStateURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/get?access_token=%s"  //获取会话状态
	UpdateKfConverseStateURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/service_state/trans?access_token=%s"  //变更会话状态
	SyncKfConverseMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/sync_msg?access_token=%s" //读取消息
	SendKfConverseMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/kf/send_msg?access_token=%s" //发送消息
)

//KefuAccount
type KefuConverse struct {
	*core.Context
}

//KefuConverse 实例化
func NewKefuConverse(context *core.Context) *KefuConverse {
	kf := new(KefuConverse)
	kf.Context = context
	return kf
}

//获取会话状态
type KfConverseStateReq struct {
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
}
type KfConverseStateRep struct {
	util.WxError
	ServiceState   int32  `json:"service_state"`
	ServicerUserid string `json:"servicer_userid"`
}
func (kf *KefuServicer) GetKfConverseState(accessToken string, req KfConverseStateReq)(result *KfConverseStateRep, err error){
	qyUrl := fmt.Sprintf(GetKfConverseStateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetKfConverseState error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//变更会话状态
type UpdateKfConverseStateReq struct {
	OpenKfid       string `json:"open_kfid"`
	ExternalUserid string `json:"external_userid"`
	ServiceState   int32  `json:"service_state"`
	ServicerUserid string `json:"servicer_userid"`
}
func (kf *KefuServicer) UpdateKfConverseState(accessToken string, req UpdateKfConverseStateReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateKfConverseStateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateKfConverseState error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//读取消息
type SyncKfConverseMsgReq struct {
	Cursor string `json:"cursor"`
	Token  string `json:"token"`
	Limit  int32  `json:"limit"`
}
type SyncKfConverseMsgRep struct {
	util.WxError
	NextCursor string `json:"next_cursor"`
	HasMore    int32  `json:"has_more"`
	MsgList    []struct {
		Msgid          string `json:"msgid"`
		OpenKfid       string `json:"open_kfid"`
		ExternalUserid string `json:"external_userid"`
		SendTime       int32  `json:"send_time"`
		Origin         int32  `json:"origin"`
		ServicerUserid string `json:"servicer_userid"`
		Msgtype        string `json:"msgtype"`
		Text           struct {
			Content string `json:"content"`
			MenuId  string `json:"menu_id"`
		} `json:"text,omitempty"`
		Image struct {
			MediaId string `json:"media_id"`
		} `json:"image,omitempty"`
		Voice struct {
			MediaId string `json:"media_id"`
		} `json:"voice,omitempty"`
		Video struct {
			MediaId string `json:"media_id"`
		} `json:"video,omitempty"`
		File struct {
			MediaId string `json:"media_id"`
		} `json:"file,omitempty"`
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Name      string  `json:"name"`
			Address   string  `json:"address"`
		} `json:"location,omitempty"`
		Link struct {
			Title  string `json:"title"`
			Desc   string `json:"desc"`
			Url    string `json:"url"`
			PicUrl string `json:"pic_url"`
		} `json:"link,omitempty"`
		BusinessCard struct {
			Userid string `json:"userid"`
		} `json:"business_card,omitempty"`
		Miniprogram struct {
			Title        string `json:"title"`
			Appid        string `json:"appid"`
			Pagepath     string `json:"pagepath"`
			ThumbMediaId string `json:"thumb_media_id"`
		} `json:"miniprogram,omitempty"`
		Msgmenu struct {
			HeadContent string `json:"head_content"`
			List        []struct {
				Type  string `json:"type"`
				Click struct {
					Id      string `json:"id"`
					Content string `json:"content"`
				} `json:"click,omitempty"`
				View struct {
					Url     string `json:"url"`
					Content string `json:"content"`
				} `json:"view,omitempty"`
				Miniprogram struct {
					Appid    string `json:"appid"`
					Pagepath string `json:"pagepath"`
					Content  string `json:"content"`
				} `json:"miniprogram,omitempty"`
			} `json:"list"`
			TailContent string `json:"tail_content"`
		} `json:"msgmenu,omitempty"`
		Event struct {
			EventType         string `json:"event_type"`
			OpenKfid          string `json:"open_kfid"`
			ExternalUserid    string `json:"external_userid"`
			Scene             string `json:"scene"`
			SceneParam        string `json:"scene_param"`
			FailMsgid         string `json:"fail_msgid"`
			FailType          int32  `json:"fail_type"`
			ServicerUserid    string `json:"servicer_userid"`
			Status            int32  `json:"status"`
			ChangeType        int32  `json:"change_type"`
			OldServicerUserid string `json:"old_servicer_userid"`
			NewServicerUserid string `json:"new_servicer_userid"`
		} `json:"event"`
	} `json:"msg_list"`
}
func (kf *KefuServicer) SyncKfConverseMsg(accessToken string, req SyncKfConverseMsgReq)(result *SyncKfConverseMsgRep, err error){
	qyUrl := fmt.Sprintf(SyncKfConverseMsgURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SyncKfConverseMsg error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送消息
type SendKfConverseMsgReq struct {
	Touser   string `json:"touser"`
	OpenKfid string `json:"open_kfid"`
	Msgid    string `json:"msgid"`
	Msgtype  string `json:"msgtype"`
	Text     struct {
		Content string `json:"content"`
	} `json:"text,omitempty"`
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image,omitempty"`
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice,omitempty"`
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"video,omitempty"`
	File struct {
		MediaId string `json:"media_id"`
	} `json:"file,omitempty"`
	Link struct {
		Title        string `json:"title"`
		Desc         string `json:"desc"`
		Url          string `json:"url"`
		ThumbMediaId string `json:"thumb_media_id"`
	} `json:"link,omitempty"`
	Miniprogram struct {
		Appid        string `json:"appid"`
		Title        string `json:"title"`
		ThumbMediaId string `json:"thumb_media_id"`
		Pagepath     string `json:"pagepath"`
	} `json:"miniprogram,omitempty"`
	Msgmenu struct {
		HeadContent string `json:"head_content"`
		List        []struct {
			Type  string `json:"type"`
			Click struct {
				Id      string `json:"id"`
				Content string `json:"content"`
			} `json:"click,omitempty"`
			View struct {
				Url     string `json:"url"`
				Content string `json:"content"`
			} `json:"view,omitempty"`
			Miniprogram struct {
				Appid    string `json:"appid"`
				Pagepath string `json:"pagepath"`
				Content  string `json:"content"`
			} `json:"miniprogram,omitempty"`
		} `json:"list"`
		TailContent string `json:"tail_content"`
	} `json:"msgmenu,omitempty"`
	Location struct {
		Name      string `json:"name"`
		Address   string `json:"address"`
		Latitude  float64  `json:"latitude"`
		Longitude float64  `json:"longitude"`
	} `json:"location,omitempty"`
}
type SendKfConverseMsgRep struct {
	util.WxError
	Msgid   string `json:"msgid"`
}
func (kf *KefuServicer) SendKfConverseMsg(accessToken string, req SendKfConverseMsgReq)(result *SendKfConverseMsgRep, err error){
	qyUrl := fmt.Sprintf(SendKfConverseMsgURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, kf.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendKfConverseMsg error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


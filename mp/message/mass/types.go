package mass

import "github.com/yijizhichang/wechat-sdk/util"

// 请求结构，根据标签进行/根据OpenID列表群发
type reqMassSendall struct {
	Msgtype string `json:"msgtype"`
	Filter  struct {
		IsToAll bool `json:"is_to_all"`
		TagID   int  `json:"tag_id"`
	} `json:"filter,omitempty"`
	Touser            []string    `json:"touser,omitempty"`
	Mpnews            commonMedia `json:"mpnews,omitempty"`
	Text              text        `json:"text,omitempty"`
	Voice             commonMedia `json:"voice,omitempty"`
	Image             commonMedia `json:"image,omitempty"`
	Mpvideo           mpvideo     `json:"mpvideo,omitempty"`
	Wxcard            wxcard      `json:"wxcard,omitempty"`
	SendIgnoreReprint int         `json:"send_ignore_reprint,omitempty"` // 图文消息被判定为转载时，是否继续群发。 1为继续群发（转载），0为停止群发。 该参数默认为0
}

type text struct {
	Content string `json:"content"`
}
type commonMedia struct {
	MediaID string `json:"media_id"`
}

type mpvideo struct {
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type wxcard struct {
	CardID string `json:"card_id"`
}

// 预览
type wxcardPreview struct {
	CardID  string `json:"card_id"`
	CardEXT string `json:"card_ext"`
}

// json后放入wxcardPreview中CardEXT字段
type CardEXT struct {
	Code      string `json:"code"`
	Openid    string `json:"openid"`
	Timestamp string `json:"timestamp"`
	Signature string `json:"signature"`
}

// 请求结构，查询群发状态
type reqMassGet struct {
	MsgID string `json:"msg_id"`
}

// 请求结构，删除
type reqMassDel struct {
	MsgID      int `json:"msg_id"`
	ArticleIDX int `json:"article_idx,omitempty"`
}

// 请求结构，预览
// towxname和touser同时赋值时，以towxname优先
type reqMassPreview struct {
	Msgtype  string        `json:"msgtype"`
	Touser   string        `json:"touser,omitempty"`
	Towxname string        `json:"towxname,omitempty"` // 微信号
	Mpnews   commonMedia   `json:"mpnews,omitempty"`
	Text     text          `json:"text,omitempty"`
	Voice    commonMedia   `json:"voice,omitempty"`
	Image    commonMedia   `json:"image,omitempty"`
	Mpvideo  commonMedia   `json:"mpvideo,omitempty"`
	Wxcard   wxcardPreview `json:"wxcard,omitempty"`
}

// 响应结构
type resMass struct {
	util.WxError
	MsgID     int `json:"msg_id"`
	MsgDataID int `json:"msg_data_id"`
}

// 响应结构，预览
type resMassPreview struct {
	util.WxError
	MsgID int `json:"msg_id"`
}

// 响应结构，查询群发状态
type resMassGet struct {
	util.WxError
	MsgID     string `json:"msg_id"`
	MsgStatus string `json:"msg_status"`
}

/**
 * @Time : 2022/7/29 13:07
 * @Author : soupzhb@gmail.com
 * @File : webhook.go
 * @Software: GoLand
 */

package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	WebhookUploadMediaURL = "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=%s&type=%s"
)

//webhook消息类型
type WebhookMsgType string
type WebhookCardType string

const (
	// webhook消息类型
	WebhookMsgTypeText         WebhookMsgType = "text"          // 文本消息
	WebhookMsgTypeMarkdown                    = "markdown"      // markdown消息
	WebhookMsgTypeImage                       = "image"         // 图片消息
	WebhookMsgTypeNews                        = "news"          // 图文消息
	WebhookMsgTypeFile                        = "file"          // 文件消息
	WebhookMsgTypeTemplateCard                = "template_card" // 模板卡片消息
)

const (
	//卡版类型
	WebhookCardTypeText WebhookCardType = "text_notice" //文本卡片
	WebhookCardTypeNews                 = "news_notice" //图文卡片
)

//Webhook 机器人消息
type Webhook struct {
	*core.Context
}

//NewWebhook 实例化
func NewWebhook(context *core.Context) *Webhook {
	m := new(Webhook)
	m.Context = context
	return m
}

//发送文本消息
type SendTextWebhookReq struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

func (m *Webhook) SendTextMessage(webhook string, req SendTextWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTextMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送markdown类型消息
type SendMarkdownWebhookReq struct {
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}

func (m *Webhook) SendMarkdownMessage(webhook string, req SendMarkdownWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMarkdownMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送图片类型
type SendImageWebhookReq struct {
	Msgtype string `json:"msgtype"`
	Image   struct {
		Base64 string `json:"base64"`
		Md5    string `json:"md5"`
	} `json:"image"`
}

func (m *Webhook) SendImageMessage(webhook string, req SendImageWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendImageMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送图文类型
type SendNewsWebhookReq struct {
	Msgtype string `json:"msgtype"`
	News    struct {
		Articles []ArticlesItem `json:"articles"`
	} `json:"news"`
}

type ArticlesItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Picurl      string `json:"picurl"`
}

func (m *Webhook) SendNewsMessage(webhook string, req SendNewsWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendNewsMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送文件类型
type SendFileWebhookReq struct {
	Msgtype string `json:"msgtype"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
}

func (m *Webhook) SendFileMessage(webhook string, req SendFileWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendFileMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送模板卡片 文本通知
type SendTemplateCardTextWebhookReq struct {
	Msgtype      string `json:"msgtype"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl   string `json:"icon_url"`
			Desc      string `json:"desc"`
			DescColor int    `json:"desc_color"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		EmphasisContent struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"emphasis_content"`
		QuoteArea struct {
			Type      int    `json:"type"`
			Url       string `json:"url"`
			Appid     string `json:"appid"`
			Pagepath  string `json:"pagepath"`
			Title     string `json:"title"`
			QuoteText string `json:"quote_text"`
		} `json:"quote_area"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int    `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int    `json:"type"`
			Url      string `json:"url,omitempty"`
			Title    string `json:"title"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int    `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
}

func (m *Webhook) SendTemplateCardTextMessage(webhook string, req SendTemplateCardTextWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTemplateCardTextMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送模板卡片 图文通知
type SendTemplateCardNewsWebhookReq struct {
	Msgtype      string `json:"msgtype"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl   string `json:"icon_url"`
			Desc      string `json:"desc"`
			DescColor int    `json:"desc_color"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		CardImage struct {
			Url         string  `json:"url"`
			AspectRatio float64 `json:"aspect_ratio"`
		} `json:"card_image"`
		ImageTextArea struct {
			Type     int    `json:"type"`
			Url      string `json:"url"`
			Title    string `json:"title"`
			Desc     string `json:"desc"`
			ImageUrl string `json:"image_url"`
		} `json:"image_text_area"`
		QuoteArea struct {
			Type      int    `json:"type"`
			Url       string `json:"url"`
			Appid     string `json:"appid"`
			Pagepath  string `json:"pagepath"`
			Title     string `json:"title"`
			QuoteText string `json:"quote_text"`
		} `json:"quote_area"`
		VerticalContentList []struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"vertical_content_list"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int    `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int    `json:"type"`
			Url      string `json:"url,omitempty"`
			Title    string `json:"title"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int    `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
}

func (m *Webhook) SendTemplateCardNewsMessage(webhook string, req SendTemplateCardNewsWebhookReq) (result *util.WxError, err error) {

	response, err := util.PostJSON(webhook, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTemplateCardNewsMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//上传文件
type UploadWebhookMediaReq struct {
	util.WxError
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (m *Webhook) UploadQyTempMedia(webhookKey, fileType, filename string) (result *UploadWebhookMediaReq, err error) {
	qyUrl := fmt.Sprintf(WebhookUploadMediaURL, webhookKey, fileType)

	response, err := util.PostFile("media", filename, qyUrl, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UploadQyTempMedia error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

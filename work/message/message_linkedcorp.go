//消息推送-互联企业消息推送
package message

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	SendQyLinkedCorpURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/message/send?access_token=%s"  //互联企业消息推送-发送应用消息
)

//MessageLinkedcorp 消息推送群聊会话
type MessageLinkedCorp struct {
	*core.Context
}

//NewMessager 实例化
func NewMessageLinkedCorp(context *core.Context) *MessageLinkedCorp {
	m := new(MessageLinkedCorp)
	m.Context = context
	return m
}

type SendLinkedCorpRep struct {
	util.WxError
	Invaliduser  []string `json:"invaliduser"`
	Invalidparty []string `json:"invalidparty"`
	Invalidtag   []string `json:"invalidtag"`
}
//互联企业消息推送-文本消息
type SendTextQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32      `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32      `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int32 `json:"safe"`
}
func (m *Message) SendTextQyLinkedCorp(accessToken string, req SendTextQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-图片消息
type SendImageQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32      `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32      `json:"agentid"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Safe int32 `json:"safe"`
}
func (m *Message) SendImageQyLinkedCorp(accessToken string, req SendImageQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendImageQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-语音消息
type SendVoiceQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32    `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32    `json:"agentid"`
	Voice   struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}
func (m *Message) SendVoiceQyLinkedCorp(accessToken string, req SendVoiceQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVoiceQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-视频消息
type SendVideoQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32    `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32    `json:"agentid"`
	Video   struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
	Safe int32 `json:"safe"`
}
func (m *Message) SendVideoQyLinkedCorp(accessToken string, req SendVideoQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVideoQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-文件消息
type SendFileQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32    `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32    `json:"agentid"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
	Safe int32 `json:"safe"`
}
func (m *Message) SendFileQyLinkedCorp(accessToken string, req SendFileQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendFileQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-文本卡片消息
type SendTextCardQyLinkedCorpReq struct {
	Touser   []string `json:"touser"`
	Toparty  []string `json:"toparty"`
	Totag    []string `json:"totag"`
	Toall    int32    `json:"toall"`
	Msgtype  string   `json:"msgtype"`
	Agentid  int32    `json:"agentid"`
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
}
func (m *Message) SendTextCardQyLinkedCorp(accessToken string, req SendTextCardQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTextCardQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-图文消息
type SendNewsQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32      `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32      `json:"agentid"`
	News    struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Url         string `json:"url"`
			Picurl      string `json:"picurl"`
			Btntxt      string `json:"btntxt"`
		} `json:"articles"`
	} `json:"news"`
}
func (m *Message) SendNewsQyLinkedCorp(accessToken string, req SendNewsQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendNewsQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-图文消息mpnews
type SendMpNewsQyLinkedCorpReq struct {
	Touser  []string `json:"touser"`
	Toparty []string `json:"toparty"`
	Totag   []string `json:"totag"`
	Toall   int32      `json:"toall"`
	Msgtype string   `json:"msgtype"`
	Agentid int32      `json:"agentid"`
	Mpnews  struct {
		Articles []struct {
			Title            string `json:"title"`
			ThumbMediaId     string `json:"thumb_media_id"`
			Author           string `json:"author"`
			ContentSourceUrl string `json:"content_source_url"`
			Content          string `json:"content"`
			Digest           string `json:"digest"`
		} `json:"articles"`
	} `json:"mpnews"`
	Safe int32 `json:"safe"`
}
func (m *Message) SendMpNewsQyLinkedCorp(accessToken string, req SendMpNewsQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMpNewsQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-markdown消息
type SendMarkdownQyLinkedCorpReq struct {
	Touser   []string `json:"touser"`
	Toparty  []string `json:"toparty"`
	Totag    []string `json:"totag"`
	Toall    int32    `json:"toall"`
	Msgtype  string   `json:"msgtype"`
	Agentid  int32    `json:"agentid"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}
func (m *Message) SendMarkdownQyLinkedCorp(accessToken string, req SendMarkdownQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMarkdownQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//互联企业消息推送-小程序通知消息
type SendMiniprogramQyLinkedCorpReq struct {
	Touser            []string `json:"touser"`
	Toparty           []string `json:"toparty"`
	Totag             []string `json:"totag"`
	Msgtype           string   `json:"msgtype"`
	MiniprogramNotice struct {
		Appid             string `json:"appid"`
		Page              string `json:"page"`
		Title             string `json:"title"`
		Description       string `json:"description"`
		EmphasisFirstItem bool   `json:"emphasis_first_item"`
		ContentItem       []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"content_item"`
	} `json:"miniprogram_notice"`
}
func (m *Message) SendMiniprogramQyLinkedCorp(accessToken string, req SendMiniprogramQyLinkedCorpReq)(result *SendLinkedCorpRep, err error){
	qyUrl := fmt.Sprintf(SendQyLinkedCorpURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMiniprogramQyLinkedCorp error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

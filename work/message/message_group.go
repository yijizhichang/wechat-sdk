//消息推送-发送消息到群聊会话
package message

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateQyAppChatURL = "https://qyapi.weixin.qq.com/cgi-bin/appchat/create?access_token=%s"  //创建群聊会话
	UpdateQyAppChatURL = "https://qyapi.weixin.qq.com/cgi-bin/appchat/update?access_token=%s"  //修改群聊会话
	GetQyAppChatURL = "https://qyapi.weixin.qq.com/cgi-bin/appchat/get?access_token=%s&chatid=%s"  //获取群聊会话
	SendQyAppChatURL = "https://qyapi.weixin.qq.com/cgi-bin/appchat/send?access_token=%s"  //群聊会话应用推送消息


)

//MessageGroup 消息推送群聊会话
type MessageGroup struct {
	*core.Context
}

//NewMessager 实例化
func NewMessageGroup(context *core.Context) *MessageGroup {
	m := new(MessageGroup)
	m.Context = context
	return m
}

//创建群聊会话
type CreateQyAppChatReq struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid"`
}
type CreateQyAppChatRep struct {
	util.WxError
	Chatid  string `json:"chatid"`
}
func (m *MessageGroup) CreateQyAppChat(accessToken string, req CreateQyAppChatReq)(result *CreateQyAppChatRep, err error){
	qyUrl := fmt.Sprintf(CreateQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改群聊会话
type UpdateQyAppChatReq struct {
	Chatid      string   `json:"chatid"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"`
	AddUserList []string `json:"add_user_list"`
	DelUserList []string `json:"del_user_list"`
}
func (m *MessageGroup) UpdateQyAppChat(accessToken string, req UpdateQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取群聊会话
type QyAppChatRep struct {
	util.WxError
	ChatInfo struct {
		Chatid   string   `json:"chatid"`
		Name     string   `json:"name"`
		Owner    string   `json:"owner"`
		Userlist []string `json:"userlist"`
	} `json:"chat_info"`
}
func (m *MessageGroup) GetQyAppChat(accessToken string, chatId string)(result *QyAppChatRep, err error){
	qyUrl := fmt.Sprintf(GetQyAppChatURL, accessToken, chatId)

	response, err := util.HTTPGet(qyUrl, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-文本消息
type SendTextQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendTextQyAppChat(accessToken string, req SendTextQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTextQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-图片消息
type SendImageQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendImageQyAppChat(accessToken string, req SendImageQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendImageQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-语音消息
type SendVoiceQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Voice   struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}
func (m *MessageGroup) SendVoiceQyAppChat(accessToken string, req SendVoiceQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVoiceQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-视频消息
type SendVideoQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Video   struct {
		MediaId     string `json:"media_id"`
		Description string `json:"description"`
		Title       string `json:"title"`
	} `json:"video"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendVideoQyAppChat(accessToken string, req SendVideoQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVideoQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-文件消息
type SendFileQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendFileQyAppChat(accessToken string, req SendFileQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendFileQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-文本卡片消息
type SendTextCardQyAppChatReq struct {
	Chatid   string `json:"chatid"`
	Msgtype  string `json:"msgtype"`
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendTextCardQyAppChat(accessToken string, req SendTextCardQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTextCardQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-图文消息
type SendNewsQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	News    struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Url         string `json:"url"`
			Picurl      string `json:"picurl"`
		} `json:"articles"`
	} `json:"news"`
	Safe int32 `json:"safe"`
}
func (m *MessageGroup) SendNewsQyAppChat(accessToken string, req SendNewsQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendNewsQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-图文消息 mpnews
type SendMpNewsQyAppChatReq struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
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
func (m *MessageGroup) SendMpNewsQyAppChat(accessToken string, req SendMpNewsQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMpNewsQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//群聊会话消息推送-markdown
type SendMarkdownQyAppChatReq struct {
	Chatid   string `json:"chatid"`
	Msgtype  string `json:"msgtype"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
}
func (m *MessageGroup) SendMarkdownQyAppChat(accessToken string, req SendMarkdownQyAppChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SendQyAppChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMarkdownQyAppChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

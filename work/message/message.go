//消息推送-发送应用消息
package message

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	QySendMessageURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s"  //发送应用消息
	QyUpdateTemplateCardMessageURL = "https://qyapi.weixin.qq.com/cgi-bin/message/update_template_card?access_token=%s"  //更新模版卡片消息
	QyRecallMessageURL = "https://qyapi.weixin.qq.com/cgi-bin/message/recall?access_token=%s"  //撤回应用消息
	GetQyMessageStatisticsURL = "https://qyapi.weixin.qq.com/cgi-bin/message/get_statistics?access_token=%s"  //查询应用消息发送统计
)

//Message 消息推送
type Message struct {
	*core.Context
}

//NewMessager 实例化
func NewMessage(context *core.Context) *Message {
	m := new(Message)
	m.Context = context
	return m
}

//发送消息
type sendMessageRep struct {
	util.WxError
	Invaliduser  string `json:"invaliduser"`
	Invalidparty string `json:"invalidparty"`
	Invalidtag   string `json:"invalidtag"`
	Msgid        string `json:"msgid"`
	ResponseCode string `json:"response_code"`
}
type updateMessageRep struct {
	util.WxError
	Invaliduser []string `json:"invaliduser"`
}
//文本消息
type SendTextMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int32 `json:"safe"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTextMessage(accessToken string, req SendTextMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//图片消息
type SendImageMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	Image   struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
	Safe                   int32 `json:"safe"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendImageMessage(accessToken string, req SendImageMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//语音消息
type SendVoiceMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	Voice   struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendVoiceMessage(accessToken string, req SendVoiceMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVoiceMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//视频消息
type SendVideoMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	Video   struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"video"`
	Safe                   int32 `json:"safe"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendVideoMessage(accessToken string, req SendVideoMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendVideoMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//文件消息
type SendFileMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	File    struct {
		MediaId string `json:"media_id"`
	} `json:"file"`
	Safe                   int32 `json:"safe"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendFileMessage(accessToken string, req SendFileMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//文本卡片消息
type SendTextCardMessageReq struct {
	Touser   string `json:"touser"`
	Toparty  string `json:"toparty"`
	Totag    string `json:"totag"`
	Msgtype  string `json:"msgtype"`
	Agentid  int32  `json:"agentid"`
	Textcard struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Url         string `json:"url"`
		Btntxt      string `json:"btntxt"`
	} `json:"textcard"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTextCardMessage(accessToken string, req SendTextCardMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTextCardMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//图文消息
type SendNewsMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
	News    struct {
		Articles []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Url         string `json:"url"`
			Picurl      string `json:"picurl"`
			Appid       string `json:"appid"`
			Pagepath    string `json:"pagepath"`
		} `json:"articles"`
	} `json:"news"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendNewsMessage(accessToken string, req SendNewsMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//图文消息mpnews
type SendMpNewsMessageReq struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int32  `json:"agentid"`
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
	Safe                   int32 `json:"safe"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendMpNewsMessage(accessToken string, req SendMpNewsMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMpNewsMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//markdown消息
type SendMarkdownMessageReq struct {
	Touser   string `json:"touser"`
	Toparty  string `json:"toparty"`
	Totag    string `json:"totag"`
	Msgtype  string `json:"msgtype"`
	Agentid  int32    `json:"agentid"`
	Markdown struct {
		Content string `json:"content"`
	} `json:"markdown"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendMarkdownMessage(accessToken string, req SendMarkdownMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//小程序通知消息
type SendMiniprogramMessageReq struct {
	Touser            string `json:"touser"`
	Toparty           string `json:"toparty"`
	Totag             string `json:"totag"`
	Msgtype           string `json:"msgtype"`
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
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendMiniprogramMessage(accessToken string, req SendMiniprogramMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendMiniprogramMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//模板卡片消息-文本通知型
type SendTemplateCardTextMessageReq struct {
	Touser       string `json:"touser"`
	Toparty      string `json:"toparty"`
	Totag        string `json:"totag"`
	Msgtype      string `json:"msgtype"`
	Agentid      int32  `json:"agentid"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		EmphasisContent struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"emphasis_content"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32  `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int32  `json:"type"`
			Title    string `json:"title"`
			Url      string `json:"url,omitempty"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int32  `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTemplateCardTextMessage(accessToken string, req SendTemplateCardTextMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//模板卡片消息-文本通知型
type SendTemplateCardNewsMessageReq struct {
	Touser       string `json:"touser"`
	Toparty      string `json:"toparty"`
	Totag        string `json:"totag"`
	Msgtype      string `json:"msgtype"`
	Agentid      int32  `json:"agentid"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		CardImage struct {
			Url         string  `json:"url"`
			AspectRatio float64 `json:"aspect_ratio"`
		} `json:"card_image"`
		VerticalContentList []struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"vertical_content_list"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32  `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int32  `json:"type"`
			Title    string `json:"title"`
			Url      string `json:"url,omitempty"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int32  `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTemplateCardNewsMessage(accessToken string, req SendTemplateCardNewsMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

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
//模板卡片消息-按钮交互型
type SendTemplateCardButtonMessageReq struct {
	Touser       string `json:"touser"`
	Toparty      string `json:"toparty"`
	Totag        string `json:"totag"`
	Msgtype      string `json:"msgtype"`
	Agentid      int32  `json:"agentid"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32  `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		CardAction struct {
			Type     int32  `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
		TaskId     string `json:"task_id"`
		ButtonList []struct {
			Text  string `json:"text"`
			Style int32  `json:"style"`
			Key   string `json:"key"`
		} `json:"button_list"`
	} `json:"template_card"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTemplateCardButtonMessage(accessToken string, req SendTemplateCardButtonMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTemplateCardButtonMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//模板卡片消息-投票选择型
type SendTemplateCardVoteMessageReq struct {
	Touser       string `json:"touser"`
	Toparty      string `json:"toparty"`
	Totag        string `json:"totag"`
	Msgtype      string `json:"msgtype"`
	Agentid      int32  `json:"agentid"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		TaskId   string `json:"task_id"`
		Checkbox struct {
			QuestionKey string `json:"question_key"`
			OptionList  []struct {
				Id        string `json:"id"`
				Text      string `json:"text"`
				IsChecked bool   `json:"is_checked"`
			} `json:"option_list"`
			Mode int `json:"mode"`
		} `json:"checkbox"`
		SubmitButton struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"submit_button"`
	} `json:"template_card"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTemplateCardVoteMessage(accessToken string, req SendTemplateCardVoteMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTemplateCardVoteMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//模板卡片消息-多选择型
type SendTemplateCardMultipleMessageReq struct {
	Touser       string `json:"touser"`
	Toparty      string `json:"toparty"`
	Totag        string `json:"totag"`
	Msgtype      string `json:"msgtype"`
	Agentid      int32  `json:"agentid"`
	TemplateCard struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		TaskId     string `json:"task_id"`
		SelectList []struct {
			QuestionKey string `json:"question_key"`
			Title       string `json:"title"`
			SelectedId  string `json:"selected_id"`
			OptionList  []struct {
				Id   string `json:"id"`
				Text string `json:"text"`
			} `json:"option_list"`
		} `json:"select_list"`
		SubmitButton struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"submit_button"`
	} `json:"template_card"`
	EnableIdTrans          int32 `json:"enable_id_trans"`
	EnableDuplicateCheck   int32 `json:"enable_duplicate_check"`
	DuplicateCheckInterval int32 `json:"duplicate_check_interval"`
}
func (m *Message) SendTemplateCardMultipleMessage(accessToken string, req SendTemplateCardMultipleMessageReq)(result *sendMessageRep, err error){
	qyUrl := fmt.Sprintf(QySendMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendTemplateCardMultipleMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新模版卡片消息-更新按钮
type UpdateTemplateCardButtonReq struct {
	Userids      []string `json:"userids"`
	Partyids     []int32  `json:"partyids"`
	Tagids       []int32  `json:"tagids"`
	Atall        int32    `json:"atall"`
	Agentid      int32    `json:"agentid"`
	ResponseCode string   `json:"response_code"`
	Button       struct {
		TaskId      string `json:"task_id"`
		ReplaceName string `json:"replace_name"`
	} `json:"button"`
}
func (m *Message) UpdateTemplateCardButton(accessToken string, req UpdateTemplateCardButtonReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardButton error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//更新模版卡片消息-文本通知型
type UpdateTemplateCardTextMessageReq struct {
	Userids        []string `json:"userids"`
	Partyids       []int32  `json:"partyids"`
	Agentid        int32    `json:"agentid"`
	ResponseCode   string   `json:"response_code"`
	OriginalTaskId string   `json:"original_task_id"`
	TemplateCard   struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		EmphasisContent struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"emphasis_content"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32  `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int32  `json:"type"`
			Title    string `json:"title"`
			Url      string `json:"url,omitempty"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int32  `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
}
func (m *Message) UpdateTemplateCardTextMessage(accessToken string, req UpdateTemplateCardTextMessageReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardTextMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//更新模版卡片消息-图文展示型
type UpdateTemplateCardNewsMessageReq struct {
	Userids        []string `json:"userids"`
	Partyids       []int32    `json:"partyids"`
	Agentid        int32      `json:"agentid"`
	ResponseCode   string   `json:"response_code"`
	OriginalTaskId string   `json:"original_task_id"`
	TemplateCard   struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		CardImage struct {
			Url         string  `json:"url"`
			AspectRatio float64 `json:"aspect_ratio"`
		} `json:"card_image"`
		VerticalContentList []struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"vertical_content_list"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32    `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		JumpList []struct {
			Type     int32    `json:"type"`
			Title    string `json:"title"`
			Url      string `json:"url,omitempty"`
			Appid    string `json:"appid,omitempty"`
			Pagepath string `json:"pagepath,omitempty"`
		} `json:"jump_list"`
		CardAction struct {
			Type     int32    `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
	} `json:"template_card"`
}
func (m *Message) UpdateTemplateCardNewsMessage(accessToken string, req UpdateTemplateCardNewsMessageReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardNewsMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//更新模版卡片消息-按钮交互型
type UpdateTemplateCardButtonInteractionMessageReq struct {
	Userids        []string `json:"userids"`
	Partyids       []int32    `json:"partyids"`
	Agentid        int32      `json:"agentid"`
	ResponseCode   string   `json:"response_code"`
	OriginalTaskId string   `json:"original_task_id"`
	TemplateCard   struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		SubTitleText          string `json:"sub_title_text"`
		HorizontalContentList []struct {
			Keyname string `json:"keyname"`
			Value   string `json:"value"`
			Type    int32    `json:"type,omitempty"`
			Url     string `json:"url,omitempty"`
			MediaId string `json:"media_id,omitempty"`
		} `json:"horizontal_content_list"`
		CardAction struct {
			Type     int32    `json:"type"`
			Url      string `json:"url"`
			Appid    string `json:"appid"`
			Pagepath string `json:"pagepath"`
		} `json:"card_action"`
		ButtonList []struct {
			Text  string `json:"text"`
			Style int32    `json:"style"`
			Key   string `json:"key"`
		} `json:"button_list"`
		ReplaceText string `json:"replace_text"`
	} `json:"template_card"`
}
func (m *Message) UpdateTemplateCardButtonInteractionMessage(accessToken string, req UpdateTemplateCardButtonInteractionMessageReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardButtonInteractionMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//更新模版卡片消息-投票选择型
type UpdateTemplateCardVoteMessageReq struct {
	Userids        []string `json:"userids"`
	Partyids       []int32    `json:"partyids"`
	Agentid        int32      `json:"agentid"`
	ResponseCode   string   `json:"response_code"`
	OriginalTaskId string   `json:"original_task_id"`
	TemplateCard   struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		Checkbox struct {
			QuestionKey string `json:"question_key"`
			OptionList  []struct {
				Id        string `json:"id"`
				Text      string `json:"text"`
				IsChecked bool   `json:"is_checked"`
			} `json:"option_list"`
			Disable bool `json:"disable"`
			Mode    int32  `json:"mode"`
		} `json:"checkbox"`
		SubmitButton struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"submit_button"`
		ReplaceText string `json:"replace_text"`
	} `json:"template_card"`
}
func (m *Message) UpdateTemplateCardVoteMessage(accessToken string, req UpdateTemplateCardVoteMessageReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardVoteMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
//更新模版卡片消息-多项选择型
type UpdateTemplateCardMultipleMessageReq struct {
	Userids        []string `json:"userids"`
	Partyids       []int32    `json:"partyids"`
	Tagids         []int32    `json:"tagids"`
	Atall          int32      `json:"atall"`
	Agentid        int32      `json:"agentid"`
	ResponseCode   string   `json:"response_code"`
	OriginalTaskId string   `json:"original_task_id"`
	TemplateCard   struct {
		CardType string `json:"card_type"`
		Source   struct {
			IconUrl string `json:"icon_url"`
			Desc    string `json:"desc"`
		} `json:"source"`
		MainTitle struct {
			Title string `json:"title"`
			Desc  string `json:"desc"`
		} `json:"main_title"`
		SelectList []struct {
			QuestionKey string `json:"question_key"`
			Title       string `json:"title"`
			SelectedId  string `json:"selected_id"`
			Disable     bool   `json:"disable"`
			OptionList  []struct {
				Id   string `json:"id"`
				Text string `json:"text"`
			} `json:"option_list"`
		} `json:"select_list"`
		SubmitButton struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		} `json:"submit_button"`
		ReplaceText string `json:"replace_text"`
	} `json:"template_card"`
}
func (m *Message) UpdateTemplateCardMultipleMessage(accessToken string, req UpdateTemplateCardMultipleMessageReq)(result *updateMessageRep, err error){
	qyUrl := fmt.Sprintf(QyUpdateTemplateCardMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateTemplateCardMultipleMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//撤回应用消息
type RecallMessageReq struct {
	Msgid string `json:"msgid"`
}
func (m *Message) QyRecallMessage(accessToken string, req RecallMessageReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(QyRecallMessageURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyRecallMessage error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//查询应用消息发送统计
type QyMessageStatisticsReq struct {
	TimeType int `json:"time_type"`
}
type qyMessageStatisticsRep struct {
	util.WxError
	Statistics []struct {
		Agentid int    `json:"agentid"`
		AppName string `json:"app_name"`
		Count   int    `json:"count"`
	} `json:"statistics"`
}
func (m *Message) GetQyMessageStatistics(accessToken string, req QyMessageStatisticsReq)(result *qyMessageStatisticsRep, err error){
	qyUrl := fmt.Sprintf(GetQyMessageStatisticsURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, m.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMessageStatistics error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


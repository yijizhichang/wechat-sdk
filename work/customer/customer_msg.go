//客户联系-消息推送
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateQyMsgTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template?access_token=%s"  //创建企业群发
	GetQyGroupMsgListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_list_v2?access_token=%s"  //获取群发记录列表
	GetQyGroupMsgTaskURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_task?access_token=%s"  //获取群发成员发送任务列表
	GetQyGroupMsgSendResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_groupmsg_send_result?access_token=%s"  //获取企业群发成员执行结果
	QySendWelcomeMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg?access_token=%s"  //发送新客户欢迎语
	CreateQyGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add?access_token=%s"  //添加入群欢迎语素材
	UpdateQyGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit?access_token=%s"  //编辑入群欢迎语素材
	GetQyGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get?access_token=%s"  //获取入群欢迎语素材
	DelQyGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del?access_token=%s"  //删除入群欢迎语素材
)

//CustomerMsg 消息推送
type CustomerMsg struct {
	*core.Context
}

//NewCustomerMsg 实例化
func NewCustomerMsg(context *core.Context) *CustomerMsg {
	cm := new(CustomerMsg)
	cm.Context = context
	return cm
}

//创建企业群发
//调用该接口并不会直接发送消息给客户/客户群，需要成员确认后才会执行发送
//同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息。
type CreateCusMsgTplReq struct {
	ChatType string `json:"chat_type"`
	ExternalUserid []string `json:"external_userid"`
	Sender string `json:"sender"`
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Attachments []attachmentItem `json:"attachments"`
}
type attachmentItem struct {
	MsgType  string  `json:"msgtype"`
	Image struct{
		MediaId	string `json:"media_id"`
		PicUrl string `json:"pic_url"`
	} `json:"image,omitempty"`
	Link struct{
		Title string `json:"title"`
		PicUrl string `json:"picurl"`
		Desc string `json:"desc"`
		Url string `json:"url"`
	} `json:"link,omitempty"`
	Miniprogram struct{
		Title string `json:"title"`
		PicMediaId string `json:"pic_media_id"`
		Appid string `json:"appid"`
		Page string `json:"page"`
	} `json:"miniprogram,omitempty"`
	Video struct{
		MediaId string `json:"media_id"`
	} `json:"video,omitempty"`
}
type CreateCusMsgTplRep struct {
	util.WxError
	FailList []string `json:"fail_list"`
	Msgid string `json:"msgid"`
}
func (cm *CustomerMsg) CreateCustomerMsgTemplate(accessToken string, req CreateCusMsgTplReq)(result *CreateCusMsgTplRep, err error){
	qyUrl := fmt.Sprintf(CreateQyMsgTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerMsgTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取群发记录列表
//企业和第三方应用可通过此接口获取企业与成员的群发记录
type CustomerGroupMsgListReq struct {
	ChatType string `json:"chat_type"`
	StartTime int64 `json:"start_time"`
	EndTime int64 `json:"end_time"`
	Creator string `json:"creator"`
	FilterType int32 `json:"filter_type"`
	Limit int32 `json:"limit"`
	Cursor string `json:"cursor"`
}
type CustomerGroupMsgList struct {
	util.WxError
	GroupMsgList []groupMsgItem `json:"group_msg_list"`
	NextCursor string `json:"next_cursor"`
}
type groupMsgItem struct {
	Msgid string `json:"msgid"`
	Creator string `json:"creator"`
	CreateTime string `json:"create_time"`
	CreateType int32 `json:"create_type"`
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Attachments []attachmentItem `json:"attachments"`
}
func (cm *CustomerMsg) GetCustomerGroupMsgList(accessToken string, req CustomerGroupMsgListReq)(result *CustomerGroupMsgList, err error){
	qyUrl := fmt.Sprintf(GetQyGroupMsgListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupMsgList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取群发成员发送任务列表
type CustomerGroupMsgTaskReq struct {
	Msgid string `json:"msgid"`
	Limit int32 `json:"limit"`
	Cursor string `json:"cursor"`
}
type CustomerGroupMsgTask struct {
	util.WxError
	TaskList []taskItem `json:"task_list"`
	NextCursor string `json:"next_cursor"`
}
type taskItem struct {
	Userid string `json:"userid"`
	Status int32 `json:"status"`
	SendTime int64 `json:"send_time"`
}
func (cm *CustomerMsg) GetCustomerGroupMsgTask(accessToken string, req CustomerGroupMsgTaskReq)(result *CustomerGroupMsgTask, err error){
	qyUrl := fmt.Sprintf(GetQyGroupMsgTaskURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerGroupMsgTask error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//发送新客户欢迎语
type SendWelcomeMsgReq struct {
	WelcomeCode string `json:"welcome_code"`
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Attachments []attachmentItem `json:"attachments"`
}
type SendWelcomeMsgRep struct {
	util.WxError
}
func (cm *CustomerMsg) SendWelcomeMsg(accessToken string, req SendWelcomeMsgReq)(result *SendWelcomeMsgRep, err error){
	qyUrl := fmt.Sprintf(QySendWelcomeMsgURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SendWelcomeMsg error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//添加入群欢迎语素材
type CreateGroupWelcomeTemplateReq struct {
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Image struct{
		MediaId	string `json:"media_id"`
		PicUrl string `json:"pic_url"`
	} `json:"image"`
	Link struct{
		Title string `json:"title"`
		PicUrl string `json:"picurl"`
		Desc string `json:"desc"`
		Url string `json:"url"`
	} `json:"link"`
	Miniprogram struct{
		Title string `json:"title"`
		PicMediaId string `json:"pic_media_id"`
		Appid string `json:"appid"`
		Page string `json:"page"`
	} `json:"miniprogram"`
	Video struct{
		MediaId string `json:"media_id"`
	} `json:"video"`
	File struct{
		MediaId	string `json:"media_id"`
	} `json:"file"`
	Agentid int64 `json:"agentid"`
	Notify int32 `json:"notify"`
}
type CreateGroupWelcomeTemplateRep struct {
	util.WxError
	TemplateId string `json:"template_id"`
}
func (cm *CustomerMsg) CreateGroupWelcomeTemplate(accessToken string, req CreateGroupWelcomeTemplateReq)(result *CreateGroupWelcomeTemplateRep, err error){
	qyUrl := fmt.Sprintf(CreateQyGroupWelcomeTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateGroupWelcomeTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改入群欢迎语素材
type UpdateGroupWelcomeTemplateReq struct {
	TemplateId string `json:"template_id"`
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Image struct{
		MediaId	string `json:"media_id"`
		PicUrl string `json:"pic_url"`
	} `json:"image"`
	Link struct{
		Title string `json:"title"`
		PicUrl string `json:"picurl"`
		Desc string `json:"desc"`
		Url string `json:"url"`
	} `json:"link"`
	Miniprogram struct{
		Title string `json:"title"`
		PicMediaId string `json:"pic_media_id"`
		Appid string `json:"appid"`
		Page string `json:"page"`
	} `json:"miniprogram"`
	Video struct{
		MediaId string `json:"media_id"`
	} `json:"video"`
	File struct{
		MediaId	string `json:"media_id"`
	} `json:"file"`
	Agentid int64 `json:"agentid"`
}
type UpdateGroupWelcomeTemplateRep struct {
	util.WxError
}
func (cm *CustomerMsg) UpdateGroupWelcomeTemplate(accessToken string, req UpdateGroupWelcomeTemplateReq)(result *UpdateGroupWelcomeTemplateRep, err error){
	qyUrl := fmt.Sprintf(UpdateQyGroupWelcomeTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateGroupWelcomeTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取入群欢迎语素材
type GroupWelcomeTemplateReq struct {
	TemplateId string `json:"template_id"`
}
type GroupWelcomeTemplateRep struct {
	util.WxError
	Text struct{
		Content string `json:"content"`
	} `json:"text"`
	Image struct{
		MediaId	string `json:"media_id"`
		PicUrl string `json:"pic_url"`
	} `json:"image"`
	Link struct{
		Title string `json:"title"`
		PicUrl string `json:"picurl"`
		Desc string `json:"desc"`
		Url string `json:"url"`
	} `json:"link"`
	Miniprogram struct{
		Title string `json:"title"`
		PicMediaId string `json:"pic_media_id"`
		Appid string `json:"appid"`
		Page string `json:"page"`
	} `json:"miniprogram"`
	Video struct{
		MediaId string `json:"media_id"`
	} `json:"video"`
	File struct{
		MediaId	string `json:"media_id"`
	} `json:"file"`
}
func (cm *CustomerMsg) GetGroupWelcomeTemplate(accessToken string, req GroupWelcomeTemplateReq)(result *GroupWelcomeTemplateRep, err error){
	qyUrl := fmt.Sprintf(GetQyGroupWelcomeTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetGroupWelcomeTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除入群欢迎语素材
type DelGroupWelcomeTemplateReq struct {
	TemplateId string `json:"template_id"`
	Agentid int64 `json:"agentid"`
}
type DelGroupWelcomeTemplateRep struct {
	util.WxError
}
func (cm *CustomerMsg) DelGroupWelcomeTemplate(accessToken string, req DelGroupWelcomeTemplateReq)(result *DelGroupWelcomeTemplateRep, err error){
	qyUrl := fmt.Sprintf(DelQyGroupWelcomeTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cm.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelGroupWelcomeTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

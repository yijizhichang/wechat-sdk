//消息推送-接收事件推送
package request

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/work/message"
	"strings"
)

// 关注事件
type SubscribeEvent struct {
	message.MsgCommon
	Event message.EventType `xml:"Event" json:"Event"` //事件类型 subscribe

	// 下面两个字段只有在扫描带参数二维码进行关注时才有值, 否则为空值!
	EventKey string `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `xml:"Ticket,omitempty"   json:"Ticket,omitempty"`   // 二维码的ticket，可用来换取二维码图片
}

func GetSubscribeEvent(msg *message.MixMessage) (subscribeEvent *SubscribeEvent) {
	subscribeEvent = new(SubscribeEvent)
	subscribeEvent.MsgCommon = msg.MsgCommon
	subscribeEvent.Event = msg.Event
	subscribeEvent.EventKey = msg.EventKey
	subscribeEvent.Ticket = msg.Ticket
	return
}

//取消关注事件
type UnsubscribeEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`                           //事件类型 unsubscribe
	EventKey string            `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 空值
}

func GetUnsubscribeEvent(msg *message.MixMessage) (unsubscribeEvent *UnsubscribeEvent) {
	unsubscribeEvent = new(UnsubscribeEvent)
	unsubscribeEvent.MsgCommon = msg.MsgCommon
	unsubscribeEvent.Event = msg.Event
	unsubscribeEvent.EventKey = msg.EventKey
	return
}

// 进入应用
type EnterAgentEvent struct {
	message.MsgCommon
	Event message.EventType `xml:"Event" json:"Event"` //事件类型 enter_agent
	EventKey string         `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 空值
}
func GetEnterAgentEvent(msg *message.MixMessage) (enterAgentEvent *EnterAgentEvent) {
	enterAgentEvent = new(EnterAgentEvent)
	enterAgentEvent.MsgCommon = msg.MsgCommon
	enterAgentEvent.Event = msg.Event
	enterAgentEvent.EventKey = msg.EventKey
	return
}

// 上报地理位置事件
type LocationEvent struct {
	message.MsgCommon
	Event     message.EventType `xml:"Event" json:"Event"`         //事件类型 LOCATION
	Latitude  float64           `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64           `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64           `xml:"Precision" json:"Precision"` // 地理位置精度
	AppType   string            `xml:"AppType" json:"AppType"`     // app类型，在企业微信固定返回wxwork，在微信不返回该字段
}

func GetLocationEvent(msg *message.MixMessage) (locationEvent *LocationEvent) {
	locationEvent = new(LocationEvent)
	locationEvent.MsgCommon = msg.MsgCommon
	locationEvent.Event = msg.Event
	locationEvent.Latitude = msg.Latitude
	locationEvent.Longitude = msg.Longitude
	locationEvent.Precision = msg.Precision
	locationEvent.AppType = msg.AppType
	return
}

//异步任务完成事件推送
type BatchJobResultEvent struct {
	message.MsgCommon
	Event     message.EventType `xml:"Event" json:"Event"`     //事件类型 batch_job_result
	JobId    string          `xml:"JobId"  json:"JobId"`       // 异步任务id
	JobType  string          `xml:"JobType"  json:"JobType"`   // 操作类型
	ErrCode  int64           `xml:"ErrCode"  json:"ErrCode"`   // 返回码
	ErrMsg   string          `xml:"ErrMsg"  json:"ErrMsg"`     // 对返回码的文本描述内容
}

func GetBatchJobResultEvent(msg *message.MixMessage) (batchJobResultEvent *BatchJobResultEvent) {
	batchJobResultEvent = new(BatchJobResultEvent)
	batchJobResultEvent.MsgCommon = msg.MsgCommon
	batchJobResultEvent.Event = msg.Event
	batchJobResultEvent.JobId = msg.JobId
	batchJobResultEvent.JobType = msg.JobType
	batchJobResultEvent.ErrCode = msg.ErrCode
	batchJobResultEvent.ErrMsg = msg.ErrMsg
	return
}

//通讯录变更事件
	//-新增成员事件
	//-新增成员事件
	//-更新成员事件
	//-删除成员事件
	//-新增部门事件
	//-更新部门事件
	//-删除部门事件
	//-标签成员变更事件


//菜单事件
//-点击菜单拉取消息的事件推送
//-点击菜单跳转链接的事件推送
type MenuEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 CLICK/VIEW
	EventKey string            `xml:"EventKey" json:"EventKey"` //事件KEY值，与自定义菜单接口中KEY值对应，如果是view，key为跳转URL
}

func GetMenuEvent(msg *message.MixMessage) (menuEvent *MenuEvent) {
	menuEvent = new(MenuEvent)
	menuEvent.MsgCommon = msg.MsgCommon
	menuEvent.Event = msg.Event
	menuEvent.EventKey = msg.EventKey
	return
}
//-扫码推事件的事件推送
//-扫码推事件且弹出“消息接收中”提示框的事件推送
type ScancodeEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 scancode_push/scancode_waitmsg
	EventKey string            `xml:"EventKey" json:"EventKey"` //事件KEY值，与自定义菜单接口中KEY值对应
	ScanCodeInfo  message.ScanCodeInfo   `xml:"ScanCodeInfo" json:"ScanCodeInfo"` //扫描信息
}
func GetScancodeEvent(msg *message.MixMessage)(scancodeEvent *ScancodeEvent){
	scancodeEvent = new(ScancodeEvent)
	scancodeEvent.MsgCommon = msg.MsgCommon
	scancodeEvent.Event = msg.Event
	scancodeEvent.EventKey = msg.EventKey
	scancodeEvent.ScanCodeInfo = msg.ScanCodeInfo
	return
}
//-弹出系统拍照发图的事件推送
//-弹出拍照或者相册发图的事件推送
//-弹出微信相册发图器的事件推送
type SysPhotoEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 pic_sysphoto/pic_photo_or_album/pic_weixin
	EventKey string            `xml:"EventKey" json:"EventKey"` //事件KEY值，与自定义菜单接口中KEY值对应
	SendPicsInfo  message.SendPicsInfo   `xml:"ScanCodeInfo" json:"ScanCodeInfo"` //图片信息
}
func GetSysPhotoEvent (msg *message.MixMessage)(sysPhotoEvent *SysPhotoEvent){
	sysPhotoEvent = new(SysPhotoEvent)
	sysPhotoEvent.Event = msg.Event
	sysPhotoEvent.EventKey = msg.EventKey
	sysPhotoEvent.SendPicsInfo = msg.SendPicsInfo
	return
}
//-弹出地理位置选择器的事件推送
type LocationSelectEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 location_select
	EventKey string            `xml:"EventKey" json:"EventKey"` //事件KEY值，与自定义菜单接口中KEY值对应
	SendLocationInfo  message.SendLocationInfo   `xml:"SendLocationInfo" json:"SendLocationInfo"` //发送的位置信息
}
func GetLocationSelectEvent(msg *message.MixMessage)(locationSelectEvent *LocationSelectEvent){
	locationSelectEvent = new(LocationSelectEvent)
	locationSelectEvent.Event = msg.Event
	locationSelectEvent.EventKey = msg.EventKey
	locationSelectEvent.SendLocationInfo = msg.SendLocationInfo
	return
}
//审批状态通知事件

//共享应用事件回调
type ShareAgentChangeEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 share_agent_change
}
func GetShareAgentChangeEvent(msg *message.MixMessage) (shareAgentChangeEvent *ShareAgentChangeEvent){
	shareAgentChangeEvent = new(ShareAgentChangeEvent)
	shareAgentChangeEvent.MsgCommon = msg.MsgCommon
	shareAgentChangeEvent.Event = msg.Event
	return
}
//模板卡片事件推送

// 获取二维码参数
func (event *SubscribeEvent) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %s 为前缀: %s", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}
//扫描带参数二维码的事件(已关注用户)
type ScanEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`       //事件类型 SCAN
	EventKey string            `xml:"EventKey" json:"EventKey"` //事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string            `xml:"Ticket"   json:"Ticket"`   //二维码的ticket，可用来换取二维码图片
}

func GetScanEvent(msg *message.MixMessage) (scanEvent *ScanEvent) {
	scanEvent = new(ScanEvent)
	scanEvent.MsgCommon = msg.MsgCommon
	scanEvent.Event = msg.Event
	scanEvent.EventKey = msg.EventKey
	scanEvent.Ticket = msg.Ticket
	return
}

//模板消息送达通知
type TemplateSendJobFinishEvent struct {
	message.MsgCommon
	Event  message.EventType `xml:"Event" json:"Event"`   //事件类型 TEMPLATESENDJOBFINISH
	MsgID  int64             `xml:"MsgId"  json:"MsgId"`  // 模板消息ID
	Status string            `xml:"Status" json:"Status"` // 发送状态
}

func GetTemplateSendJobFinishEvent(msg *message.MixMessage) (templateSendJobFinishEvent *TemplateSendJobFinishEvent) {
	templateSendJobFinishEvent = new(TemplateSendJobFinishEvent)
	templateSendJobFinishEvent.MsgCommon = msg.MsgCommon
	templateSendJobFinishEvent.Event = msg.Event
	templateSendJobFinishEvent.MsgID = msg.MsgID
	templateSendJobFinishEvent.Status = msg.Status
	return
}

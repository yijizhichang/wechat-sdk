//接收事件推送
package request

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/message"
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

// 上报地理位置事件
type LocationEvent struct {
	message.MsgCommon
	Event     message.EventType `xml:"Event" json:"Event"`         //事件类型 LOCATION
	Latitude  float64           `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64           `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64           `xml:"Precision" json:"Precision"` // 地理位置精度
}

func GetLocationEvent(msg *message.MixMessage) (locationEvent *LocationEvent) {
	locationEvent = new(LocationEvent)
	locationEvent.MsgCommon = msg.MsgCommon
	locationEvent.Event = msg.Event
	locationEvent.Latitude = msg.Latitude
	locationEvent.Longitude = msg.Longitude
	locationEvent.Precision = msg.Precision
	return
}

//点击菜单事件
type MenuEvent struct {
	message.MsgCommon
	Event    message.EventType `xml:"Event" json:"Event"`         //事件类型 CLICK/VIEW
	EventKey string            `xml:"Longitude" json:"Longitude"` //事件KEY值，与自定义菜单接口中KEY值对应，如果是view，key为跳转URL
}

func GetMenuEvent(msg *message.MixMessage) (menuEvent *MenuEvent) {
	menuEvent = new(MenuEvent)
	menuEvent.MsgCommon = msg.MsgCommon
	menuEvent.Event = msg.Event
	menuEvent.EventKey = msg.EventKey
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

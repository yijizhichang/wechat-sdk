//消息推送-接收消息与事件格式定义
package message

import "encoding/xml"

const (
	// 普通消息类型
	MsgTypeText       MsgType = "text"                      // 文本消息
	MsgTypeImage              = "image"                     // 图片消息
	MsgTypeVoice              = "voice"                     // 语音消息
	MsgTypeVideo              = "video"                     // 视频消息
	MsgTypeShortVideo         = "shortvideo"                // 小视频消息
	MsgTypeLocation           = "location"                  // 地理位置消息
	MsgTypeLink               = "link"                      // 链接消息
	MsgTypeMusic              = "music"                     //音乐消息[限回复]
	MsgTypeNews               = "news"                      //图文消息[限回复]
	MsgTypeTransfer           = "transfer_customer_service" //消息转发到客服
	MsgTypeEvent              = "event"                     //事件推送消息
	MsgTypeNothing            = "nothing"                   //回复空或success
)

const (
	//普通事件类型
	EventSubscribe             EventType = "subscribe"             //订阅
	EventUnsubscribe                     = "unsubscribe"           //取消订阅
	EventScan                            = "SCAN"                  //已经关注的用户扫描带参数二维码事件
	EventLocation                        = "LOCATION"              //上报地理位置事件
	EventClick                           = "CLICK"                 //点击菜单拉取消息时的事件推送
	EventView                            = "VIEW"                  //点击菜单跳转链接时的事件推送
	EventScancodePush                    = "scancode_push"         //扫码推事件的事件推送
	EventScancodeWaitmsg                 = "scancode_waitmsg"      //扫码推事件且弹出“消息接收中”提示框的事件推送
	EventPicSysphoto                     = "pic_sysphoto"          //弹出系统拍照发图的事件推送
	EventPicPhotoOrAlbum                 = "pic_photo_or_album"    //弹出拍照或者相册发图的事件推送
	EventPicWeixin                       = "pic_weixin"            //弹出微信相册发图器的事件推送
	EventLocationSelect                  = "location_select"       //弹出地理位置选择器的事件推送
	EventTemplateSendJobFinish           = "TEMPLATESENDJOBFINISH" //发送模板消息推送通知
	EventEnterAgent                      = "enter_agent"           //进入应用
	EventKfMsgOrEvent                    = "kf_msg_or_event"       //客服会话
)

//消息类型
type MsgType string

//事件类型
type EventType string

//图片事件
type EventPic string

//消息公共字段
type MsgCommon struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
	AgentID      int64    `xml:"AgentID,omitempty"`
	Token        string   `xml:"Token,omitempty"`
}

//安全模式下的消息
type EncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml"  json:"-"`
	ToUserName   string   `xml:"ToUserName" json:"ToUserName"`
	AgentID      int64    `xml:"AgentID" json:"AgentID"`
	EncryptedMsg string   `xml:"Encrypt"  json:"Encrypt"`
}

//需要返回的消息体
type ResponseEncryptedXMLMsg struct {
	XMLName      struct{} `xml:"xml"  json:"-"`
	EncryptedMsg string   `xml:"Encrypt"     json:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature" json:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"    json:"TimeStamp"`
	Nonce        string   `xml:"Nonce"        json:"Nonce"`
}

//存放所有微信发送过来的消息和事件
type ScanCodeInfo struct {
	ScanType   string `xml:"ScanType"`
	ScanResult string `xml:"ScanResult"`
}

type SendPicsInfo struct {
	Count   int32      `xml:"Count"`
	PicList []EventPic `xml:"PicList>item"`
}

type SendLocationInfo struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}

type MixMessage struct {
	MsgCommon

	//基本消息
	MsgID        int64   `xml:"MsgId"`
	Content      string  `xml:"Content"`
	Recognition  string  `xml:"Recognition"`
	PicURL       string  `xml:"PicUrl"`
	MediaID      string  `xml:"MediaId"`
	Format       string  `xml:"Format"`
	ThumbMediaID string  `xml:"ThumbMediaId"`
	LocationX    float64 `xml:"Location_X"`
	LocationY    float64 `xml:"Location_Y"`
	Scale        float64 `xml:"Scale"`
	Label        string  `xml:"Label"`
	Title        string  `xml:"Title"`
	Description  string  `xml:"Description"`
	URL          string  `xml:"Url"`
	AppType      string  `xml:"AppType"`

	//事件相关
	Event       EventType `xml:"Event"`
	EventKey    string    `xml:"EventKey"`
	Ticket      string    `xml:"Ticket"`
	Latitude    float64   `xml:"Latitude"`
	Longitude   float64   `xml:"Longitude"`
	Precision   float64   `xml:"Precision"`
	MenuID      string    `xml:"MenuId"`
	Status      string    `xml:"Status"`
	SessionFrom string    `xml:"SessionFrom"`
	JobId       string    `xml:"JobId"`
	JobType     string    `xml:"JobType"`
	ErrCode     int64     `xml:"ErrCode"`
	ErrMsg      string    `xml:"ErrMsg"`

	ScanCodeInfo ScanCodeInfo `xml:"ScanCodeInfo"`
	SendPicsInfo SendPicsInfo `xml:"SendPicsInfo"`
	SendLocationInfo SendLocationInfo `xml:"SendLocationInfo"`
}

//设置接收对象
func (msg *MsgCommon) SetToUserName(toUserName string) {
	msg.ToUserName = toUserName
}

//设置来源
func (msg *MsgCommon) SetFromUserName(fromUserName string) {
	msg.FromUserName = fromUserName
}

//设置创建时间戳
func (msg *MsgCommon) SetCreateTime(createTime int64) {
	msg.CreateTime = createTime
}

//设置消息类型
func (msg *MsgCommon) SetMsgType(msgType MsgType) {
	msg.MsgType = msgType
}

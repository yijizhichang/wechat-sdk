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
	EventChangeExternalContact           = "change_external_contact"  //客户事件
	EventChangeExternalChat              = "change_external_chat"  //客户群事件
	EventChangeExternalTag               = "change_external_tag"  //企业客户标签事件
	EventBatchJobResult                  = "batch_job_result" //异步任务完成事件推送
	EventChangeContact                   = "change_contact"  //部门变更通知
)

const(
	//事件 ChangeType
	ChangeTypeAddExternalContact         ChangeType = "add_external_contact"  //添加企业客户事件
	ChangeTypeEditExternalContact                   = "edit_external_contact"  //编辑企业客户事件
	ChangeTypeAddHalfExternalContact                = "add_half_external_contact"  //外部联系人免验证添加成员事件
	ChangeTypeDelExternalContact                    = "del_external_contact"  //删除企业客户事件
	ChangeTypeDelFollowUser                         = "del_follow_user"  //删除跟进成员事件
	ChangeTypeTransferFail                          = "transfer_fail"  //客户接替失败事件
	ChangeTypeCreate                                = "create"  //客户群创建事件
	ChangeTypeUpdate                                = "update"  //客户群变更事件
	ChangeTypeDismiss                               = "dismiss"  //客户群解散事件
	ChangeTypeDelete                                = "delete"  //企业客户标签删除事件
	ChangeTypeShuffle                               = "shuffle"  //企业客户标签重排事件
	ChangeTypeCreateParty                           = "create_party"  //新增部门事件
	ChangeTypeUpdateParty                           = "update_party"  //更新部门事件
	ChangeTypeDeleteParty                           = "delete_party"  //删除部门事件
)

//消息类型
type MsgType string

//事件类型
type EventType string
type ChangeType string

//图片事件
type EventPic string

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
	PicList []PicMd5Sum `xml:"PicList>item"`
}
type PicMd5Sum struct {
	PicMd5Sum string `xml:"PicMd5Sum"`
}

type SendLocationInfo struct {
	LocationX float64 `xml:"Location_X"`
	LocationY float64 `xml:"Location_Y"`
	Scale     float64 `xml:"Scale"`
	Label     string  `xml:"Label"`
	Poiname   string  `xml:"Poiname"`
}

type BatchJob struct {
	JobId string `xml:"JobId"`
	JobType string `xml:"JobType"`
	ErrCode int64 `xml:"ErrCode"`
	ErrMsg string `xml:"ErrMsg"`
}

//消息公共字段
type MsgCommon struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      MsgType  `xml:"MsgType"`
	AgentID      int64    `xml:"AgentID,omitempty"`
	Token        string   `xml:"Token,omitempty"`
	Event        EventType  `xml:"Event,omitempty"`
}
//通讯录员工，部门，卡片，审批状态通知事件 todo
type MixMessage struct {
	MsgCommon
	//客户联系消息字段
	ChangeType     ChangeType   `xml:"ChangeType"`
	UserID         string   `xml:"UserID"`
	ExternalUserID string   `xml:"ExternalUserID"`
	State          string   `xml:"State,omitempty"`
	WelcomeCode    string   `xml:"WelcomeCode,omitempty"`
	Source         string   `xml:"Source,omitempty"`
	FailReason     string   `xml:"FailReason,omitempty"`
	ChatId         string   `xml:"ChatId,omitempty"`
	UpdateDetail   string   `xml:"UpdateDetail,omitempty"`
	JoinScene      int64    `xml:"JoinScene,omitempty"`
	QuitScene      int64    `xml:"QuitScene,omitempty"`
	MemChangeCnt   int64    `xml:"MemChangeCnt,omitempty"`
	Id             string   `xml:"Id,omitempty"` //标签或标签组的ID/部门Id
	TagType        string   `xml:"TagType,omitempty"`  //创建标签时，此项为tag，创建标签组时，此项为tag_group
	StrategyId     int64    `xml:"StrategyId,omitempty"`  //标签或标签组所属的规则组id，只回调给“客户联系”应用


	//基本消息
	MsgID        int64   `xml:"MsgId,omitempty"`
	Content      string  `xml:"Content,omitempty"`
	PicURL       string  `xml:"PicUrl,omitempty"`
	MediaID      string  `xml:"MediaId,omitempty"`
	Format       string  `xml:"Format,omitempty"`
	ThumbMediaID string  `xml:"ThumbMediaId,omitempty"`
	LocationX    float64 `xml:"Location_X,omitempty"`
	LocationY    float64 `xml:"Location_Y,omitempty"`
	Scale        float64 `xml:"Scale,omitempty"`
	Label        string  `xml:"Label,omitempty"`
	AppType      string  `xml:"AppType,omitempty"`
	Title        string  `xml:"Title,omitempty"`
	Description  string  `xml:"Description,omitempty"`
	URL          string  `xml:"Url,omitempty"`

	//通讯录-部门变更  //Id 部门Id与其他共用
	Name 		string `xml:"Name"`
	ParentId 	string `xml:"ParentId"`
	Order       int32  `xml:"Order"`

	//事件相关
	EventKey    string    `xml:"EventKey,omitempty"`
	Ticket      string    `xml:"Ticket,omitempty"`
	Latitude    float64   `xml:"Latitude,omitempty"`
	Longitude   float64   `xml:"Longitude,omitempty"`
	Precision   float64   `xml:"Precision"`
	MenuID      string    `xml:"MenuId"`
	Status      string    `xml:"Status"`
	SessionFrom string    `xml:"SessionFrom"`
	JobId       string    `xml:"JobId"`
	JobType     string    `xml:"JobType"`
	ErrCode     int64     `xml:"ErrCode"`
	ErrMsg      string    `xml:"ErrMsg"`


	BatchJob  BatchJob `xml:"BatchJob,omitempty"`
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

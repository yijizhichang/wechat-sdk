//消息推送-接收消息与事件
package request

import "github.com/yijizhichang/wechat-sdk/work/message"

//文本消息
type Text struct {
	message.MsgCommon
	MsgID   int64  `xml:"MsgId"   json:"MsgId"`   // 消息id, 64位整型
	Content string `xml:"Content" json:"Content"` // 文本消息内容
}

func GetText(msg *message.MixMessage) (text *Text) {
	text = new(Text)
	text.MsgCommon = msg.MsgCommon
	text.MsgID = msg.MsgID
	text.Content = msg.Content
	return
}

//图片消息
type Image struct {
	message.MsgCommon
	MsgID   int64  `xml:"MsgId" json:"MsgId"`     // 消息id，64位整型
	MediaID string `xml:"MediaId" json:"MediaId"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	PicURL  string `xml:"PicUrl"   json:"PicUrl"` // 图片链接（由系统生成）
}

func GetImage(msg *message.MixMessage) (image *Image) {
	image = new(Image)
	image.MsgCommon = msg.MsgCommon
	image.MsgID = msg.MsgID
	image.PicURL = msg.PicURL
	image.MediaID = msg.MediaID
	return
}

//语音消息
type Voice struct {
	message.MsgCommon
	MsgID       int64  `xml:"MsgId" json:"MsgId"`                                 // 消息id，64位整型
	MediaID     string `xml:"MediaId" json:"MediaId"`                             // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format      string `xml:"Format"   json:"Format"`                             // 语音格式，如amr，speex等
}

func GetVoice(msg *message.MixMessage) (voice *Voice) {
	voice = new(Voice)
	voice.MsgCommon = msg.MsgCommon
	voice.MsgID = msg.MsgID
	voice.MediaID = msg.MediaID
	voice.Format = msg.Format
	return
}

//视频消息
type Video struct {
	message.MsgCommon
	MsgID        int64  `xml:"MsgId" json:"MsgId"`                 // 消息id，64位整型
	MediaID      string `xml:"MediaId" json:"MediaId"`             // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaID string `xml:"ThumbMediaId"   json:"ThumbMediaId"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据
}

func GetVideo(msg *message.MixMessage) (video *Video) {
	video = new(Video)
	video.MsgCommon = msg.MsgCommon
	video.MsgID = msg.MsgID
	video.MediaID = msg.MediaID
	video.ThumbMediaID = msg.ThumbMediaID
	return
}

// 小视频消息
type ShortVideo struct {
	message.MsgCommon
	MsgID        int64  `xml:"MsgId"        json:"MsgId"`        // 消息id, 64位整型
	MediaID      string `xml:"MediaId"      json:"MediaId"`      // 视频消息媒体id, 可以调用多媒体文件下载接口拉取数据.
	ThumbMediaID string `xml:"ThumbMediaId" json:"ThumbMediaId"` // 视频消息缩略图的媒体id, 可以调用多媒体文件下载接口拉取数据.
}

func GetShortVideo(msg *message.MixMessage) (shortVideo *ShortVideo) {
	shortVideo = new(ShortVideo)
	shortVideo.MsgCommon = msg.MsgCommon
	shortVideo.MsgID = msg.MsgID
	shortVideo.MediaID = msg.MediaID
	shortVideo.ThumbMediaID = msg.ThumbMediaID
	return
}

// 地理位置消息
type Location struct {
	message.MsgCommon
	MsgID     int64   `xml:"MsgId"      json:"MsgId"`      // 消息id, 64位整型
	LocationX float64 `xml:"Location_X" json:"Location_X"` // 地理位置纬度
	LocationY float64 `xml:"Location_Y" json:"Location_Y"` // 地理位置经度
	Scale     float64 `xml:"Scale"      json:"Scale"`      // 地图缩放大小
	Label     string  `xml:"Label"      json:"Label"`      // 地理位置信息
	AppType   string  `xml:"AppType"    json:"AppType"`    //app类型，在企业微信固定返回wxwork，在微信不返回该字段
}

func GetLocation(msg *message.MixMessage) (location *Location) {
	location = new(Location)
	location.MsgCommon = msg.MsgCommon
	location.MsgID = msg.MsgID
	location.LocationX = msg.LocationX
	location.LocationY = msg.LocationY
	location.Scale = msg.Scale
	location.Label = msg.Label
	location.AppType = msg.AppType
	return
}

// 链接消息
type Link struct {
	message.MsgCommon
	MsgID       int64  `xml:"MsgId"       json:"MsgId"`       // 消息id, 64位整型
	Title       string `xml:"Title"       json:"Title"`       // 消息标题
	Description string `xml:"Description" json:"Description"` // 消息描述
	URL         string `xml:"Url"         json:"Url"`         // 消息链接
	PicUrl      string `xml:"PicUrl"      json:"PicUrl"`      // 封面缩略图的url
}

func GetLink(msg *message.MixMessage) (link *Link) {
	link = new(Link)
	link.MsgCommon = msg.MsgCommon
	link.MsgID = msg.MsgID
	link.Title = msg.Title
	link.Description = msg.Description
	link.URL = msg.URL
	link.PicUrl = msg.PicURL
	return
}


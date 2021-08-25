//被动回复消息格式
package response

import (
	"errors"
	"github.com/yijizhichang/wechat-sdk/work/message"
)

//无效的回复
var ErrInvalidReply = errors.New("无效的回复消息")

//不支持的回复类型
var ErrUnsupportReply = errors.New("不支持的回复消息")

//消息回复
type Reply struct {
	MsgType message.MsgType
	MsgData interface{}
}

//文本
type Text struct {
	message.MsgCommon
	Content string `xml:"Content"  json:"Content"` //回复文本消息内容  必填
}

func NewText(content string) (text *Text) {
	text = new(Text)
	text.Content = content
	return
}

//图片
type Image struct {
	message.MsgCommon
	Image struct {
		MediaId string `xml:"MediaId"  json:"MediaId"` //通过素材管理接口上传多媒体文件得到 MediaId  必填
	} `xml:"Image" json:"Image"`
}

func NewImage(mediaId string) (image *Image) {
	image = new(Image)
	image.Image.MediaId = mediaId
	return
}

//语音
type Voice struct {
	message.MsgCommon
	Voice struct {
		MediaId string `xml:"MediaId"  json:"MediaId"` //通过素材管理接口上传多媒体文件得到 MediaId  必填
	} `xml:"Voice" json:"Voice"`
}

func NewVoice(mediaId string) (voice *Voice) {
	voice = new(Voice)
	voice.Voice.MediaId = mediaId
	return
}

//视频
type Video struct {
	message.MsgCommon
	Video struct {
		MediaId     string `xml:"MediaId"  json:"MediaId"`
		Title       string `xml:"Title"  json:"Title"`
		Description string `xml:"Description" json:"Description"`
	} `xml:"Video"  json:"Video"`
}

func NewVideo(mediaId, title, description string) (video *Video) {
	video = new(Video)
	video.Video.MediaId = mediaId
	video.Video.Title = title
	video.Video.Description = description
	return
}

//音乐
type Music struct {
	message.MsgCommon
	Music struct {
		Title        string `xml:"Title"  json:"Title"`              //音乐标题  非必填
		Description  string `xml:"Description" json:"Description"`   //音乐描述  非必填
		MusicURL     string `xml:"MusicUrl"  json:"MusicUrl"`        //音乐链接  非必填
		HQMusicURL   string `xml:"HQMusicUrl" json:"HQMusicUrl"`     //高质量音乐链接，WIFI环境优先使用该链接播放音乐  非必填
		ThumbMediaId string `xml:"ThumbMediaId" json:"ThumbMediaId"` //缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id  必填
	} `xml:"Music"  json:"Music"`
}

func NewMusic(title, description, musicURL, hqMusicURL, thumbMediaId string) (music *Music) {
	music = new(Music)
	music.Music.Title = title
	music.Music.Description = description
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = hqMusicURL
	music.Music.ThumbMediaId = thumbMediaId
	return
}

//图文
type Article struct {
	Title       string `xml:"Title,omitempty"  json:"Title,omitempty"`             // 图文消息标题	 必填
	Description string `xml:"Description,omitempty"  json:"Description,omitempty"` // 图文消息描述	   必填
	PicURL      string `xml:"PicUrl,omitempty"  json:"PicUrl,omitempty"`           // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200	  必填
	URL         string `xml:"Url,omitempty"  json:"Url,omitempty"`                 // 点击图文消息跳转链接	   必填
}

//单个图文信息
func NewArticle(title, description, picURL, url string) (article *Article) {
	article = new(Article)
	article.Title = title
	article.Description = description
	article.PicURL = picURL
	article.URL = url
	return
}

//多条图文
type News struct {
	message.MsgCommon
	ArticleCount int        `xml:"ArticleCount"  json:"ArticleCount"`                  // 图文消息个数, 限制为8条以内		必填
	Articles     []*Article `xml:"Articles>item,omitempty"  json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数量超限, 则将会无响应		必填
}

func NewNews(articles []*Article) (news *News) {
	news = new(News)
	news.ArticleCount = len(articles)
	news.Articles = articles
	return
}


//更新点击用户的按钮点击文案
type CardButton struct {
	message.MsgCommon
	Button struct{
		ReplaceName string `xml:"ReplaceName" json:"ReplaceName"`
	} `xml:"Button" json:"button"`
}
func NewCardButton(replaceName string)(cardButton *CardButton){
	cardButton = new(CardButton)
	cardButton.Button.ReplaceName = replaceName
	return
}

//更新点击用户的整张卡片
type CardWhole struct {
	message.MsgCommon
	TemplateCard struct {
		CardType string `xml:"CardType"`
		Source   struct {
			IconUrl string `xml:"IconUrl"`
			Desc    string `xml:"Desc"`
		} `xml:"Source"`
		MainTitle struct {
			Title string `xml:"Title"`
			Desc  string `xml:"Desc"`
		} `xml:"MainTitle"`
		SubTitleText          string `xml:"SubTitleText"`
		HorizontalContentList []struct {
			KeyName string `xml:"KeyName"`
			Value   string `xml:"Value"`
			Type    string `xml:"Type"`
			URL     string `xml:"Url"`
		} `xml:"HorizontalContentList"`
		JumpList struct {
			Title string `xml:"Title"`
			Type  string `xml:"Type"`
			URL   string `xml:"Url"`
		} `xml:"JumpList"`
		CardAction struct {
			Title string `xml:"Title"`
			Type  string `xml:"Type"`
			URL   string `xml:"Url"`
		} `xml:"CardAction"`
		EmphasisContent struct {
			Title string `xml:"Title"`
			Desc  string `xml:"Desc"`
		} `xml:"EmphasisContent"`
	} `xml:"TemplateCard"`
}
func NewCardWhole(cardWholeReq CardWhole)(cardWhole *CardWhole){
	cardWhole = new(CardWhole)
	cardWhole.TemplateCard = cardWholeReq.TemplateCard
	return
}

//图文展示型

//按钮交互型

//投票选择型

//多项选择型
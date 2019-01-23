package custom

//消息公共字段
type CustomMsgCommon struct {
	ToUser   	string   `json:"touser"`
	MsgType 	string   `json:"msgtype"`
	CustomService  struct{
		KfAccount   string    `json:"kf_account,omitempty"`
	}  `json:"customservice,omitempty"`
}

//文本
type Text struct {
	CustomMsgCommon
	Text struct{
		Content string `json:"content"`
	} `json:"text"`

}

func NewText(content string) (text *Text) {
	text = new(Text)
	text.Text.Content = content
	return
}

//PS 发送文本消息时，支持插入跳小程序的文字链   文本内容<a href="http://www.qq.com" data-miniprogram-appid="appid" data-miniprogram-path="pages/index/index">点击跳小程序</a>
/*说明：
1.data-miniprogram-appid 项，填写小程序appid，则表示该链接跳小程序；
2.data-miniprogram-path项，填写小程序路径，路径与app.json中保持一致，可带参数；
3.对于不支持data-miniprogram-appid 项的客户端版本，如果有herf项，则仍然保持跳href中的网页链接；
4.data-miniprogram-appid对应的小程序必须与公众号有绑定关系。*/


//图片
type Image struct {
	CustomMsgCommon
	Image struct {
		MediaId string `json:"media_id"` //通过素材管理接口上传多媒体文件得到 MediaId  必填
	} `json:"image"`
}

func NewImage(mediaId string) (image *Image) {
	image = new(Image)
	image.Image.MediaId = mediaId
	return
}

//语音
type Voice struct {
	CustomMsgCommon
	Voice struct {
		MediaId string `json:"media_id"` //通过素材管理接口上传多媒体文件得到 MediaId  必填
	} `json:"voice"`
}

func NewVoice(mediaId string) (voice *Voice) {
	voice = new(Voice)
	voice.Voice.MediaId = mediaId
	return
}

//视频
type Video struct {
	CustomMsgCommon
	Video struct {
		MediaId     	string `json:"media_id"`
		ThumbMediaId    string `json:"thumb_media_id"`
		Title       	string `json:"title"`
		Description 	string `json:"description"`
	} `json:"video"`
}

func NewVideo(mediaId, thumbMediaId, title, description string) (video *Video) {
	video = new(Video)
	video.Video.MediaId = mediaId
	video.Video.ThumbMediaId = thumbMediaId
	video.Video.Title = title
	video.Video.Description = description
	return
}

//音乐
type Music struct {
	CustomMsgCommon
	Music struct {
		Title        string `json:"title"`              //音乐标题  非必填
		Description  string `json:"description"`   //音乐描述  非必填
		MusicURL     string `json:"musicurl"`        //音乐链接  非必填
		HQMusicURL   string `json:"hqmusicurl"`     //高质量音乐链接，WIFI环境优先使用该链接播放音乐  非必填
		ThumbMediaId string `json:"thumb_media_id"` //缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id  必填
	} `json:"music"`
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
	Title       string `json:"title"`           // 图文消息标题	 必填
	Description string `json:"description"` 	// 图文消息描述	   必填
	PicURL      string `json:"picurl"`          // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200	  必填
	URL         string `json:"url"`             // 点击图文消息跳转链接	   必填
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

//发送图文消息（点击跳转到外链） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
type News struct {
	CustomMsgCommon
	News struct{
		Articles []*Article `json:"articles"`
	}	`json:"news"`
}

func NewNews(articles []*Article) (news *News) {
	news = new(News)
	news.News.Articles = articles
	return
}

//mpnews
//发送图文消息（点击跳转到图文消息页面） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
type MpNews struct {
	CustomMsgCommon
	MpNews struct {
		MediaId string `json:"media_id"` //通过素材管理接口上传多媒体文件得到 MediaId  必填
	} `json:"mpnews"`
}

func NewMpNews(mediaId string) (mpnews *MpNews) {
	mpnews = new(MpNews)
	mpnews.MpNews.MediaId = mediaId
	return
}

//发送卡券
//特别注意客服消息接口投放卡券仅支持非自定义Code码和导入code模式的卡券的卡券，详情请见：是否自定义code码。https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1451025056&anchor=2.2.2
type Card struct {
	CustomMsgCommon
	WxCard struct {
		CardId string `json:"card_id"`
	} `json:"wxcard"`
}

func NewCard(cardId string) (card *Card) {
	card = new(Card)
	card.WxCard.CardId = cardId
	return
}

//发送小程序卡片（要求小程序与公众号已关联）
type MiniProgramPage struct {
	CustomMsgCommon
	MiniProgramPage struct {
		Title string `json:"title"`
		AppId string `json:"appid"`
		PagePath string `json:"pagepath"`
		ThumbMediaId string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
}

func NewMiniProgramPage(title, appId, pagePath, thumbMediaId string) (miniProgramPage *MiniProgramPage) {
	miniProgramPage = new(MiniProgramPage)
	miniProgramPage.MiniProgramPage.Title = title
	miniProgramPage.MiniProgramPage.AppId = appId
	miniProgramPage.MiniProgramPage.PagePath = pagePath
	miniProgramPage.MiniProgramPage.ThumbMediaId = thumbMediaId
	return
}
package media

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	MediaUploadTempURL       = "https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"          //新增临时素材
	MediaGetTempURL          = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"         //获取临时素材
	MediaGetHqVoiceURL       = "https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"   //高清语音素材获取接口
	MediaAddNewsPermanentURL = "https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=%s"             //新增永久图文素材
	MediaUploadImgURL        = "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"               //上传图文消息内的图片获取URL
	MediaAddMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s" //新增其他类型永久素材
	MediaGetMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=%s"         //获取永久素材
	MediaDelMaterialURL      = "https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=%s"         //删除永久素材
	MediaUpdateNewsURL       = "https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=%s"          //修改永久图文素材
	MediaGetMaterialCountURL = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=%s"    //获取素材总数
	MediaBatchgetMaterialURL = "https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=%s"    //获取素材列表
)

//Media 素材管理
type Media struct {
	*core.Context
}

//NewMedia 实例化
func NewMedia(context *core.Context) *Media {
	media := new(Media)
	media.Context = context
	return media
}

//新增临时素材
type uploadTempMediaResult struct {
	util.WxError
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int32  `json:"created_at"`
}

func (media *Media) UploadTempMedia(fileType, fileName string) (result uploadTempMediaResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaUploadTempURL, accessToken, fileType)

	response, err := util.PostFile("media", fileName, wxUrl, media.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("UploadTempMedia error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("新增临时素材错误", err)
	}
	return
}

//获取临时素材 返回下载地址，作另存处理
//正常返回可下载素材，如果是视频消息则返回{"video_url":DOWN_URL}
func (media *Media) GetTempMediaUrl(mediaId string) (mediaURL string, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	mediaURL = fmt.Sprintf(MediaGetTempURL, accessToken, mediaId)
	return
}

//高清语音素材获取接口
//公众号可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。
func (media *Media) GetMediaGetHqVoiceUrl(mediaId string) (mediaURL string, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	mediaURL = fmt.Sprintf(MediaGetHqVoiceURL, accessToken, mediaId)
	return
}

//上传图文消息内的图片获取URL  图文素材中详情中的图片使用
type mediaUploadImgResult struct {
	util.WxError
	ImgUrl string `json:"url"`
}

func (media *Media) MediaUploadImg(fileName string) (result mediaUploadImgResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaUploadImgURL, accessToken)

	response, err := util.PostFile("media", fileName, wxUrl, media.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("MediaUploadImg error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("上传图文消息内的图片错误", err)
	}
	return
}

//新增永久图文素材
type Article struct {
	Title              string `json:"title"`                 //标题
	ThumbMediaId       string `json:"thumb_media_id"`        //图文消息的封面图片素材id（必须是永久mediaID）
	Author             string `json:"author"`                //作者
	Digest             string `json:"digest"`                //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	ShowCoverPic       int32  `json:"show_cover_pic"`        //是否显示封面，0为false，即不显示，1为true，即显示
	Content            string `json:"content"`               //图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
	ContentSourceUrl   string `json:"content_source_url"`    //图文消息的原文地址，即点击“阅读原文”后的URL
	NeedOpenComment    uint32 `json:"need_open_comment"`     //Uint32 是否打开评论，0不打开，1打开
	OnlyFansCanComment uint32 `json:"only_fans_can_comment"` //Uint32 是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}

//单个图文信息
func NewArticle(title, thumbMediaId, author, digest string, showCoverPic int32, content, contentSourceUrl string, needOpenComment, onlyFansCanComment uint32) (article *Article) {
	article = new(Article)
	article.Title = title
	article.ThumbMediaId = thumbMediaId
	article.Author = author
	article.Digest = digest
	article.ShowCoverPic = showCoverPic
	article.Content = content
	article.ContentSourceUrl = contentSourceUrl
	article.NeedOpenComment = needOpenComment
	article.OnlyFansCanComment = onlyFansCanComment
	return
}

type News struct {
	Articles []*Article `json:"articles"`
}

func NewNews(articles []*Article) (news *News) {
	news = new(News)
	news.Articles = articles
	return
}

type addNewsPermanentResult struct {
	util.WxError
	MediaId string `json:"media_id"`
}
//若新增的是多图文素材，则此处应有几段articles结构，最多8段
func (media *Media) AddNewsPermanent(news *News) (result addNewsPermanentResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaAddNewsPermanentURL, accessToken)

	newsJson, _ := json.Marshal(news)
	media.WXLog.Debug("永久图文素材上传内容", news, "json", string(newsJson))

	response, err := util.PostJSON(wxUrl, news, media.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("AddNewsPermanent error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("新增永久图文素材错误", err)
	}
	return
}

//上传永久素材  视频素材需要单独上传
type addMaterialMediaResult struct {
	util.WxError
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

func (media *Media) AddMaterialMedia(fileType, fileName string) (result addMaterialMediaResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaAddMaterialURL, accessToken, fileType)

	response, err := util.PostFile("media", fileName, wxUrl, media.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("AddMaterialMedia error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("上传永久素材错误", err)
	}
	return
}

//上传永久素材-视频
type addVideoMedia struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}
type addVideoMediaResult struct {
	util.WxError
	MediaId string `json:"media_id"`
	Url     string `json:"url"`
}

func (media *Media) AddVideoMedia(title, introduction, fileName string) (result addVideoMediaResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaAddMaterialURL, accessToken, "video")

	videoDesc := new(addVideoMedia)

	videoDesc.Title = title
	videoDesc.Introduction = introduction

	videoDescJson, err := json.Marshal(videoDesc)
	if err != nil {
		return
	}

	fields := []util.MultipartFormField{
		{
			IsFile:    true,
			Fieldname: "media",
			Filename:  fileName,
		},
		{
			IsFile:    false,
			Fieldname: "description",
			Value:     videoDescJson,
		},
	}

	media.WXLog.Debug("永久视频素材上传内容", fields)

	response, err := util.PostMultipartForm(fields, wxUrl, media.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("AddVideoMedia error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("上传永久视频素材错误", err)
	}
	return
}

//获取永久素材

//永久图文素材
type newsMediaInfoResult struct {
	util.WxError
	NewsItem []*newsItem `json:"news_item"`
}

type mediaInfo struct {
	MediaId string `json:"media_id"`
}

func (media *Media) GetNewsMediaInfo(mediaId string) (result *newsMediaInfoResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaGetMaterialURL, accessToken)

	postData := new(mediaInfo)
	postData.MediaId = mediaId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetNewsMediaInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("获取永久图文素材错误", err)
	}
	return
}

//永久视频素材
type videoMediaInfoResult struct {
	util.WxError
	Title       string `json:"title"`
	Description string `json:"description"`
	DownUrl     string `json:"down_url"`
}

func (media *Media) GetVideoMediaInfo(mediaId string) (result *videoMediaInfoResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaGetMaterialURL, accessToken)

	postData := new(mediaInfo)
	postData.MediaId = mediaId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetVideoMediaInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("获取永久视频素材错误", err)
	}
	return
}

//永久其它素材 其他类型的素材消息，则响应的直接为素材的内容，开发者可以自行保存为文件

func (media *Media) GetOtherMediaInfo(mediaId string) (result []byte, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaGetMaterialURL, accessToken)

	postData := new(mediaInfo)
	postData.MediaId = mediaId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}
	result = response
	return
}

//删除永久素材
type delMaterialMedia struct {
	MediaId string `json:"media_id"`
}

func (media *Media) DelMaterialMedia(mediaId string) (result util.WxError, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaDelMaterialURL, accessToken)

	postData := new(delMaterialMedia)
	postData.MediaId = mediaId

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("DelMaterialMedia error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("删除永久素材错误", err)
	}
	return
}

//修改永久图文素材
type UpdateNewsMedia struct {
	MediaId  string `json:"media_id"` //要修改的图文消息的id
	Index    int32  `json:"index"`    //要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
	Articles struct {
		Title            	string 		`json:"title"`
		ThumbMediaId     	string 		`json:"thumb_media_id"`
		Author           	string 		`json:"author"`
		Digest           	string 		`json:"digest"`
		ShowCoverPic     	int32  		`json:"show_cover_pic"`
		Content          	string 		`json:"content"`
		ContentSourceUrl 	string 		`json:"content_source_url"`
		NeedOpenComment  	uint32		`json:"need_open_comment"`
		OnlyFansCanComment 	uint32		`json:"only_fans_can_comment"`
	} `json:"articles"`
}

func (media *Media) UpdateNewsMedia(postData *UpdateNewsMedia) (result util.WxError, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaUpdateNewsURL, accessToken)

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("UpdateNewsMedia error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("修改永久图文素材错误", err)
	}
	return
}

//获取永久素材总数
type materialMediaCount struct {
	util.WxError
	VoiceCount int32 `json:"voice_count"`
	VideoCount int32 `json:"video_count"`
	ImageCount int32 `json:"image_count"`
	NewsCount  int32 `json:"news_count"`
}

func (media *Media) GetMaterialMediaCount() (result materialMediaCount, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaGetMaterialCountURL, accessToken)

	response, err := util.HTTPGet(wxUrl, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetMaterialMediaCount error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("获取永久素材总数错误", err)
	}
	return
}

//获取永久素材列表

//获取图文永久素材列表

type getMediaList struct {
	Type   string `json:"type"`
	Offset int32  `json:"offset"`
	Count  int32  `json:"count"`
}

type newsMediaListResult struct {
	util.WxError
	TotalCount int32   `json:"total_count"`
	ItemCount  int32   `json:"item_count"`
	Item       []*item `json:"item"`
}

type item struct {
	MediaId string `json:"media_id"`
	Content struct {
		NewsItem []*newsItem `json:"news_item"`
	} `json:"content"`
	UpdateTime int32 `json:"update_time"`
}

type newsItem struct {
	Title            	string 		`json:"title"`          //标题
	ThumbMediaId     	string 		`json:"thumb_media_id"` //图文消息的封面图片素材id（必须是永久mediaID）
	Author           	string 		`json:"author"`         //作者
	Digest           	string 		`json:"digest"`         //图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
	ShowCoverPic     	int32  		`json:"show_cover_pic"` //是否显示封面，0为false，即不显示，1为true，即显示
	Content          	string 		`json:"content"`        //图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
	Url              	string 		`json:"url"`
	ContentSourceUrl 	string 		`json:"content_source_url"` //图文消息的原文地址，即点击“阅读原文”后的URL
	NeedOpenComment  	uint32		`json:"need_open_comment"`  //是否打开评论，0不打开，1打开
	OnlyFansCanComment 	uint32		`json:"only_fans_can_comment"`  //是否粉丝才可评论，0所有人可评论，1粉丝才可评论
}

//图文列表
func (media *Media) GetNewsMediaList(offset, count int32) (result *newsMediaListResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaBatchgetMaterialURL, accessToken)

	postData := new(getMediaList)
	postData.Type = "news"
	postData.Offset = offset
	postData.Count = count

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetNewsMediaList error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("获取图文永久素材列表错误", err)
	}
	return
}

//获取其他类型（图片、语音、视频）
type otherMediaListResult struct {
	util.WxError
	TotalCount int32        `json:"total_count"`
	ItemCount  int32        `json:"item_count"`
	Item       []*itemOther `json:"item"`
}

type itemOther struct {
	MediaId    string `json:"media_id"`
	Name       string `json:"name"`
	UpdateTime int32  `json:"update_time"`
	Url        string `json:"url"`
}

//其它列表
func (media *Media) GetOtherMediaList(mediaType string, offset, count int32) (result *otherMediaListResult, err error) {
	accessToken, err := media.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(MediaBatchgetMaterialURL, accessToken)

	postData := new(getMediaList)
	postData.Type = mediaType
	postData.Offset = offset
	postData.Count = count

	response, err := util.PostJSON(wxUrl, postData, media.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetOtherMediaList error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		media.WXLog.Error("获取其它永久素材列表错误", err)
	}
	return
}

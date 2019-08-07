package mass

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
)

// 预览接口
func (m *Mass) MassPreview(opts ...massPreviewOption) (res resMassPreview, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := new(reqMassPreview)
	for _, f := range opts {
		f(req)
	}

	uri := massPreviewUrl + commonToken + accessToken
	response, err := util.PostJSON(uri, req, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", res.ErrCode, res.ErrMsg)
	}
	return
}

type massPreviewOption func(req *reqMassPreview)

// 接收消息用户对应该公众号的openid
func WithPreviewTouserOption(touser string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Touser = touser
	}
}

// 接收消息用户对应该公众号的微信号
func WithPreviewTowxnameOption(towxname string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Towxname = towxname
	}
}

// 图文消息
func WithPreviewMpnewsOption(mediaID string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = mpnewsType
		req.Mpnews.MediaID = mediaID
	}
}

// 文本消息
func WithPreviewTextOption(content string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = textType
		req.Text.Content = content
	}
}

// 语音/音频
func WithPreviewVoiceOption(mediaID string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = voiceType
		req.Voice.MediaID = mediaID
	}
}

// 图片
func WithPreviewImageOption(mediaID string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = imageType
		req.Image.MediaID = mediaID
	}
}

// 视频
func WithPreviewMpvideoOption(mediaID string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = mpvideoType
		req.Mpvideo.MediaID = mediaID
	}
}

// 卡券中传入的cardEXT字符串，需要对CardEXT进行json后传入
func WithPreviewWxcardOption(cardID, cardEXT string) massPreviewOption {
	return func(req *reqMassPreview) {
		req.Msgtype = wxcardType
		req.Wxcard.CardID = cardID
		req.Wxcard.CardEXT = cardEXT
	}
}

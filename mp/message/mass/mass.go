package mass

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	baseUrl          = "https://api.weixin.qq.com/cgi-bin/message/mass"
	commonToken      = "?access_token="
	massSendallUrl   = baseUrl + "/sendall"
	massSendUrl      = baseUrl + "/send"
	massDelUrl       = baseUrl + "/delete"
	massPreviewUrl   = baseUrl + "/preview"
	massGetStatusUrl = baseUrl + "/get"
	massSpeedGetUrl  = baseUrl + "/speed/get"
	massSpeedSetUrl  = baseUrl + "/speed/set"

	massFlag = "群发消息"

	mpnewsType  = "mpnews"
	textType    = "text"
	voiceType   = "voice"
	imageType   = "image"
	mpvideoType = "mpvideo" // 视频
	wxcardType  = "wxcard"  // 卡券消息
)

type Mass struct {
	*core.Context
}

func NewMass(context *core.Context) *Mass {
	mass := new(Mass)
	mass.Context = context
	return mass
}

// 根据标签进行群发
func (m *Mass) MassSendall(opts ...massOption) (res resMass, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := new(reqMassSendall)
	for _, f := range opts {
		f(req)
	}

	uri := massSendallUrl + commonToken + accessToken
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
		m.WXLog.Error(massFlag, err)
	}
	return
}

// 根据OpenID列表群发
func (m *Mass) MassSend(opts ...massOption) (res resMass, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := new(reqMassSendall)
	for _, f := range opts {
		f(req)
	}

	uri := massSendUrl + commonToken + accessToken
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
		m.WXLog.Error(massFlag, err)
	}
	return
}

// 删除群发
func (m *Mass) MassDel(msgID, articleIDX int) (err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := reqMassDel{
		MsgID:      msgID,
		ArticleIDX: articleIDX,
	}

	uri := massDelUrl + commonToken + accessToken
	response, err := util.PostJSON(uri, req, m.ProxyUrl)
	if err != nil {
		return
	}
	var res util.WxError
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", res.ErrCode, res.ErrMsg)
		m.WXLog.Error(massFlag, err)
	}
	return
}

// 查看群发状态
func (m *Mass) MassGet(msgID string) (res resMassGet, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := reqMassGet{MsgID: msgID}

	uri := massGetStatusUrl + commonToken + accessToken
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
		m.WXLog.Error(massFlag, err)
	}
	return
}

//控制群发速度
type MassSpeedGetRes struct{
	util.WxError
	Speed	int		`json:"speed"`
	RealSpeed   int		`json:"realspeed"`
}

func (m *Mass) MassSpeedGet() (res MassSpeedGetRes, err error)  {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	uri := massSpeedGetUrl + commonToken + accessToken
	response, err := util.HTTPGet(uri, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", res.ErrCode, res.ErrMsg)
		m.WXLog.Error(massFlag, err)
	}
	return
}

//控制群发速度
type reqMassSpeedSet struct{
	Speed	int		`json:"speed"`
}
func (m *Mass) MassSpeedSet(speed int) (res util.WxError, err error)  {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	req := reqMassSpeedSet{Speed: speed}

	uri := massSpeedSetUrl + commonToken + accessToken
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
		m.WXLog.Error(massFlag, err)
	}
	return
}


type massOption func(req *reqMassSendall)

// 根据标签进行列表群发
func WithFilterOption(isToAll bool, tagId int) massOption {
	return func(req *reqMassSendall) {
		req.Filter.IsToAll = isToAll
		req.Filter.TagID = tagId
	}
}

// 根据OpenID列表群发
func WithTouserOption(openids []string) massOption {
	return func(req *reqMassSendall) {
		req.Touser = openids
	}
}

// 图文消息
func WithMpnewsOption(mediaID string, SendIgnoreReprint int) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = mpnewsType
		req.Mpnews.MediaID = mediaID
		req.SendIgnoreReprint = SendIgnoreReprint
	}
}

// 文本
func WithTextOption(content string) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = textType
		req.Text.Content = content
	}
}

// 语音/音频
func WithVoiceOption(mediaID string) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = voiceType
		req.Voice.MediaID = mediaID
	}
}

// 图片
func WithImageOption(mediaID string) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = imageType
		req.Image.MediaID = mediaID
	}
}

// 视频
func WithMpvideoOption(mediaID string) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = mpvideoType
		req.Mpvideo.MediaID = mediaID
	}
}

// 卡券消息
func WithWxcardOption(CardID string) massOption {
	return func(req *reqMassSendall) {
		req.Msgtype = wxcardType
		req.Wxcard.CardID = CardID
	}
}

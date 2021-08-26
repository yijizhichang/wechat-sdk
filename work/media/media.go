//客户联系-素材管理
package media

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	UploadQyTempMediaURL = "https://qyapi.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s"  //上传临时素材
	UploadQyImgMediaURL = "https://qyapi.weixin.qq.com/cgi-bin/media/uploadimg?access_token=%s"  //上传图片
	GetQyTempMediaURL = "https://qyapi.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"  //获取临时素材
	GetQyCustomerRemarkURL = "https://qyapi.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=%s&media_id=%s"  //获取高清语音素材
)

//Media 客户管理
type Media struct {
	*core.Context
}

//NewCustomer 实例化
func NewMedia(context *core.Context) *Media {
	m := new(Media)
	m.Context = context
	return m
}

//上传临时素材
type UploadQyTempMediaReq struct {
	util.WxError
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}
func (m *Media) UploadQyTempMedia(accessToken, fileType, fieldname, filename string)(result *UploadQyTempMediaReq, err error){
	qyUrl := fmt.Sprintf(UploadQyTempMediaURL, accessToken, fileType)

	response, err := util.PostFile(fieldname, filename ,qyUrl, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UploadQyTempMedia error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//上传图片
type UploadQyImgMediaReq struct {
	util.WxError
	Url     string `json:"url"`
}
func (m *Media) UploadQyImgMedia(accessToken, fieldname, filename string)(result *UploadQyImgMediaReq, err error){
	qyUrl := fmt.Sprintf(UploadQyImgMediaURL, accessToken)

	response, err := util.PostFile(fieldname, filename ,qyUrl, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UploadQyImgMedia error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取临时素材下载地址
func (m *Media) GetQyTempMediaURL(accessToken, mediaId string)(downloadUrl string, err error){
	downloadUrl = fmt.Sprintf(GetQyTempMediaURL, accessToken, mediaId)
	return
}

//获取获取高清语音素材下载地址
func (m *Media) GetQyCustomerRemarkURL(accessToken, mediaId string)(downloadUrl string, err error){
	downloadUrl = fmt.Sprintf(GetQyCustomerRemarkURL, accessToken, mediaId)
	return
}

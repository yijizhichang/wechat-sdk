//客户联系-企业服务人员管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCustomerFollowUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list?access_token=%s"  //获取配置了客户联系功能的成员列表
	CreateCustomerContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way?access_token=%s"  //配置客户联系「联系我」方式
	GetCustomerContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way?access_token=%s"  //获取企业已配置的「联系我」方式
	UpdateCustomerContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way?access_token=%s"  //更新企业已配置的「联系我」方式
	DelCustomerContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way?access_token=%s"  //删除企业已配置的「联系我」方式
	CloseCustomerTempChatURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat?access_token=%s"  //结束临时会话
)

//CustomerFollow 客户管理
type CustomerFollow struct {
	*core.Context
}

//NewCustomer 实例化
func NewCustomerFollow(context *core.Context) *CustomerFollow {
	cfw := new(CustomerFollow)
	cfw.Context = context
	return cfw
}

//获取配置了客户联系功能的成员列表
type followUserList struct {
	util.WxError
	FollowUser []string `json:"follow_user"`
}
func (cfw *CustomerFollow) GetCustomerFollowUserList(accessToken string)(result *followUserList, err error){
	qyUrl := fmt.Sprintf(GetCustomerFollowUserListURL, accessToken)

	response, err := util.HTTPGet(qyUrl, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerFollowUserList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//配置客户联系「联系我」方式
type CreateCustomerContactWayReq struct {
	Type          int32    `json:"type"`
	Scene         int32    `json:"scene"`
	Style         int32    `json:"style"`
	Remark        string   `json:"remark"`
	SkipVerify    bool     `json:"skip_verify"`
	State         string   `json:"state"`
	User          []string `json:"user"`
	Party         []int32  `json:"party"`
	IsTemp        bool     `json:"is_temp"`
	ExpiresIn     int32    `json:"expires_in"`
	ChatExpiresIn int32    `json:"chat_expires_in"`
	Unionid       string   `json:"unionid"`
	Conclusions   struct {
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		Image struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
		Link struct {
			Title  string `json:"title"`
			Picurl string `json:"picurl"`
			Desc   string `json:"desc"`
			Url    string `json:"url"`
		} `json:"link"`
		Miniprogram struct {
			Title      string `json:"title"`
			PicMediaId string `json:"pic_media_id"`
			Appid      string `json:"appid"`
			Page       string `json:"page"`
		} `json:"miniprogram"`
	} `json:"conclusions"`
}
type CreateCustomerContactWayRep struct {
	util.WxError
	ConfigId string `json:"config_id"`
	QrCode   string `json:"qr_code"`
}
func (cfw *CustomerFollow) CreateCustomerContactWay(accessToken string, req CreateCustomerContactWayReq)(result *CreateCustomerContactWayRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerContactWayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerContactWay error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取企业已配置的「联系我」方式
type CustomerContactWayReq struct {
	ConfigId string `json:"config_id"`
}
type CustomerContactWayRep struct {
	util.WxError
	ContactWay struct {
		ConfigId      string   `json:"config_id"`
		Type          int32    `json:"type"`
		Scene         int32    `json:"scene"`
		Style         int32    `json:"style"`
		Remark        string   `json:"remark"`
		SkipVerify    bool     `json:"skip_verify"`
		State         string   `json:"state"`
		QrCode        string   `json:"qr_code"`
		User          []string `json:"user"`
		Party         []int32  `json:"party"`
		IsTemp        bool     `json:"is_temp"`
		ExpiresIn     int32    `json:"expires_in"`
		ChatExpiresIn int32    `json:"chat_expires_in"`
		Unionid       string   `json:"unionid"`
		Conclusions   struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
			Image struct {
				PicUrl string `json:"pic_url"`
			} `json:"image"`
			Link struct {
				Title  string `json:"title"`
				Picurl string `json:"picurl"`
				Desc   string `json:"desc"`
				Url    string `json:"url"`
			} `json:"link"`
			Miniprogram struct {
				Title      string `json:"title"`
				PicMediaId string `json:"pic_media_id"`
				Appid      string `json:"appid"`
				Page       string `json:"page"`
			} `json:"miniprogram"`
		} `json:"conclusions"`
	} `json:"contact_way"`
}
func (cfw *CustomerFollow) GetCustomerContactWay(accessToken string, req CustomerContactWayReq)(result *CustomerContactWayRep, err error){
	qyUrl := fmt.Sprintf(GetCustomerContactWayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerContactWay error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新企业已配置的「联系我」方式
type UpdateCustomerContactWayReq struct {
	ConfigId      string   `json:"config_id"`
	Remark        string   `json:"remark"`
	SkipVerify    bool     `json:"skip_verify"`
	Style         int32    `json:"style"`
	State         string   `json:"state"`
	User          []string `json:"user"`
	Party         []int32  `json:"party"`
	ExpiresIn     int32    `json:"expires_in"`
	ChatExpiresIn int32    `json:"chat_expires_in"`
	Unionid       string   `json:"unionid"`
	Conclusions   struct {
		Text struct {
			Content string `json:"content"`
		} `json:"text"`
		Image struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
		Link struct {
			Title  string `json:"title"`
			Picurl string `json:"picurl"`
			Desc   string `json:"desc"`
			Url    string `json:"url"`
		} `json:"link"`
		Miniprogram struct {
			Title      string `json:"title"`
			PicMediaId string `json:"pic_media_id"`
			Appid      string `json:"appid"`
			Page       string `json:"page"`
		} `json:"miniprogram"`
	} `json:"conclusions"`
}
func (cfw *CustomerFollow) UpdateCustomerContactWay(accessToken string, req UpdateCustomerContactWayReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateCustomerContactWayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateCustomerContactWay error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除企业已配置的「联系我」方式
type DelCustomerContactWayReq struct {
	ConfigId string `json:"config_id"`
}
func (cfw *CustomerFollow) DelCustomerContactWay(accessToken string, req DelCustomerContactWayReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelCustomerContactWayURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCustomerContactWay error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//结束临时会话
type CloseCustomerTempChatReq struct {
	Userid         string `json:"userid"`
	ExternalUserid string `json:"external_userid"`
}
func (cfw *CustomerFollow) CloseCustomerTempChat(accessToken string, req CloseCustomerTempChatReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(CloseCustomerTempChatURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cfw.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CloseCustomerTempChat error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

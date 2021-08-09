//客户联系-客户管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetQyCustomerListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list?access_token=%s&userid=%s"  //获取客户列表
	GetQyCustomerViewURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s"  //获取客户详情
	GetQyCustomerViewBatchURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/batch/get_by_user?access_token=%s"  //批量获取客户详情
	UpdateQyCustomerRemark = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token=%s"  //修改客户备注信息
)

//Customer 客户管理
type Customer struct {
	*core.Context
}

//NewCustomer 实例化
func NewCustomer(context *core.Context) *Customer {
	cus := new(Customer)
	cus.Context = context
	return cus
}

//获取客户列表
type cusList struct {
	util.WxError
	ExternalUserid []string `json:"external_userid"`
}
func (c *Customer) GetQyCustomerList(accessToken, userid string)(result *cusList, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerListURL, accessToken, userid)

	response, err := util.HTTPGet(qyUrl, c.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取客户详情
type cusView struct {
	util.WxError
	ExternalContact struct{
		ExternalUserid  string  `json:"external_userid"`
		Name string `json:"name"`
		Position string `json:"position"`
		Avatar string `json:"avatar"`
		CorpName string `json:"corp_name"`
		CorpFullName string `json:"corp_full_name"`
		Type int64 `json:"type"`
		Gender int64 `json:"gender"`
		Unionid string `json:"unionid"`
		ExternalProfile struct{
			ExternalAttr []externalAttrItem `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowUser []followUserItem `json:"follow_user"`
	NextCursor string `json:"next_cursor"`
}

type externalAttrItem struct {
	Type int64 `json:"type"`
	Name string `json:"name"`
	Text struct{
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct{
		Url string `json:"url"`
		Title string `json:"title"`
	} `json:"web,omitempty"`
	Miniprogram struct{
		Appid string `json:"appid"`
		Pagepath string `json:"pagepath"`
		Title string `json:"title"`
	} `json:"miniprogram,omitempty"`
}

type followUserItem struct {
	Userid string `json:"userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	Createtime int64 `json:"createtime"`
	Tags []tagItem `json:"tags"`
	RemarkCorpName string `json:"remark_corp_name"`
	RemarkMobiles []string `json:"remark_mobiles"`
	OperUserid string `json:"oper_userid"`
	AddWay int64 `json:"add_way"`
	State string `json:"state"`
}

type tagItem struct {
	GroupName string `json:"group_name"`
	TagName string `json:"tag_name"`
	TagId string `json:"tag_id"`
	Type int64 `json:"type"`
}
func (c *Customer) GetQyCustomerView(accessToken, externalUserid, cursor string)(result *cusView, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerViewURL, accessToken, externalUserid, cursor)

	response, err := util.HTTPGet(qyUrl, c.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//批量获取客户详情
type cusViewBatch struct {
	util.WxError
	ExternalContactList []externalContactItem `json:"external_contact_list"`
}

type externalContactItem struct {
	ExternalContact struct{
		ExternalUserid  string  `json:"external_userid"`
		Name string `json:"name"`
		Position string `json:"position"`
		Avatar string `json:"avatar"`
		CorpName string `json:"corp_name"`
		CorpFullName string `json:"corp_full_name"`
		Type int64 `json:"type"`
		Gender int64 `json:"gender"`
		Unionid string `json:"unionid"`
		ExternalProfile struct{
			ExternalAttr []externalAttrItem `json:"external_attr"`
		} `json:"external_profile"`
	} `json:"external_contact"`
	FollowInfo followInfoItem `json:"follow_info"`
}

type followInfoItem struct {
	Userid string `json:"userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	Createtime int64 `json:"createtime"`
	Tags []tagItem `json:"tags"`
	RemarkCorpName string `json:"remark_corp_name"`
	RemarkMobiles []string `json:"remark_mobiles"`
	OperUserid string `json:"oper_userid"`
	AddWay int64 `json:"add_way"`
	State string `json:"state"`
}

type CusViewBatchReq struct {
	Userid  string  `json:"userid"`
	Cursor  string  `json:"cursor"`
	Limit   int     `json:"limit"`
}
func (c *Customer) GetQyCustomerViewBatch(accessToken, req CusViewBatchReq)(result *cusViewBatch, err error){
	qyUrl := fmt.Sprintf(GetQyCustomerViewBatchURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyCustomerViewBatch error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//修改客户备注信息
type UpdateCusRemarkReq struct {
	Userid  string 	`json:"userid"`
	ExternalUserid string `json:"external_userid"`
	Remark string `json:"remark"`
	Description string `json:"description"`
	RemarkCompany string `json:"remark_company"`
	RemarkMobiles []string `json:"remark_mobiles"`
	RemarkPicMediaid string `json:"remark_pic_mediaid"`
}

func (c *Customer) UpdateQyCustomerRemark(accessToken, req UpdateCusRemarkReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateQyCustomerRemark, accessToken)

	response, err := util.PostJSON(qyUrl, req, c.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateQyCustomerRemark error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
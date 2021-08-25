//客户联系-客户标签管理
package customer

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCustomerTagListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list?access_token=%s"  //获取企业标签库
	CreateCustomerTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag?access_token=%s"  //添加企业客户标签
	UpdateCustomerTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag?access_token=%s"  //编辑企业客户标签
	DelCustomerTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag?access_token=%s"  //删除企业客户标签
	MarkCustomerTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag?access_token=%s"  //编辑客户企业标签
	GetCustomerStrategyTagListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_strategy_tag_list?access_token=%s"  //获取指定规则组下的企业客户标签
	CreateCustomerStrategyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_strategy_tag?access_token=%s"  //为指定规则组创建企业客户标签
	UpdateCustomerStrategyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_strategy_tag?access_token=%s"  //编辑指定规则组下的企业客户标签
	DelCustomerStrategyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_strategy_tag?access_token=%s"  //删除指定规则组下的企业客户标签
)

//CustomerTag 客户管理
type CustomerTag struct {
	*core.Context
}

//NewCustomer 实例化
func NewCustomerTag(context *core.Context) *CustomerTag {
	cst := new(CustomerTag)
	cst.Context = context
	return cst
}

//获取企业标签库
type CusTagReq struct {
	TagId []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
}
type cusTagList struct {
	util.WxError
	TagGroup []tagGroupItem `json:"tag_group"`
}
type tagGroupItem struct {
	GroupId string `json:"group_id"`
	GroupName string `json:"group_name"`
	CreateTime int64 `json:"create_time"`
	Order int32 `json:"order"`
	Deleted bool `json:"deleted"`
	Tag []cusTagItem `json:"tag"`
}
type cusTagItem struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreateTime int64 `json:"create_time"`
	Order int32 `json:"order"`
	Deleted bool `json:"deleted"`
}
func (ct *CustomerTag) GetCustomerTagList(accessToken string, req CusTagReq)(result *cusTagList, err error){
	qyUrl := fmt.Sprintf(GetCustomerTagListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerTagList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//添加企业客户标签
type CreateCusTagReq struct {
	GroupId string `json:"group_id"`
	GroupName string `json:"group_name"`
	Order int32 `json:"order"`
	Tag []AddTag `json:"tag"`
	Agentid int32 `json:"agentid"`
}
type AddTag struct {
	Name string `json:"name"`
	Order int32 `json:"order"`
}
type CreateCusTagRep struct {
	util.WxError
	TagGroup struct{
		GroupId string `json:"group_id"`
		GroupName string `json:"group_name"`
		CreateTime int64 `json:"create_time"`
		Order int32 `json:"order"`
		Tag []AddTagRep `json:"tag"`
	} `json:"tag_group"`
}
type AddTagRep struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CreateTime int64 `json:"create_time"`
	Order int32 `json:"order"`
}
func (ct *CustomerTag) CreateCustomerTag(accessToken string, req CreateCusTagReq)(result *CreateCusTagRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑企业客户标签
type UpdateCusTagReq struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Order int32 `json:"order"`
	Agentid int32 `json:"agentid"`
}
func (ct *CustomerTag) UpdateCustomerTag(accessToken string, req UpdateCusTagReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateCustomerTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateCustomerTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除企业客户标签
type DelCusTagReq struct {
	TagId []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
	Agentid int32 `json:"agentid"`
}
func (ct *CustomerTag) DelCustomerTag(accessToken string, req DelCusTagReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelCustomerTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCustomerTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑客户企业标签
type MarkTagReq struct {
	Userid string `json:"userid"`
	ExternalUserid string `json:"external_userid"` //请确保external_userid是userid的外部联系人
	AddTag []string `json:"add_tag"`
	RemoveTag []string `json:"remove_tag"`
}
func (ct *CustomerTag) MarkCustomerTag(accessToken string, req MarkTagReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(MarkCustomerTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("MarkCustomerTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取指定规则组下的企业客户标签
type CustomerStrategyTagReq struct {
	StrategyId int32    `json:"strategy_id"`
	TagId      []string `json:"tag_id"`
	GroupId    []string `json:"group_id"`
}
type customerStrategyTagList struct {
	util.WxError
	TagGroup []struct {
		GroupId    string `json:"group_id"`
		GroupName  string `json:"group_name"`
		CreateTime int64  `json:"create_time"`
		Order      int32  `json:"order"`
		StrategyId int32  `json:"strategy_id"`
		Tag        []struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			CreateTime int64  `json:"create_time"`
			Order      int32  `json:"order"`
		} `json:"tag"`
	} `json:"tag_group"`
}
func (ct *CustomerTag) GetCustomerStrategyTagList(accessToken string, req CustomerStrategyTagReq)(result *customerStrategyTagList, err error){
	qyUrl := fmt.Sprintf(GetCustomerStrategyTagListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCustomerStrategyTagList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//为指定规则组创建企业客户标签
type CreateCustomerStrategyTagReq struct {
	StrategyId int32  `json:"strategy_id"`
	GroupId    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	Order      int    `json:"order"`
	Tag        []struct {
		Name  string `json:"name"`
		Order int32  `json:"order"`
	} `json:"tag"`
}
type createCustomerStrategyTagRep struct {
	util.WxError
	TagGroup struct {
		GroupId    string `json:"group_id"`
		GroupName  string `json:"group_name"`
		CreateTime int32  `json:"create_time"`
		Order      int32  `json:"order"`
		Tag        []struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			CreateTime int32  `json:"create_time"`
			Order      int32  `json:"order"`
		} `json:"tag"`
	} `json:"tag_group"`
}
func (ct *CustomerTag) CreateCustomerStrategyTag(accessToken string, req CreateCustomerStrategyTagReq)(result *createCustomerStrategyTagRep, err error){
	qyUrl := fmt.Sprintf(CreateCustomerStrategyTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCustomerStrategyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//编辑指定规则组下的企业客户标签
type UpdateCustomerStrategyTagReq struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Order int32  `json:"order"`
}
func (ct *CustomerTag) UpdateCustomerStrategyTag(accessToken string, req UpdateCustomerStrategyTagReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(UpdateCustomerStrategyTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateCustomerStrategyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除指定规则组下的企业客户标签
type DelCustomerStrategyTagReq struct {
	TagId   []string `json:"tag_id"`
	GroupId []string `json:"group_id"`
}
func (ct *CustomerTag) DelCustomerStrategyTag(accessToken string, req DelCustomerStrategyTagReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelCustomerStrategyTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCustomerStrategyTagURL error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

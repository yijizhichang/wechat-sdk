//通讯录管理-标签管理
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetCompanyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/list?access_token=%s"  //获取标签列表
	CreateCompanyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=%s"  //创建标签
	UpdateCompanyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=%s"  //更新标签
	DelCompanyTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=%s&tagid=%d"  //删除标签
	GetCompanyTagUserURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=%s&tagid=%d"  //获取标签成员
	CreateCompanyTagUserURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=%s"  //增加标签成员
	DelCompanyTagUserURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=%s"  //删除标签成员
)

//CompanyTag 标签管理
type CompanyTag struct {
	*core.Context
}

//NewCompanyTag 实例化
func NewCompanyTag(context *core.Context) *CompanyTag {
	ct := new(CompanyTag)
	ct.Context = context
	return ct
}

//获取标签列表
type CompanyTagList struct {
	util.WxError
	Taglist []struct {
		Tagid   int32  `json:"tagid"`
		Tagname string `json:"tagname"`
	} `json:"taglist"`
}
func (ct *CompanyTag) GetCompanyTag (accessToken string) (result *CompanyTagList, err error) {
	qyUrl := fmt.Sprintf(GetCompanyTagURL, accessToken)

	response, err := util.HTTPGet(qyUrl, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCompanyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//创建标签
type CreateCompanyTagReq struct {
	Tagname string `json:"tagname"`
	Tagid   int32  `json:"tagid"`
}
type CreateCompanyTagRep struct {
	util.WxError
	Tagid   int32    `json:"tagid"`
}
func (ct *CompanyTag) CreateCompanyTag (accessToken string, req CreateCompanyTagReq) (result *CreateCompanyTagRep, err error) {
	qyUrl := fmt.Sprintf(CreateCompanyTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCompanyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新标签
type UpdateCompanyTagReq struct {
	Tagname string `json:"tagname"`
	Tagid   int32  `json:"tagid"`
}
func (ct *CompanyTag) UpdateCompanyTag (accessToken string, req UpdateCompanyTagReq) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(UpdateCompanyTagURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateCompanyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除标签
func (ct *CompanyTag) DelCompanyTag (accessToken string, tagId int32) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(DelCompanyTagURL, accessToken, tagId)

	response, err := util.HTTPGet(qyUrl, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCompanyTag error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取标签成员
type CompanyTagUserList struct {
	util.WxError
	Tagname  string `json:"tagname"`
	Userlist []struct {
		Userid string `json:"userid"`
		Name   string `json:"name"`
	} `json:"userlist"`
	Partylist []int32 `json:"partylist"`
}
func (ct *CompanyTag) GetCompanyTagUser (accessToken string, tagId int32) (result *CompanyTagUserList, err error) {
	qyUrl := fmt.Sprintf(GetCompanyTagUserURL, accessToken, tagId)

	response, err := util.HTTPGet(qyUrl, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetCompanyTagUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//增加标签成员
type CreateCompanyTagUserReq struct {
	Tagid     int32      `json:"tagid"`
	Userlist  []string   `json:"userlist"`
	Partylist []int32    `json:"partylist"`
}
type CreateCompanyTagUserRep struct {
	util.WxError
	Invalidlist  string `json:"invalidlist,omitempty"`
	Invalidparty []int32 `json:"invalidparty,omitempty"`
}
func (ct *CompanyTag) CreateCompanyTagUser (accessToken string, req CreateCompanyTagUserReq) (result *CreateCompanyTagUserRep, err error) {
	qyUrl := fmt.Sprintf(CreateCompanyTagUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateCompanyTagUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除标签成员
type DelCompanyTagUserReq struct {
	Tagid     int32    `json:"tagid"`
	Userlist  []string `json:"userlist"`
	Partylist []int32  `json:"partylist"`
}
type DelCompanyTagUserRep struct {
	util.WxError
	Invalidlist  string `json:"invalidlist,omitempty"`
	Invalidparty []int32 `json:"invalidparty,omitempty"`
}
func (ct *CompanyTag) DelCompanyTagUser (accessToken string, req DelCompanyTagUserReq) (result *DelCompanyTagUserRep, err error) {
	qyUrl := fmt.Sprintf(DelCompanyTagUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ct.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelCompanyTagUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//通讯录管理-互联企业
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetLinkedCorpPermListURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/agent/get_perm_list?access_token=%s"  //获取应用的可见范围
	GetLinkedCorpUserInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/get?access_token=%s"  //获取互联企业成员详细信息
	GetLinkedCorpUserSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/simplelist?access_token=%s"  //获取互联企业部门成员
	GetLinkedCorpUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/user/list?access_token=%s"  //获取互联企业部门成员详情
	GetLinkedCorpDepartmentListURL = "https://qyapi.weixin.qq.com/cgi-bin/linkedcorp/department/list?access_token=%s"  //获取互联企业部门列表


)

//CompanyLinkedCorp
type CompanyLinkedCorp struct {
	*core.Context
}

//NewCompanyLinkedCorp 实例化
func NewCompanyLinkedCorp(context *core.Context) *CompanyLinkedCorp {
	clc := new(CompanyLinkedCorp)
	clc.Context = context
	return clc
}

//获取应用的可见范围
type LinkedCorpPermReq struct {

}
type LinkedCorpPermRep struct {
	util.WxError
	Userids       []string `json:"userids"`
	DepartmentIds []string `json:"department_ids"`
}
func (clc *CompanyLinkedCorp) GetLinkedCorpPermList (accessToken string, req LinkedCorpPermReq) (result *LinkedCorpPermRep, err error) {
	qyUrl := fmt.Sprintf(GetLinkedCorpPermListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, clc.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLinkedCorpPermList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取互联企业成员详细信息
type LinkedCorpUserInfoReq struct {
	Userid string `json:"userid"`
}
type LinkedCorpUserInfoRep struct {
	util.WxError
	UserInfo struct {
		Userid     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		Mobile     string   `json:"mobile"`
		Telephone  string   `json:"telephone"`
		Email      string   `json:"email"`
		Position   string   `json:"position"`
		Corpid     string   `json:"corpid"`
		Extattr    struct {
			Attrs []struct {
				Name  string `json:"name"`
				Value string `json:"value,omitempty"`
				Type  int32  `json:"type"`
				Text  struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
			} `json:"attrs"`
		} `json:"extattr"`
	} `json:"user_info"`
}
func (clc *CompanyLinkedCorp) GetLinkedCorpUserInfo (accessToken string, req LinkedCorpUserInfoReq) (result *LinkedCorpUserInfoRep, err error) {
	qyUrl := fmt.Sprintf(GetLinkedCorpUserInfoURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, clc.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLinkedCorpUserInfo error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取互联企业部门成员
type LinkedCorpUserSimpleListReq struct {
	DepartmentId string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}
type LinkedCorpUserSimpleListRep struct {
	util.WxError
	Userlist []struct {
		Userid     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		Corpid     string   `json:"corpid"`
	} `json:"userlist"`
}
func (clc *CompanyLinkedCorp) GetLinkedCorpUserSimpleList (accessToken string, req LinkedCorpUserSimpleListReq) (result *LinkedCorpUserSimpleListRep, err error) {
	qyUrl := fmt.Sprintf(GetLinkedCorpUserSimpleListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, clc.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLinkedCorpUserSimpleList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取互联企业部门成员详情
type LinkedCorpUserListReq struct {
	DepartmentId string `json:"department_id"`
	FetchChild   bool   `json:"fetch_child"`
}
type LinkedCorpUserListRep struct {
	util.WxError
	Userlist []struct {
		Userid     string   `json:"userid"`
		Name       string   `json:"name"`
		Department []string `json:"department"`
		Mobile     string   `json:"mobile"`
		Telephone  string   `json:"telephone"`
		Email      string   `json:"email"`
		Position   string   `json:"position"`
		Corpid     string   `json:"corpid"`
		Extattr    struct {
			Attrs []struct {
				Name  string `json:"name"`
				Value string `json:"value,omitempty"`
				Type  int32  `json:"type"`
				Text  struct {
					Value string `json:"value"`
				} `json:"text,omitempty"`
				Web struct {
					Url   string `json:"url"`
					Title string `json:"title"`
				} `json:"web,omitempty"`
			} `json:"attrs"`
		} `json:"extattr"`
	} `json:"userlist"`
}
func (clc *CompanyLinkedCorp) GetLinkedCorpUserList (accessToken string, req LinkedCorpUserListReq) (result *LinkedCorpUserListRep, err error) {
	qyUrl := fmt.Sprintf(GetLinkedCorpUserListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, clc.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLinkedCorpUserList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取互联企业部门列表
type LinkedCorpDepartmentListReq struct {
	DepartmentId string `json:"department_id"`
}
type LinkedCorpDepartmentListRep struct {
	util.WxError
	DepartmentList []struct {
		DepartmentId   string `json:"department_id"`
		DepartmentName string `json:"department_name"`
		Parentid       string `json:"parentid"`
		Order          int32  `json:"order"`
	} `json:"department_list"`
}
func (clc *CompanyLinkedCorp) GetLinkedCorpDepartmentList (accessToken string, req LinkedCorpDepartmentListReq) (result *LinkedCorpDepartmentListRep, err error) {
	qyUrl := fmt.Sprintf(GetLinkedCorpDepartmentListURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, clc.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetLinkedCorpDepartmentList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

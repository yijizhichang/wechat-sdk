//通讯录管理-部门管理
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=%s"  //创建部门
	UpdateDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=%s"  //更新部门
	DelDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=%s&id=%d"  //删除部门
	GetDepartmentListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=%s&id=%d"  //获取部门列表
)

//Department 客户管理
type Department struct {
	*core.Context
}

//NewDepartment 实例化
func NewDepartment(context *core.Context) *Department {
	dep := new(Department)
	dep.Context = context
	return dep
}

//创建部门
type CreateDepartmentReq struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	NameEn string `json:"name_en"`
	Parentid int32 `json:"parentid"`
	Order int32 `json:"order"`
}

type createDepartmentRep struct {
	util.WxError
	Id int32 `json:"id"`
}
func (d *Department) CreateDepartment (accessToken string, req CreateDepartmentReq) (result *createDepartmentRep, err error) {
	qyUrl := fmt.Sprintf(CreateDepartmentURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateDepartment error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//更新部门
func (d *Department) UpdateDepartment (accessToken string, req CreateDepartmentReq) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(UpdateDepartmentURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("UpdateDepartment error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除部门
func (d *Department) DelDepartment (accessToken string, id int32) (result *util.WxError, err error) {
	qyUrl := fmt.Sprintf(DelDepartmentURL, accessToken, id)

	response, err := util.HTTPGet(qyUrl, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelDepartment error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取部门列表
type departmentListRep struct {
	util.WxError
	Department []departmentItem `json:"department"`
}
type departmentItem struct {
	Id  int32  `json:"id"`
	Name string `json:"name"`
	NameEn string `json:"name_en"`
	Parentid int32 `json:"parentid"`
	Order int32 `json:"order"`
}
func (d *Department) GetDepartmentList (accessToken string, id int32) (result *departmentListRep, err error) {
	qyUrl := fmt.Sprintf(GetDepartmentListURL, accessToken, id)

	response, err := util.HTTPGet(qyUrl, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDepartmentList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
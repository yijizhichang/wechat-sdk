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
	GetDepartmentSimpleListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/simplelist?access_token=%s&id=%d"  //获取子部门ID列表
	GetDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/get?access_token=%s&id=%d"  //获取单个部门详情

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
	Id int32 `json:"id,omitempty"`
	Name string `json:"name"`
	NameEn string `json:"name_en"`
	Parentid int32 `json:"parentid"`
	Order int32 `json:"order"`
}

type CreateDepartmentRep struct {
	util.WxError
	Id int32 `json:"id"`
}
func (d *Department) CreateDepartment (accessToken string, req CreateDepartmentReq) (result *CreateDepartmentRep, err error) {
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
type DepartmentListRep struct {
	util.WxError
	Department []departmentItem `json:"department"`
}
type departmentItem struct {
	Id  int32  `json:"id"`
	Name string `json:"name"`
	NameEn string `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	Parentid int32 `json:"parentid"`
	Order int32 `json:"order"`
}
func (d *Department) GetDepartmentList (accessToken string, id int32) (result *DepartmentListRep, err error) {
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


//获取子部门ID列表
type DepartmentSimpleListRep struct {
	util.WxError
	DepartmentId []struct {
		Id       int32 `json:"id"`
		Parentid int32 `json:"parentid"`
		Order    int32 `json:"order"`
	} `json:"department_id"`
}
func (d *Department) GetDepartmentSimpleList (accessToken string, id int32) (result *DepartmentSimpleListRep, err error) {
	qyUrl := fmt.Sprintf(GetDepartmentSimpleListURL, accessToken, id)

	response, err := util.HTTPGet(qyUrl, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDepartmentSimpleList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取单个部门详情
type DepartmentRep struct {
	util.WxError
	Department struct {
		Id               int32    `json:"id"`
		Name             string   `json:"name"`
		NameEn           string   `json:"name_en"`
		DepartmentLeader []string `json:"department_leader"`
		Parentid         int32    `json:"parentid"`
		Order            int32    `json:"order"`
	} `json:"department"`
}
func (d *Department) GetDepartment (accessToken string, id int32) (result *DepartmentRep, err error) {
	qyUrl := fmt.Sprintf(GetDepartmentURL, accessToken, id)

	response, err := util.HTTPGet(qyUrl, d.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetDepartment error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}
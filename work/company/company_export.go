//通讯录管理-异步导出接口
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	QyExportSimpleUserURL = "https://qyapi.weixin.qq.com/cgi-bin/export/simple_user?access_token=%s"  //导出成员
	QyExportUserURL = "https://qyapi.weixin.qq.com/cgi-bin/export/user?access_token=%s"  //导出成员详情
	QyExportDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/export/department?access_token=%s"  //导出部门
	QyExportTagUserURL = "https://qyapi.weixin.qq.com/cgi-bin/export/taguser?access_token=%s"  //导出标签成员
	QyExportResultURL = "https://qyapi.weixin.qq.com/cgi-bin/export/get_result?access_token=%s&jobid=%s"  //获取导出结果
)

//CompanyExport 客户管理
type CompanyExport struct {
	*core.Context
}

//NewDepartment 实例化
func NewCompanyExport(context *core.Context) *CompanyExport {
	ce := new(CompanyExport)
	ce.Context = context
	return ce
}
//导出成员
type ExportSimpleUserReq struct {
	EncodingAeskey string `json:"encoding_aeskey"`
	BlockSize      int32  `json:"block_size"`
}
type ExportSimpleUserRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (ce *CompanyExport) QyExportSimpleUser (accessToken string, req ExportSimpleUserReq) (result *ExportSimpleUserRep, err error) {
	qyUrl := fmt.Sprintf(QyExportSimpleUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ce.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyExportSimpleUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//导出成员详情
type ExportUserReq struct {
	EncodingAeskey string `json:"encoding_aeskey"`
	BlockSize      int    `json:"block_size"`
}
type ExportUserRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (ce *CompanyExport) QyExportUser (accessToken string, req ExportUserReq) (result *ExportUserRep, err error) {
	qyUrl := fmt.Sprintf(QyExportUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ce.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyExportUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//导出部门
type ExportDepartmentReq struct {
	EncodingAeskey string `json:"encoding_aeskey"`
	BlockSize      int64    `json:"block_size"`
}
type ExportDepartmentRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (ce *CompanyExport) QyExportDepartment (accessToken string, req ExportDepartmentReq) (result *ExportDepartmentRep, err error) {
	qyUrl := fmt.Sprintf(QyExportDepartmentURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ce.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyExportDepartment error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//导出标签成员
type ExportTagUserReq struct {
	Tagid          int32  `json:"tagid"`
	EncodingAeskey string `json:"encoding_aeskey"`
	BlockSize      int32  `json:"block_size"`
}
type ExportTagUserRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (ce *CompanyExport) QyExportTagUser (accessToken string, req ExportTagUserReq) (result *ExportTagUserRep, err error) {
	qyUrl := fmt.Sprintf(QyExportTagUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, ce.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyExportTagUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取导出结果
type ExportResultRep struct {
	util.WxError
	Status   int32    `json:"status"`
	DataList []struct {
		Url  string      `json:"url"`
		Size interface{} `json:"size"`
		Md5  string      `json:"md5"`
	} `json:"data_list"`
}
func (ce *CompanyExport) QyExportResult (accessToken string, jobId string) (result *ExportResultRep, err error) {
	qyUrl := fmt.Sprintf(QyExportResultURL, accessToken, jobId)

	response, err := util.HTTPGet(qyUrl, ce.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("QyExportResult error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

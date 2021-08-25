//通讯录管理-异步批量接口
package company

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateBatchSyncUserURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/syncuser?access_token=%s"  //增量更新成员
	CreateBatchReplaceUserURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceuser?access_token=%s"  //全量覆盖成员
	CreateBatchReplacePartyURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty?access_token=%s"  //全量覆盖部门
	GetBatchResultURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult?access_token=%s&jobid=%s"  //获取异步任务结果
)

//CompanyBatch
type CompanyBatch struct {
	*core.Context
}

//NewCompanyBatch 实例化
func NewCompanyBatch(context *core.Context) *CompanyBatch {
	cb := new(CompanyBatch)
	cb.Context = context
	return cb
}

//增量更新成员
type CreateBatchSyncUserReq struct {
	MediaId  string `json:"media_id"`
	ToInvite bool   `json:"to_invite"`
	Callback struct {
		Url            string `json:"url"`
		Token          string `json:"token"`
		Encodingaeskey string `json:"encodingaeskey"`
	} `json:"callback"`
}
type createBatchSyncUserRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (cb *CompanyBatch) CreateBatchSyncUser (accessToken string, req CreateBatchSyncUserReq) (result *createBatchSyncUserRep, err error) {
	qyUrl := fmt.Sprintf(CreateBatchSyncUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cb.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateBatchSyncUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//全量覆盖成员
type CreateBatchReplaceUserReq struct {
	MediaId  string `json:"media_id"`
	ToInvite bool   `json:"to_invite"`
	Callback struct {
		Url            string `json:"url"`
		Token          string `json:"token"`
		Encodingaeskey string `json:"encodingaeskey"`
	} `json:"callback"`
}
type createBatchReplaceUserRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (cb *CompanyBatch) CreateBatchReplaceUser (accessToken string, req CreateBatchReplaceUserReq) (result *createBatchReplaceUserRep, err error) {
	qyUrl := fmt.Sprintf(CreateBatchReplaceUserURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cb.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateBatchReplaceUser error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//全量覆盖部门
type CreateBatchReplacePartyReq struct {
	MediaId  string `json:"media_id"`
	Callback struct {
		Url            string `json:"url"`
		Token          string `json:"token"`
		Encodingaeskey string `json:"encodingaeskey"`
	} `json:"callback"`
}
type createBatchReplacePartyRep struct {
	util.WxError
	Jobid   string `json:"jobid"`
}
func (cb *CompanyBatch) CreateBatchReplaceParty (accessToken string, req CreateBatchReplacePartyReq) (result *createBatchReplacePartyRep, err error) {
	qyUrl := fmt.Sprintf(CreateBatchReplacePartyURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, cb.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateBatchReplaceParty error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取异步任务结果
type batchResultRep struct {
	util.WxError
	Status     int32    `json:"status"`
	Type       string `json:"type"`
	Total      int32    `json:"total"`
	Percentage int32    `json:"percentage"`
	Result     []struct {
		Userid  string `json:"userid,omitempty"`
		Errcode int32    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Action  int32    `json:"action,omitempty"`
		Partyid int32    `json:"partyid,omitempty"`
	} `json:"result"`
}
func (cb *CompanyBatch) GetBatchResult (accessToken string, jobid string) (result *batchResultRep, err error) {
	qyUrl := fmt.Sprintf(GetBatchResultURL, accessToken, jobid)

	response, err := util.HTTPGet(qyUrl, cb.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetBatchResult error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

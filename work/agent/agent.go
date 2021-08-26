//应用管理
package agent

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	GetQyAgentViewURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/get?access_token=%s&agentid=%d"  //获取指定的应用详情
	GetQyAgentListURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/list?access_token=%s"  //获取access_token对应的应用列表
	SetQyAgentURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/set?access_token=%s"  //设置应用

)

type Agent struct {
	*core.Context
}

func NewAgent(context *core.Context) *Agent {
	a := new(Agent)
	a.Context = context
	return a
}

//获取指定的应用详情
type QyAgentViewRep struct {
	util.WxError
	Agentid        int32  `json:"agentid"`
	Name           string `json:"name"`
	SquareLogoUrl  string `json:"square_logo_url"`
	Description    string `json:"description"`
	AllowUserinfos struct {
		User []struct {
			Userid string `json:"userid"`
		} `json:"user"`
	} `json:"allow_userinfos"`
	AllowPartys struct {
		Partyid []int32 `json:"partyid"`
	} `json:"allow_partys"`
	AllowTags struct {
		Tagid []int32 `json:"tagid"`
	} `json:"allow_tags"`
	Close              int32  `json:"close"`
	RedirectDomain     string `json:"redirect_domain"`
	ReportLocationFlag int32  `json:"report_location_flag"`
	Isreportenter      int32  `json:"isreportenter"`
	HomeUrl            string `json:"home_url"`
}
func (a *Agent) GetQyAgentView(accessToken string, agentId int32)(result *QyAgentViewRep, err error){
	qyUrl := fmt.Sprintf(GetQyAgentViewURL, accessToken, agentId)

	response, err := util.HTTPGet(qyUrl, a.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyAgentView error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取access_token对应的应用列表
type QyAgentList struct {
	util.WxError
	Agentlist []struct {
		Agentid       int32  `json:"agentid"`
		Name          string `json:"name"`
		SquareLogoUrl string `json:"square_logo_url"`
	} `json:"agentlist"`
}
func (a *Agent) GetQyAgentList(accessToken string)(result *QyAgentList, err error){
	qyUrl := fmt.Sprintf(GetQyAgentListURL, accessToken)

	response, err := util.HTTPGet(qyUrl, a.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyAgentList error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//设置应用
type SetQyAgentReq struct {
	Agentid            int32  `json:"agentid"`
	ReportLocationFlag int32  `json:"report_location_flag"`
	LogoMediaid        string `json:"logo_mediaid"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	RedirectDomain     string `json:"redirect_domain"`
	Isreportenter      int32  `json:"isreportenter"`
	HomeUrl            string `json:"home_url"`
}
func (a *Agent) SetQyAgent(accessToken string, req SetQyAgentReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SetQyAgentURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, a.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SetQyAgent error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//应用管理-设置工作台自定义展示
package agent

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	SetWorkbenchTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_template?access_token=%s"  //设置应用在工作台展示的模版
	GetWorkbenchTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/get_workbench_template?access_token=%s"  //获取应用在工作台展示的模版
	SetWorkbenchDataURL = "https://qyapi.weixin.qq.com/cgi-bin/agent/set_workbench_data?access_token=%s"  //设置应用在用户工作台展示的数据

)

type AgentWorkbench struct {
	*core.Context
}

func NewAgentWorkbench(context *core.Context) *AgentWorkbench {
	aw := new(AgentWorkbench)
	aw.Context = context
	return aw
}

//设置应用在工作台展示的模版
type SetWorkbenchTemplateReq struct {
	Agentid int32  `json:"agentid"`
	Type    string `json:"type"`
	Image   struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"image,omitempty"`
	Keydata struct{
		Items []struct {
			Key      string `json:"key"`
			Data     string `json:"data"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"keydata,omitempty"`
	List struct{
		Items []struct {
			Title    string `json:"title"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"list,omitempty"`
	Webview struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"webview,omitempty"`
	ReplaceUserData bool `json:"replace_user_data"`
}
func (aw *AgentWorkbench) SetWorkbenchTemplate(accessToken string, req SetWorkbenchTemplateReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SetWorkbenchTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, aw.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SetWorkbenchTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取应用在工作台展示的模版
type WorkbenchTemplateReq struct {
	Agentid int `json:"agentid"`
}
type WorkbenchTemplateRep struct {
	util.WxError
	Agentid int32  `json:"agentid"`
	Type    string `json:"type"`
	Image   struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"image,omitempty"`
	Keydata struct{
		Items []struct {
			Key      string `json:"key"`
			Data     string `json:"data"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"keydata,omitempty"`
	List struct{
		Items []struct {
			Title    string `json:"title"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"list,omitempty"`
	Webview struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"webview,omitempty"`
	ReplaceUserData bool `json:"replace_user_data"`
}
func (aw *AgentWorkbench) GetWorkbenchTemplate(accessToken string, req WorkbenchTemplateReq)(result *WorkbenchTemplateRep, err error){
	qyUrl := fmt.Sprintf(GetWorkbenchTemplateURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, aw.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetWorkbenchTemplate error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//设置应用在用户工作台展示的数据
type SetWorkbenchDataReq struct {
	Agentid int32  `json:"agentid"`
	Userid  string `json:"userid"`
	Type    string `json:"type"`
	Image   struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"image,omitempty"`
	Keydata struct{
		Items []struct {
			Key      string `json:"key"`
			Data     string `json:"data"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"keydata,omitempty"`
	List struct{
		Items []struct {
			Title    string `json:"title"`
			JumpUrl  string `json:"jump_url"`
			Pagepath string `json:"pagepath"`
		} `json:"items"`
	} `json:"list,omitempty"`
	Webview struct {
		Url      string `json:"url"`
		JumpUrl  string `json:"jump_url"`
		Pagepath string `json:"pagepath"`
	} `json:"webview,omitempty"`
}
func (aw *AgentWorkbench) SetWorkbenchData(accessToken string, req SetWorkbenchDataReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(SetWorkbenchDataURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, aw.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("SetWorkbenchData error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


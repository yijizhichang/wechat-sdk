//应用管理-自定义菜单
package agent

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
)

const (
	CreateQyMenuURL = "https://qyapi.weixin.qq.com/cgi-bin/menu/create?access_token=%s&agentid=%d"  //创建菜单
	GetQyMenuURL = "https://qyapi.weixin.qq.com/cgi-bin/menu/get?access_token=%s&agentid=%d"  //获取菜单
	DelQyMenuURL = "https://qyapi.weixin.qq.com/cgi-bin/menu/delete?access_token=%s&agentid=%d"  //删除菜单

)

type AgentMenu struct {
	*core.Context
}

func NewAgentMenu(context *core.Context) *AgentMenu {
	am := new(AgentMenu)
	am.Context = context
	return am
}

//创建菜单
type CreateQyMenuReq struct {
	Button []struct {
		Name string `json:"name"`
		Type string `json:"type,omitempty"`
		Key  string `json:"key,omitempty"`
		SubButton []struct {
			Type      string        `json:"type"`
			Name      string        `json:"name"`
			Key       string        `json:"key,omitempty"`
			SubButton []interface{} `json:"sub_button,omitempty"`
			Pagepath  string        `json:"pagepath,omitempty"`
			Appid     string        `json:"appid,omitempty"`
			Url  string `json:"url,omitempty"`
		} `json:"sub_button,omitempty"`

	} `json:"button"`
}
func (am *AgentMenu) CreateQyMenu(accessToken string, req CreateQyMenuReq)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(CreateQyMenuURL, accessToken)

	response, err := util.PostJSON(qyUrl, req, am.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("CreateQyMenu error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//获取菜单
type QyMenuRep struct {
	util.WxError
	Button []struct {
		Name string `json:"name"`
		Type string `json:"type,omitempty"`
		Key  string `json:"key,omitempty"`
		SubButton []struct {
			Type      string        `json:"type"`
			Name      string        `json:"name"`
			Key       string        `json:"key,omitempty"`
			SubButton []interface{} `json:"sub_button,omitempty"`
			Pagepath  string        `json:"pagepath,omitempty"`
			Appid     string        `json:"appid,omitempty"`
			Url  string `json:"url,omitempty"`
		} `json:"sub_button,omitempty"`

	} `json:"button"`
}
func (am *AgentMenu) GetQyMenu(accessToken string, agentId int32)(result *QyMenuRep, err error){
	qyUrl := fmt.Sprintf(GetQyMenuURL, accessToken, agentId)

	response, err := util.HTTPGet(qyUrl, am.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetQyMenu error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}

//删除菜单
func (am *AgentMenu) DelQyMenu(accessToken string, agentId int32)(result *util.WxError, err error){
	qyUrl := fmt.Sprintf(DelQyMenuURL, accessToken, agentId)

	response, err := util.HTTPGet(qyUrl, am.ProxyUrl)

	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("DelQyMenu error : errcode=%d , errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	return
}


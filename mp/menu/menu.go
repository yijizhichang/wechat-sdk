package menu

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	menuFlag      = "设置微信菜单"
	baseURL       = "https://api.weixin.qq.com/cgi-bin"
	menuBaseURL   = baseURL + "/menu"
	commonToken   = "?access_token="
	menuCreateURL = menuBaseURL + "/create"
	menuGetURL    = menuBaseURL + "/get"
	menuDeleteURL = menuBaseURL + "/delete"
	// 个性化菜单
	menuAddConditionalURL    = menuBaseURL + "/addconditional"
	menuDeleteConditionalURL = menuBaseURL + "/delconditional"
	menuTryMatchURL          = menuBaseURL + "/trymatch"

	menuSelfMenuInfoURL = baseURL + "/get_current_selfmenu_info"
)

type Menu struct {
	*core.Context
}

type MenuButton struct {
	SubButton struct {
		List []MenuButton `json:"list"`
	} `json:"sub_button,omitempty"`
	Type string `json:"type"`
	Name string `json:"name"`
	Key  string `json:"key"`
	URL  string `json:"url,omitempty"`
}

func NewMenu(context *core.Context) *Menu {
	menu := new(Menu)
	menu.Context = context
	return menu
}

//SetMenu 设置按钮
func (m *Menu) SetMenu(buttons ...*Button) (result util.WxError, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}
	uri := menuCreateURL + commonToken + accessToken
	response, err := util.PostJSON(uri, &menuReq{Button: buttons}, m.ProxyUrl)
	if err != nil {
		return
	}
	var WxError util.WxError
	err = json.Unmarshal(response, &WxError)
	if err != nil {
		return
	}
	if WxError.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", WxError.ErrCode, WxError.ErrMsg)
	}
	return
}

//GetMenu 获取菜单配置
func (m *Menu) GetMenu() (resMenu MenuRes, err error) {
	var accessToken string
	accessToken, err = m.GetAccessToken()
	if err != nil {
		return
	}
	uri := menuGetURL + commonToken + accessToken
	var response []byte
	response, err = util.HTTPGet(uri, m.ProxyUrl)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &resMenu)
	if err != nil {
		return
	}
	if resMenu.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", resMenu.ErrCode, resMenu.ErrMsg)
		return
	}
	return
}

//DeleteMenu 删除菜单
func (m *Menu) DeleteMenu() error {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return err
	}
	uri := menuDeleteURL + commonToken + accessToken
	response, err := util.HTTPGet(uri, m.ProxyUrl)
	if err != nil {
		return err
	}
	var WxError util.WxError
	err = json.Unmarshal(response, &WxError)
	if err != nil {
		return err
	}
	if WxError.ErrCode != 0 {
		return fmt.Errorf("errcode-%d,errmsg-%s", WxError.ErrCode, WxError.ErrMsg)
	}
	return nil
}

// 添加个性化菜单
type AddConditionalResult struct{
	util.WxError
	MenuId	int32	`json:"menuid"`
}
func (m *Menu) AddConditional(buttons []*Button, matchRule *MatchRule) (result AddConditionalResult, err error) {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return
	}

	uri := menuAddConditionalURL + commonToken + accessToken
	reqMenu := &menuReq{
		Button:    buttons,
		MatchRule: matchRule,
	}

	response, err := util.PostJSON(uri, reqMenu, m.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("AddConditional error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
	}
	return
}

// 删除个性化菜单
func (m *Menu) DeleteConditional(menuid int64) error {
	accessToken, err := m.GetAccessToken()
	if err != nil {
		return err
	}

	uri := menuDeleteConditionalURL + commonToken + accessToken
	reqDelConditional := &reqDelConditional{
		Menuid: menuid,
	}

	response, err := util.PostJSON(uri, reqDelConditional, m.ProxyUrl)
	if err != nil {
		return err
	}
	var wxError util.WxError
	err = json.Unmarshal(response, &wxError)
	if err != nil {
		return err
	}
	if wxError.ErrCode != 0 {
		return fmt.Errorf("errcode-%d,errmsg-%s", wxError.ErrCode, wxError.ErrMsg)
	}
	return nil
}

// 测试个性化菜单匹配结果
func (m *Menu) MenuTryMatch(userID string) (buttons []Button, err error) {
	var accessToken string
	accessToken, err = m.GetAccessToken()
	if err != nil {
		return
	}
	uri := menuTryMatchURL + commonToken + accessToken
	reqMenuTryMatch := &reqMenuTryMatch{userID}
	var response []byte
	response, err = util.PostJSON(uri, reqMenuTryMatch, m.ProxyUrl)
	if err != nil {
		return
	}
	//fmt.Println("---------", string(response))
	var resMenuTryMatch resMenuTryMatch
	err = json.Unmarshal(response, &resMenuTryMatch)
	if err != nil {
		return
	}
	if resMenuTryMatch.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", resMenuTryMatch.ErrCode, resMenuTryMatch.ErrMsg)
		return
	}
	buttons = resMenuTryMatch.Menu.Button
	return
}

// 获取自定义菜单配置接口
func (m *Menu) GetSelfMenuInfo() (resSelfMenuInfo ResSelfMenuInfo, err error) {
	var accessToken string
	accessToken, err = m.GetAccessToken()
	if err != nil {
		return
	}
	uri := menuSelfMenuInfoURL + commonToken + accessToken
	var response []byte
	response, err = util.HTTPGet(uri, m.ProxyUrl)
	if err != nil {
		return
	}
	//fmt.Println("---------", string(response))
	err = json.Unmarshal(response, &resSelfMenuInfo)
	if err != nil {
		return
	}
	if resSelfMenuInfo.ErrCode != 0 {
		err = fmt.Errorf("errcode-%d,errmsg-%s", resSelfMenuInfo.ErrCode, resSelfMenuInfo.ErrMsg)
		return
	}
	return
}

// 设置菜单请求数据
type menuReq struct {
	Button    []*Button  `json:"button,omitempty"`
	MatchRule *MatchRule `json:"matchrule,omitempty"`
}

// 删除个性化菜单请求数据
type reqDelConditional struct {
	Menuid int64 `json:"menuid"`
}

// 菜单匹配请求
type reqMenuTryMatch struct {
	UserID string `json:"user_id"`
}

// 查询菜单的返回数据
type MenuRes struct {
	util.WxError
	Menu struct {
		Button []Button `json:"button"`
		Menuid int64    `json:"menuid"`
	} `json:"menu"`
	Conditionalmenu []resConditionalMenu `json:"conditionalmenu"`
}

// 个性化菜单返回结果
type resConditionalMenu struct {
	Button    []Button  `json:"button"`
	MatchRule MatchRule `json:"matchrule"`
	MenuID    int64     `json:"menuid"`
}

// 菜单匹配请求结果
type resMenuTryMatch struct {
	util.WxError
	Menu struct {
		Button []Button `json:"button"`
	} `json:"menu"`
}

// 个性化菜单规则
type MatchRule struct {
	GroupID            string `json:"group_id,omitempty"`
	Sex                string `json:"sex,omitempty"`
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	City               string `json:"city,omitempty"`
	ClientPlatformType string `json:"client_platform_type,omitempty"`
	Language           string `json:"language,omitempty"`
}

// 自定义菜单配置返回结果
type ResSelfMenuInfo struct {
	util.WxError
	IsMenuOpen   int32 `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []SelfMenuButton `json:"button"`
	} `json:"selfmenu_info"`
}

// 自定义菜单配置详情
type SelfMenuButton struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	URL       string `json:"url,omitempty"`
	Value     string `json:"value,omitempty"`
	SubButton struct {
		List []SelfMenuButton `json:"list"`
	} `json:"sub_button,omitempty"`
	// 不是通过api设置的菜单会返回
	NewsInfo struct {
		List []News `json:"list"`
	} `json:"news_info,omitempty"`
}

//ButtonNew 图文消息菜单
type News struct {
	Title      string `json:"title"`       // 图文消息的标题
	Digest     string `json:"digest"`      // 摘要
	Author     string `json:"author"`      // 作者
	ShowCover  int8   `json:"show_cover"`  // 是否显示封面，0为不显示，1为显示
	CoverURL   string `json:"cover_url"`   // 封面图片的URL
	ContentURL string `json:"content_url"` // 正文的URL
	SourceURL  string `json:"source_url"`  // 原文的URL，若置空则无查看原文入口
}

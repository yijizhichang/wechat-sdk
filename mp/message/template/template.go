//模板消息包
package template

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/util"
)

const (
	SetIndustryURL           = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"         //设置所属行业
	GetIndustryURL           = "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s"             //获取设置的行业信息
	AddTemplateURL           = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s"         //获得模板ID
	GetAllPrivateTemplateURL = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s" //获取模板列表
	DelPrivateTemplateURL    = "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s"     //删除模板
	SendTemplateURL          = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"             //发送模板消息
)

//Template 模板消息
type Template struct {
	*core.Context
}

//NewTemplate 实例化
func NewTemplate(context *core.Context) *Template {
	tpl := new(Template)
	tpl.Context = context
	return tpl
}

//设置所属行业
type setIndustry struct {
	IndustryId1 int64 `json:"industry_id1"`
	IndustryId2 int64 `json:"industry_id2"`
}

func (tpl *Template) SetIndustry(industryId1, industryId2 int64) (result util.WxError, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(SetIndustryURL, accessToken)

	postData := new(setIndustry)
	postData.IndustryId1 = industryId1
	postData.IndustryId2 = industryId2

	response, err := util.PostJSON(wxUrl, postData, tpl.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("SetIndustry error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("设置所属行业错误", err)
	}
	return
}

//获取设置的行业信息
type Industry struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}
type getIndustry struct {
	util.WxError
	PrimaryIndustry   Industry `json:"primary_industry"`
	SecondaryIndustry Industry `json:"secondary_industry"`
}

func (tpl *Template) GetIndustry() (result *getIndustry, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetIndustryURL, accessToken)
	response, err := util.HTTPGet(wxUrl, tpl.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetIndustry error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("获取设置的行业信息错误", err)
	}
	return
}

//获得模板ID
type tplIdShort struct {
	TemplateIdShort string `json:"template_id_short"`
}

type addTemplate struct {
	util.WxError
	TemplateID string `json:"template_id"`
}

func (tpl *Template) GetAddTemplate(templateIdShort string) (result *addTemplate, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(AddTemplateURL, accessToken)
	postData := new(tplIdShort)
	postData.TemplateIdShort = templateIdShort

	response, err := util.PostJSON(wxUrl, postData, tpl.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetAddTemplate error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("获得模板ID错误", err)
	}
	return
}

//获取模板列表
type template struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type templateList struct {
	util.WxError
	TemplateList []template `json:"template_list"`
}

func (tpl *Template) GetAllPrivateTemplate() (result *templateList, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(GetAllPrivateTemplateURL, accessToken)
	response, err := util.HTTPGet(wxUrl, tpl.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("GetAllPrivateTemplate error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("获取模板列表错误", err)
	}
	return
}

//删除模板
type templateId struct {
	TemplateID string `json:"template_id"`
}

func (tpl *Template) DelPrivateTemplate(tplId string) (result *util.WxError, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(DelPrivateTemplateURL, accessToken)
	postData := new(templateId)
	postData.TemplateID = tplId

	response, err := util.PostJSON(wxUrl, postData, tpl.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("DelPrivateTemplate error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("删除模板错误", err)
		return
	}
	return
}

//发送模板消息
type Message struct {
	ToUser     string               `json:"touser"`          // 必须, 接受者OpenID
	TemplateID string               `json:"template_id"`     // 必须, 模版ID
	URL        string               `json:"url,omitempty"`   // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	Color      string               `json:"color,omitempty"` // 可选, 整个消息的颜色, 可以不设置
	Data       map[string]*DataItem `json:"data"`            // 必须, 模板数据

	MiniProgram struct {
		AppID    string `json:"appid"`    //所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
		PagePath string `json:"pagepath"` //所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar）
	} `json:"miniprogram"` //可选,跳转至小程序地址
}

//模版内 .DATA 的值
type DataItem struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

type resTemplateSend struct {
	util.WxError
	MsgID int64 `json:"msgid"`
}

//发送模板消息
func (tpl *Template) Send(msg *Message) (result *resTemplateSend, err error) {
	accessToken, err := tpl.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(SendTemplateURL, accessToken)

	response, err := util.PostJSON(wxUrl, msg, tpl.ProxyUrl)
	if err != nil {
		return
	}
	tpl.WXLog.Debug("发送模板消息内容", msg, "代理地址", tpl.ProxyUrl)

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	tpl.WXLog.Debug("发送模板消息返回结果", result, "代理地址", tpl.ProxyUrl)

	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("DelPrivateTemplate error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		tpl.WXLog.Error("发送模板消息错误", err, "代理地址", tpl.ProxyUrl)
		return
	}
	return
}

//生成带参数的二维码
package account

import (
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"fmt"
	"encoding/json"
)

const (
	ActionNameInt		=	"QR_SCENE"			//二维码类型，QR_SCENE为临时的整型参数值
	ActionNameStr		=	"QR_STR_SCENE"		//QR_STR_SCENE为临时的字符串参数值
	ActionNameLimitInt  =	"QR_LIMIT_SCENE"	//QR_LIMIT_SCENE为永久的整型参数值
	ActionNameLimitStr	=	"QR_LIMIT_STR_SCENE"	//QR_LIMIT_STR_SCENE为永久的字符串参数值
	CreateQrCodeURL 	= 	"https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"  	//生成带参数的二维码
	ShowQrCodeURL		=	"https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s"		//通过ticket换取二维码
	ShortURL			=	"https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s"	//长链接转短链接接口
)


//Account 账户管理
type Account struct {
	*core.Context
}

//NewAccount 实例化
func NewAccount(context *core.Context) *Account {
	account := new(Account)
	account.Context = context
	return account
}


//生成二维码  场景值ID = 数值
type reqCreateQrCodeSceneId struct{
	ActionName		string		`json:"action_name"`
	ExpireSeconds 	int32		`json:"expire_seconds"`
	ActionInfo	struct{
			Scene	struct{
				SceneId		int32	`json:"scene_id"`
			}	`json:"scene"`
	}	`json:"action_info"`
}

type reqCreateQrCodeSceneStr struct{
	ActionName		string		`json:"action_name"`
	ExpireSeconds 	int32		`json:"expire_seconds,omitempty"`
	ActionInfo	struct{
		Scene	struct{
			SceneStr		string	`json:"scene_str"`
		}	`json:"scene"`
	}	`json:"action_info"`
}

type qrCodeResult struct{
	util.WxError
	Ticket   		string		`json:"ticket"`
	ExpireSeconds 	int32		`json:"expire_seconds"`
	Url             string		`json:"url"`
}

//创建二维码 scene_id
func (account *Account) CreateQrCodeSceneId(isLimit bool, sceneId int32, expireSeconds int32)(result qrCodeResult, err error)  {
	accessToken, err := account.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(CreateQrCodeURL, accessToken)

	postData := new(reqCreateQrCodeSceneId)

	if isLimit == true {
		postData.ActionName = ActionNameLimitInt
	}else{
		postData.ActionName = ActionNameInt
		postData.ExpireSeconds = expireSeconds
	}
	postData.ActionInfo.Scene.SceneId = sceneId


	response, err := util.PostJSON(wxUrl, postData, account.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("CreateQrCodeSceneId error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		account.WXLog.Error("创建二维码（SceneId）错误", err)
	}
	return
}

//创建二维码 scene_str
func (account *Account) CreateQrCodeSceneStr(isLimit bool, sceneStr string, expireSeconds int32)(result qrCodeResult, err error)  {
	accessToken, err := account.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(CreateQrCodeURL, accessToken)

	postData := new(reqCreateQrCodeSceneStr)

	if isLimit == true {
		postData.ActionName = ActionNameLimitStr
	}else{
		postData.ActionName = ActionNameStr
		postData.ExpireSeconds = expireSeconds
	}
	postData.ActionInfo.Scene.SceneStr = sceneStr


	response, err := util.PostJSON(wxUrl, postData, account.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("CreateQrCodeSceneStr error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		account.WXLog.Error("创建二维码（SceneStr）错误", err)
	}
	return
}

//通过ticket换取二维码
func (account *Account) GetQrCodeUrl(ticket string) (Url string)  {
	Url = fmt.Sprintf(ShowQrCodeURL, ticket)
	return
}


//长链接转短链接
type reqShortUrl struct{
	Action    	string	`json:"action"`
	LongUrl		string	`json:"long_url"`

}

type shortUrlResult struct{
	util.WxError
	ShortUrl   string	`json:"short_url"`
}
func (account *Account) ShortUrl(longUrl string) (result shortUrlResult, err error)  {
	accessToken, err := account.GetAccessToken()
	if err != nil {
		return
	}
	wxUrl := fmt.Sprintf(ShortURL, accessToken)

	postData := new(reqShortUrl)
	postData.Action	= "long2short"
	postData.LongUrl = longUrl


	response, err := util.PostJSON(wxUrl, postData, account.ProxyUrl)
	if err != nil {
		return
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		//打印日志
		err = fmt.Errorf("ShortUrl error : errcode=%v , errmsg=%v", result.ErrCode, string(result.ErrMsg))
		account.WXLog.Error("长链接转短链接错误", err)
	}
	return
}


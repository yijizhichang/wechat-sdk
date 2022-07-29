package wechat

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/account"
	"github.com/yijizhichang/wechat-sdk/mp/client"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/mp/custom"
	"github.com/yijizhichang/wechat-sdk/mp/jssdk"
	"github.com/yijizhichang/wechat-sdk/mp/media"
	"github.com/yijizhichang/wechat-sdk/mp/menu"
	"github.com/yijizhichang/wechat-sdk/mp/message/mass"
	"github.com/yijizhichang/wechat-sdk/mp/message/template"
	"github.com/yijizhichang/wechat-sdk/mp/oauth2"
	"github.com/yijizhichang/wechat-sdk/mp/server"
	"github.com/yijizhichang/wechat-sdk/mp/user"
	"github.com/yijizhichang/wechat-sdk/util/cache"
	"net/http"
	"sync"
)

//Wechat结构体
type Wechat struct {
	Context *core.Context
}

//用户wechat配置
type Config struct {
	AppID            string      //开发者ID(AppID)
	AppSecret        string      //开发者PWD AppSecret
	Token            string      //令牌(Token)
	EncodingAESKey   string      //消息加解密密钥 EncodingAESKey
	PayMchId         string      //支付 - 商户 ID
	PayNotifyUrl     string      //支付 - 接受微信支付结果通知的接口地址
	PayKey           string      //支付 - 商户后台设置的支付 key
	Cache            cache.Cache //缓存
	ThirdAccessToken bool        //是否共享其它accessToken，非appID获取 默认false
	ProxyUrl         string      //缓存配置文件
}

//实例化wechat
func NewWechat(cfg *Config) *Wechat {
	context := new(core.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *core.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.PayMchId = cfg.PayMchId
	context.PayKey = cfg.PayKey
	context.PayNotifyUrl = cfg.PayNotifyUrl
	context.Cache = cfg.Cache
	context.ProxyUrl = cfg.ProxyUrl
	context.ThirdAccessToken = cfg.ThirdAccessToken
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}

// GetServer 消息管理
func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

// GetResponseServer 消息不直接发给微信，而是返回给调用应用，由应用返回给微信
func (wc *Wechat) GetResponseServer(req *http.Request) *server.Server {
	wc.Context.Request = req
	return server.NewServer(wc.Context)
}

//获取access_token
func (wc *Wechat) GetAccessToken() string {
	accessToken, err := wc.Context.GetAccessToken()
	if err != nil {
		fmt.Printf("mp GetAccessToken Err:%+v", err)
	}
	return accessToken
}

//Oauth2网页授权
func (wc *Wechat) GetOauth() *oauth2.Oauth {
	return oauth2.NewOauth(wc.Context)
}

func (wc *Wechat) GetClient() *client.Client {
	return client.NewClient(wc.Context)
}

//模板消息管理
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Context)
}

//菜单管理
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

//群发管理
func (wc *Wechat) GetMass() *mass.Mass {
	return mass.NewMass(wc.Context)
}

//用户管理
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

//客服管理
func (wc *Wechat) GetCustom() *custom.Custom {
	return custom.NewCustom(wc.Context)
}

//素材管理
func (wc *Wechat) GetMedia() *media.Media {
	return media.NewMedia(wc.Context)
}

//账户管理
func (wc *Wechat) GetAccount() *account.Account {
	return account.NewAccount(wc.Context)
}

//账户管理
func (wc *Wechat) GetJSSDK() *jssdk.JSAPISDK {
	return jssdk.NewJSSDK(wc.Context)
}

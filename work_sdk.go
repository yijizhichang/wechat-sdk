package wechat

import (
	"github.com/yijizhichang/wechat-sdk/util/cache"
	"github.com/yijizhichang/wechat-sdk/work/core"
	"github.com/yijizhichang/wechat-sdk/work/customer"
	"github.com/yijizhichang/wechat-sdk/work/company"
	"sync"
)

//企微结构
type QyWechat struct {
	Context *core.Context
}

//企微配置
type QyConfig struct {
	CorpID        		string  // 企业ID
	CorpSecret    		string  // 应用的凭证密钥; 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
	RasPrivateKey 		string  // 消息加密私钥
	ThirdAccessToken 	bool    //是用其他应用生成的access_token

	Cache      cache.Cache //缓存
	ProxyUrl   string      //代理地址
}

//实例化qy wechat
func NewQyWechat(cfg *QyConfig) *QyWechat {
	context := new(core.Context)
	copyConfigToQyContext(cfg, context)
	return &QyWechat{context}
}

func copyConfigToQyContext(cfg *QyConfig, context *core.Context) {
	context.CorpID = cfg.CorpID
	context.CorpSecret = cfg.CorpSecret
	context.RasPrivateKey = cfg.RasPrivateKey
	context.ThirdAccessToken = cfg.ThirdAccessToken

	context.Cache = cfg.Cache
	context.ProxyUrl = cfg.ProxyUrl

	context.SetAccessTokenLock(new(sync.RWMutex))
}


//获取企微对应应用的access_token
func (qw *QyWechat) GetQyAccessToken(corpSecret string) (accessToken string, err error) {
	accessToken, err = qw.Context.GetAccessToken(corpSecret)
	return
}

//客户联系-客户管理
func (qw *QyWechat) GetCustomer() *customer.Customer {
	return customer.NewCustomer(qw.Context)
}
//客户联系-客户标签管理
func (qw *QyWechat) GetCustomerTag() *customer.CustomerTag {
	return customer.NewCustomerTag(qw.Context)
}

//通讯录管理-部门管理
func (qw *QyWechat) GetDepartment() *company.Department {
	return company.NewDepartment(qw.Context)
}
//通讯录管理-部门成员管理
func (qw *QyWechat) GetDepartmentUser() *company.DepartmentUser {
	return company.NewDepartmentUser(qw.Context)
}



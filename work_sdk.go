package wechat

import (
	"github.com/yijizhichang/wechat-sdk/util/cache"
	"github.com/yijizhichang/wechat-sdk/work/agent"
	"github.com/yijizhichang/wechat-sdk/work/company"
	"github.com/yijizhichang/wechat-sdk/work/core"
	"github.com/yijizhichang/wechat-sdk/work/customer"
	"github.com/yijizhichang/wechat-sdk/work/kefu"
	"github.com/yijizhichang/wechat-sdk/work/media"
	"github.com/yijizhichang/wechat-sdk/work/message"
	"github.com/yijizhichang/wechat-sdk/work/oa"
	"github.com/yijizhichang/wechat-sdk/work/oauth2"
	"github.com/yijizhichang/wechat-sdk/work/server"
	"github.com/yijizhichang/wechat-sdk/work/tools"
	"github.com/yijizhichang/wechat-sdk/work/webhook"
	"net/http"
	"sync"
)

//企微结构
type QyWechat struct {
	Context *core.Context
}

//企微配置
type QyConfig struct {
	CorpID           string      // 企业ID
	CorpSecret       string      // 应用的凭证密钥; 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
	RasPrivateKey    string      // 消息加密私钥
	Token            string      // 令牌(Token)
	EncodingAESKey   string      // 消息加解密密钥 EncodingAESKey
	ThirdAccessToken bool        //是用其他应用生成的access_token
	Debug            bool        //为true时会打印一些调试信息
	Cache            cache.Cache //缓存
	ProxyUrl         string      //代理地址
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
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.RasPrivateKey = cfg.RasPrivateKey
	context.ThirdAccessToken = cfg.ThirdAccessToken
	context.Debug = cfg.Debug
	context.Cache = cfg.Cache
	context.ProxyUrl = cfg.ProxyUrl

	context.SetAccessTokenLock(new(sync.RWMutex))
}

// GetQyServer 消息管理
func (qw *QyWechat) GetQyServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	qw.Context.Request = req
	qw.Context.Writer = writer
	return server.NewServer(qw.Context)
}

// GetQyResponseServer 消息不直接发给微信，而是返回给调用应用，由应用返回给微信
func (qw *QyWechat) GetQyResponseServer(req *http.Request) *server.Server {
	qw.Context.Request = req
	return server.NewServer(qw.Context)
}

//获取企微对应应用的access_token
func (qw *QyWechat) GetQyAccessToken(corpSecret string) (accessToken string, err error) {
	accessToken, err = qw.Context.GetAccessToken(corpSecret)
	return
}

//--客户联系--//
//客户联系-企业服务人员管理
func (qw *QyWechat) GetCustomerFollow() *customer.CustomerFollow {
	return customer.NewCustomerFollow(qw.Context)
}

//客户联系-客户管理
func (qw *QyWechat) GetCustomer() *customer.Customer {
	return customer.NewCustomer(qw.Context)
}

//客户联系-客户标签管理
func (qw *QyWechat) GetCustomerTag() *customer.CustomerTag {
	return customer.NewCustomerTag(qw.Context)
}

//客户联系-在职继承
func (qw *QyWechat) GetCustomerOnTransfer() *customer.CustomerOnTransfer {
	return customer.NewCustomerOnTransfer(qw.Context)
}

//客户联系-离职继承
func (qw *QyWechat) GetCustomerOffTransfer() *customer.CustomerOffTransfer {
	return customer.NewCustomerOffTransfer(qw.Context)
}

//客户联系-客户群管理
func (qw *QyWechat) GetCustomerGroup() *customer.CustomerGroup {
	return customer.NewCustomerGroup(qw.Context)
}

//客户联系-朋友圈
func (qw *QyWechat) GetCustomerMoment() *customer.CustomerMoment {
	return customer.NewCustomerMoment(qw.Context)
}

//客户联系-消息推送
func (qw *QyWechat) GetCustomerMsg() *customer.CustomerMsg {
	return customer.NewCustomerMsg(qw.Context)
}

//客户联系-统计管理
func (qw *QyWechat) GetCustomerData() *customer.CustomerData {
	return customer.NewCustomerData(qw.Context)
}

//--通讯录管理--//
//通讯录管理-部门管理
func (qw *QyWechat) GetDepartment() *company.Department {
	return company.NewDepartment(qw.Context)
}

//通讯录管理-部门成员管理
func (qw *QyWechat) GetDepartmentUser() *company.DepartmentUser {
	return company.NewDepartmentUser(qw.Context)
}

//通讯录管理-标签管理
func (qw *QyWechat) GetCompanyTag() *company.CompanyTag {
	return company.NewCompanyTag(qw.Context)
}

//通讯录管理-异步批量接口
func (qw *QyWechat) GetCompanyBatch() *company.CompanyBatch {
	return company.NewCompanyBatch(qw.Context)
}

//通讯录管理-互联企业
func (qw *QyWechat) GetCompanyLinkedCorp() *company.CompanyLinkedCorp {
	return company.NewCompanyLinkedCorp(qw.Context)
}

//通讯录管理-异步导出接口
func (qw *QyWechat) GetCompanyExport() *company.CompanyExport {
	return company.NewCompanyExport(qw.Context)
}

//--微信客服--//
//微信客服-客服账号管理
func (qw *QyWechat) GetKefuAccount() *kefu.KefuAccount {
	return kefu.NewKefuAccount(qw.Context)
}

//微信客服-接待人员管理
func (qw *QyWechat) GetKefuServicer() *kefu.KefuServicer {
	return kefu.NewKefuServicer(qw.Context)
}

//微信客服-会话分配与消息收发
func (qw *QyWechat) GetKefuConverse() *kefu.KefuConverse {
	return kefu.NewKefuConverse(qw.Context)
}

//微信客服-升级服务配置/客户基本信息获取
func (qw *QyWechat) GetKefu() *kefu.Kefu {
	return kefu.NewKefu(qw.Context)
}

//--身份验证--//
func (qw *QyWechat) GetOauth2() *oauth2.Oauth2 {
	return oauth2.NewOauth2(qw.Context)
}

//--应用管理--//
//应用管理-应用设置
func (qw *QyWechat) GetAgent() *agent.Agent {
	return agent.NewAgent(qw.Context)
}

//应用管理-自定义菜单
func (qw *QyWechat) GetAgentMenu() *agent.AgentMenu {
	return agent.NewAgentMenu(qw.Context)
}

//应用管理-设置工作台自定义展示
func (qw *QyWechat) GetAgentWorkbench() *agent.AgentWorkbench {
	return agent.NewAgentWorkbench(qw.Context)
}

//--消息推送--//
//消息推送-发送消息/更新模板卡片/撤回应用消息
func (qw *QyWechat) GetMessage() *message.Message {
	return message.NewMessage(qw.Context)
}

//消息推送-发送消息到群聊会话
func (qw *QyWechat) GetMessageGroup() *message.MessageGroup {
	return message.NewMessageGroup(qw.Context)
}

//消息推送-互联企业消息推送
func (qw *QyWechat) GetMessageLinkedCorp() *message.MessageLinkedCorp {
	return message.NewMessageLinkedCorp(qw.Context)
}

//消息推送-家校消息推送
//todo

//--素材管理--//
func (qw *QyWechat) GetMedia() *media.Media {
	return media.NewMedia(qw.Context)
}

//--OA--//
//会议室
func (qw *QyWechat) GetMeetingroom() *oa.MeetingRoom {
	return oa.NewMeetingRoom(qw.Context)
}

//--效率工具--//

//直播
func (qw *QyWechat) GetLiving() *tools.Living {
	return tools.NewLiving(qw.Context)
}

//日历
func (qw *QyWechat) GetCalendar() *tools.Calendar {
	return tools.NewCalendar(qw.Context)
}

//机器人管理
func (qw *QyWechat) GetWebhook() *webhook.Webhook {
	return webhook.NewWebhook(qw.Context)
}

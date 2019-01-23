package wechat

import (
	"github.com/yijizhichang/wechat-sdk/mp/account"
	"github.com/yijizhichang/wechat-sdk/mp/client"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/mp/custom"
	"github.com/yijizhichang/wechat-sdk/mp/media"
	"github.com/yijizhichang/wechat-sdk/mp/menu"
	"github.com/yijizhichang/wechat-sdk/mp/message/mass"
	"github.com/yijizhichang/wechat-sdk/mp/message/template"
	"github.com/yijizhichang/wechat-sdk/mp/oauth2"
	"github.com/yijizhichang/wechat-sdk/mp/server"
	"github.com/yijizhichang/wechat-sdk/mp/user"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/util/cache"
	flog "github.com/yijizhichang/wechat-sdk/util/log"
	"net/http"
	"sync"
)

//微信SDK打印日志默认配置
const (
	WxLogConfigLogLevel    = 4                  //日志级别 =0 ALL; =1 DEBUG; =2 INFO; =3 WARN; =4 ERROR; =5 FATAL; =6 ALERT; =7 OFF;  注意：测试可以设置DEBUG;线上设置INFO 或 ERROR
	WxLogConfigIsConsole   = false              //是否输出到控制台
	WxLogConfigIsFile      = false              //是否写文件
	WxLogConfigFilePath    = "./util/debug/"    //文件路径
	WxLogConfigFilename    = "wechat"           //文件名称
	WxLogConfigFileSuffix  = "log"              //文件后缀
	WxLogConfigFileMaxSize = 1024 * 1024 * 1024 //单个日志文件大小 单位B, 1024 * 1024 * 1024 为1G

	WxCacheFilePath = "./util/debug/cache.txt" //缓存文件目录
	WxRedisAddr     = "127.0.0.1:6380"         //redis Addr
	WxRedisPassword = "your redis pwd"         //redis pwd
)

//Wechat结构体
type Wechat struct {
	Context *core.Context
}

//redis缓存配置
type RedisConfig struct {
	Addr     string //reids 地址 127.0.0.1:6381
	Password string //PWD  your redis pwd
}

//redis集群配置
type RedisClusterConfig struct {
	Addr     []string //redisCluster 地址集合
	Password string   //PWD  your redis pwd
}

type CacheConfig struct {
	FilePath           string              //缓存目录,CacheModel 为file时有效
	RedisConfig        *RedisConfig        //redis配置 CacheModel 为redis时有效
	RedisClusterConfig *RedisClusterConfig //redis配置 CacheModel 为redisCluster时有效
}

type FlogConfig struct {
	LogLevel    flog.LEVEL //日志级别 =0 ALL; =1 DEBUG; =2 INFO; =3 WARN; =4 ERROR; =5 FATAL; =6 ALERT; =7 OFF;  注意：测试可以设置DEBUG;线上设置INFO 或 ERROR
	IsConsole   bool       //是否输出到控制台
	IsFile      bool       //是否写文件
	FilePath    string     //文件日志路径
	Filename    string     //文件名称
	FileSuffix  string     //文件后缀
	FileMaxSize int64      //单个日志文件大小 单位B, 1024 * 1024 * 1024 为1G
}

//用户wechat配置
type Config struct {
	AppID            string               //开发者ID(AppID)
	AppSecret        string               //开发者PWD AppSecret
	Token            string               //令牌(Token)
	EncodingAESKey   string               //消息加解密密钥 EncodingAESKey
	PayMchId         string               //支付 - 商户 ID
	PayNotifyUrl     string               //支付 - 接受微信支付结果通知的接口地址
	PayKey           string               //支付 - 商户后台设置的支付 key
	CacheModel       string               //缓存模式
	Cache            cache.Cache          //缓存
	WXLog            flog.LoggerInterface //日志模式
	ThirdAccessToken bool                 //是否共享其它accessToken，非appID获取 默认false
	FlogConfig       *FlogConfig          //文件日志配置
	CacheConfig      *CacheConfig         //缓存配置文件
	ProxyUrl         string               //缓存配置文件
}

//实例化wechat
func NewWechat(cfg *Config) *Wechat {

	//日志默认值
	util.LevelDefault(&cfg.FlogConfig.LogLevel, WxLogConfigLogLevel)
	util.BoolDefault(&cfg.FlogConfig.IsConsole, WxLogConfigIsConsole)
	util.BoolDefault(&cfg.FlogConfig.IsFile, WxLogConfigIsFile)
	util.StrDefault(&cfg.FlogConfig.FilePath, WxLogConfigFilePath)
	util.StrDefault(&cfg.FlogConfig.Filename, WxLogConfigFilename)
	util.StrDefault(&cfg.FlogConfig.FileSuffix, WxLogConfigFileSuffix)
	util.Int64Default(&cfg.FlogConfig.FileMaxSize, WxLogConfigFileMaxSize)

	//文件缓存默认值
	util.StrDefault(&cfg.CacheConfig.FilePath, WxCacheFilePath)
	util.StrDefault(&cfg.CacheConfig.RedisConfig.Addr, WxRedisAddr)
	util.StrDefault(&cfg.CacheConfig.RedisConfig.Password, WxRedisPassword)
	util.ArrayDefault(&cfg.CacheConfig.RedisClusterConfig.Addr, []string{"127.0.0.1:6380", "127.0.0.2:6380"})
	util.StrDefault(&cfg.CacheConfig.RedisClusterConfig.Password, "your redis pwd")

	//默认打日志配置
	logFile := flog.GetLogger()
	logFile.SetConfig(flog.LoggerConfig{
		LogLevel:    cfg.FlogConfig.LogLevel,
		IsConsole:   cfg.FlogConfig.IsConsole,
		IsFile:      cfg.FlogConfig.IsFile,
		FilePath:    cfg.FlogConfig.FilePath,
		Filename:    cfg.FlogConfig.Filename,
		FileSuffix:  cfg.FlogConfig.FileSuffix,
		FileMaxSize: cfg.FlogConfig.FileMaxSize,
	})
	cfg.WXLog = logFile

	//缓存方式
	var cacheModel cache.Cache
	switch cfg.CacheModel {
	case "file":
		cacheModel, _ = cache.NewFile(cfg.CacheConfig.FilePath)
	case "redis":
		cacheModel, _ = cache.NewRedis(&cache.Redis{
			Addr:     cfg.CacheConfig.RedisConfig.Addr,
			Password: cfg.CacheConfig.RedisConfig.Password,
		})
	case "redisCluster":
		cacheModel, _ = cache.NewRedisCluster(&cache.RedisCluster{
			Addrs:    cfg.CacheConfig.RedisClusterConfig.Addr,
			Password: cfg.CacheConfig.RedisClusterConfig.Password,
		})
	default:
		cacheModel, _ = cache.NewFile(cfg.CacheConfig.FilePath)
	}
	cfg.Cache = cacheModel

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
	context.CacheModel = cfg.CacheModel
	context.Cache = cfg.Cache
	context.WXLog = cfg.WXLog
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

//获取access_token
func (wc *Wechat) GetAccessToken() string {
	accessToken, _ := wc.Context.GetAccessToken()
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

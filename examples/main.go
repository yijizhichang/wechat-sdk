package main

import (
	"flag"
	"fmt"
	"github.com/yijizhichang/wechat-sdk"
	"github.com/yijizhichang/wechat-sdk/examples/example"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/examples/cache"
	wxCache "github.com/yijizhichang/wechat-sdk/util/cache"
	"net/http"
)

// 测试参数
// go run example.go -appid='your appdi' -appsecret='your appsecret' -token='your token' -port=':80'
var appid = flag.String("appid", "your AppID", "appid")
var appsecret = flag.String("appsecret", "your AppSecret", "appsecret")
var token = flag.String("token", "your Token", "token")
var aeskey = flag.String("asekey", "your EncodingAesKey", "asekey")
var port = flag.String("port", "80", "port")

// 参数配置
func getWxConfig(cacheModel wxCache.Cache) *wechat.Config {
	config := &wechat.Config{
		AppID:            *appid,     // 开发者ID(AppID)
		AppSecret:        *appsecret, // 开发者PWD AppSecret
		Token:            *token,     // 令牌(Token)
		EncodingAESKey:   *aeskey,    // 消息加解密密钥 EncodingAESKey
		PayMchId:         "",         // 支付 - 商户 ID
		PayNotifyUrl:     "",         // 支付 - 接受微信支付结果通知的接口地址
		PayKey:           "",         // 支付 - 商户后台设置的支付 key
		Cache:            cacheModel, // 缓存方式 可选 file,redis,redisCluster 来实现cache的接口
		ThirdAccessToken: false,      // 是否使用第三方accessToken
		ProxyUrl:         "",         // 代理
	}
	return config
}

func main() {
	// NewWechat
	flag.Parse()
	//缓存 -file模式
	fileCache,_ := cache.NewFile("./debug/cache/cache.txt")

	//缓存 -redis
	//redisClusterCache, _ := cache.NewRedisCluster(&cache.RedisCluster{
	//	Addrs:    []string{"127.0.0.1:6379","127.0.0.1:6380"},
	//	Password: "123456",
	//})

	wxconf.WechatClient = wechat.NewWechat(getWxConfig(fileCache))

	fmt.Println(*appid, *appsecret, *token, *port)

	// 测试样例
	// example.SendTemplateMsg()  //模板消息发送
	// example.MassManage() //群发

	// example.CustomManage()  //客服管理
	// example.CustomSession()  //客服会话
	// example.CustomMsg()  	//聊天记录
	// example.SendCustomMsg()  //客服消息发送

	// example.UserTagManage()  //用户标签管理
	// example.UserRemark()  //用户备注
	// example.UserInfo()  //用户基本信息
	// example.UserList()  //用户列表
	// example.UserBlack()  //用户黑名单管理

	// example.Comment()   //评论管理

	// example.MenuCreate()   //创建菜单
	// example.MenuAddConditional()   //创建个性化菜单
	// example.MenuGet()   //获取所有菜单
	// example.MenuDel()  //删除所有菜单
	// example.MenuDelConditional()  //删除个性化菜单
	// example.MenuTryMatch() //个性化菜单匹配测试
	// example.GetSelfMenuInfo() //获取自定义菜单配置

	// example.AccountManage() //二维码管理
	// example.GetJSSign() //jssdk配置

	//http.HandleFunc("/wechat/server", example.Serve)                             // server 服务
	http.HandleFunc("/wechat/server", example.ResponseServe)               // server 服务 数据先返回应用
	http.HandleFunc("/customHeadImg", example.CustomHeadImg)               // 设置客服头像
	http.HandleFunc("/uploadTempMedia", example.UploadTempMedia)           // 上传临时素材
	http.HandleFunc("/uploadNewsPermanent", example.UploadNewsPermanent)   // 上传永久图文素材
	http.HandleFunc("/uploadVideoPermanent", example.UploadVideoPermanent) // 上传永久视频素材
	http.HandleFunc("/getMaterialMediaList", example.GetMaterialMediaList) // 获取永久素材列表
	http.HandleFunc("/getMaterialMediaInfo", example.GetMaterialMediaInfo) // 获取永久素材info
	http.HandleFunc("/delMaterialMedia", example.DelMaterialMedia)         // 获取永久素材
	http.HandleFunc("/updateNewsMedia", example.UpdateNewsMedia)           // 修改图文永久素材

	err := http.ListenAndServe("127.0.0.1:"+*port, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}

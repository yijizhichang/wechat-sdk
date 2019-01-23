package main

import (
	"flag"
	"fmt"
	"github.com/yijizhichang/wechat-sdk"
	"github.com/yijizhichang/wechat-sdk/examples/example"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"net/http"
)

//测试参数
//go run example.go -appid='your appdi' -appsecret='your appsecret' -token='your token' -cache='file' -port=':80'
var appid = flag.String("appid", "your AppID", "appid")
var appsecret = flag.String("appsecret", "your AppSecret", "appsecret")
var token = flag.String("token", "your Token", "token")
var cache = flag.String("cache", "file", "cache")
var aeskey = flag.String("asekey", "your EncodingAesKey", "asekey")
var port = flag.String("port", "80", "port")

//参数配置
func getWxConfig() *wechat.Config {
	config := &wechat.Config{
		AppID:            *appid,     //开发者ID(AppID)
		AppSecret:        *appsecret, //开发者PWD AppSecret
		Token:            *token,     //令牌(Token)
		EncodingAESKey:   *aeskey,    //消息加解密密钥 EncodingAESKey
		PayMchId:         "",         //支付 - 商户 ID
		PayNotifyUrl:     "",         //支付 - 接受微信支付结果通知的接口地址
		PayKey:           "",         //支付 - 商户后台设置的支付 key
		CacheModel:       *cache,     //缓存方式 默认为file，可选 file,redis,redisCluster
		ThirdAccessToken: false,      //是否使用第三方accessToken
		ProxyUrl:         "",         //代理
		CacheConfig: &wechat.CacheConfig{ //缓存配置
			FilePath: "./debug/cache/cache.txt", //缓存文件路径  CacheModel = "file" 时有效
			RedisConfig: &wechat.RedisConfig{
				Addr:     "127.0.0.1:6370", //Redis 地址，CacheModel = "redis" 时有效
				Password: "your redis pwd", //Redis PWD
			},
			RedisClusterConfig: &wechat.RedisClusterConfig{
				Addr:     []string{"127.0.0.1:6370", "127.0.0.1:6370"}, //RedisCluster 地址，CacheModel = "redisCluster" 时有效
				Password: "your redis pwd",                             //RedisCluster PWD
			},
		},
		FlogConfig: &wechat.FlogConfig{
			LogLevel:    1,                  //日志级别 =0 ALL; =1 DEBUG; =2 INFO; =3 WARN; =4 ERROR; =5 FATAL; =6 ALERT; =7 OFF;  注意：测试可以设置DEBUG;线上设置INFO 或 ERROR
			IsConsole:   true,               //是否输出到控制台
			IsFile:      true,               //是否写文件
			FilePath:    "./debug/log/",     //文件日志路径
			Filename:    "wechat-sdk",       //文件名称
			FileSuffix:  "txt",              //文件后缀
			FileMaxSize: 1024 * 1024 * 1024, //单个日志文件大小 单位B, 1024 * 1024 * 1024 为1G
		},
	}
	return config
}

func main() {
	//NewWechat
	wxconf.WechatClient = wechat.NewWechat(getWxConfig())

	flag.Parse()
	fmt.Println(*appid, *appsecret, *token, *cache, *port)

	//测试样例
	//example.SendTemplateMsg()  //模板消息发送
	//example.MassManage() //群发

	//example.CustomManage()  //客服管理
	//example.CustomSession()  //客服会话
	//example.CustomMsg()  	//聊天记录
	//example.SendCustomMsg()  //客服消息发送

	//example.UserTagManage()  //用户标签管理
	//example.UserRemark()  //用户备注
	//example.UserInfo()  //用户基本信息
	//example.UserList()  //用户列表
	//example.UserBlack()  //用户黑名单管理

	//example.Comment()   //评论管理

	//example.MenuCreate()   //创建菜单
	//example.MenuAddConditional()   //创建个性化菜单
	//example.MenuGet()   //获取所有菜单
	//example.MenuDel()  //删除所有菜单
	//example.MenuDelConditional()  //删除个性化菜单
	//example.MenuTryMatch() //个性化菜单匹配测试
	//example.GetSelfMenuInfo() //获取自定义菜单配置

	//example.AccountManage() //二维码管理

	http.HandleFunc("/serve", example.Serve)                               //server 服务
	http.HandleFunc("/customHeadImg", example.CustomHeadImg)               //设置客服头像
	http.HandleFunc("/uploadTempMedia", example.UploadTempMedia)           //上传临时素材
	http.HandleFunc("/uploadNewsPermanent", example.UploadNewsPermanent)   //上传永久图文素材
	http.HandleFunc("/uploadVideoPermanent", example.UploadVideoPermanent) //上传永久视频素材
	http.HandleFunc("/getMaterialMediaList", example.GetMaterialMediaList) //获取永久素材列表
	http.HandleFunc("/getMaterialMediaInfo", example.GetMaterialMediaInfo) //获取永久素材info
	http.HandleFunc("/delMaterialMedia", example.DelMaterialMedia)         //获取永久素材
	http.HandleFunc("/updateNewsMedia", example.UpdateNewsMedia)           //修改图文永久素材

	err := http.ListenAndServe("127.0.0.1:"+*port, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}

}

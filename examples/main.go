package main

import (
	"flag"
	"fmt"
	"github.com/yijizhichang/wechat-sdk"
	"github.com/yijizhichang/wechat-sdk/examples/cache"
	"github.com/yijizhichang/wechat-sdk/examples/example"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	wxCache "github.com/yijizhichang/wechat-sdk/util/cache"
	"net/http"
)

// 测试参数
// go run example.go -appid='your appdi' -appsecret='your appsecret' -token='your token' -port='80'
var appid = flag.String("appid", "your AppID", "appid")
var appsecret = flag.String("appsecret", "your AppSecret", "appsecret")
var token = flag.String("token", "your Token", "token")
var aeskey = flag.String("asekey", "your EncodingAesKey", "asekey")
var port = flag.String("port", "80", "port")

//qw
// go run main.go -cropid='your cropid' -token='your token' -aeskey='your aeskey' -cropsecret='your cropsecret' -rpkey='your rpkey' -port='80'
// go run main.go -cropid='111' -cropsecret='222' -rpkey='333'
var cropid = flag.String("cropid", "your CorpID", "cropid")
var cropsecret = flag.String("cropsecret", "your CorpSecret", "cropsecret")
var rpkey = flag.String("rpkey", "your RasPrivateKey", "rpkey")

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

func getQyWechatConfig(fileCache cache.FileClient) *wechat.QyConfig {
	config := &wechat.QyConfig{
		CorpID:           *cropid,     // 企业ID
		CorpSecret:       *cropsecret, // 应用的凭证密钥; 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
		RasPrivateKey:    *rpkey,      // 消息加密私钥
		Token:            *token,      // 令牌(Token)
		EncodingAESKey:   *aeskey,     // 消息加解密密钥 EncodingAESKey
		ThirdAccessToken: false,       //是用其他应用生成的access_token
		Cache:            fileCache,   //缓存
		ProxyUrl:         "",          //代理地址
	}
	return config
}

func main() {
	// NewWechat
	flag.Parse()
	//缓存 -file模式
	//fileCache,_ := cache.NewFile("./cache/cache.txt")

	//缓存 -redis
	//redisClusterCache, _ := cache.NewRedisCluster(&cache.RedisCluster{
	//	Addrs:    []string{"127.0.0.1:6379","127.0.0.1:6380"},
	//	Password: "123456",
	//})

	//微信公众号
	//wxconf.WechatClient = wechat.NewWechat(getWxConfig(fileCache))
	//fmt.Println(*appid, *appsecret, *token, *port)

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
	//http.HandleFunc("/wechat/server", example.ResponseServe)               // server 服务 数据先返回应用
	//http.HandleFunc("/customHeadImg", example.CustomHeadImg)               // 设置客服头像
	//http.HandleFunc("/uploadTempMedia", example.UploadTempMedia)           // 上传临时素材
	//http.HandleFunc("/uploadNewsPermanent", example.UploadNewsPermanent)   // 上传永久图文素材
	//http.HandleFunc("/uploadVideoPermanent", example.UploadVideoPermanent) // 上传永久视频素材
	//http.HandleFunc("/getMaterialMediaList", example.GetMaterialMediaList) // 获取永久素材列表
	//http.HandleFunc("/getMaterialMediaInfo", example.GetMaterialMediaInfo) // 获取永久素材info
	//http.HandleFunc("/delMaterialMedia", example.DelMaterialMedia)         // 获取永久素材
	//http.HandleFunc("/updateNewsMedia", example.UpdateNewsMedia)           // 修改图文永久素材

	/*****企业微信*****/
	file, err := cache.NewFile("./cache/cache.txt")
	if err != nil {
		fmt.Printf("cache.NewFile Err: %+v", err)
	}
	fileCache := cache.NewFileClient(file)
	wxconf.QyWechatClient = wechat.NewQyWechat(getQyWechatConfig(fileCache))
	fmt.Printf("qy param:", *cropid, *token, *aeskey, *cropsecret, *rpkey)

	//测试
	//example.QyAccessToken()  //获取企业access_token

	token, err := wxconf.QyWechatClient.GetQyAccessToken(*cropsecret)
	if err != nil {
		fmt.Printf("GetQyAccessToken Err: %+v", err)
	}
	fmt.Println("token=%s", token)
	//example.QyGetCustomerTagList(token)  //获取企业标签库
	//example.QyCreateCustomerTag(token)  //创建企业标签
	//example.QyUpdateCustomerTag(token)  //修改企业标签
	//example.QyDelCustomerTag(token) //删除企业标签
	//example.QyMarkTag(token) //编辑客户标签

	//example.QyGetDepartment(token) //获取部门列表
	//example.QyGetDepartmentSimpleUserList(token) //获取部门成员
	//example.QyGetDepartmentUserList(token) //获取部门成员详细

	//example.QyGetCustomerList(token) // 获取指定客户下的外部联系人external_userid
	//example.QyGetCustomerView(token) // 获取客户详情

	http.HandleFunc("/uploadTempMedia", example.WebhookUpload) // 上传webhook文件

	http.HandleFunc("/qw/server", example.QyServe) // server 服务
	//http.HandleFunc("/wechat/qy/server", example.QyResponseServe)               // server 服务 数据先返回应用

	err = http.ListenAndServe("127.0.0.1:"+*port, nil)
	if err != nil {
		fmt.Printf("start server error , err=%v", err)
	}
}

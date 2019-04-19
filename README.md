# wechat SDK for Go

## 概述
| 模块    | 描述                     |
|--------:|:-------------------------|
| mp      | 微信公众平台 SDK         |
| util    | 公共文件                 |
| examples| Demo文件                 |

## 安装
go get -u github.com/yijizhichang/wechat-sdk

## 快速开始

微信交互的运行流程:微信服务器把用户消息转到我们的自有服务器（虚线返回部分） 后的处理过程

```go
                                 +-----------------+                       +---------------+
+----------+                     |                 |    POST/GET/PUT       |               |
|          | ------------------> |                 | ------------------->  |               |
|   user   |                     |  wechat server  |                       |  your server  |
|          | < - - - - - - - - - |                 |                       |               |
+----------+                     |                 | <- - - - - - - - - -  |               |
                                 +-----------------+                       +---------------+
```
完成服务器端验证与接收响应用户发送的消息Demo:

```go
	//配置微信参数

	config := &wechat.Config{
    		AppID:            "your appId",  //开发者ID(AppID)
    		AppSecret:        "your appSecret",	//开发者PWD AppSecret
    		Token:            "your token",	//令牌(Token)
    		EncodingAESKey:   "your encoding aes key",		//消息加解密密钥 EncodingAESKey
    		PayMchId:         "",       //支付 - 商户 ID
    		PayNotifyUrl:     "",       //支付 - 接受微信支付结果通知的接口地址
    		PayKey:           "",       //支付 - 商户后台设置的支付 key
    		CacheModel:       "file",   //缓存方式 默认为file，可选 file,redis,redisCluster
    		ThirdAccessToken:  false,	//是否使用第三方accessToken
    		ProxyUrl:          "",		//代理
    		CacheConfig: &wechat.CacheConfig{  //缓存配置
    			FilePath: "./util/debug2/cache.txt",   //缓存文件路径  CacheModel = "file" 时有效
    			RedisConfig: &wechat.RedisConfig{
    				Addr: "127.0.0.1:6370",	  	//Redis 地址，CacheModel = "redis" 时有效
    				Password: "your redis pwd",   		//Redis PWD 
    			},
    			RedisClusterConfig: &wechat.RedisClusterConfig{
    				Addr: []string{"127.0.0.1:6370", "127.0.0.1:6370"},		//RedisCluster 地址，CacheModel = "redisCluster" 时有效
    				Password: "your redis pwd",  									//RedisCluster PWD 
    			},
    		},
    		FlogConfig: &wechat.FlogConfig{
    			LogLevel: 	  1,					//日志级别 =0 ALL; =1 DEBUG; =2 INFO; =3 WARN; =4 ERROR; =5 FATAL; =6 ALERT; =7 OFF;  注意：测试可以设置DEBUG;线上设置INFO 或 ERROR
    			IsConsole:    true,					//是否输出到控制台
    			IsFile:       true,					//是否写文件
    			FilePath:     "./util/debug/",		//文件日志路径
    			Filename:     "wechat-sdk",			//文件名称
    			FileSuffix:   "txt",				//文件后缀
    			FileMaxSize:  1024*1024*1024,		//单个日志文件大小 单位B, 1024 * 1024 * 1024 为1G
    		},
    	}

	wc := wechat.NewWechat(config)

    server := wc.GetServer(request, responseWriter)  // 传入request和responseWriter

    //设置接收消息的处理方法
    server.SetMessageHandler(func(msg message.MixMessage) *response.Reply {
        reStr = response.NewText("回复微信")
        return &response.Reply{MsgType: msgType, MsgData: reStr}
    })
    server.Serve()
    server.Send()
```
详细Demo：[examples/example/serve.go](examples/example/serve.go)

## 配置说明

```go
    config := &wechat.Config{
            AppID:            "your appId",  //开发者ID(AppID)
            AppSecret:        "your appSecret",	//开发者PWD AppSecret
            Token:            "your token",	//令牌(Token)
            EncodingAESKey:   "your encoding aes key",		//消息加解密密钥 EncodingAESKey
            PayMchId:         "",       //支付 - 商户 ID
            PayNotifyUrl:     "",       //支付 - 接受微信支付结果通知的接口地址
            PayKey:           "",       //支付 - 商户后台设置的支付 key
            CacheModel:       "file",   //缓存方式 默认为file，可选 file,redis,redisCluster
            ThirdAccessToken:  false,	//是否使用第三方accessToken
            ProxyUrl:          "",		//代理
            CacheConfig: &wechat.CacheConfig{  //缓存配置
                FilePath: "./util/debug2/cache.txt",   //缓存文件路径  CacheModel = "file" 时有效
                RedisConfig: &wechat.RedisConfig{
                    Addr: "127.0.0.1:6370",	  	//Redis 地址，CacheModel = "redis" 时有效
                    Password: "your redis pwd",   		//Redis PWD 
                },
                RedisClusterConfig: &wechat.RedisClusterConfig{
                    Addr: []string{"127.0.0.1:6370", "127.0.0.1:6370"},		//RedisCluster 地址，CacheModel = "redisCluster" 时有效
                    Password: "your redis pwd",  									//RedisCluster PWD 
                },
            },
            FlogConfig: &wechat.FlogConfig{
                LogLevel: 	  1,					//日志级别 =0 ALL; =1 DEBUG; =2 INFO; =3 WARN; =4 ERROR; =5 FATAL; =6 ALERT; =7 OFF;  注意：测试可以设置DEBUG;线上设置INFO 或 ERROR
                IsConsole:    true,					//是否输出到控制台
                IsFile:       true,					//是否写文件
                FilePath:     "./util/debug/",		//文件日志路径
                Filename:     "wechat-sdk",			//文件名称
                FileSuffix:   "txt",				//文件后缀
                FileMaxSize:  1024*1024*1024,		//单个日志文件大小 单位B, 1024 * 1024 * 1024 为1G
            },
        }
```
##### CacheModel 说明：
Cache主要用来保存全局的，access_token，可以选file,redis,redisCluster等模式

CacheModel="file"
CacheConfig中的FilePath，需要填写日志目录

CacheModel="redis"
CacheConfig中的RedisConfig，需要填redis相关配置

CacheModel="redisCluster"
CacheConfig中的RedisClusterConfig，需要填写redisCluster相关配置

##### ThirdAccessToken 说明：
正常设置为false 即可
如果遇到下面场景，设置为true
共享其它项目已有的access_token，当前项目不再从微信更新维护access_token,只是利用其它项目中已存在access_token，来发消息或其它操作时，可设置为true
如果设置为true,需要从其它项目中把access_token取到，并设置保存到缓存中，这样通过实例化后，调用发消息或其它操作时，默认取ThirdAccessToken
设置方法：`wc.Context.SetThirdAccessToken("your third_access_token",600)`,缓存时间单位为秒

##### ProxyUrl 说明:
如果需要设置代理,请填写自己的代理地址。格式："http://10.10.10.10:8080/"

##### FlogConfig 说明：
配置sdk相关日志信息，日志记录与微信的交互过程中的一些调试，报错信息。
日志级别：测试可以设置DEBUG;线上设置INFO 或 ERROR



## 基本的API调用

- [消息管理](mp/message/README.md)
- [客服管理](mp/custom/README.md)
- [群发消息](mp/message/mass/README.md)
- [素材管理](mp/media/README.md)
- [图文消息留言管理](mp/media/README.md)
- [用户管理](mp/user/README.md)
- [菜单管理](mp/menu/README.md)
- [微信网页授权](mp/oauth2/README.md)
- [账户管理](mp/account/README.md)
- [jssdk](mp/jssdk/README.md)

## 常见问题

- [被动回复微信消息失败](FAQ.md#被动回复微信消息失败)

```go
    该公众号暂时无法提供服务，请稍后再试
```
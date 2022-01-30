# 企业微信

## 快速开始

```go
	//企业微信配置
    qyConfigWechat := &wechat.QyConfig{
        CorpID:        		"your qyCorpID",  // 企业ID
        CorpSecret:    		"",  // 应用的凭证密钥; 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
        RasPrivateKey: 		"your qyRasPrivateKey",  // 消息加密私钥
        Token:              "your qyToken",    // 令牌(Token)
        EncodingAESKey:     "your qyEncodingAESKey",    // 消息加解密密钥 EncodingAESKey
        ThirdAccessToken: 	false,    //是用其他应用生成的access_token
        Debug: false, //true sdk会打印一些调试信息
        Cache:      cacheModel, //缓存方式 可选 file,redis,redisCluster 来实现cache的接口
        ProxyUrl:   "your http_proxy",      //代理地址
    
    }
```
##### CacheModel 说明：
Cache主要用来保存全局的，access_token，可以选file,redis,redisCluster等模式

## 基本的API调用

- [通讯录管理](work/company/README.md)
- [客户联系](work/customer/README.md)
- [微信客服](work/kefu/README.md)
- [身份验证](work/oauth/README.md)
- [应用管理](work/agent/README.md)
- [消息推送](work/message/README.md)
- [素材管理](work/media/README.md)
- [OA]()
- [效率工具](work/tools/README.md)


## 常见问题

- [access_token错误](FAQ.md#access_token错误)

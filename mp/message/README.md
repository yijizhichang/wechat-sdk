## 消息管理

- [消息管理](#消息管理)
	- [接收消息](#接收普通消息和事件)
	    - 接收普通消息
        - 接收事件推送
	- [被动回复消息](#被动回复消息)
		- 回复文本消息
		- 回复图片消息
		- 回复视频消息
		- 回复音乐消息
		- 回复图文消息
		- 将消息转发到客服
	- [模板消息](#模板消息)



## 消息管理

通过`wechat.GetServer(request,responseWriter)`获取到server对象之后

调用`SetMessageHandler(func(msg message.MixMessage){})`设置消息的处理函数，函数参数为message.MixMessage 结构如下：

```go
//存放所有微信发送过来的消息和事件
type MixMessage struct {
	MsgCommon

	//基本消息
	MsgID        	int64   	`xml:"MsgId"`
	Content      	string  	`xml:"Content"`
	Recognition  	string  	`xml:"Recognition"`
	PicURL       	string  	`xml:"PicUrl"`
	MediaID      	string  	`xml:"MediaId"`
	Format       	string  	`xml:"Format"`
	ThumbMediaID 	string  	`xml:"ThumbMediaId"`
	LocationX    	float64 	`xml:"Location_X"`
	LocationY    	float64 	`xml:"Location_Y"`
	Scale        	float64 	`xml:"Scale"`
	Label        	string  	`xml:"Label"`
	Title        	string  	`xml:"Title"`
	Description  	string  	`xml:"Description"`
	URL          	string  	`xml:"Url"`

	//事件相关
	Event     		EventType 	`xml:"Event"`
	EventKey  		string    	`xml:"EventKey"`
	Ticket    		string    	`xml:"Ticket"`
	Latitude  		float64    	`xml:"Latitude"`
	Longitude 		float64    	`xml:"Longitude"`
	Precision 		float64    	`xml:"Precision"`
	MenuID    		string    	`xml:"MenuId"`
	Status    		string    	`xml:"Status"`
	SessionFrom     string  	`xml:"SessionFrom"`

	ScanCodeInfo struct {
		ScanType   		string 		`xml:"ScanType"`
		ScanResult 		string 		`xml:"ScanResult"`
	} `xml:"ScanCodeInfo"`

	SendPicsInfo struct {
		Count   	int32      		`xml:"Count"`
		PicList 	[]EventPic 		`xml:"PicList>item"`
	} `xml:"SendPicsInfo"`

	SendLocationInfo struct {
		LocationX 		float64 	`xml:"Location_X"`
		LocationY 		float64 	`xml:"Location_Y"`
		Scale     		float64 	`xml:"Scale"`
		Label     		string  	`xml:"Label"`
		Poiname   		string  	`xml:"Poiname"`
	} `xml:"SendLocationInfo"`
}

```

具体参数请参考微信文档：[接收普通消息
](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140453)

### 接收普通消息和事件
```go
server.SetMessageHandler(func(v message.MixMessage) *message.Reply {

		//根据微信回调时的消息类型，来相应获取对应消息明细
		switch msg.MsgCommon.MsgType {

			//消息类型
			case "text":
				reMsg = request.GetText(&msg)
			case "image":
				reMsg = request.GetImage(&msg)
			case "voice":
				reMsg = request.GetVoice(&msg)
			case "video":
				reMsg = request.GetVideo(&msg)
			case "shortvideo":
				reMsg = request.GetShortVideo(&msg)
			case "location":
				reMsg = request.GetLocation(&msg)
			case "link":
				reMsg = request.GetLink(&msg)

			//事件类型
			case "event":
				switch msg.Event {
    				case "subscribe":
    					reMsg = request.GetSubscribeEvent(&msg)
    				case "unsubscribe":
    					reMsg = request.GetUnsubscribeEvent(&msg)
    				case "SCAN":
    					reMsg = request.GetScanEvent(&msg)
    				case "CLICK","VIEW":
    					reMsg = request.GetMenuEvent(&msg)
    				case "TEMPLATESENDJOBFINISH":
    					reMsg = request.GetTemplateSendJobFinishEvent(&msg)
				}
		}

		log.Println("消息明细：",reMsg)
}


```



### 被动回复消息

回复消息需要返回 `*response.Reply` 对象结构体如下：

```go
type Reply struct{
	MsgType message.MsgType  //消息类型
	MsgData interface{}     //消息内容
}

```
注意：`return nil`表示什么也不做

####  回复文本消息
```go
	text := response.NewText("回复文本消息")
	return &response.Reply{message.MsgTypeText, text}
```
####  回复图片消息
```go
image :=response.NewImage("mediaID")
return &response.Reply{message.MsgTypeVideo, image}
```
####  回复语音消息
```go
voice :=response.NewVoice("mediaID")
return &response.Reply{message.MsgTypeVoice, voice}
```
####  回复视频消息
```go
video := response.NewVideo("mediaID", "视频标题", "视频描述")
return &response.Reply{message.MsgTypeVideo, video}
```
####  回复音乐消息
```go
music := response.NewMusic("title", "description", "musicURL", "hQMusicURL", "thumbMediaID")
return &response.Reply{message.MsgTypeMusic,music}
```
**字段说明：**

Title:音乐标题
Description:音乐描述
MusicURL:音乐链接
HQMusicUrl：高质量音乐链接，WIFI环境优先使用该链接播放音乐
ThumbMediaId：缩略图的媒体id，通过素材管理接口上传多媒体文件，得到的id

####  回复图文消息

```go
articles := make([]*message.Article, 1)
article := response.NewArticle("图文标题","图文消息描述","https://www.baidu.com/img/bd_logo1.png","https://www.baidu.com/")
articles[0] = article
news := response.NewNews(articles)
return &response.Reply{message.MsgTypeNews,news}
```
**字段说明：**

Title：图文消息标题
Description：图文消息描述
PicUrl	：图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
Url	：点击图文消息跳转链接


####  将消息转发到客服

```go
	transferKf := response.NewTransferKf("")  //可指定客服账号
	return &response.Reply{message.MsgTypeTransfer, transferKf}
```

## 模板消息

```go
    config := &wechat.Config{
        AppID:          "your appId",
        AppSecret:      "your appSecret",
        Token:          "your token",
        EncodingAESKey: "your encoding aes key",
        CacheModel: 	"file",  //缓存方式 file,redis,redisCluster
        ThirdAccessToken: false,
    }

    wc := wechat.NewWechat(config)

	//如果ThirdAccessToken=true,则需要把其它项目中的access_token设置到缓存中
	//注意：为避免access_token失效，需要自己维护定时从其它项目中获取设置到缓存中
	//wc.Context.SetThirdAccessToken("your third_access_token",600) //缓存时间单位为秒

	//实例化发送模板消息
	tpl := wc.GetTemplate()


    //模板内容： {{first.DATA}} 投资标的：{{keyword1.DATA}} 当期收益：{{keyword2.DATA}} {{remark.DATA}}

    //模板消息样例
	first := &template.DataItem{
		"亲爱的xxx，您有一笔消费哦~",
		"#FF4040",
	}
	remark:=&template.DataItem{
		"感谢您的使用,祝您生活愉快~",
		"#FF00FF",
	}
	keyword1 := &template.DataItem{
		"500大洋",
		"#CAFF70",
	}
	keyword2:=&template.DataItem{
		"手机",
		"#9932CC",
	}

	msgTpl := new(template.Message)
	msgTpl.ToUser ="oZZwr0REd0cbVxxxQQmicxS3FrI0"
	msgTpl.TemplateID ="CYM5ydWPBFvx__SDFFQN_DwHH8pjO3o1RyjvYqCTdc"
	msgTpl.Data = make(map[string]*template.DataItem)
	msgTpl.Data["first"] = first
	msgTpl.Data["keyword1"] = keyword1
	msgTpl.Data["keyword2"] = keyword2
	msgTpl.Data["remark"] = remark


    //发送模板消息
	re,_ := tpl.Send(msgTpl)
	log.Println("返回发送模板消息的结果：",re)
```

详细Demo：[examples/example/template.go](examples/example/template.go)
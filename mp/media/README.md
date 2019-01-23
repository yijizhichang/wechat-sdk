## 素材管理

- [素材管理](#素材管理)
	- [临时素材](#临时素材)
	    - 新增临时素材
        - 获取临时素材
	- [永久素材](#永久素材)
		- 新增永久素材
		- 获取永久素材
		- 删除永久素材
		- 修改永久图文素材
		- 获取素材总数
		- 获取素材列表
	- [图文消息留言管理](#图文消息留言管理)
	    - 打开已群发文章评论
	    - 关闭已群发文章评论
	    - 查看指定文章的评论数据
	    - 将评论标记精选
	    - 将评论取消精选
	    - 删除评论
	    - 回复评论
	    - 删除回复

## 素材管理

```go
公众号经常有需要用到一些临时性的多媒体素材的场景，
例如在使用接口特别是发送消息时，对多媒体文件、多媒体消息的获取和调用等操作，是通过media_id来进行的。
素材管理接口对所有认证的订阅号和服务号开放。


```

具体参数请参考微信文档：[素材管理](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1444738726)


## 临时素材

##### 上传临时素材
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()

	#上传临时素材  类型，文件路径  用FORM表单方式上传一个多媒体文件
	//支持类型： 图片（image）,语音（voice）,视频（video）,缩略图（thumb）
    re,err :=mda.UploadTempMedia("image","./util/upload/test.jpg")

```

##### 获取临时素材
```go
    //config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()

    #获取临时素材  返回下载地址,根据需要下载对应文件
    downUrl,err := mda.GetTempMediaUrl("wqeadsfofqweasdfnwefadf")

```

## 永久素材

##### 新增永久素材
```go

    #因不同类型的素材所需要字段不一样,所以分多个方法实现

    #新增永久素材  图片（image）,语音（voice）,缩略图（thumb）
    //类型，文件路径  用FORM表单方式上传一个多媒体文件
    re,err :=mda.AddMaterialMedia("image","./util/upload/test.jpg")

    #新增视频永久素材 视频（video）
    //标题，视频描述，视频文件  用FORM表单方式上传一个多媒体文件
    re,err :=mda.AddVideoMedia("title", "video desc","./util/upload/test.mp4")


    #上传图文素材  //arList最多8条
	ar := media.NewArticle("title", "media_id", "author", "digest", 0, "test content ...", "http://wx.qq.com", 0, 0)
	var arList []*media.Article
	arList = append(arList, ar)

	news := media.NewNews(arList)

	re2,err :=mda.AddNewsPermanent(news)

	//参数说明
	title	是	标题
    thumb_media_id	是	图文消息的封面图片素材id（必须是永久mediaID）
    author	否	作者
    digest	否	图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空。如果本字段为没有填写，则默认抓取正文前64个字。
    show_cover_pic	是	是否显示封面，0为false，即不显示，1为true，即显示
    content	是	图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS,涉及图片url必须来源 "上传图文消息内的图片获取URL"接口获取。外部图片url将被过滤。
    content_source_url	是	图文消息的原文地址，即点击“阅读原文”后的URL
    need_open_comment	否	Uint32 是否打开评论，0不打开，1打开
    only_fans_can_comment	否	Uint32 是否粉丝才可评论，0所有人可评论，1粉丝才可评论


    #上传图文消息内的图片获取URL
    //如果内容详情中有图片素材，则利用此方法处理
    re,err := mda.MediaUploadImg("./util/upload/test.jpg")
    re.ImgUrl  为图片链接，可插入到详情中使用



```

##### 获取永久素材

```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()

    #获取图文素材 返回是对应结构体类型
    re,err :=mda.GetNewsMediaInfo(media_id)

    //返回参数说明
    {
        "news_item":
        [
            {
            "title":TITLE,
            "thumb_media_id"::THUMB_MEDIA_ID,
            "show_cover_pic":SHOW_COVER_PIC(0/1),
            "author":AUTHOR,
            "digest":DIGEST,
            "content":CONTENT,
            "url":URL,
            "content_source_url":CONTENT_SOURCE_URL
            },
            //多图文消息有多篇文章
        ]
    }


    #获取视频素材 返回是对应结构体类型
    re,err :=mda.GetVideoMediaInfo(media_id)

    //返回参数说明
    {
      "title":TITLE,
      "description":DESCRIPTION,
      "down_url":DOWN_URL,
    }


    #获取其它素材 image,voice,thumb 返回素材的内容，开发者可自行保存为文件
    re,err :=mda.GetOtherMediaInfo(media_id)
	ioutil.WriteFile("./debug/down/test.jpg", re, 0666)   //直接保存为对应类型的文件


```


##### 删除永久素材

```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()

	#删除永久素材
    re,err :=mda.DelMaterialMedia(media_id)

```

##### 修改永久图文素材
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()


    #修改图文素材
	updata := new(media.UpdateNewsMedia)
	updata.MediaId = "0DGmfM0mkFDafFtgztW0nRW9M9JKQNDl-VRDCvsldQk"
	updata.Index = 0
	updata.Articles.Title = "new title"
	updata.Articles.Content ="new content"
	updata.Articles.Author = "new author"
	updata.Articles.ThumbMediaId = "0DGmfM0mkFDafFtgztW0naWpRTDuQdT3urOSTjrjm7g"
	updata.Articles.Digest = "new digest"
	updata.Articles.ContentSourceUrl = "http://newwx.qq.com"
	updata.Articles.ShowCoverPic = 1
	re, err := mda.UpdateNewsMedia(updata)

    //参数说明
    media_id	是	要修改的图文消息的id
    index	    是	要更新的文章在图文消息中的位置（多图文消息时，此字段才有意义），第一篇为0
    title	    是	标题
    thumb_media_id	是	图文消息的封面图片素材id（必须是永久mediaID）
    author	    是	作者
    digest	    是	图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空
    show_cover_pic	是	是否显示封面，0为false，即不显示，1为true，即显示
    content	    是	图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS
    content_source_url	是	图文消息的原文地址，即点击“阅读原文”后的URL
```


##### 获取素材总数
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()


    #获取素材总数
    re,err :=media.GetMaterialMediaCount()

    //返回参数说明
    voice_count	语音总数量
    video_count	视频总数量
    image_count	图片总数量
    news_count	图文总数量
```


##### 获取素材列表
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	mda := wc.GetMedia()

	#获取永久图文消息素材列表  返回对应结构体
	re2,err := mda.GetNewsMediaList(0,4)

	//返回参数说明
	{
       "total_count": TOTAL_COUNT,
       "item_count": ITEM_COUNT,
       "item": [{
           "media_id": MEDIA_ID,
           "content": {
               "news_item": [{
                   "title": TITLE,
                   "thumb_media_id": THUMB_MEDIA_ID,
                   "show_cover_pic": SHOW_COVER_PIC(0 / 1),
                   "author": AUTHOR,
                   "digest": DIGEST,
                   "content": CONTENT,
                   "url": URL,
                   "content_source_url": CONTETN_SOURCE_URL
               },
               //多图文消息会在此处有多篇文章
               ]
            },
            "update_time": UPDATE_TIME
        },
        //可能有多个图文消息item结构
      ]
    }


    #其他类型（图片、语音、视频）  返回对应结构体
    re,err := mda.GetOtherMediaList

    //返回参数说明
    {
       "total_count": TOTAL_COUNT,
       "item_count": ITEM_COUNT,
       "item": [{
           "media_id": MEDIA_ID,
           "name": NAME,
           "update_time": UPDATE_TIME,
           "url":URL
       },
       //可能会有多个素材
       ]
    }
```


## 图文消息留言管理

##### 打开已群发文章评论
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err := media.OpenComment(2247483665,1)

    参数	        是否必须	    类型	    说明
    msg_data_id	    是	    Uint32	群发返回的msg_data_id
    index	          否	    Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文

```

##### 关闭已群发文章评论
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err := media.CloseComment(2247483665,1)

    参数	        是否必须	    类型	    说明
    msg_data_id	    是	    Uint32	群发返回的msg_data_id
    index	          否	    Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文

```

##### 查看指定文章的评论数据
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err := media.GetCommentList(2247483665,0,10,0,0)

    参数	是否必须	类型	    说明
msg_data_id      是	Uint32	群发返回的msg_data_id
    index	    否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文
    begin	    是	Uint32	起始位置
    count	    是	Uint32	获取数目（>=50会被拒绝）
    type 	    是	Uint32	type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论参数	是否必须	类型	说明

```

##### 将评论标记精选
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err :=  media.MarkElectComment(2247483665,0,1)

    参数	    是否必须	类型	说明
msg_data_id	    是	Uint32	群发返回的msg_data_id
    index  	    否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	Uint32	用户评论id
```

##### 将评论取消精选
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err :=  media.UnMarkElectComment(2247483665,0,1)

    参数	    是否必须	类型	说明
msg_data_id	    是	Uint32	群发返回的msg_data_id
    index  	    否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	Uint32	用户评论id
```

##### 删除评论
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err :=  media.DeleteComment(2247483665,0,1)

    参数	    是否必须	类型	说明
msg_data_id	    是	Uint32	群发返回的msg_data_id
    index  	    否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	Uint32	用户评论id
```

##### 回复评论
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err :=  media.ReplayComment(2247483665,0,1,"你最美")

    参数	    是否必须	类型	说明
msg_data_id    	是	Uint32	群发返回的msg_data_id
    index      	否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	Uint32	评论id
    content    	是	string	回复内容
```


##### 删除回复
```go
    #config配置文件省略...
	wc := wechat.NewWechat(config)
	media := wxconf.WechatClient.GetMedia()

    re, err :=  media.DeleteReplayComment(2247483665,0,1)

    参数	    是否必须	类型	说明
msg_data_id    	是	Uint32	群发返回的msg_data_id
    index      	否	Uint32	多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
user_comment_id	是	Uint32	评论id
```

详细Demo：[examples/example/media.go](examples/example/media.go)

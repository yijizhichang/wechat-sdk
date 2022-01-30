## 素材管理

- [素材管理](#素材管理)
    - [上传获取素材](#上传获取素材)
        - 上传临时素材
        - 上传图片
        - 获取临时素材
        - 获取高清语音素材
        
## 素材管理

```go
在使用企业微信API接口中，往往开发者需要使用自定义的资源，比如发送本地图片消息，设置通讯录自定义头像等。
为了实现同一资源文件，一次上传可以多次使用，这里提供了素材管理接口：以media_id来标识资源文件，实现文件的上传与下载。
```

具体参数请参考微信文档：[素材管理](https://developer.work.weixin.qq.com/document/path/91054)

## 上传获取素材
```go
    //config配置文件省略...
    qw := wechat.NewQyWechat(config)
    cli := qw.GetMedia()
    
    //上传临时素材
    re,err := cli.UploadQyTempMedia(accessToken, fileType, fieldname, filename)
    //上传图片
    re,err := cli.UploadQyImgMedia(accessToken, fieldname, filename)
    //获取临时素材
    re,err := cli.GetQyTempMediaURL(accessToken, mediaId)
    //获取高清语音素材
    re,err := cli.GetQyCustomerRemarkURL(accessToken, mediaId)
```

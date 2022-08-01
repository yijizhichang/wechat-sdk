/**
 * @Time : 2022/7/29 15:02
 * @Author : soupzhb@gmail.com
 * @File : qw_webhook_upload.go
 * @Software: GoLand
 */

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"io"
	"net/http"
	"os"
)

//上传临时素材
func WebhookUpload(rw http.ResponseWriter, req *http.Request) {
	//从请求当中判断方法
	if req.Method == "GET" {
		io.WriteString(rw, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传类型</label>"+":"+
			"<select name='type'><option value ='image'>图片</option><option value ='voice'>语音</option><option value ='video'>视频</option><option value ='thumb'>缩略图</option></select>"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传图片\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, head, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		upType := req.FormValue("type")
		fmt.Println("上传类型：", upType)

		//创建文件
		fW, err := os.Create("/tmp/upload/" + head.Filename)
		if err != nil {
			fmt.Println("文件创建失败")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("文件保存失败")
			return
		}

		key := "9691a6f0-701f-425c-9ff2-fd7187b6f108"
		//webhook 上传管理
		webhookClient := wxconf.QyWechatClient.GetWebhook()
		re, err := webhookClient.UploadQyTempMedia(key, "file", "/tmp/upload/"+head.Filename)

		fmt.Println("上传webhook文件：", re, "err:", err)

	}

}

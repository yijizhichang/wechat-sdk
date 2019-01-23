package example

import (
	"encoding/json"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/mp/media"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//上传临时素材
func UploadTempMedia(rw http.ResponseWriter, req *http.Request) {
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
		fW, err := os.Create("./debug/upload/" + head.Filename)
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

		//素材管理
		media := wxconf.WechatClient.GetMedia()
		re, err := media.UploadTempMedia(upType, "./debug/upload/"+head.Filename)
		fmt.Println("上传临时素材：", re, "err:", err)

		//获取临时素材下载地址
		murl, err := media.GetTempMediaUrl(re.MediaId)

		fmt.Println("临时素材下载地址：", murl, "err:", err)

		io.WriteString(rw, head.Filename+" 保存成功")
	}

}

//上传永久图文素材
func UploadNewsPermanent(rw http.ResponseWriter, req *http.Request) {
	//从请求当中判断方法
	if req.Method == "GET" {
		io.WriteString(rw, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传图片</label>"+":"+
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

		//创建文件
		fW, err := os.Create("./debug/upload/" + head.Filename)
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

		//素材管理
		media2 := wxconf.WechatClient.GetMedia()
		re, err := media2.AddMaterialMedia("image", "./debug/upload/"+head.Filename)
		fmt.Println("上传图文永久素材：", re, "err:", err)

		//上传图文素材
		ar := media.NewArticle("今日头条", re.MediaId, "author", "digest", 0, "test content ...", "http://wx.qq.com", 1, 0)
		var arList []*media.Article
		arList = append(arList, ar)

		ar2 := media.NewArticle("今日尾条", re.MediaId, "author", "digest", 0, "test content ...", "http://wx.qq.com", 1, 0)
		arList = append(arList, ar2)

		news := media.NewNews(arList)

		re2, err := media2.AddNewsPermanent(news)
		fmt.Println("上传图文素材：", re2, "err:", err)

		io.WriteString(rw, head.Filename+" 保存成功")
	}

}

//上传永久视频素材
func UploadVideoPermanent(rw http.ResponseWriter, req *http.Request) {
	//从请求当中判断方法
	if req.Method == "GET" {
		io.WriteString(rw, "<html><head><title>上传</title></head>"+
			"<body><form action='#' method=\"post\" enctype=\"multipart/form-data\">"+
			"<label>上传视频</label>"+":"+
			"<input type=\"file\" name='file'  /><br/><br/>    "+
			"<label><input type=\"submit\" value=\"上传视频\"/></label></form></body></html>")
	} else {
		//获取文件内容 要这样获取
		file, head, err := req.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		//创建文件
		fW, err := os.Create("./debug/upload/" + head.Filename)
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

		//素材管理
		media2 := wxconf.WechatClient.GetMedia()
		re, err := media2.AddVideoMedia("title", "video des", "./debug/upload/"+head.Filename)
		fmt.Println("上传永久视频素材：", re, "err:", err)

		io.WriteString(rw, head.Filename+" 保存成功")
	}

}

//获取图文永久素材列表
func GetMaterialMediaList(rw http.ResponseWriter, req *http.Request) {
	media := wxconf.WechatClient.GetMedia()
	re, err := media.GetMaterialMediaCount()
	fmt.Println("获取素材总量：", re, "err:", err)

	re2, err := media.GetNewsMediaList(0, 4)
	re2Json, err := json.Marshal(re2)
	fmt.Println("获取图文列表：", re2, "err:", err, "re2Json", string(re2Json))

	re3, err := media.GetOtherMediaList("video", 0, 4)
	re3Json, err := json.Marshal(re3)
	fmt.Println("获取视频列表：", re3, "err:", err, "re3Json", string(re3Json))
}

//获取图文永久素材info
func GetMaterialMediaInfo(rw http.ResponseWriter, req *http.Request) {
	media_id := req.FormValue("media_id")
	media_type := req.FormValue("media_type")

	media := wxconf.WechatClient.GetMedia()
	if media_type == "news" {
		re, err := media.GetNewsMediaInfo(media_id)
		reJson, err := json.Marshal(re)
		fmt.Println("获取图文info：", re, "err:", err, "reJson", string(reJson))
	} else if media_type == "video" {
		re, err := media.GetVideoMediaInfo(media_id)
		reJson, err := json.Marshal(re)
		fmt.Println("获取视频info：", re, "err:", err, "reJson", string(reJson))
	} else {
		re, err := media.GetOtherMediaInfo(media_id)
		reJson, err := json.Marshal(re)
		ioutil.WriteFile("./debug/down/test.jpg", re, 0666)
		fmt.Println("获取并保存其它素材：", re, "err:", err, "reJson", string(reJson))
	}
}

//删除永久素材
func DelMaterialMedia(rw http.ResponseWriter, req *http.Request) {
	media_id := req.FormValue("media_id")
	media := wxconf.WechatClient.GetMedia()
	re, err := media.DelMaterialMedia(media_id)
	fmt.Println("获取素材总量：", re, "err:", err)
}

//修改永久图文素材

func UpdateNewsMedia(rw http.ResponseWriter, req *http.Request) {
	media2 := wxconf.WechatClient.GetMedia()
	updata := new(media.UpdateNewsMedia)
	updata.MediaId = "0DGmfM0mkFDafFtgztW0nRW9M9JKQNDl-XXXXXXX"
	updata.Index = 0
	updata.Articles.Title = "new title"
	updata.Articles.Content = "new content"
	updata.Articles.Author = "new author"
	updata.Articles.ThumbMediaId = "0DGmfM0mkFDafFtgztW0naWpRTDuQdXXXXXXXXX"
	updata.Articles.Digest = "new digest"
	updata.Articles.ContentSourceUrl = "http://newwx.qq.com"
	updata.Articles.ShowCoverPic = 1
	re, err := media2.UpdateNewsMedia(updata)

	fmt.Println("修改永久图文素材：", re, "err:", err)
}

//评论管理
func Comment() {
	media := wxconf.WechatClient.GetMedia()

	//群发
	/*	m := wxconf.WechatClient.GetMass()
		res, err := m.MassSendall(
			mass.WithFilterOption(false, 101),
			mass.WithMpnewsOption("0DGmfM0mkFDafFtgztW0nQGly5N6SRbuwkYXXXXXXX",0),
		)
		fmt.Println("群发：", res, "err:", err)*/

	//打开已群发文章评论     {{0 send job submission success} 3147483649 2247483665}
	/*	re2, err := media.OpenComment(uint32(2247483665),1)
		fmt.Println("评论管理：", re2, err)*/

	//关闭已群发文章评论
	/*	re2, err := media.CloseComment(uint32(2247483665),1)
		fmt.Println("评论管理：", re2, err)*/

	//查看指定文章的评论数据
	/*	re2, err := media.GetCommentList(2247483665,0,10,0,0)
		fmt.Println("评论管理：", re2, err,re2.Comment[0].Content)*/

	//将评论标记精选
	/*	re2, err := media.MarkElectComment(2247483665,0,1)
		fmt.Println("评论管理：", re2, err,)*/

	//将评论取消精选
	/*	re2, err := media.UnMarkElectComment(2247483665,0,1)
		fmt.Println("评论管理：", re2, err,)*/

	//删除评论
	/*	re2, err := media.DeleteComment(2247483665,0,1)
		fmt.Println("评论管理：", re2, err)*/

	//回复评论
	/*	re2, err := media.ReplayComment(2247483665,0,1,"你最美")
		fmt.Println("评论管理：", re2, err)*/

	//删除回复
	re2, err := media.DeleteReplayComment(2247483665, 0, 1)
	fmt.Println("评论管理：", re2, err)
}

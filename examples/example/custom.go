package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/mp/custom"
	"io"
	"net/http"
	"os"
)

//客服管理
func CustomManage() {

	kf := wxconf.WechatClient.GetCustom()

	//re,err := kf.AddKfAccount("zhangsan33@test", "小宜33", )
	//re,err := kf.UpdateKfAccount("zhangsan@test2", "小宜aaaa")
	re, err := kf.DelKfAccount("zhangsan33@test")
	//re,err := kf.GetKfList()
	//fmt.Println(re.KfList[0].KfAccount,re.KfList[0].KfHeadImgUrl,re.KfList[0].KfNickName)
	//fmt.Println(re.KfList[1].KfAccount,re.KfList[1].KfHeadImgUrl,re.KfList[1].KfNickName)
	//re,err := kf.GetKfOnlineList()
	//fmt.Println(re.KfOnlineList[0].KfAccount,re.KfOnlineList[0].Status,re.KfOnlineList[0].AcceptedCase)
	//re, err := kf.InviteWorker("zhangsan33@test","whf-1020")
	fmt.Println("返回结果：", re, "Err:", err)
}

//上传客服头像
func CustomHeadImg(rw http.ResponseWriter, req *http.Request) {
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

		//kf
		kf := wxconf.WechatClient.GetCustom()
		re, err := kf.SetHeadImgURL("zhangsan@test", "./debug/upload/"+head.Filename)
		fmt.Println("设置微信头像：", re, "err:", err)

		io.WriteString(rw, head.Filename+" 保存成功")
	}

}

//会话控制
func CustomSession() {
	kf := wxconf.WechatClient.GetCustom()
	re, err := kf.CreateKfSession("zhangsan@test", "abcd1234abcd1234abcd1234")
	//re, err := kf.CloseKfSession("zhangsan@test","abcd1234abcd1234abcd1234")
	//re, err := kf.GetKfSession("abcd1234abcd1234abcd1234")
	//re, err := kf.GetKfSessionList("zhangsan@test")
	//re, err := kf.GetWaitCaseList()
	fmt.Println("会话控制返回结果：", re, "Err:", err)
}

//获取聊天记录
func CustomMsg() {
	kf := wxconf.WechatClient.GetCustom()
	re, err := kf.GetMsgList("2018-11-16 00:10:00", "2018-11-16 23:10:00", 1, 1000)
	fmt.Println("会话控制返回结果：", re, "Err:", err)
}

//客服消息发送
func SendCustomMsg() {
	kf := wxconf.WechatClient.GetCustom()

	//文本消息
	text := custom.NewText("客服文本消息发送测试")
	text.ToUser = "abcd1234abcd1234abcd1234"
	text.MsgType = "text"
	text.CustomService.KfAccount = "zhangsan@test"

	re, err := kf.SendMsgByKf(text)

	//图片消息
	/*	img := custom.NewImage("123124435245")
		img.ToUser = "abcd1234abcd1234abcd1234"
		img.MsgType = "image"
		img.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(img)
		fmt.Println(re,err)*/

	//语音
	/*	voice := custom.NewVoice("mediald")
		voice.ToUser = "abcd1234abcd1234abcd1234"
		voice.MsgType = "voice"
		voice.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(voice)*/

	//视频
	/*	video := custom.NewVideo("mediald", "thumbMediald", "title","description")
		video.ToUser = "abcd1234abcd1234abcd1234"
		video.MsgType = "video"
		video.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(video)*/

	//音乐
	/*	music := custom.NewMusic("title", "des	cription", "musicURL", "HQMusicUrl", "thumbMediald")
		music.ToUser = "abcd1234abcd1234abcd1234"
		music.MsgType = "video"
		music.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(music)*/

	//图文  发送图文消息（点击跳转到外链） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
	/*ar := custom.NewArticle("图文消息", "我是一条图文消息", "https://www.baidu.com/img/bd_logo1.png", "")
	var newsList []*custom.Article
	newsList = append(newsList, ar)
	ars := custom.NewNews(newsList)
	ars.ToUser = "abcd1234abcd1234abcd1234"
	ars.MsgType = "news"
	ars.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(ars)*/

	//mpnews  发送图文消息（点击跳转到图文消息页面） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
	/*mpnews := custom.NewMpNews("mediald")
	mpnews.ToUser = "abcd1234abcd1234abcd1234"
	mpnews.MsgType = "video"
	mpnews.CustomService.KfAccount = "zhangsan@test"
	re,err := kf.SendMsgByKf(mpnews)*/

	//卡券
	/*	card := custom.NewCard("cardid")
		card.ToUser = "abcd1234abcd1234abcd1234"
		card.MsgType = "video"
		card.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(card)*/

	//小程序卡片
	/*	minipg := custom.NewMiniProgramPage("title","appid","pagepath","thumbmediald")
		minipg.ToUser = "abcd1234abcd1234abcd1234"
		minipg.MsgType = "video"
		minipg.CustomService.KfAccount = "zhangsan@test"
		re,err := kf.SendMsgByKf(minipg)*/

	fmt.Println(re, err)
}

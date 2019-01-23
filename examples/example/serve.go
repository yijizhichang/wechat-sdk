package example

import (
	"net/http"
	"github.com/yijizhichang/wechat-sdk/mp/message"
	"github.com/yijizhichang/wechat-sdk/mp/message/callback/response"
	"github.com/yijizhichang/wechat-sdk/mp/message/callback/request"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

//与微信交互服务地址
func Serve(rw http.ResponseWriter, req *http.Request) {

	// 传入request和responseWriter
	server := wxconf.WechatClient.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *response.Reply {
		var reStr interface{}
		var msgType message.MsgType
		var getCon interface{}

		//根据微信回调时的消息类型，来相应获取对应消息明细
		switch msg.MsgCommon.MsgType {
		//消息类型
		case "text":
			getCon = request.GetText(&msg)
		case "image":
			getCon = request.GetImage(&msg)
		case "voice":
			getCon = request.GetVoice(&msg)
		case "video":
			getCon = request.GetVideo(&msg)
		case "shortvideo":
			getCon = request.GetShortVideo(&msg)
		case "location":
			getCon = request.GetLocation(&msg)
		case "link":
			getCon = request.GetLink(&msg)
			//事件类型
		case "event":
			switch msg.Event {
			case "subscribe":
				getCon = request.GetSubscribeEvent(&msg)
			case "unsubscribe":
				getCon = request.GetUnsubscribeEvent(&msg)
			case "SCAN":
				getCon = request.GetScanEvent(&msg)
			case "CLICK", "VIEW":
				getCon = request.GetMenuEvent(&msg)
			case "TEMPLATESENDJOBFINISH":
				getCon = request.GetTemplateSendJobFinishEvent(&msg)
			}
		}
		fmt.Println("test消息明细：", getCon)

		//根据业务需求,被动回复微信消息
		switch msg.Content {
		case "1":
			reStr = response.NewText("回复文件消息")
			msgType = message.MsgTypeText
		case "2":
			reStr = response.NewImage("9999999999")
			msgType = message.MsgTypeImage
		case "3":
			reStr = response.NewVoice("9999999999")
			msgType = message.MsgTypeVoice
		case "4":
			reStr = response.NewVideo("999999999", "视频", "我是一条视频信息")
			msgType = message.MsgTypeVideo
		case "5":
			ar := response.NewArticle("图文消息", "我是一条图文消息", "https://www.baidu.com/img/bd_logo1.png", "https://www.baidu.com/")
			var newsList []*response.Article
			newsList = append(newsList, ar)
			reStr = response.NewNews(newsList)
			fmt.Println("图文消息：", reStr)
			msgType = message.MsgTypeNews
		default:
			reStr = response.NewText("默认回复你消息吧")
			msgType = message.MsgTypeText
		}

		//转发到客服
/*		if msg.MsgCommon.MsgType == "text"{
			reStr = response.NewTransferKf("")
			msgType = message.MsgTypeTransfer
		}*/

		return &response.Reply{MsgType: msgType, MsgData: reStr}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}



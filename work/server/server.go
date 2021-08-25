//回调配置，接收和返回企业微信信息
package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/util"
	"github.com/yijizhichang/wechat-sdk/work/core"
	"github.com/yijizhichang/wechat-sdk/work/message"
	"github.com/yijizhichang/wechat-sdk/work/message/callback/response"
	"reflect"
	"runtime/debug"
	"strconv"
)

//Server struct
type Server struct {
	*core.Context
	openID         string
	messageHandler func(message.MixMessage) *response.Reply

	requestRawXMLMsg  []byte
	requestMsg        message.MixMessage
	responseRawXMLMsg []byte
	responseMsg       interface{}

	isSafeMode bool
	random     []byte
	nonce      string
	timestamp  int64
	wxcrypt    *util.WXBizMsgCrypt
}

//NewServer init
func NewServer(context *core.Context) *Server {
	srv := new(Server)
	srv.Context = context
	srv.isSafeMode = true //默认为true
	srv.wxcrypt = util.NewWXBizMsgCrypt(context.Token, context.EncodingAESKey, context.CorpID, util.XmlType)
	return srv
}

//处理企业微信的请求消息
func (srv *Server) Serve() error {
	// post 方式无echostr参数, 是不是不用验证？ todo
	echoStr, cryptErr := srv.Validate()
	if cryptErr != nil {
		return fmt.Errorf("请求校验失败,wxErrCode=%v, wxErrMsg=%v",cryptErr.ErrCode, cryptErr.ErrMsg)
	}

	//echostr 存在的话是GET URL 验证，post 方式，无此参数
	_, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(string(echoStr))
		return nil
	}

    //处理 wxPostData 时参数解析
	response, err := srv.handleRequest()
	if err != nil {
		return err
	}
	return srv.buildResponse(response)
}

//处理微信的请求消息,并返回给应用
func (srv *Server) ResponseServe() (str string, contentType string, echostrExist bool, err error) {
	echoStr, cryptErr := srv.Validate()
	if cryptErr != nil {
		return"","", false,  fmt.Errorf("请求校验失败,wxErrCode=%v, wxErrMsg=%v",cryptErr.ErrCode, cryptErr.ErrMsg)
	}

    //echostr 是GET URL 验证时，才会有的参数
	_, exists := srv.GetQuery("echostr")
	if exists {
		return string(echoStr),"text/plain; charset=utf-8", true, nil
	}

	//处理 wxPostData 时参数解析
	response, err := srv.handleRequest()
	if err != nil {
		return
	}

	return "","", false, srv.buildResponse(response)
}

//验证微信消息真实性
func (srv *Server) Validate() ([]byte, *util.CryptError) {
	verifyTimestamp := srv.Query("timestamp")  //时间戳(timestamp)
	verifyNonce := srv.Query("nonce")  //随机数字串(nonce)
	verifyMsgSign := srv.Query("msg_signature")  //包括消息体签名(msg_signature)
	verifyEchoStr := srv.Query("echostr") //随机加密字符串(echostr)
	return srv.wxcrypt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
}

//处理微信的请求
func (srv *Server) handleRequest() (reply *response.Reply, err error) {
	//安全模式判断
	/*srv.isSafeMode = false
	encryptType := srv.Query("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}*/

	//set openID
	//srv.openID = srv.Query("openid")

	var msg interface{}
	msg, err = srv.getMessage()
	if err != nil {
		return
	}
	mixMessage, success := msg.(message.MixMessage)
	if !success {
		err = errors.New("消息类型转换失败")
	}
	srv.requestMsg = mixMessage
	reply = srv.messageHandler(mixMessage)
	return
}

//GetOpenID return openID
func (srv *Server) GetOpenID() string {
	return srv.openID
}

//解析微信的消息
func (srv *Server) getMessage() (interface{}, error) {
	var rawXMLMsgBytes []byte
	//if srv.isSafeMode { //企微默认为加密方式传输
		reqTimestamp := srv.Query("timestamp")  //时间戳(timestamp)
		reqNonce := srv.Query("nonce")  //随机数字串(nonce)
		reqMsgSign := srv.Query("msg_signature")  //包括消息体签名(msg_signature)
		reqData, err := srv.PostData()  //post请求的密文数据
		if err != nil {
			return nil, fmt.Errorf("获取postData失败,err=%v", err)
		}
		msg, cryptErr := srv.wxcrypt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData)
		if cryptErr != nil {
			return nil, fmt.Errorf("从body中解析xml失败,wxErrCode=%v, wxErrMsg=%v",cryptErr.ErrCode, cryptErr.ErrMsg)
		}

		srv.requestRawXMLMsg = msg
		return srv.parseRequestMessage(rawXMLMsgBytes)
	//}
}

//xmlMsg to structMsg
func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg message.MixMessage, err error) {
	msg = message.MixMessage{}
	err = xml.Unmarshal(rawXMLMsgBytes, &msg)
	return
}

//设置用户自定义的回调方法
func (srv *Server) SetMessageHandler(handler func(message.MixMessage) *response.Reply) {
	srv.messageHandler = handler
}

func (srv *Server) buildResponse(reply *response.Reply) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("buildResponse error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		return nil
	}

	msgType := reply.MsgType

	switch msgType {
	case message.MsgTypeText:
	case message.MsgTypeImage:
	case message.MsgTypeVoice:
	case message.MsgTypeVideo:
	case message.MsgTypeMusic:
	case message.MsgTypeNews:
	case message.MsgTypeTransfer:
	case message.MsgTypeNothing:
		return
	default:
		err = response.ErrUnsupportReply
		return
	}

	msgData := reply.MsgData


	value := reflect.ValueOf(msgData)
	//msgData must be a ptr 企微是否需要此判断？ todo
	/*kind := value.Kind().String()
	if "ptr" != kind {
		return response.ErrUnsupportReply
	}*/

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.requestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.requestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

	params[0] = reflect.ValueOf(srv.requestMsg.AgentID)
	value.MethodByName("SetAgentID").Call(params)

	params[0] = reflect.ValueOf(msgType)
	value.MethodByName("SetMsgType").Call(params)

	params[0] = reflect.ValueOf(util.GetCurTs())
	value.MethodByName("SetCreateTime").Call(params)

	srv.responseMsg = msgData
	srv.responseRawXMLMsg, err = xml.Marshal(msgData)
	return
}

//消息发送给微信
func (srv *Server) Send() (err error) {
	replyMsg := srv.responseMsg

	if replyMsg == nil {
		srv.String("")
		return
	}

	//if srv.isSafeMode {
		//如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		reqTimestamp := strconv.FormatInt(timestamp, 10)
		encryptMsg, cryptErr := srv.wxcrypt.EncryptMsg(string(srv.responseRawXMLMsg), reqTimestamp, srv.nonce)
		if cryptErr != nil {
			return fmt.Errorf("回复企微加密失败,wxErrCode=%v, wxErrMsg=%v",cryptErr.ErrCode, cryptErr.ErrMsg)
		}

/*
		//安全模式下对消息进行加密
		var encryptedMsg []byte
		encryptedMsg, err = util.EncryptMsg(srv.random, srv.responseRawXMLMsg, srv.AppID, srv.EncodingAESKey)
		if err != nil {
			return
		}
		//如果获取不到timestamp nonce 则自己生成
		timestamp := srv.timestamp
		timestampStr := strconv.FormatInt(timestamp, 10)
		msgSignature := util.Signature(srv.Token, timestampStr, srv.nonce, string(encryptedMsg))
		replyMsg = message.ResponseEncryptedXMLMsg{
			EncryptedMsg: string(encryptedMsg),
			MsgSignature: msgSignature,
			Timestamp:    timestamp,
			Nonce:        srv.nonce,
		}
 */
	//}

	return srv.XML(string(encryptMsg))

}


//消息发送给应用，由应用发给微信
func (srv *Server) ResponseSend() (str string, contentType string, err error) {
	replyMsg := srv.responseMsg
	if replyMsg == nil {
		return "success", "text/plain; charset=utf-8", nil
	}

	timestamp := srv.timestamp
	reqTimestamp := strconv.FormatInt(timestamp, 10)
	encryptMsg, cryptErr := srv.wxcrypt.EncryptMsg(string(srv.responseRawXMLMsg), reqTimestamp, srv.nonce)
	if cryptErr != nil {
		return "", "text/plain; charset=utf-8",fmt.Errorf("回复企微加密失败,wxErrCode=%v, wxErrMsg=%v",cryptErr.ErrCode, cryptErr.ErrMsg)
	}

	s,err := srv.ResponseXML(string(encryptMsg))
	if err != nil {
		return
	}
	return s, "application/xml; charset=utf-8", nil
}


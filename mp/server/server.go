package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yijizhichang/wechat-sdk/mp/core"
	"github.com/yijizhichang/wechat-sdk/mp/message"
	"github.com/yijizhichang/wechat-sdk/mp/message/callback/response"
	"github.com/yijizhichang/wechat-sdk/util"
	"io/ioutil"
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
}

//NewServer init
func NewServer(context *core.Context) *Server {
	srv := new(Server)
	srv.Context = context
	return srv
}

//处理微信的请求消息
func (srv *Server) Serve() error {
	if !srv.Validate() {
		return fmt.Errorf("请求校验失败")
	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		srv.String(echostr)
		return nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		return err
	}

	return srv.buildResponse(response)
}

//处理微信的请求消息,并返回给应用
func (srv *Server) ResponseServe() (str string, contentType string, echostrExist bool, err error) {
	if !srv.Validate() {
		return "","", false, fmt.Errorf("请求校验失败,(返回给应用)")

	}

	echostr, exists := srv.GetQuery("echostr")
	if exists {
		return echostr,"text/plain; charset=utf-8", true, nil
	}

	response, err := srv.handleRequest()
	if err != nil {
		return
	}

	return "","", false, srv.buildResponse(response)
}

//验证微信消息真实性
func (srv *Server) Validate() bool {
	timestamp := srv.Query("timestamp")
	nonce := srv.Query("nonce")
	signature := srv.Query("signature")
	return signature == util.Signature(srv.Token, timestamp, nonce)
}

//处理微信的请求
func (srv *Server) handleRequest() (reply *response.Reply, err error) {
	//安全模式判断
	srv.isSafeMode = false
	encryptType := srv.Query("encrypt_type")
	if encryptType == "aes" {
		srv.isSafeMode = true
	}

	//set openID
	srv.openID = srv.Query("openid")

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
	var err error
	if srv.isSafeMode {
		var encryptedXMLMsg message.EncryptedXMLMsg
		if err := xml.NewDecoder(srv.Request.Body).Decode(&encryptedXMLMsg); err != nil {
			srv.WXLog.Error("从body中解析xml失败,err", err)
			return nil, fmt.Errorf("从body中解析xml失败,err=%v", err)
		}

		//验证消息签名
		timestamp := srv.Query("timestamp")
		srv.timestamp, err = strconv.ParseInt(timestamp, 10, 32)
		if err != nil {
			return nil, err
		}
		nonce := srv.Query("nonce")
		srv.nonce = nonce
		msgSignature := srv.Query("msg_signature")
		msgSignatureGen := util.Signature(srv.Token, timestamp, nonce, encryptedXMLMsg.EncryptedMsg)
		if msgSignature != msgSignatureGen {
			srv.WXLog.Error("消息不合法，验证签名失败")
			return nil, fmt.Errorf("消息不合法，验证签名失败")
		}

		//解密
		srv.random, rawXMLMsgBytes, err = util.DecryptMsg(srv.AppID, encryptedXMLMsg.EncryptedMsg, srv.EncodingAESKey)
		if err != nil {
			srv.WXLog.Error("消息解密失败, err", err)
			return nil, fmt.Errorf("消息解密失败, err=%v", err)
		}
	} else {
		rawXMLMsgBytes, err = ioutil.ReadAll(srv.Request.Body)
		if err != nil {
			srv.WXLog.Error("从body中解析xml失败, err", err)
			return nil, fmt.Errorf("从body中解析xml失败, err=%v", err)
		}
	}

	srv.requestRawXMLMsg = rawXMLMsgBytes

	srv.WXLog.Debug("解析微信消息内容", string(rawXMLMsgBytes))

	return srv.parseRequestMessage(rawXMLMsgBytes)
}

//xmlMsg to structMsg
func (srv *Server) parseRequestMessage(rawXMLMsgBytes []byte) (msg message.MixMessage, err error) {
	srv.WXLog.Debug("解析微信的消息", string(rawXMLMsgBytes))
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
			err = fmt.Errorf("panic error: %v\n%s", e, debug.Stack())
		}
	}()
	if reply == nil {
		return nil
	}

	msgType := reply.MsgType

	srv.WXLog.Debug("被动回复微信消息类型：", msgType)
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
	//msgData must be a ptr
	kind := value.Kind().String()
	if "ptr" != kind {
		return response.ErrUnsupportReply
	}

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(srv.requestMsg.FromUserName)
	value.MethodByName("SetToUserName").Call(params)

	params[0] = reflect.ValueOf(srv.requestMsg.ToUserName)
	value.MethodByName("SetFromUserName").Call(params)

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
	srv.WXLog.Debug("被动回复微信消息内容-加密前：", srv.responseMsg)

	if replyMsg == nil {
		srv.String("")
		return
	}

	if srv.isSafeMode {
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
	}

	srv.XML(replyMsg)
	return
}


//消息发送给应用，由应用发给微信
func (srv *Server) ResponseSend() (str string, contentType string, err error) {
	replyMsg := srv.responseMsg
	srv.WXLog.Debug("被动回复微信消息内容-加密前：（返回给应用）", srv.responseMsg)

	if replyMsg == nil {
		return "success", "text/plain; charset=utf-8", nil
	}

	if srv.isSafeMode {
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
	}

	s := srv.ResponseXML(replyMsg)
	return s, "application/xml; charset=utf-8", nil
}

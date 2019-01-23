package core

import (
	"encoding/xml"
	"github.com/yijizhichang/wechat-sdk/util/cache"
	flog "github.com/yijizhichang/wechat-sdk/util/log"
	"net/http"
	"sync"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

type Context struct {
	AppID            string
	AppSecret        string
	Token            string
	EncodingAESKey   string
	PayMchId         string
	PayNotifyUrl     string
	PayKey           string
	ThirdAccessToken bool

	CacheModel string      //缓存模式
	Cache      cache.Cache //缓存
	WXLog      flog.LoggerInterface
	ProxyUrl   string     //代理地址

	Writer  http.ResponseWriter
	Request *http.Request

	accessTokenLock *sync.RWMutex //读写锁 同一个AppID一个
	jsApiTicketLock *sync.RWMutex //读写锁 同一个AppID一个
}

//获取RUL参数值
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	return value
}

//获取URL参数值,并返回是否存在key
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//设置lock
func (ctx *Context) SetJsAPITicketLock(lock *sync.RWMutex) {
	ctx.jsApiTicketLock = lock
}

func (ctx *Context) SetAccessTokenLock(lock *sync.RWMutex) {
	ctx.accessTokenLock = lock
}

//获取lock
func (ctx *Context) GetJsAPITicketLock() *sync.RWMutex {
	return ctx.jsApiTicketLock
}

func (ctx *Context) GetAccessTokenLock() *sync.RWMutex {
	return ctx.accessTokenLock
}

//render from bytes
func (ctx *Context) Render(bytes []byte) {
	ctx.Writer.WriteHeader(200)
	_, err := ctx.Writer.Write(bytes)
	if err != nil {
		panic(err)
	}
}

//render from string
func (ctx *Context) String(str string) {
	writeContextType(ctx.Writer, plainContentType)
	ctx.Render([]byte(str))
}

//render to xml
func (ctx *Context) XML(obj interface{}) {
	writeContextType(ctx.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}

	ctx.WXLog.Debug("被动回复微信消息内容", string(bytes))
	ctx.Render(bytes)
}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

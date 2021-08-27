//企微配置
package core

import (
	"encoding/xml"
	"github.com/yijizhichang/wechat-sdk/util/cache"
	"io/ioutil"
	"net/http"
	"sync"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

type Context struct {
	CorpID        string  // 企业ID
	CorpSecret    string  // 应用的凭证密钥; 每个应用有独立的secret，获取到的access_token只能本应用使用，所以每个应用的access_token应该分开来获取
	Token            string
	EncodingAESKey   string
	RasPrivateKey string  // 消息加密私钥
	ThirdAccessToken bool  //是用其他应用生成的access_token
	Debug  bool  //为true时会打印一些调试信息

	Cache      cache.Cache //缓存
	ProxyUrl   string      //代理地址


	Writer  http.ResponseWriter
	Request *http.Request

	accessTokenLock *sync.RWMutex //读写锁 同一个AppID一个
}

//获取RUL参数值
func (ctx *Context) Query(key string) string {
	value, _ := ctx.GetQuery(key)
	//unValue,_:=url.PathUnescape(value)
	return value
}

//发取URL参数值,并返回是否存在key
func (ctx *Context) GetQuery(key string) (string, bool) {
	req := ctx.Request
	if values, ok := req.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}

//获取post参数
func (ctx *Context) PostData()([]byte, error){
	req := ctx.Request
	return ioutil.ReadAll(req.Body)
}

//设置lock
func (ctx *Context) SetAccessTokenLock(lock *sync.RWMutex) {
	ctx.accessTokenLock = lock
}

//获取lock
func (ctx *Context) GetAccessTokenLock() *sync.RWMutex {
	return ctx.accessTokenLock
}

//render from bytes
func (ctx *Context) Render(bytes []byte)(err error){
	ctx.Writer.WriteHeader(200)
	_, err = ctx.Writer.Write(bytes)
	return
}

//render from string
func (ctx *Context) String(str string) {
	writeContextType(ctx.Writer, plainContentType)
	ctx.Render([]byte(str))
}

//render to xml
func (ctx *Context) XML(obj interface{}) (err error) {
	writeContextType(ctx.Writer, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		return
	}
	err = ctx.Render(bytes)
	return
}

func (ctx *Context) ResponseXML(obj interface{}) (s string, err error) {
	bytes, err := xml.Marshal(obj)
	s = string(bytes)
	return
}

func writeContextType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}


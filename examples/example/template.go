package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/mp/message/template"
)

//发送模板消息
func SendTemplateMsg() {
	//NewWechat
	tpl := wxconf.WechatClient.GetTemplate()

	//模板参数
	first := &template.DataItem{
		"亲爱的xxx，您本月的收益情况为",
		"#FF4040",
	}
	remark := &template.DataItem{
		"感谢您的使用,祝您生活愉快~",
		"#FF00FF",
	}
	keyword1 := &template.DataItem{
		"500大洋",
		"#CAFF70",
	}
	keyword2 := &template.DataItem{
		"手机",
		"#9932CC",
	}

	msgTest2 := new(template.Message)
	msgTest2.ToUser = "abcd1234abcd1234abcd1234"
	msgTest2.TemplateID = "MYl2aWAlESuuCVZ-KyOzxWFlUZ0Gp95fO2mXXXXXXX"
	msgTest2.Data = make(map[string]*template.DataItem)
	msgTest2.Data["first"] = first
	msgTest2.Data["keyword1"] = keyword1
	msgTest2.Data["keyword2"] = keyword2
	msgTest2.Data["remark"] = remark

	re, _ := tpl.Send(msgTest2)
	fmt.Println("返回结果：", re)
}

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

func MassManage() {
	ms := wxconf.WechatClient.GetMass()

	//根据标签进行群发【订阅号与服务号认证后均可用】

	//文本
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithTextOption("hello tag1"),
	)*/

	//图文
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithMpnewsOption("asdfafasdfsdfasf",0),
	)*/

	//语音/音频
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithVoiceOption("asdfafasdfsdfasf"),
	)*/

	//图片
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithImageOption("asdfafasdfsdfasf"),
	)*/

	//视频
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithMpvideoOption("asdfafasdfsdfasf"),
	)*/

	//视频
	/*	res, err := ms.MassSendall(
		mass.WithFilterOption(true, 1),
		mass.WithWxcardOption("asdfafasdfsdfasf"),
	)*/

	//根据OpenID列表群发【订阅号不可用，服务号认证后可用】
	/*	res, err := ms.MassSend(
		mass.WithTouserOption([]string{"abcd1234abcd1234abcd1234", "abcd1234abcd1234abcd1235"}),
		mass.WithTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)*/

	//预览接口【订阅号与服务号认证后均可用】
	//开发者可通过该接口发送消息给指定用户，在手机端查看消息的样式和排版。为了满足第三方平台开发者的需求，在保留对openID预览能力的同时，增加了对指定微信号发送预览的能力，但该能力每日调用次数有限制（100次），请勿滥用。

	//指定openid
	/*res, err := ms.MassPreview(
		mass.WithPreviewTouserOption("abcd1234abcd1234abcd1234"),
		mass.WithPreviewTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)*/

	//指定微信号
	/*	res, err := ms.MassPreview(
		mass.WithPreviewTowxnameOption("wxname"),
		mass.WithPreviewTextOption("hello openid134"),  //根据发的素材不同,发对应内容
	)*/

	//删除群发
	//err := ms.MassDel(30124,2)

	//查询状态
	//res,err := ms.MassGet("301234")

	//控制群发速度
	res, err := ms.MassSpeedSet(2) //群发速度的级别0-4，是一个0到4的整数，数字越大表示群发速度越慢

	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}
	fmt.Println(res)
}

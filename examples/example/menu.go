package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/mp/menu"
	"encoding/json"
)


//创建菜单
func MenuCreate()  {

	mu := wxconf.WechatClient.GetMenu()
	//二级菜单列表
	subMenuList1 := menu.SetButton(
		menu.WithClickButton("赞我们一下","V1001_GOOD"),
		menu.WithViewButton("搜一下","http://www.soso.com/"),
		menu.WithLocationSelectButton("上报位置","wz2039_fdei"),
		menu.WithMiniprogramButton("跳转小程序","http://mp.weixin.qq.com","wx286b93c14bbf93aa","pages/lunar/index"),
	)
	subMenuList2 := menu.SetButton(
		menu.WithScanCodeWaitMsgButton("扫码带提示","rselfmenu_0_0"),
		menu.WithScanCodePushButton("扫码推事件","rselfmenu_0_1"),
	)

	//一级菜单列表
	parentMenu1 := menu.SetButton(
		menu.WithSubButton("菜单1",subMenuList1),
		menu.WithSubButton("菜单2",subMenuList2),
		menu.WithPicPhotoOrAlbumButton("菜单3","22222"),
	)


	//创建菜单
	res, err := mu.SetMenu(parentMenu1...)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}

	fmt.Println("设置菜单返回", res)
}

//获取菜单
func MenuGet(){
	mu := wxconf.WechatClient.GetMenu()
	res, err := mu.GetMenu()
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}

	jsonStr,err :=json.Marshal(res)

	fmt.Println("查询菜单返回", res,string(jsonStr))

}

//删除菜单 
func MenuDel()  {
	mu := wxconf.WechatClient.GetMenu()
	err := mu.DeleteMenu()
	fmt.Println("查询菜单返回", err)
}

//添加个性化菜单
func MenuAddConditional(){

	mu := wxconf.WechatClient.GetMenu()

	//二级菜单列表
	subMenuList1 := menu.SetButton(
		menu.WithClickButton("我是个性菜单","V1001_GOODdfd"),
		menu.WithLocationSelectButton("报告位置","wz2039_fdeidffd"),
	)
	//一级菜单列表
	parentMenu1 := menu.SetButton(
		menu.WithClickButton("今日歌曲","vmufddf"),
		menu.WithSubButton("菜单1",subMenuList1),
	)

	//匹配规则
	matchRule := menu.MatchRule{
		GroupID:"2",
	}

	//创建个性菜单
	res,err := mu.AddConditional(parentMenu1, &matchRule)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}

	fmt.Println("创建个性菜单返回", res, err)
}

//删除个性化菜单
func MenuDelConditional() {
	mu := wxconf.WechatClient.GetMenu()
	err := mu.DeleteConditional(420025855)
	if err != nil {
		fmt.Printf("err=%v", err)
		return
	}
	fmt.Println("删除个性化菜单返回", err)
}

//个性化菜单测试
func MenuTryMatch(){
	mu := wxconf.WechatClient.GetMenu()
	res, err := mu.MenuTryMatch("pipi520733")
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}

	jsonStr,err :=json.Marshal(res)

	fmt.Println("查询菜单返回", res,string(jsonStr))

}

//获取自定义菜单配置
func GetSelfMenuInfo() {
	mu := wxconf.WechatClient.GetMenu()
	res, err := mu.GetSelfMenuInfo()
	if err != nil {
		fmt.Printf("err=%v", err)
		return
	}
	fmt.Println(res)
}
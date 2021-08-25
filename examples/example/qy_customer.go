/**
 * @Time: 2021/8/9 11:23 上午
 * @Author: soupzhb@gmail.com
 * @File: qy_customer.go
 * @Software: GoLand
 */

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

func QyGetCustomerList(token string)  {
	cus := wxconf.QyWechatClint.GetCustomer()
	res, err := cus.GetQyCustomerList(token, "xinyu888")
	if err != nil {
		fmt.Printf("Err: %+v",err)
	}
	fmt.Printf("Res: %+v",res)
}

func QyGetCustomerView(token string)  {
	cus := wxconf.QyWechatClint.GetCustomer()
	res, err := cus.GetQyCustomerView(token, "wmt0zCEAAAfyx96crCFRHwLxqtR3_5yA","") //wmt0zCEAAAfyx96crCFRHwLxqtR3_3yA
	if err != nil {
		fmt.Printf("Err: %+v",err)
	}
	fmt.Printf("Res: %+v",res)
}



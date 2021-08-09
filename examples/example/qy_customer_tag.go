/**
 * @Time: 2021/8/9 11:25 上午
 * @Author: soupzhb@gmail.com
 * @File: qy_customer_tag.go
 * @Software: GoLand
 */

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
	"github.com/yijizhichang/wechat-sdk/work/customer"
)

func QyGetCustomerTagList(token string){

	customerTag := wxconf.QyWechatClint.GetCustomerTag()
	res, err := customerTag.GetCustomerTagList(token, customer.CusTagReq{TagId:[]string{}, GroupId: []string{}})
	if err != nil {
		fmt.Printf("QyGetCustomerTagList GetCustomerTagList Err: %+v",err)
	}

	fmt.Printf("QyGetCustomerTagList GetCustomerTagList Res: %+v",res)
}

func QyCreateCustomerTag(token string){
	customerTag := wxconf.QyWechatClint.GetCustomerTag()
	req := new(customer.CreateCusTagReq)
	req.GroupName = "技术Mark"
	addTag1 := customer.AddTag{
		"m1",
		2,
	}
	addTag2 := customer.AddTag{
		"m2",
		1,
	}
	req.Tag = append(req.Tag, addTag1, addTag2)
	res, err := customerTag.CreateCustomerTag(token, *req)
	if err != nil {
		fmt.Printf("QyCreateCustomerTag CreateCustomerTag Err: %+v",err)
	}

	fmt.Printf("QyCreateCustomerTag CreateCustomerTag Res: %+v",res)
}

func QyUpdateCustomerTag(token string){
	customerTag := wxconf.QyWechatClint.GetCustomerTag()
	req := new(customer.UpdateCusTagReq)
	req.Id = "ett0zCEAAAqlT6DRVQUh6CJnUhD2mZTA"
	req.Name = "m3"
	req.Order = 5


	res, err := customerTag.UpdateCustomerTag(token, *req)
	if err != nil {
		fmt.Printf("QyUpdateCustomerTag UpdateCustomerTag Err: %+v",err)
	}

	fmt.Printf("QyUpdateCustomerTag UpdateCustomerTag Res: %+v",res)
}

func QyDelCustomerTag(token string){
	customerTag := wxconf.QyWechatClint.GetCustomerTag()
	req := new(customer.DelCusTagReq)
	req.TagId = []string{"ett0zCEAAAqlT6DRVQUh6CJnUhD2mZTA"}


	res, err := customerTag.DelCustomerTag(token, *req)
	if err != nil {
		fmt.Printf("QyDelCustomerTag DelCustomerTag Err: %+v",err)
	}

	fmt.Printf("QyDelCustomerTag DelCustomerTag Res: %+v",res)
}

func QyMarkTag(token string){
	customerTag := wxconf.QyWechatClint.GetCustomerTag()
	req := new(customer.MarkTagReq)
	req.Userid = "xinyu8521"
	req.ExternalUserid = "wmt0zCEAAAfyx96crCFRHwLxqtR3_3yA"
	req.AddTag = []string{"ett0zCEAAAw3i1L_kh9gQIIc6A-XZ6tQ"}


	res, err := customerTag.MarkCustomerTag(token, *req)
	if err != nil {
		fmt.Printf("QyMarkTag MarkCustomerTag Err: %+v",err)
	}

	fmt.Printf("QyMarkTag MarkCustomerTag Res: %+v",res)
}



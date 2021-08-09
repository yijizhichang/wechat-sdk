/**
 * @Time: 2021/8/9 2:08 下午
 * @Author: soupzhb@gmail.com
 * @File: qy_company_department.go
 * @Software: GoLand
 */

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

func QyGetDepartment(token string){

	department := wxconf.QyWechatClint.GetDepartment()
	res, err := department.GetDepartmentList(token,0)
	if err != nil {
		fmt.Printf("QyGetDepartment GetDepartmentList Err: %+v",err)
	}

	fmt.Printf("QyGetDepartment GetDepartmentList Res: %+v",res)
}


func QyGetDepartmentSimpleUserList(token string){

	departmentUser := wxconf.QyWechatClint.GetDepartmentUser()
	res, err := departmentUser.GetDepartmentSimpleUserList(token,7,0)
	if err != nil {
		fmt.Printf("QyGetDepartmentSimpleUserList GetDepartmentSimpleUserList Err: %+v",err)
	}

	fmt.Printf("QyGetDepartmentSimpleUserList GetDepartmentSimpleUserList Res: %+v",res)
}

func QyGetDepartmentUserList(token string){

	departmentUser := wxconf.QyWechatClint.GetDepartmentUser()
	res, err := departmentUser.GetDepartmentUserList(token,7,0)
	if err != nil {
		fmt.Printf("QyGetDepartmentUserList GetDepartmentUserList Err: %+v",err)
	}

	fmt.Printf("QyGetDepartmentUserList GetDepartmentUserList Res: %+v",res)
}




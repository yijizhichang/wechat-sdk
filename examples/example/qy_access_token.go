/**
 * @Time: 2021/8/5 6:30 下午
 * @Author: soupzhb@gmail.com
 * @File: qy_access_token.go
 * @Software: GoLand
 */

package example

import (
	"fmt"
	"github.com/yijizhichang/wechat-sdk/examples/wxconf"
)

func QyAccessToken() {
	token, err := wxconf.QyWechatClient.GetQyAccessToken("VyMmm3hy5n7Q5t23tozkraUtuOt_sdfsadfsfsDCXGEQ")
	token2, err2 := wxconf.QyWechatClient.GetQyAccessToken("2rLl96BjFobh7Y_lD6sdfsfdsfsfsfsdjWY3QwIu4")
	fmt.Printf("企业微token:", token, err)
	fmt.Printf("企业微token2:", token2, err2)

}

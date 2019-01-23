//时间日期相关方法
package util

import "time"

//获取当前时间戳
func GetCurTs() int64 {
	return time.Now().Unix()
}

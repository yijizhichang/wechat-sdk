package util

import (
	flog "github.com/yijizhichang/wechat-sdk/util/log"
)

func StrDefault(str *string, defaultStr string) {
	if *str == "" {
		*str = defaultStr
	}
}

func IntDefault(str *int, defaultInt int) {
	if *str == 0 {
		*str = defaultInt
	}
}

func Int8Default(str *int8, defaultInt int8) {
	if *str == 0 {
		*str = defaultInt
	}
}

func Int64Default(str *int64, defaultInt64 int64) {
	if *str == 0 {
		*str = defaultInt64
	}
}

func BoolDefault(str *bool, defaultBool bool) {
	if *str == false {
		*str = defaultBool
	}
}

func LevelDefault(str *flog.LEVEL, defaultLevel flog.LEVEL)  {
	if *str == 0 {
		*str = 4  //默认Error级别
	}
}

func ArrayDefault(str *[]string, defaultArray []string) {
	if len(*str)==0 {
		*str = defaultArray
	}
}

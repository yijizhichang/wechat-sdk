package util

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
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

func ArrayDefault(str *[]string, defaultArray []string) {
	if len(*str) == 0 {
		*str = defaultArray
	}
}

func URLKVString(m map[string]string) (s string) {
	kArr := make([]string, len(m))
	for k, _ := range m {
		kArr = append(kArr, k)
	}
	sort.Strings(kArr)

	for i := 0; i < len(kArr); i++ {
		s = kArr[i] + `=` + m[kArr[i]] + `&`
	}
	s = strings.TrimRight(s, `&`)
	return
}

func RandomString(lenth int) string {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	b := bytes.Buffer{}
	b.WriteString(s)
	str := b.String()
	strLen := len(str)
	if strLen <= 0 {
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	b = bytes.Buffer{}
	for i := 0; i < lenth; i++ {
		b.WriteByte(str[rand.Intn(strLen)])
	}
	return b.String()
}

func SHA1(str string) string {
	sha := sha1.New()
	io.WriteString(sha, str)
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func Mkdir(dir string) (e error) {
	_, er := os.Stat(dir)
	b := er == nil || os.IsExist(er)
	if !b {
		if err := os.MkdirAll(dir, 0750); err != nil {
			if os.IsPermission(err) {
				e = err
			}
		}
	}
	return
}
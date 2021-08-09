/**
 * @Time: 2021/8/5 5:01 下午
 * @Author: soupzhb@gmail.com
 * @File: file2.go
 * @Software: GoLand
 */

package cache

import "time"

type FileClient interface {
	Get(k string) (v string, err error)
	Set(k string, v interface{}, expir ...time.Duration) (err error)
}

type fileClient struct {
	File *File
}

func (fc *fileClient) Get(k string) (v string, err error) {
	v,err = fc.File.Get(k)
	return
}

func (fc *fileClient) Set(k string, v interface{}, expir ...time.Duration) (err error) {
	err = fc.File.Set(k,v,expir[0])
	return
}

func NewFileClient(file *File) FileClient {
	return &fileClient{
		File: file,
	}
}




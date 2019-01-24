package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	FILENIL = "file:nil"
)

var (
	timeLocation, _ = time.LoadLocation("Asia/Chongqing") //当地时间
)

type File struct {
	fileClient   *os.File
	data         *fileString
	FileFullPath string
	mux          *sync.RWMutex
	interval     time.Duration // 自动检测key过去时间间隔
}

// 文件中json串结构
type fileString struct {
	Str map[string]baseAttr `json:"str"`
}

type baseAttr struct {
	Value      interface{} `json:"value"`
	Expiration int64       `json:"expiration"`
}

func NewFile(fileFullPath string) (*File, error) {
	var err error
	file := new(File)
	file.interval = time.Second * 2
	file.FileFullPath = absolutePath(fileFullPath)
	file.fileClient, err = os.OpenFile(file.FileFullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0750)
	file.mux = new(sync.RWMutex)
	file.data = new(fileString)
	file.loadFile()
	if file.data.Str == nil {
		file.data.Str = make(map[string]baseAttr)
	}
	go file.runCheckExpired()
	return file, err
}

func (f *File) Get(key string) (interface{}, error) {
	is, err := f.IsExist(key)
	if err != nil {
		return nil, errors.New(FILENIL)
	}
	if is {
		v, ok := f.data.Str[key]
		if ok {
			return v.Value, nil
		}
	}
	return nil, errors.New(FILENIL)
}

func (f *File) Set(key string, val interface{}, t time.Duration) error {
	val = changeType(val)
	it := nowTimeAdd2Nano(t)
	if t == 0 {
		it = 0
	}

	f.data.Str[key] = baseAttr{
		Expiration: it,
		Value:      val,
	}
	return f.writeFile()
}

func (f *File) IsExist(key string) (bool, error) {
	v, ok := f.data.Str[key]
	if ok {
		if f.isExpired(v) {
			f.Delete(key)
		} else {
			return true, nil
		}
	}
	// 去文件中查一次
	fs, err := f.getFromFile()
	if err != nil {
		return false, err
	}
	// 更新map
	v, ok = f.data.Str[key]
	if ok {
		if f.isExpired(v) {
			f.Delete(key)
		} else {
			f.data.Str[key] = fs.Str[key]
			return true, nil
		}
	}
	return false, errors.New(FILENIL)
}

func (f *File) Delete(key string) (bool, error) {
	_, ok := f.data.Str[key]
	if ok {
		go delete(f.data.Str, key)
	}
	err := f.writeFile()
	if err != nil {
		return false, err
	}
	return true, nil
}

// 将k/v加载到内存
func (f *File) loadFile() error {
	fileStr, err := f.getFromFile()
	if err != nil {
		ba := make(map[string]baseAttr)
		f.data = &fileString{Str: ba}
	} else {
		f.data = &fileStr
	}
	return err
}

func (f *File) writeFile() error {
	b, err := json.Marshal(f.data)
	if err != nil {
		return err
	}
	f.mux.Lock()
	defer f.mux.Unlock()
	err = os.Truncate(f.FileFullPath, 0)
	if err != nil {
		return err
	}
	_, err = f.fileClient.Write(b)
	return err
}

func (f *File) getFromFile() (fileStr fileString, err error) {
	tmp := new(fileString)
	if fileStr.Str == nil {
		f.mux.Lock()
		fileStr.Str = make(map[string]baseAttr)
		f.mux.Unlock()
	}
	b, err := f.readFile()
	if len(b) == 0 {
		return
	}
	err = json.Unmarshal(b, &tmp)
	if err != nil {
		return
	}

	f.mux.Lock()
	fileStr = *tmp
	f.mux.Unlock()
	return
}

func (f *File) readFile() (b []byte, err error) {
	f.mux.Lock()
	defer f.mux.Unlock()

	b, err = ioutil.ReadAll(f.fileClient)
	if err != nil {
		return
	}
	return
}

func (f *File) isExpired(bs baseAttr) bool {
	if bs.Expiration == 0 {
		return false
	}
	if getNowTimeNano() > bs.Expiration {
		return true
	} else {
		return false
	}
}

// 定时检测key是否失效，失效及删除key（文件、内存）;同步没有失效key;
// 文件->内存
func (f *File) runCheckExpired() {
	ticker := time.NewTicker(f.interval)
	for {
		select {
		case <-ticker.C:
			f.checkExpired()
		}
	}
}

func (f *File) checkExpired() {
	var ok bool
	// 去文件中查一次
	fs, err := f.getFromFile()
	if err != nil {
		return
	}
	for key, val := range fs.Str {
		f.mux.Lock()
		if f.isExpired(val) {
			_, ok = f.data.Str[key]
			if ok {
				delete(f.data.Str, key)
			}
		} else {
			f.data.Str[key] = val
		}
		f.mux.Unlock()
	}
}

func getNowTimeNano() int64 {
	return time.Now().In(timeLocation).UnixNano()
}

func nowTimeAdd2Nano(t time.Duration) int64 {
	return time.Now().In(timeLocation).Add(t).UnixNano()
}

//获取绝对路径
func absolutePath(dir string) (fp string) {
	fp, _ = filepath.Abs(dir)
	return
}

func changeType(val interface{}) interface{} {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		val = fmt.Sprintf("%d", val)
	case string:
	case bool:
		val = fmt.Sprintf("%v", val)
	case float64:
		val = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case float32:
		val = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 64)
	}
	return val
}

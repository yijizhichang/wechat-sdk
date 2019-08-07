package cache

import (
	"testing"
	"time"
)

func TestNewFile(t *testing.T) {

	f, err := NewFile("./token.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("data:%+v-%+v", f.data, time.Second)

	err = f.Set("a", 1, time.Second*20)
	if err != nil {
		t.Fatal(err)
	}
	v, err := f.Get("a")
	t.Logf("%#v", v)
	if err != nil {
		t.Fatal(err)
	}
	//go func() {
	//	for {
	//		v, err = f.Get("a")
	//		if err != nil {
	//			if err.Error() != FILENIL {
	//				t.Fatal(err)
	//			}
	//		}
	//		t.Log(v)
	//		time.Sleep(time.Second)
	//	}
	//}()
	//time.Sleep(time.Second * 5)
	//v, err = f.Get("a")
	//if err != nil {
	//	if err.Error() != FILENIL {
	//		t.Fatal(err)
	//	}
	//}
}

func TestTime(t *testing.T) {
	t.Logf("%#v", changeType(1))

	var s map[int]int

	t.Log(s==nil)
}

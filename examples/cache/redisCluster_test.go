package cache

import (
	"testing"
	"time"
)

func TestNewRedisCluster(t *testing.T) {
	r, err := NewRedisCluster(&RedisCluster{
		Addrs:    []string{"10.141.6.87:6379", "10.141.6.88:6380", "10.141.6.89:6381"},
		Password: "wR68543q71",
	})
	if err != nil {
		t.Fatal("redis初始化出错", err)
	}

	err = r.Set("aaa", 20.10, time.Second*10)
	t.Logf("%#v\n", err)

	v, err := r.Get("aaa")

	t.Logf("%#v---%#v", v, err)
}

package cache

import (
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	r, _ := NewRedis(&Redis{
		Addr:     "127.0.0.1:6381",
		Password: "your redis pwd",
	})
	r.Set("hello", "a", time.Second)
	a, err := r.Get("hello")
	t.Log(a, err)
}

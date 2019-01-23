package cache

import "time"

type Cache interface {
	Get(string) (interface{}, error)
	Set(string, interface{}, time.Duration) error
	IsExist(string) (bool, error)
	Delete(string) (bool, error)
}

package cache

import "time"

type Cache interface {
	Get(k string) (v string, err error)
	Set(k string, v interface{}, expir ...time.Duration) (err error)
	//Get(string) (interface{}, error)
	//Set(string, interface{}, time.Duration) error
	//IsExist(string) (bool, error)
	//Delete(string) (bool, error)
}

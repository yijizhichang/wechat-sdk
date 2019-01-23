package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	redisClient *redis.Client
	Addr        string
	Password    string
}

func NewRedis(redis2 *Redis) (*Redis, error) {
	var err error
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redis2.Addr,
		Password: redis2.Password,
	})
	redis2.redisClient = redisClient
	_, err = redisClient.Ping().Result()
	return redis2, err
}

func (r *Redis) Get(key string) (interface{}, error) {
	return r.redisClient.Get(key).Result()
}

func (r *Redis) Set(key string, val interface{}, t time.Duration) error {
	return r.redisClient.Set(key, val, t).Err()
}

func (r *Redis) IsExist(key string) (bool, error) {
	i, err := r.redisClient.Exists(key).Result()
	if err != nil || i == 0 {
		return false, err
	}
	return true, err
}

func (r *Redis) Delete(key string) (bool, error) {
	i, err := r.redisClient.Del(key).Result()
	if err != nil || i == 0 {
		return false, err
	}
	return true, err
}

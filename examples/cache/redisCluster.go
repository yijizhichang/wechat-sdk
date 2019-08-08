package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type RedisCluster struct {
	redisClusterClient *redis.ClusterClient
	Addrs              []string
	Password           string
}

func NewRedisCluster(redis2 *RedisCluster) (*RedisCluster, error) {
	var err error
	clusterRedisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redis2.Addrs,
		Password: redis2.Password,
	})
	redis2.redisClusterClient = clusterRedisClient
	_, err = clusterRedisClient.Ping().Result()
	return redis2, err
}

func (r *RedisCluster) Get(key string) (string, error) {
	return r.redisClusterClient.Get(key).Result()
}

func (r *RedisCluster) Set(key string, val interface{}, t ...time.Duration) error {
	var exp time.Duration
	if len(t) == 1 {
		exp = t[0]
	}
	return r.redisClusterClient.Set(key, val, exp).Err()
}

func (r *RedisCluster) IsExist(key string) (bool, error) {
	i, err := r.redisClusterClient.Exists(key).Result()
	if err != nil || i == 0 {
		return false, err
	}
	return true, err
}

func (r *RedisCluster) Delete(key string) (bool, error) {
	i, err := r.redisClusterClient.Del(key).Result()
	if err != nil || i == 0 {
		return false, err
	}
	return true, err
}

package cache

import (
	"app-api/boot"
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

const CachePrefix = "cache:"

type RedisKeys struct {
	Key string
	TTL time.Duration
}

// 获取广告缓存key TTL
func GetAdKeys(space string) RedisKeys {
	return RedisKeys{Key: CachePrefix + "ad:space:" + space, TTL: time.Minute * 5}
}

// 获取缓存
func Get(key string) (value string, err error) {
	value, err = boot.Redis.Get(context.Background(), key).Result()
	if err == redis.Nil {
		err = errors.New(key + " does not exist")
		return "", err
	} else if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

// 设置缓存
func Set(key string, value string, expire time.Duration) (err error) {
	err = boot.Redis.Set(context.Background(), key, value, expire).Err()
	return
}

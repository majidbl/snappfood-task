package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"task/config"
)

var RedisCLi *redis.Client
var once sync.Once

func SetUpRedis(config config.Config) error {
	var redisErr error
	once.Do(func() {
		RedisCLi = redis.NewClient(&redis.Options{
			Addr:     config.DBSource,
			Username: "",
			Password: "",
			DB:       0,
		})

		redisErr = RedisCLi.Ping(context.Background()).Err()
	})

	return redisErr
}

func NewRedisCli() *redis.Client {
	return RedisCLi
}

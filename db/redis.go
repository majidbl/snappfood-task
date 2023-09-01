package db

import (
	"context"

	"github.com/go-redis/redis/v8"

	"task/util"
)

var RedisCLi *redis.Client

func SetUpRedis(config util.Config) error {
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

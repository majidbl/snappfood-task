package queue

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

type RedisQueue[T any] struct {
	Data T
	cli  *redis.Client
	key  string
}

func NewRedisQueue[T any](cli *redis.Client, key string) Queue[T] {
	if key == "" {
		log.Warn("redis queue requires key name")
	}

	return &RedisQueue[T]{
		key: key,
		cli: cli,
	}
}

func (r RedisQueue[T]) Enqueue(ctx context.Context, data T) error {
	// Convert the order to JSON and enqueue it in Redis.
	b, err := r.ToString(data)
	if err != nil {
		return err
	}

	err = r.cli.RPush(ctx, "orders", b).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r RedisQueue[T]) Dequeue(ctx context.Context) (T, error) {
	// Dequeue an order from Redis and convert it back to Order struct.
	var data T
	dataStr, err := r.cli.LPop(ctx, r.key).Result()
	if err != nil {
		return data, err
	}

	data, err = r.ToData(dataStr)
	if err != nil {
		log.Printf("Failed to parse order: %v", err)
		return data, err
	}

	return data, nil
}

func (r RedisQueue[T]) ToString(data T) (string, error) {
	// Convert Order struct to a string in JSON format.
	// You can use any serialization method here.
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (r RedisQueue[T]) ToData(str string) (T, error) {
	// Parse the JSON string back to an Order struct.
	var data T
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

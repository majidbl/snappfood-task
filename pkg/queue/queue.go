package queue

import (
	"context"
	"task/db"
	"task/util"

	"github.com/labstack/gommon/log"
)

const (
	InMemory = "memory"
	Redis    = "redis"
)

type Queue[T any] interface {
	Enqueue(ctx context.Context, data T) error
	Dequeue(ctx context.Context) (T, error)
}

const (
	EmptyQueue = "queue is empty"
)

func New[T any](config util.Config) Queue[T] {
	switch config.QueueType {
	case InMemory:
		return NewMemoryQueue[T]()
	case Redis:
		return NewRedisQueue[T](db.NewRedisCli(), config.OrderQueueKey)
	}

	log.Warn("Queue type is default [memory]")
	return NewMemoryQueue[T]()
}

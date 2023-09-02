package queue

import (
	"context"
	"sync"
	"task/models"

	"github.com/labstack/gommon/log"

	"task/storage/redis"
	"task/util"
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
		return NewRedisQueue[T](redis.NewRedisCli(), config.OrderQueueKey)
	}

	log.Warn("Queue type is default [memory]")
	return NewMemoryQueue[T]()
}

var OrderQueueManger Queue[models.Order]
var once sync.Once

func SetUpQueueManager(config util.Config) {
	once.Do(func() {
		OrderQueueManger = New[models.Order](config)
	})
}

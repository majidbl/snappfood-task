package queue

import (
	"context"
	"fmt"
	"sync"
)

type MemoryQueue[T any] struct {
	data  []T
	mutex sync.Mutex
}

func NewMemoryQueue[T any]() Queue[T] {
	return &MemoryQueue[T]{}
}

func (oq *MemoryQueue[T]) Enqueue(ctx context.Context, data T) error {
	oq.mutex.Lock()
	defer oq.mutex.Unlock()
	oq.data = append(oq.data, data)
	return nil
}

func (oq *MemoryQueue[T]) Dequeue(ctx context.Context) (T, error) {
	var t T

	oq.mutex.Lock()
	defer oq.mutex.Unlock()

	if len(oq.data) == 0 {
		return t, fmt.Errorf(EmptyQueue)
	}

	order := oq.data[0]
	oq.data = oq.data[1:]
	return order, nil
}

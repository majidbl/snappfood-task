package queue

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestRedisQueueEnqueueDequeue(t *testing.T) {
	// Create a Redis client.
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default DB
	})

	// Initialize the Redis queue.
	key := "test_queue"
	queue := NewRedisQueue[int](client, key)
	ctx := context.TODO()

	// Test Enqueue
	err := queue.Enqueue(ctx, 1)
	if err != nil {
		t.Errorf("Enqueue returned an error: %v", err)
	}

	// Test Dequeue
	data, err := queue.Dequeue(ctx)
	if err != nil {
		t.Errorf("Dequeue returned an error: %v", err)
	}
	if data != 1 {
		t.Errorf("Dequeue didn't return the expected data. Got %v, expected 1", data)
	}
}

func TestRedisQueueEmptyQueueError(t *testing.T) {
	// Create a Redis client.
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
		DB:   0,                // Default DB
	})

	// Initialize the Redis queue.
	key := "test_queue"
	queue := NewRedisQueue[int](client, key)
	ctx := context.TODO()

	// Test Dequeue from an empty queue
	_, err := queue.Dequeue(ctx)
	expectedError := "redis: nil"
	if err == nil || err.Error() != expectedError {
		t.Errorf("Dequeue didn't return the expected error for an empty queue. Got: %v, Expected: %v", err, expectedError)
	}
}

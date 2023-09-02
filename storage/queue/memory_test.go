package queue

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestMemoryQueueEnqueueDequeue(t *testing.T) {
	queue := NewMemoryQueue[int]()
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

	// Test Dequeue from an empty queue
	_, err = queue.Dequeue(ctx)
	expectedError := fmt.Sprintf(EmptyQueue)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Dequeue didn't return the expected error for an empty queue. Got: %v, Expected: %v", err, expectedError)
	}
}

func TestMemoryQueueConcurrentEnqueueDequeue(t *testing.T) {
	queue := NewMemoryQueue[int]()
	ctx := context.TODO()
	var wg sync.WaitGroup

	// Enqueue data concurrently
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := queue.Enqueue(ctx, i)
			if err != nil {
				t.Errorf("Enqueue returned an error: %v", err)
			}
		}(i)
	}

	// Dequeue data concurrently
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data, err := queue.Dequeue(ctx)
			if err != nil {
				t.Errorf("Dequeue returned an error: %v", err)
			}
			if data < 1 || data > 100 {
				t.Errorf("Dequeue returned unexpected data: %v", data)
			}
		}()
	}

	wg.Wait()
}

func TestMemoryQueueEmptyQueueError(t *testing.T) {
	queue := NewMemoryQueue[int]()
	ctx := context.TODO()

	// Test Dequeue from an empty queue
	_, err := queue.Dequeue(ctx)
	expectedError := fmt.Sprintf(EmptyQueue)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Dequeue didn't return the expected error for an empty queue. Got: %v, Expected: %v", err, expectedError)
	}
}

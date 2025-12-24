package structures

import (
	"testing"
)

func TestQueue_NewQueue(t *testing.T) {
	queue := NewQueue()
	if queue == nil {
		t.Error("NewQueue should not return nil")
	}
}

func TestQueue_EnqueueDequeue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	
	if queue.Peek() != "A" {
		t.Errorf("Expected 'A' at front, got %s", queue.Peek())
	}
	
	dequeued := queue.Dequeue()
	if dequeued != "A" {
		t.Errorf("Expected 'A' dequeued, got %s", dequeued)
	}
	
	if queue.Dequeue() != "B" {
		t.Errorf("Expected 'B' dequeued")
	}
	
	if queue.Dequeue() != "[QUEUE_EMPTY]" {
		t.Errorf("Expected [QUEUE_EMPTY], got %s", queue.Dequeue())
	}
}

func TestQueue_PeekEmpty(t *testing.T) {
	queue := NewQueue()
	if queue.Peek() != "[QUEUE_EMPTY]" {
		t.Errorf("Expected [QUEUE_EMPTY] for empty queue, got %s", queue.Peek())
	}
}

func TestQueue_Print(t *testing.T) {
	queue := NewQueue()
	queue.Print() // Пустая очередь
	
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Print() // Непустая очередь
}

func TestQueue_IsEmpty(t *testing.T) {
	queue := NewQueue()
	if !queue.IsEmpty() {
		t.Error("New queue should be empty")
	}
	
	queue.Enqueue("A")
	if queue.IsEmpty() {
		t.Error("Queue with element should not be empty")
	}
	
	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeue")
	}
}

func TestQueue_Clear(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue("A")
	queue.Enqueue("B")
	queue.Clear()
	
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after Clear")
	}
}

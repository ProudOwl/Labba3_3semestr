package structures

import (
	"fmt"
	"strings"
)

type QueueNode struct {
	Value string
	Next  *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val string) {
	node := &QueueNode{Value: val}
	if q.tail == nil {
		q.head, q.tail = node, node
		return
	}
	q.tail.Next = node
	q.tail = node
}

func (q *Queue) Dequeue() string {
	if q.head == nil {
		return "[QUEUE_EMPTY]"
	}
	val := q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	return val
}

func (q *Queue) Peek() string {
	if q.head == nil {
		return "[QUEUE_EMPTY]"
	}
	return q.head.Value
}

func (q *Queue) Print() {
	if q.head == nil {
		fmt.Println("Очередь пуста.")
		return
	}
	var values []string
	current := q.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	fmt.Printf("Очередь: %s\n", strings.Join(values, " -> "))
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Clear() {
	q.head, q.tail = nil, nil
}

package structures

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Value string
	Next  *ListNode
}

type List struct {
	head *ListNode
}

func NewList() *List {
	return &List{}
}

func (l *List) AddHead(val string) {
	l.head = &ListNode{Value: val, Next: l.head}
}

func (l *List) AddTail(val string) {
	if l.head == nil {
		l.head = &ListNode{Value: val}
		return
	}
	current := l.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &ListNode{Value: val}
}

func (l *List) DeleteHead() {
	if l.head == nil {
		return
	}
	l.head = l.head.Next
}

func (l *List) DeleteTail() {
	if l.head == nil {
		return
	}
	if l.head.Next == nil {
		l.head = nil
		return
	}
	current := l.head
	for current.Next.Next != nil {
		current = current.Next
	}
	current.Next = nil
}

func (l *List) DeleteByValue(val string) {
	if l.head == nil {
		return
	}
	if l.head.Value == val {
		l.head = l.head.Next
		return
	}
	current := l.head
	for current.Next != nil && current.Next.Value != val {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
}

func (l *List) Find(val string) bool {
	current := l.head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *List) PrintForward() {
	if l.head == nil {
		fmt.Println("Список пуст.")
		return
	}
	var values []string
	current := l.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	fmt.Printf("Список: %s\n", strings.Join(values, " -> "))
}

func (l *List) PrintBackward() {
	if l.head == nil {
		fmt.Println("Список пуст.")
		return
	}
	values := l.collectValues()
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	fmt.Printf("Список (обратный): %s\n", strings.Join(values, " <- "))
}

func (l *List) collectValues() []string {
	var values []string
	current := l.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func (l *List) Clear() {
	l.head = nil
}

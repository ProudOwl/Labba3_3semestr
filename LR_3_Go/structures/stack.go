package structures

import (
	"fmt"
	"strings"
)

type StackNode struct {
	Value string
	Next  *StackNode
}

type Stack struct {
	top *StackNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val string) {
	s.top = &StackNode{Value: val, Next: s.top}
}

func (s *Stack) Pop() string {
	if s.top == nil {
		return "[STACK_EMPTY]"
	}
	val := s.top.Value
	s.top = s.top.Next
	return val
}

func (s *Stack) Peek() string {
	if s.top == nil {
		return "[STACK_EMPTY]"
	}
	return s.top.Value
}

func (s *Stack) Print() {
	if s.top == nil {
		fmt.Println("Стек пуст.")
		return
	}
	var values []string
	current := s.top
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	fmt.Printf("Стек (верх -> низ): %s\n", strings.Join(values, " -> "))
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Clear() {
	s.top = nil
}

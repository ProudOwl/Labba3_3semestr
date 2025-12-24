package structures

import (
	"testing"
)

func TestStack_NewStack(t *testing.T) {
	stack := NewStack()
	if stack == nil {
		t.Error("NewStack should not return nil")
	}
}

func TestStack_PushPop(t *testing.T) {
	stack := NewStack()
	stack.Push("A")
	stack.Push("B")
	
	if stack.Peek() != "B" {
		t.Errorf("Expected 'B' at top, got %s", stack.Peek())
	}
	
	popped := stack.Pop()
	if popped != "B" {
		t.Errorf("Expected 'B' popped, got %s", popped)
	}
	
	if stack.Pop() != "A" {
		t.Errorf("Expected 'A' popped")
	}
	
	if stack.Pop() != "[STACK_EMPTY]" {
		t.Errorf("Expected [STACK_EMPTY], got %s", stack.Pop())
	}
}

func TestStack_PeekEmpty(t *testing.T) {
	stack := NewStack()
	if stack.Peek() != "[STACK_EMPTY]" {
		t.Errorf("Expected [STACK_EMPTY] for empty stack, got %s", stack.Peek())
	}
}

func TestStack_Print(t *testing.T) {
	stack := NewStack()
	stack.Print() // Пустой стек
	
	stack.Push("A")
	stack.Push("B")
	stack.Print() // Непустой стек
}

func TestStack_IsEmpty(t *testing.T) {
	stack := NewStack()
	if !stack.IsEmpty() {
		t.Error("New stack should be empty")
	}
	
	stack.Push("A")
	if stack.IsEmpty() {
		t.Error("Stack with element should not be empty")
	}
	
	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after pop")
	}
}

func TestStack_Clear(t *testing.T) {
	stack := NewStack()
	stack.Push("A")
	stack.Push("B")
	stack.Clear()
	
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after Clear")
	}
}

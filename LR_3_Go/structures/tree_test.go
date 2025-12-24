package structures

import (
	"testing"
)

func TestTree_Clear(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	
	tree.Clear()
	
	if tree.Search(50) || tree.Search(30) || tree.Search(70) {
		t.Error("All elements should be cleared")
	}
	
	// Clear пустого дерева
	tree2 := NewBinaryTree()
	tree2.Clear() // Не должно падать
	
	// Проверяем Print после Clear
	tree.Print()
}

func TestTree_IsEmpty(t *testing.T) {
	tree := NewBinaryTree()
	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
	
	tree.Insert(50)
	if tree.IsEmpty() {
		t.Error("Tree with elements should not be empty")
	}
	
	tree.Clear()
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after Clear")
	}
}

func TestTree_Size(t *testing.T) {
	tree := NewBinaryTree()
	if tree.Size() != 0 {
		t.Errorf("New tree size should be 0, got %d", tree.Size())
	}
	
	tree.Insert(50)
	if tree.Size() != 1 {
		t.Errorf("Tree size should be 1, got %d", tree.Size())
	}
	
	tree.Insert(30)
	tree.Insert(70)
	if tree.Size() != 3 {
		t.Errorf("Tree size should be 3, got %d", tree.Size())
	}
	
	tree.Remove(30)
	if tree.Size() != 2 {
		t.Errorf("Tree size should be 2 after removal, got %d", tree.Size())
	}
	
	tree.Clear()
	if tree.Size() != 0 {
		t.Errorf("Tree size should be 0 after Clear, got %d", tree.Size())
	}
}

func TestTree_EdgeCases(t *testing.T) {
	tree := NewBinaryTree()
	
	// Удаление несуществующего элемента
	tree.Remove(100) // Не должно падать
	
	// Удаление из пустого дерева
	tree2 := NewBinaryTree()
	tree2.Remove(50) // Не должно падать
	
	// Поиск в пустом дереве
	if tree2.Search(50) {
		t.Error("Search in empty tree should return false")
	}
}

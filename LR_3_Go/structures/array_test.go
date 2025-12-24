package structures

import (
	"testing"
)

func TestArray_AddHead(t *testing.T) {
	arr := NewArray()
	arr.AddHead("B")
	arr.AddHead("A")
	
	if arr.Length() != 2 {
		t.Errorf("Expected length 2, got %d", arr.Length())
	}
	if arr.GetAt(0) != "A" {
		t.Errorf("Expected 'A' at index 0, got %s", arr.GetAt(0))
	}
	if arr.GetAt(1) != "B" {
		t.Errorf("Expected 'B' at index 1, got %s", arr.GetAt(1))
	}
}

func TestArray_ReplaceAt(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	arr.ReplaceAt(1, "NewB")
	
	if arr.GetAt(1) != "NewB" {
		t.Errorf("Expected 'NewB' at index 1, got %s", arr.GetAt(1))
	}
	
	// Замена с неверным индексом
	arr.ReplaceAt(10, "X") // Не должно падать
	if arr.Length() != 3 {
		t.Errorf("Length should not change after invalid replace")
	}
}

func TestArray_Clear(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	arr.Clear()
	
	if arr.Length() != 0 {
		t.Errorf("Expected length 0 after Clear, got %d", arr.Length())
	}
	
	// Clear пустого массива
	arr2 := NewArray()
	arr2.Clear() // Не должно падать
	if arr2.Length() != 0 {
		t.Errorf("Empty array should stay empty after Clear")
	}
}

func TestArray_EdgeCases(t *testing.T) {
	arr := NewArray()
	
	// Тестируем граничные случаи
	arr.AddAt(-1, "X") // Неверный индекс - ничего не должно произойти
	if arr.Length() != 0 {
		t.Errorf("Invalid AddAt should not add elements")
	}
	
	arr.AddAt(100, "Y") // Индекс больше размера
	if arr.Length() != 0 {
		t.Errorf("Out of bounds AddAt should not add elements")
	}
	
	arr.DeleteAt(-1) // Неверный индекс удаления
	arr.DeleteAt(100) // Индекс больше размера
	
	// GetAt с неверным индексом
	if arr.GetAt(-1) != "[INVALID_INDEX]" {
		t.Errorf("Expected [INVALID_INDEX] for index -1, got %s", arr.GetAt(-1))
	}
	if arr.GetAt(0) != "[INVALID_INDEX]" {
		t.Errorf("Expected [INVALID_INDEX] for empty array, got %s", arr.GetAt(0))
	}
}

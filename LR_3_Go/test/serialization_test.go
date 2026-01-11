package main

import (
	"os"
	"testing"
	"go-data-structures/structures"
)

func TestArraySerialization(t *testing.T) {
	arr := structures.NewArray()
	arr.AddTail("Москва")
	arr.AddTail("Лондон")
	arr.AddTail("Токио")
	
	// Сохранение
	err := arr.SaveToFile("test_array.txt")
	if err != nil {
		t.Fatalf("Failed to save array: %v", err)
	}
	
	// Загрузка
	arr2 := structures.NewArray()
	err = arr2.LoadFromFile("test_array.txt")
	if err != nil {
		t.Fatalf("Failed to load array: %v", err)
	}
	
	if arr2.Length() != 3 {
		t.Errorf("Expected length 3, got %d", arr2.Length())
	}
	
	// Очистка
	os.Remove("test_array.txt")
}

package structures

import (
	"os"
	"testing"
)

func TestArray_Serialize(t *testing.T) {
	// Создаем и заполняем массив
	arr := NewArray()
	arr.AddTail("Москва")
	arr.AddTail("Лондон")
	arr.AddTail("Токио")
	arr.AddTail("Париж")
	
	// Сохраняем
	filename := "test_array.txt"
	defer os.Remove(filename)
	
	err := arr.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}
	
	// Загружаем в новый массив
	arr2 := NewArray()
	err = arr2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}
	
	// Проверяем
	if arr2.Length() != 4 {
		t.Errorf("Expected length 4, got %d", arr2.Length())
	}
	if arr2.GetAt(0) != "Москва" {
		t.Errorf("Expected 'Москва' at index 0, got %s", arr2.GetAt(0))
	}
	if arr2.GetAt(3) != "Париж" {
		t.Errorf("Expected 'Париж' at index 3, got %s", arr2.GetAt(3))
	}
}

func TestArray_SerializeEmpty(t *testing.T) {
	// Пустой массив
	arr := NewArray()
	
	filename := "test_empty_array.txt"
	defer os.Remove(filename)
	
	err := arr.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed for empty array: %v", err)
	}
	
	arr2 := NewArray()
	err = arr2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed for empty array: %v", err)
	}
	
	if arr2.Length() != 0 {
		t.Errorf("Expected empty array, got length %d", arr2.Length())
	}
}

func TestArray_SerializeMissingFile(t *testing.T) {
	arr := NewArray()
	err := arr.LoadFromFile("non_existent_file.txt")
	if err == nil {
		t.Error("Expected error for missing file")
	}
}

func TestList_Serialize(t *testing.T) {
	// Создаем и заполняем список
	list := NewList()
	list.AddTail("Берлин")
	list.AddTail("Мадрид")
	list.AddTail("Рим")
	
	// Сохраняем
	filename := "test_list.txt"
	defer os.Remove(filename)
	
	err := list.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}
	
	// Загружаем в новый список
	list2 := NewList()
	err = list2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}
	
	// Проверяем
	if !list2.Find("Берлин") {
		t.Error("Expected to find 'Берлин'")
	}
	if !list2.Find("Мадрид") {
		t.Error("Expected to find 'Мадрид'")
	}
	if !list2.Find("Рим") {
		t.Error("Expected to find 'Рим'")
	}
}

func TestList_SerializeEmpty(t *testing.T) {
	list := NewList()
	
	filename := "test_empty_list.txt"
	defer os.Remove(filename)
	
	err := list.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed for empty list: %v", err)
	}
	
	list2 := NewList()
	err = list2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed for empty list: %v", err)
	}
	
	// Должен быть пустым
	list2.PrintForward() // Не должно падать
}

func TestBinaryTree_SerializeBinary(t *testing.T) {
	// Создаем дерево
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)
	
	// Сохраняем в бинарный файл
	filename := "test_tree.bin"
	defer os.Remove(filename)
	
	err := tree.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("SaveToBinaryFile failed: %v", err)
	}
	
	// Загружаем в новое дерево
	tree2 := NewBinaryTree()
	err = tree2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("LoadFromBinaryFile failed: %v", err)
	}
	
	// Проверяем
	if !tree2.Search(50) {
		t.Error("Expected to find 50")
	}
	if !tree2.Search(30) {
		t.Error("Expected to find 30")
	}
	if !tree2.Search(70) {
		t.Error("Expected to find 70")
	}
	if !tree2.Search(20) {
		t.Error("Expected to find 20")
	}
	if !tree2.Search(80) {
		t.Error("Expected to find 80")
	}
	if tree2.Search(100) {
		t.Error("Should not find 100")
	}
}

func TestBinaryTree_SerializeEmpty(t *testing.T) {
	tree := NewBinaryTree()
	
	filename := "test_empty_tree.bin"
	defer os.Remove(filename)
	
	err := tree.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("SaveToBinaryFile failed for empty tree: %v", err)
	}
	
	tree2 := NewBinaryTree()
	err = tree2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("LoadFromBinaryFile failed for empty tree: %v", err)
	}
	
	if tree2.Search(0) {
		t.Error("Empty tree should not find anything")
	}
}

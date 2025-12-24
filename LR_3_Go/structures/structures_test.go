package structures

import (
	"testing"
)

func TestArray_AddAt(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("C")
	arr.AddAt(1, "B")
	
	if arr.Length() != 3 {
		t.Errorf("Expected length 3, got %d", arr.Length())
	}
	if arr.GetAt(1) != "B" {
		t.Errorf("Expected 'B' at index 1, got %s", arr.GetAt(1))
	}
}

func TestArray_DeleteHead(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("B")
	arr.DeleteHead()
	
	if arr.Length() != 1 {
		t.Errorf("Expected length 1, got %d", arr.Length())
	}
	if arr.GetAt(0) != "B" {
		t.Errorf("Expected 'B' at index 0, got %s", arr.GetAt(0))
	}
}

func TestArray_DeleteTail(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("B")
	arr.DeleteTail()
	
	if arr.Length() != 1 {
		t.Errorf("Expected length 1, got %d", arr.Length())
	}
	if arr.GetAt(0) != "A" {
		t.Errorf("Expected 'A' at index 0, got %s", arr.GetAt(0))
	}
}

func TestArray_DeleteAt(t *testing.T) {
	arr := NewArray()
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	arr.DeleteAt(1)
	
	if arr.Length() != 2 {
		t.Errorf("Expected length 2, got %d", arr.Length())
	}
	if arr.GetAt(1) != "C" {
		t.Errorf("Expected 'C' at index 1, got %s", arr.GetAt(1))
	}
}

func TestArray_Print(t *testing.T) {
	arr := NewArray()
	arr.AddTail("Test")
	// Этот тест проверяет что Print не падает
	arr.Print()
}

func TestDList_AddHead(t *testing.T) {
	dlist := NewDList()
	dlist.AddHead("A")
	dlist.AddHead("B")
	
	// Проверяем что не падает
	dlist.PrintForward()
}

func TestDList_AddTail(t *testing.T) {
	dlist := NewDList()
	dlist.AddTail("A")
	dlist.AddTail("B")
	
	// Проверяем что не падает
	dlist.PrintForward()
	dlist.PrintBackward()
}

func TestDList_DeleteByValue(t *testing.T) {
	dlist := NewDList()
	dlist.AddTail("A")
	dlist.AddTail("B")
	dlist.AddTail("C")
	dlist.DeleteByValue("B")
	
	if dlist.Find("B") {
		t.Error("Expected 'B' to be deleted")
	}
}

func TestHash_ChainHash(t *testing.T) {
	hash := NewChainHash(10)
	hash.Insert("key1", "value1")
	hash.Insert("key2", "value2")
	
	if hash.Find("key1") != "value1" {
		t.Errorf("Expected 'value1' for key1, got %s", hash.Find("key1"))
	}
	
	hash.Insert("key1", "updated")
	if hash.Find("key1") != "updated" {
		t.Errorf("Expected 'updated' for key1, got %s", hash.Find("key1"))
	}
	
	if !hash.Delete("key2") {
		t.Error("Expected to delete key2")
	}
	
	if hash.Find("key2") != "" {
		t.Error("key2 should not be found after deletion")
	}
	
	// Проверяем Print
	hash.Print()
}

func TestHash_OpenHash(t *testing.T) {
	hash := NewOpenHash(5)
	hash.Insert("k1", "v1")
	hash.Insert("k2", "v2")
	
	if hash.Find("k1") != "v1" {
		t.Errorf("Expected 'v1' for k1, got %s", hash.Find("k1"))
	}
	
	hash.Delete("k1")
	if hash.Find("k1") != "" {
		t.Error("k1 should not be found after deletion")
	}
	
	// Тест повторного использования удаленного слота
	hash.Insert("k3", "v3")
	if hash.Find("k3") != "v3" {
		t.Errorf("Expected 'v3' for k3, got %s", hash.Find("k3"))
	}
	
	// Проверяем Print
	hash.Print()
}

func TestTree_InsertSearch(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)
	
	if !tree.Search(30) {
		t.Error("Expected to find 30")
	}
	if !tree.Search(80) {
		t.Error("Expected to find 80")
	}
	if tree.Search(100) {
		t.Error("Should not find 100")
	}
	
	// Проверяем Print
	tree.Print()
}

func TestTree_Remove(t *testing.T) {
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	
	tree.Remove(30)
	if tree.Search(30) {
		t.Error("30 should have been removed")
	}
	
	tree.Remove(50)
	if tree.Search(50) {
		t.Error("50 should have been removed")
	}
	
	if !tree.Search(70) {
		t.Error("70 should still exist")
	}
}

func TestSerialize_Array(t *testing.T) {
	arr := NewArray()
	arr.AddTail("Test1")
	arr.AddTail("Test2")
}

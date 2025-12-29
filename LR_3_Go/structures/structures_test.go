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

func TestHash_EdgeCases(t *testing.T) {
	hash1 := NewChainHash(1)
	hash1.Insert("test", "value")
	hash1.Insert("test2", "value2")
	if hash1.Find("test") != "value" {
		t.Error("Should find inserted value")
	}
	hash1.Print()
	
	hash2 := NewOpenHash(2)
	hash2.Insert("a", "1")
	hash2.Insert("b", "2")
	hash2.Insert("c", "3")
	hash2.Print()
	
	hash3 := NewChainHash(3)
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		hash3.Insert(key, "value"+string(rune('0'+i)))
	}
	
	for i := 0; i < 10; i++ {
		key := string(rune('a' + i))
		expected := "value" + string(rune('0'+i))
		if hash3.Find(key) != expected {
			t.Errorf("Expected %s for key %s, got %s", expected, key, hash3.Find(key))
		}
	}
	hash3.Print()
	
	hash4 := NewChainHash(5)
	hash4.Insert("k1", "v1")
	hash4.Insert("k2", "v2")
	hash4.Delete("k1")
	hash4.Delete("k2")
	if hash4.Find("k1") != "" || hash4.Find("k2") != "" {
		t.Error("All elements should be deleted")
	}
	hash4.Print()
}

//Тест для проверки логики удаления
func TestOpenHash_DeleteLogic(t *testing.T) {
	hash := NewOpenHash(5)
	
	hash.Insert("a", "1")
	hash.Insert("b", "2")
	hash.Insert("c", "3")
	
	if !hash.Delete("b") {
		t.Error("Should delete existing key")
	}
	
	idx := hash.hash("b")
	foundDeleted := false
	
	for i := 0; i < hash.cap; i++ {
		currentIdx := (idx + i) % hash.cap
		if hash.table[currentIdx].Deleted && hash.table[currentIdx].Key == "b" {
			foundDeleted = true
			break
		}
		if !hash.table[currentIdx].Used && !hash.table[currentIdx].Deleted {
			break
		}
	}
	
	if !foundDeleted {
		t.Error("Slot should be marked as Deleted")
	}
	
	if hash.Delete("nonexistent") {
		t.Error("Should return false for non-existent key")
	}
	
	if hash.Delete("b") {
		t.Error("Should return false for already deleted key")
	}
}

func TestOpenHash_DeleteAllAndReinsert(t *testing.T) {
	hash := NewOpenHash(3)
	
	hash.Insert("k1", "v1")
	hash.Insert("k2", "v2")
	hash.Insert("k3", "v3")
	
	hash.Delete("k1")
	hash.Delete("k2")
	hash.Delete("k3")
	
	for i := 0; i < hash.cap; i++ {
		if !hash.table[i].Deleted {
			t.Errorf("Slot %d should be marked as Deleted", i)
		}
	}
	
	hash.Insert("new1", "val1")
	hash.Insert("new2", "val2")
	hash.Insert("new3", "val3")
	
	if hash.Find("new1") != "val1" || hash.Find("new2") != "val2" || hash.Find("new3") != "val3" {
		t.Error("Should find all reinserted values")
	}
}

// Удаление из середины кластера
func TestOpenHash_DeleteInMiddleOfCluster(t *testing.T) {
	hash := NewOpenHash(5)
	
	hash.Insert("a", "1")
	hash.Insert("aa", "2")
	hash.Insert("aaa", "3")
	
	if !hash.Delete("aa") {
		t.Error("Should delete from middle of cluster")
	}
	
	if hash.Find("a") != "1" {
		t.Error("Should find a")
	}
	if hash.Find("aaa") != "3" {
		t.Error("Should find aaa")
	}
	if hash.Find("aa") != "" {
		t.Error("Should not find deleted aa")
	}
	
	hash.Insert("bb", "4")
	if hash.Find("bb") != "4" {
		t.Error("Should find bb in reused slot")
	}
}

// Проверка что не зацикливается при поиске
func TestOpenHash_DeleteNotFoundLoops(t *testing.T) {
	hash := NewOpenHash(3)
	
	hash.Insert("k1", "v1")
	hash.Insert("k2", "v2")
	hash.Insert("k3", "v3")
	
	hash.Delete("k1")
	hash.Delete("k2")
	hash.Delete("k3")
	
	if hash.Delete("nonexistent") {
		t.Error("Should return false for non-existent key in full deleted table")
	}
	
	if hash.Find("nonexistent") != "" {
		t.Error("Should return empty string for non-existent key")
	}
}

func TestOpenHash_DeleteThenFind(t *testing.T) {
	hash := NewOpenHash(10)

	keys := []string{"apple", "banana", "cherry", "date", "elderberry"}
	for i, key := range keys {
		hash.Insert(key, string(rune('A'+i)))
	}
	
	for i, key := range keys {
		expected := string(rune('A' + i))
		if hash.Find(key) != expected {
			t.Errorf("Expected %s for %s, got %s", expected, key, hash.Find(key))
		}
	}
	
	hash.Delete("banana")
	hash.Delete("date")
	
	if hash.Find("banana") != "" || hash.Find("date") != "" {
		t.Error("Deleted keys should not be found")
	}
	
	if hash.Find("apple") != "A" || hash.Find("cherry") != "C" || hash.Find("elderberry") != "E" {
		t.Error("Non-deleted keys should still be found")
	}
	
	hash.Insert("blueberry", "F")
	hash.Insert("fig", "G")
	
	if hash.Find("blueberry") != "F" || hash.Find("fig") != "G" {
		t.Error("New keys should be found")
	}
}

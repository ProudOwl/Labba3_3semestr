package structures

import (
	"testing"
)

func TestNewArray(t *testing.T) {
	arr := NewArray()
	
	if arr.Length() != 0 {
		t.Errorf("New array should have length 0, got %d", arr.Length())
	}
	
	if arr.Capacity() != 0 {
		t.Errorf("New array should have capacity 0, got %d", arr.Capacity())
	}
	
	if !arr.IsEmpty() {
		t.Error("New array should be empty")
	}
}

func TestAddHead(t *testing.T) {
	arr := NewArray()
	
	arr.AddHead("A")
	if arr.Length() != 1 {
		t.Errorf("Length should be 1 after AddHead, got %d", arr.Length())
	}
	if arr.GetAt(0) != "A" {
		t.Errorf("Expected 'A' at index 0, got %s", arr.GetAt(0))
	}
	
	arr.AddHead("B")
	if arr.Length() != 2 {
		t.Errorf("Length should be 2, got %d", arr.Length())
	}
	if arr.GetAt(0) != "B" {
		t.Errorf("Expected 'B' at index 0, got %s", arr.GetAt(0))
	}
	if arr.GetAt(1) != "A" {
		t.Errorf("Expected 'A' at index 1, got %s", arr.GetAt(1))
	}
}

func TestAddTail(t *testing.T) {
	arr := NewArray()
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	if arr.Length() != 3 {
		t.Errorf("Expected length 3, got %d", arr.Length())
	}
	
	expected := []string{"A", "B", "C"}
	for i, exp := range expected {
		actual := arr.GetAt(i)
		if actual != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, actual)
		}
	}
}

func TestAddAt(t *testing.T) {
	arr := NewArray()
	
	arr.AddTail("A")
	arr.AddTail("C")
	
	arr.AddAt(1, "B")
	
	expected := []string{"A", "B", "C"}
	if arr.Length() != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), arr.Length())
	}
	
	for i, exp := range expected {
		actual := arr.GetAt(i)
		if actual != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, actual)
		}
	}
	
	arr.AddAt(0, "Start")
	if arr.GetAt(0) != "Start" {
		t.Errorf("AddAt(0) failed, got %s", arr.GetAt(0))
	}
	
	arr.AddAt(arr.Length(), "End")
	if arr.GetAt(arr.Length()-1) != "End" {
		t.Errorf("AddAt(end) failed, got %s", arr.GetAt(arr.Length()-1))
	}
}

func TestDeleteHead(t *testing.T) {
	arr := NewArray()
	
	arr.DeleteHead()
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	arr.DeleteHead()
	
	if arr.Length() != 2 {
		t.Errorf("Length after DelHead should be 2, got %d", arr.Length())
	}
	
	if arr.GetAt(0) != "B" {
		t.Errorf("After DelHead, index 0 should be 'B', got %s", arr.GetAt(0))
	}
	
	if arr.GetAt(1) != "C" {
		t.Errorf("After DelHead, index 1 should be 'C', got %s", arr.GetAt(1))
	}
	
	arr.DeleteHead()
	arr.DeleteHead()
	
	if !arr.IsEmpty() {
		t.Error("Array should be empty after deleting all elements")
	}
	
	arr.DeleteHead()
}

func TestDeleteTail(t *testing.T) {
	arr := NewArray()
	
	arr.DeleteTail()
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	arr.DeleteTail()
	
	if arr.Length() != 2 {
		t.Errorf("Length after DelTail should be 2, got %d", arr.Length())
	}
	
	if arr.GetAt(0) != "A" {
		t.Errorf("Index 0 should still be 'A', got %s", arr.GetAt(0))
	}
	
	if arr.GetAt(1) != "B" {
		t.Errorf("Index 1 should be 'B', got %s", arr.GetAt(1))
	}
	
	if arr.GetAt(2) != "[INVALID_INDEX]" {
		t.Errorf("Index 2 should be invalid, got %s", arr.GetAt(2))
	}
}

func TestDeleteAt(t *testing.T) {
	arr := NewArray()
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	arr.AddTail("D")
	
	arr.DeleteAt(1)
	
	expected := []string{"A", "C", "D"}
	if arr.Length() != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), arr.Length())
	}
	
	for i, exp := range expected {
		actual := arr.GetAt(i)
		if actual != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, actual)
		}
	}
	
	arr.DeleteAt(0)
	if arr.GetAt(0) != "C" {
		t.Errorf("After DelAt(0), index 0 should be 'C', got %s", arr.GetAt(0))
	}
	
	arr.DeleteAt(arr.Length() - 1)
	if arr.Length() != 1 {
		t.Errorf("Length should be 1, got %d", arr.Length())
	}
	
	arr.DeleteAt(-1)
	arr.DeleteAt(100)
	
	if arr.Length() != 1 {
		t.Errorf("Invalid DelAt should not change length, got %d", arr.Length())
	}
}

func TestGetAt(t *testing.T) {
	arr := NewArray()
	
	if arr.GetAt(0) != "[INVALID_INDEX]" {
		t.Errorf("GetAt(0) on empty array should return [INVALID_INDEX], got %s", arr.GetAt(0))
	}
	
	if arr.GetAt(-1) != "[INVALID_INDEX]" {
		t.Errorf("GetAt(-1) should return [INVALID_INDEX], got %s", arr.GetAt(-1))
	}
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	if arr.GetAt(0) != "A" {
		t.Errorf("GetAt(0) should return 'A', got %s", arr.GetAt(0))
	}
	
	if arr.GetAt(2) != "C" {
		t.Errorf("GetAt(2) should return 'C', got %s", arr.GetAt(2))
	}
	
	if arr.GetAt(3) != "[INVALID_INDEX]" {
		t.Errorf("GetAt(3) should return [INVALID_INDEX], got %s", arr.GetAt(3))
	}
	
	if arr.GetAt(-5) != "[INVALID_INDEX]" {
		t.Errorf("GetAt(-5) should return [INVALID_INDEX], got %s", arr.GetAt(-5))
	}
}

func TestReplaceAt(t *testing.T) {
	arr := NewArray()
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	arr.ReplaceAt(1, "NewB")
	
	if arr.GetAt(1) != "NewB" {
		t.Errorf("ReplaceAt failed, got %s at index 1", arr.GetAt(1))
	}
	
	arr.ReplaceAt(0, "NewA")
	if arr.GetAt(0) != "NewA" {
		t.Errorf("ReplaceAt at index 0 failed")
	}
	
	arr.ReplaceAt(2, "NewC")
	if arr.GetAt(2) != "NewC" {
		t.Errorf("ReplaceAt at last index failed")
	}
	
	oldLength := arr.Length()
	arr.ReplaceAt(-1, "X")
	arr.ReplaceAt(100, "Y")
	
	if arr.Length() != oldLength {
		t.Errorf("Invalid ReplaceAt changed length from %d to %d", oldLength, arr.Length())
	}
}

func TestEnsureCapacity(t *testing.T) {
	arr := NewArray()
	
	if arr.Capacity() != 0 {
		t.Errorf("Initial capacity should be 0, got %d", arr.Capacity())
	}
	
	arr.AddTail("A")
	if arr.Capacity() != 4 {
		t.Errorf("After first add, capacity should be 4, got %d", arr.Capacity())
	}
	
	arr.AddTail("B")
	arr.AddTail("C")
	arr.AddTail("D")
	
	arr.AddTail("E")
	if arr.Capacity() != 8 {
		t.Errorf("Capacity should double to 8 when adding 5th element, got %d", arr.Capacity())
	}
	
	for i := 0; i < 4; i++ {
		arr.AddTail("X")
	}
	
	if arr.Capacity() != 16 {
		t.Errorf("Capacity should be 16 for 9+ elements, got %d", arr.Capacity())
	}
}

func TestClear(t *testing.T) {
	arr := NewArray()
	
	for i := 0; i < 10; i++ {
		arr.AddTail("Item")
	}
	
	arr.Clear()
	
	if arr.Length() != 0 {
		t.Errorf("After Clear, length should be 0, got %d", arr.Length())
	}
	
	if !arr.IsEmpty() {
		t.Error("After Clear, array should be empty")
	}
	
	if arr.Capacity() != 0 {
		t.Errorf("After Clear, capacity should be 0, got %d", arr.Capacity())
	}
	
	arr2 := NewArray()
	arr2.Clear()
	
	arr.AddTail("NewItem")
	if arr.Length() != 1 {
		t.Errorf("Should be able to add after Clear, length = %d", arr.Length())
	}
}

func TestContains(t *testing.T) {
	arr := NewArray()
	
	if arr.Contains("A") {
		t.Error("Empty array should not contain anything")
	}
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	
	if !arr.Contains("A") {
		t.Error("Array should contain 'A'")
	}
	
	if !arr.Contains("B") {
		t.Error("Array should contain 'B'")
	}
	
	if !arr.Contains("C") {
		t.Error("Array should contain 'C'")
	}
	
	if arr.Contains("X") {
		t.Error("Array should not contain 'X'")
	}
	
	if arr.Contains("") {
		t.Error("Array should not contain empty string")
	}
	
	arr.DeleteAt(1)
	if arr.Contains("B") {
		t.Error("Array should not contain 'B' after deletion")
	}
}

func TestIndexOf(t *testing.T) {
	arr := NewArray()
	
	if arr.IndexOf("A") != -1 {
		t.Errorf("IndexOf on empty array should return -1, got %d", arr.IndexOf("A"))
	}
	
	arr.AddTail("A")
	arr.AddTail("B")
	arr.AddTail("C")
	arr.AddTail("B")
	
	if arr.IndexOf("A") != 0 {
		t.Errorf("IndexOf('A') should return 0, got %d", arr.IndexOf("A"))
	}
	
	if arr.IndexOf("B") != 1 {
		t.Errorf("IndexOf('B') should return first occurrence (1), got %d", arr.IndexOf("B"))
	}
	
	if arr.IndexOf("C") != 2 {
		t.Errorf("IndexOf('C') should return 2, got %d", arr.IndexOf("C"))
	}
	
	if arr.IndexOf("X") != -1 {
		t.Errorf("IndexOf('X') should return -1, got %d", arr.IndexOf("X"))
	}
	
	arr.DeleteAt(0)
	if arr.IndexOf("A") != -1 {
		t.Errorf("IndexOf('A') after deletion should return -1, got %d", arr.IndexOf("A"))
	}
}

func TestToSlice(t *testing.T) {
	arr := NewArray()
	
	emptySlice := arr.ToSlice()
	if len(emptySlice) != 0 {
		t.Errorf("ToSlice on empty array should return empty slice, got %v", emptySlice)
	}
	
	expected := []string{"A", "B", "C", "D"}
	for _, val := range expected {
		arr.AddTail(val)
	}
	
	result := arr.ToSlice()
	
	if len(result) != len(expected) {
		t.Errorf("Slice length mismatch: expected %d, got %d", len(expected), len(result))
	}
	
	for i, exp := range expected {
		if result[i] != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, result[i])
		}
	}
	
	result[0] = "MODIFIED"
	if arr.GetAt(0) == "MODIFIED" {
		t.Error("Modifying returned slice should not affect array")
	}
}

func TestComplexScenario(t *testing.T) {
	arr := NewArray()
	
	arr.AddTail("A")      
	arr.AddHead("Start")  
	arr.AddAt(1, "B")    
	arr.AddTail("End")   
	
	expected := []string{"Start", "B", "A", "End"}
	if arr.Length() != len(expected) {
		t.Fatalf("Expected length %d, got %d", len(expected), arr.Length())
	}
	
	for i, exp := range expected {
		if arr.GetAt(i) != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, arr.GetAt(i))
		}
	}
	
	arr.DeleteAt(2)          
	arr.DeleteHead()         
	arr.DeleteTail()         
	
	if arr.Length() != 1 {
		t.Errorf("After deletions, length should be 1, got %d", arr.Length())
	}
	
	if arr.GetAt(0) != "B" {
		t.Errorf("After deletions, element should be 'B', got %s", arr.GetAt(0))
	}
	
	arr.ReplaceAt(0, "NewB")
	if arr.GetAt(0) != "NewB" {
		t.Errorf("Replace failed, got %s", arr.GetAt(0))
	}
	
	arr.Clear()
	if !arr.IsEmpty() {
		t.Error("Array should be empty after Clear")
	}
	
	arr.AddTail("NewStart")
	if arr.Length() != 1 || arr.GetAt(0) != "NewStart" {
		t.Errorf("Reuse after Clear failed: length=%d, element=%s", arr.Length(), arr.GetAt(0))
	}
}

func TestEdgeCases(t *testing.T) {
	arr := NewArray()
	
	arr.AddAt(-1, "X")
	arr.AddAt(100, "Y")
	if arr.Length() != 0 {
		t.Errorf("Invalid AddAt should not add elements, length=%d", arr.Length())
	}
	
	arr.DeleteHead()
	arr.DeleteTail()
	arr.DeleteAt(0)
	arr.ReplaceAt(0, "Z")
	
	if arr.Length() != 0 {
		t.Errorf("Operations on empty array should not change length, got %d", arr.Length())
	}
	
	for i := 0; i < 1000; i++ {
		arr.AddTail("Item")
	}
	
	if arr.Length() != 1000 {
		t.Errorf("Length after 1000 adds should be 1000, got %d", arr.Length())
	}
	
	if arr.Capacity() < 1000 {
		t.Errorf("Capacity should be >= 1000, got %d", arr.Capacity())
	}
	
	for i := 0; i < 500; i++ {
		arr.DeleteHead()
	}
	
	if arr.Length() != 500 {
		t.Errorf("Length after deletions should be 500, got %d", arr.Length())
	}
	
	for i := 0; i < arr.Length(); i++ {
		if arr.GetAt(i) != "Item" {
			t.Errorf("Element %d should be 'Item', got %s", i, arr.GetAt(i))
		}
	}
}

func BenchmarkArrayAddTail(b *testing.B) {
	arr := NewArray()
	for i := 0; i < b.N; i++ {
		arr.AddTail("item")
	}
}

func BenchmarkArrayAddHead(b *testing.B) {
	arr := NewArray()
	for i := 0; i < b.N; i++ {
		arr.AddHead("item")
	}
}

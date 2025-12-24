package structures

import (
	"testing"
)

func TestDList_DeleteHead(t *testing.T) {
	dlist := NewDList()
	dlist.AddTail("A")
	dlist.AddTail("B")
	dlist.AddTail("C")
	
	dlist.DeleteHead()
	
	if dlist.Find("A") {
		t.Error("'A' should be deleted")
	}
	if !dlist.Find("B") {
		t.Error("'B' should still exist")
	}
	
	// Удаление из списка с одним элементом
	dlist2 := NewDList()
	dlist2.AddTail("Single")
	dlist2.DeleteHead()
	if dlist2.Find("Single") {
		t.Error("Single element should be deleted")
	}
	
	// Удаление из пустого списка
	dlist3 := NewDList()
	dlist3.DeleteHead() // Не должно падать
}

func TestDList_DeleteTail(t *testing.T) {
	dlist := NewDList()
	dlist.AddTail("A")
	dlist.AddTail("B")
	dlist.AddTail("C")
	
	dlist.DeleteTail()
	
	if dlist.Find("C") {
		t.Error("'C' should be deleted")
	}
	if !dlist.Find("B") {
		t.Error("'B' should still exist")
	}
	
	// Удаление из списка с одним элементом
	dlist2 := NewDList()
	dlist2.AddTail("Single")
	dlist2.DeleteTail()
	if dlist2.Find("Single") {
		t.Error("Single element should be deleted")
	}
	
	// Удаление из пустого списка
	dlist3 := NewDList()
	dlist3.DeleteTail() // Не должно падать
}

func TestDList_Clear(t *testing.T) {
	dlist := NewDList()
	dlist.AddTail("A")
	dlist.AddTail("B")
	dlist.AddTail("C")
	
	dlist.Clear()
	
	if dlist.Find("A") || dlist.Find("B") || dlist.Find("C") {
		t.Error("All elements should be cleared")
	}
	
	// Clear пустого списка
	dlist2 := NewDList()
	dlist2.Clear() // Не должно падать
	
	// Проверяем что Print работает после Clear
	dlist.PrintForward()
	dlist.PrintBackward()
}

package structures

import (
	"testing"
)

func TestList_NewList(t *testing.T) {
	list := NewList()
	if list == nil {
		t.Error("NewList should not return nil")
	}
}

func TestList_AddHead(t *testing.T) {
	list := NewList()
	list.AddHead("A")
	list.AddHead("B")
	
	// Должен быть порядок B -> A
	list.PrintForward() // Проверяем что не падает
}

func TestList_AddTail(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	
	// Должен быть порядок A -> B
	list.PrintForward()
}

func TestList_DeleteHead(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	list.DeleteHead()
	
	// Должен остаться только B
	list.PrintForward()
	
	// Удаление из пустого списка
	list2 := NewList()
	list2.DeleteHead() // Не должно падать
}

func TestList_DeleteTail(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	list.DeleteTail()
	
	// Должен остаться только A
	list.PrintForward()
	
	// Удаление из пустого списка
	list2 := NewList()
	list2.DeleteTail() // Не должно падать
}

func TestList_DeleteByValue(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	list.AddTail("C")
	
	list.DeleteByValue("B")
	
	// Проверяем что B удален
	list.PrintForward()
	
	// Удаление несуществующего значения
	list.DeleteByValue("Z") // Не должно падать
}

func TestList_Find(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	
	if !list.Find("A") {
		t.Error("Should find 'A'")
	}
	
	if list.Find("Z") {
		t.Error("Should not find 'Z'")
	}
	
	// Поиск в пустом списке
	list2 := NewList()
	if list2.Find("anything") {
		t.Error("Empty list should not find anything")
	}
}

func TestList_PrintBackward(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	list.AddTail("C")
	
	list.PrintBackward() // Не должно падать
}

func TestList_Clear(t *testing.T) {
	list := NewList()
	list.AddTail("A")
	list.AddTail("B")
	list.Clear()
	
	// После очистки список должен быть пустым
	list.PrintForward() // Должен написать "Список пуст."
}

package structures

import (
	"fmt"
	"strings"
)

type Array struct {
	data     []string
	size     int 
	capacity int     
}

func NewArray() *Array {
	return &Array{
		data:     make([]string, 0),
		size:     0,
		capacity: 0,
	}
}

func (a *Array) ensureCapacity(newSize int) {
	if newSize <= a.capacity {
		return
	}

	newCapacity := 4
	if a.capacity != 0 {
		newCapacity = a.capacity * 2
	}
	
	for newCapacity < newSize {
		newCapacity *= 2
	}

	newData := make([]string, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	
	a.data = newData
	a.capacity = newCapacity
}

func (a *Array) AddHead(val string) {
	a.ensureCapacity(a.size + 1)
	
	for i := a.size; i > 0; i-- {
		a.data[i] = a.data[i-1]
	}
	
	a.data[0] = val
	a.size++
}

func (a *Array) AddTail(val string) {
	a.ensureCapacity(a.size + 1)
	a.data[a.size] = val
	a.size++
}

func (a *Array) AddAt(idx int, val string) {
	if idx < 0 || idx > a.size {
		return
	}
	
	a.ensureCapacity(a.size + 1)
	
	for i := a.size; i > idx; i-- {
		a.data[i] = a.data[i-1]
	}
	
	a.data[idx] = val
	a.size++
}

func (a *Array) DeleteHead() {
	if a.size == 0 {
		return
	}
	
	for i := 0; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	
	a.size--
	a.data[a.size] = ""
}

func (a *Array) DeleteTail() {
	if a.size == 0 {
		return
	}
	a.size--
	a.data[a.size] = ""
}

func (a *Array) DeleteAt(idx int) {
	if idx < 0 || idx >= a.size {
		return
	}
	
	for i := idx; i < a.size-1; i++ {
		a.data[i] = a.data[i+1]
	}
	
	a.size--
	a.data[a.size] = "" 
}

func (a *Array) GetAt(idx int) string {
	if idx < 0 || idx >= a.size {
		return "[INVALID_INDEX]"
	}
	return a.data[idx]
}

func (a *Array) ReplaceAt(idx int, val string) {
	if idx < 0 || idx >= a.size {
		return
	}
	a.data[idx] = val
}

func (a *Array) Length() int {
	return a.size
}

func (a *Array) Print() {
	if a.size == 0 {
		fmt.Println("Массив пуст.")
		return
	}
	
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Массив [%d/%d]: ", a.size, a.capacity))
	
	for i := 0; i < a.size; i++ {
		sb.WriteString(a.data[i])
		if i < a.size-1 {
			sb.WriteString(" -> ")
		}
	}
	
	fmt.Println(sb.String())
}

func (a *Array) Clear() {
	a.data = make([]string, 0)
	a.size = 0
	a.capacity = 0
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) Contains(val string) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == val {
			return true
		}
	}
	return false
}

func (a *Array) IndexOf(val string) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == val {
			return i
		}
	}
	return -1
}

func (a *Array) ToSlice() []string {
	result := make([]string, a.size)
	for i := 0; i < a.size; i++ {
		result[i] = a.data[i]
	}
	return result
}

func (a *Array) Capacity() int {
	return a.capacity
}

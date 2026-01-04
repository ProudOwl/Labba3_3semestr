package structures

import (
	"fmt"
	"strings"
)

type Array struct {
	data []string
}

func NewArray() *Array {
	return &Array{data: make([]string, 0)}
}

func (a *Array) AddHead(val string) {
	a.data = append([]string{val}, a.data...)
}

func (a *Array) AddTail(val string) {
	a.data = append(a.data, val)
}

func (a *Array) AddAt(idx int, val string) {
	if idx < 0 || idx > len(a.data) {
		return
	}
	a.data = append(a.data[:idx], append([]string{val}, a.data[idx:]...)...)
}

func (a *Array) DeleteHead() {
	if len(a.data) == 0 {
		return
	}
	a.data = a.data[1:]
}

func (a *Array) DeleteTail() {
	if len(a.data) == 0 {
		return
	}
	a.data = a.data[:len(a.data)-1]
}

func (a *Array) DeleteAt(idx int) {
	if idx < 0 || idx >= len(a.data) {
		return
	}
	a.data = append(a.data[:idx], a.data[idx+1:]...)
}

func (a *Array) GetAt(idx int) string {
	if idx < 0 || idx >= len(a.data) {
		return "[INVALID_INDEX]"
	}
	return a.data[idx]
}

func (a *Array) ReplaceAt(idx int, val string) {
	if idx < 0 || idx >= len(a.data) {
		return
	}
	a.data[idx] = val
}

func (a *Array) Length() int {
	return len(a.data)
}

func (a *Array) Print() {
	if len(a.data) == 0 {
		fmt.Println("Массив пуст.")
		return
	}
	fmt.Printf("Массив [%d]: %s\n", len(a.data), strings.Join(a.data, " -> "))
}

func (a *Array) Clear() {
	a.data = make([]string, 0)
}

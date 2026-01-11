package structures

import (
	"fmt"
	"strings"
)

type DListNode struct {
	Value string
	Next  *DListNode
	Prev  *DListNode
}

type DList struct {
	head *DListNode
	tail *DListNode
}

func NewDList() *DList {
	return &DList{}
}

func (d *DList) AddHead(val string) {
	node := &DListNode{Value: val, Next: d.head}
	if d.head != nil {
		d.head.Prev = node
	}
	d.head = node
	if d.tail == nil {
		d.tail = node
	}
}

func (d *DList) AddTail(val string) {
	node := &DListNode{Value: val, Prev: d.tail}
	if d.tail != nil {
		d.tail.Next = node
	}
	d.tail = node
	if d.head == nil {
		d.head = node
	}
}

func (d *DList) DeleteHead() {
	if d.head == nil {
		return
	}
	if d.head == d.tail {
		d.head, d.tail = nil, nil
		return
	}
	d.head = d.head.Next
	d.head.Prev = nil
}

func (d *DList) DeleteTail() {
	if d.tail == nil {
		return
	}
	if d.head == d.tail {
		d.head, d.tail = nil, nil
		return
	}
	d.tail = d.tail.Prev
	d.tail.Next = nil
}

func (d *DList) DeleteByValue(val string) {
	current := d.head
	for current != nil {
		if current.Value == val {
			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {
				d.head = current.Next
			}
			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
				d.tail = current.Prev
			}
			return
		}
		current = current.Next
	}
}

func (d *DList) Find(val string) bool {
	current := d.head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

func (d *DList) PrintForward() {
	if d.head == nil {
		fmt.Println("Список пуст.")
		return
	}
	var values []string
	current := d.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	fmt.Printf("Список: %s\n", strings.Join(values, " <-> "))
}

func (d *DList) PrintBackward() {
	if d.tail == nil {
		fmt.Println("Список пуст.")
		return
	}
	var values []string
	current := d.tail
	for current != nil {
		values = append(values, current.Value)
		current = current.Prev
	}
	fmt.Printf("Список (обратный): %s\n", strings.Join(values, " <-> "))
}

func (d *DList) Clear() {
	d.head, d.tail = nil, nil
}

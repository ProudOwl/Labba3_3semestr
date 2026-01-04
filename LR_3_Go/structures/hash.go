package structures

import (
	"fmt"
)

type ChainHashNode struct {
	Key  string
	Val  string
	Next *ChainHashNode
}

type ChainHash struct {
	table []*ChainHashNode
	cap   int
}

func NewChainHash(capacity int) *ChainHash {
	table := make([]*ChainHashNode, capacity)
	return &ChainHash{table: table, cap: capacity}
}

func (h *ChainHash) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % h.cap
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (h *ChainHash) Insert(key, value string) {
	idx := h.hash(key)
	head := h.table[idx]
	current := head

	for current != nil {
		if current.Key == key {
			current.Val = value
			return
		}
		current = current.Next
	}

	newNode := &ChainHashNode{Key: key, Val: value, Next: head}
	h.table[idx] = newNode
}

func (h *ChainHash) Delete(key string) bool {
	idx := h.hash(key)
	current := h.table[idx]
	var prev *ChainHashNode

	for current != nil {
		if current.Key == key {
			if prev != nil {
				prev.Next = current.Next
			} else {
				h.table[idx] = current.Next
			}
			return true
		}
		prev = current
		current = current.Next
	}
	return false
}

func (h *ChainHash) Find(key string) string {
	idx := h.hash(key)
	current := h.table[idx]

	for current != nil {
		if current.Key == key {
			return current.Val
		}
		current = current.Next
	}
	return ""
}

func (h *ChainHash) Print() {
	hasItems := false
	for i := 0; i < h.cap; i++ {
		if h.table[i] != nil {
			hasItems = true
			fmt.Printf("[%d]: ", i)
			current := h.table[i]
			for current != nil {
				fmt.Printf("%s->%s", current.Key, current.Val)
				if current.Next != nil {
					fmt.Print(" -> ")
				}
				current = current.Next
			}
			fmt.Print("  ")
		}
	}
	if !hasItems {
		fmt.Print("пусто")
	}
	fmt.Println()
}

type OpenHashEntry struct {
	Key     string
	Val     string
	Used    bool
	Deleted bool
}

type OpenHash struct {
	table []OpenHashEntry
	cap   int
}

func NewOpenHash(capacity int) *OpenHash {
	table := make([]OpenHashEntry, capacity)
	return &OpenHash{table: table, cap: capacity}
}

func (h *OpenHash) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash*31 + int(ch)) % h.cap
	}
	if hash < 0 {
		hash = -hash
	}
	return hash
}

func (h *OpenHash) Insert(key, value string) {
	idx := h.hash(key)
	start := idx

	for {
		if !h.table[idx].Used && !h.table[idx].Deleted {
			h.table[idx] = OpenHashEntry{Key: key, Val: value, Used: true}
			return
		}
		if h.table[idx].Used && h.table[idx].Key == key {
			h.table[idx].Val = value
			return
		}
		if h.table[idx].Deleted {
			h.table[idx] = OpenHashEntry{Key: key, Val: value, Used: true}
			return
		}

		idx = (idx + 1) % h.cap
		if idx == start {
			fmt.Println("Таблица переполнена")
			return
		}
	}
}

func (h *OpenHash) Delete(key string) bool {
	idx := h.hash(key)
	start := idx

	for {
		if h.table[idx].Used && h.table[idx].Key == key {
			h.table[idx].Used = false
			h.table[idx].Deleted = true
			return true
		}
		if !h.table[idx].Used && !h.table[idx].Deleted {
			return false
		}
		idx = (idx + 1) % h.cap
		if idx == start {
			return false
		}
	}
}

func (h *OpenHash) Find(key string) string {
	idx := h.hash(key)
	start := idx

	for {
		if h.table[idx].Used && h.table[idx].Key == key {
			return h.table[idx].Val
		}
		if !h.table[idx].Used && !h.table[idx].Deleted {
			return ""
		}
		idx = (idx + 1) % h.cap
		if idx == start {
			return ""
		}
	}
}

func (h *OpenHash) Print() {
	hasItems := false
	for i := 0; i < h.cap; i++ {
		if h.table[i].Used {
			hasItems = true
			fmt.Printf("[%d]%s->%s  ", i, h.table[i].Key, h.table[i].Val)
		} else if h.table[i].Deleted {
			hasItems = true
			fmt.Printf("[%d](удалено)  ", i)
		}
	}
	if !hasItems {
		fmt.Print("пусто")
	}
	fmt.Println()
}

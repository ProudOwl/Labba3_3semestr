package structures

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
	size int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) {
	node := &TreeNode{Data: value}
	if t.root == nil {
		t.root = node
		t.size = 1
		return
	}

	// Реализация как полного бинарного дерева
	nodes := []*TreeNode{t.root}
	for len(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		if current.Left == nil {
			current.Left = node
			break
		} else {
			nodes = append(nodes, current.Left)
		}

		if current.Right == nil {
			current.Right = node
			break
		} else {
			nodes = append(nodes, current.Right)
		}
	}
	t.size++
}

func (t *BinaryTree) Search(value int) bool {
	return t.searchRecursive(t.root, value)
}

func (t *BinaryTree) searchRecursive(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if node.Data == value {
		return true
	}
	return t.searchRecursive(node.Left, value) || t.searchRecursive(node.Right, value)
}

func (t *BinaryTree) Remove(value int) {
	if t.root == nil {
		fmt.Println("Нельзя удалить из пустого дерева.")
		return
	}

	// Находим узел для удаления и последний узел
	var target, last *TreeNode
	nodes := []*TreeNode{t.root}
	for len(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		if current.Data == value {
			target = current
		}

		if current.Left != nil {
			nodes = append(nodes, current.Left)
		}
		if current.Right != nil {
			nodes = append(nodes, current.Right)
		}

		last = current
	}

	if target == nil {
		fmt.Printf("Элемент %d не найден.\n", value)
		return
	}

	if target == last {
		t.root = nil
		t.size = 0
		return
	}

	target.Data = last.Data
	t.removeLastNode(last)
	t.size--
}

func (t *BinaryTree) removeLastNode(last *TreeNode) {
	nodes := []*TreeNode{t.root}
	for len(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]

		if current.Left == last {
			current.Left = nil
			return
		} else if current.Left != nil {
			nodes = append(nodes, current.Left)
		}

		if current.Right == last {
			current.Right = nil
			return
		} else if current.Right != nil {
			nodes = append(nodes, current.Right)
		}
	}
}

func (t *BinaryTree) Print() {
	if t.root == nil {
		fmt.Println("Дерево пусто.")
		return
	}

	levels := int(math.Log2(float64(t.size))) + 1
	t.printRecursive(t.root, 0, levels)
}

func (t *BinaryTree) printRecursive(node *TreeNode, level, maxLevel int) {
	if node == nil || level >= maxLevel {
		return
	}

	t.printRecursive(node.Right, level+1, maxLevel)

	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}
	fmt.Println(node.Data)

	t.printRecursive(node.Left, level+1, maxLevel)
}

func (t *BinaryTree) Clear() {
	t.root = nil
	t.size = 0
}

func (t *BinaryTree) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree) Size() int {
	return t.size
}

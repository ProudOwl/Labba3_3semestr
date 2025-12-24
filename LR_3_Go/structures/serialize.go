package structures

import (
	"bufio"
	"encoding/binary"
	"os"
	"strconv"
)

// ========== Array ==========
func (a *Array) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	// Пишем размер
	writer.WriteString(strconv.Itoa(len(a.data)) + "\n")
	
	// Пишем данные
	for _, val := range a.data {
		writer.WriteString(val + "\n")
	}
	
	return writer.Flush()
}

func (a *Array) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	// Читаем размер
	if !scanner.Scan() {
		return nil
	}
	
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	// Очищаем текущие данные
	a.data = make([]string, 0, size)
	
	// Читаем данные
	for i := 0; i < size && scanner.Scan(); i++ {
		a.data = append(a.data, scanner.Text())
	}
	
	return nil
}

// ========== List ==========
func (l *List) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	// Считаем количество элементов
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.Next
	}
	
	// Пишем количество
	writer.WriteString(strconv.Itoa(count) + "\n")
	
	// Пишем элементы
	current = l.head
	for current != nil {
		writer.WriteString(current.Value + "\n")
		current = current.Next
	}
	
	return writer.Flush()
}

func (l *List) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	// Читаем количество
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	// Очищаем текущий список
	l.head = nil
	
	// Читаем и добавляем элементы
	for i := 0; i < count && scanner.Scan(); i++ {
		if l.head == nil {
			l.head = &ListNode{Value: scanner.Text()}
		} else {
			current := l.head
			for current.Next != nil {
				current = current.Next
			}
			current.Next = &ListNode{Value: scanner.Text()}
		}
	}
	
	return nil
}

// ========== BinaryTree (бинарная сериализация) ==========
func (t *BinaryTree) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.serializeNode(file, t.root)
}

func (t *BinaryTree) serializeNode(file *os.File, node *TreeNode) error {
	if node == nil {
		// Маркер для nil узла
		marker := int32(-1)
		return binary.Write(file, binary.LittleEndian, marker)
	}
	
	// Пишем значение узла
	if err := binary.Write(file, binary.LittleEndian, int32(node.Data)); err != nil {
		return err
	}
	
	// Рекурсивно сериализуем левое и правое поддеревья
	if err := t.serializeNode(file, node.Left); err != nil {
		return err
	}
	
	return t.serializeNode(file, node.Right)
}

func (t *BinaryTree) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	t.root = t.deserializeNode(file)
	return nil
}

func (t *BinaryTree) deserializeNode(file *os.File) *TreeNode {
	var value int32
	if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
		return nil
	}
	
	// -1 = маркер nil узла
	if value == -1 {
		return nil
	}
	
	node := &TreeNode{Data: int(value)}
	node.Left = t.deserializeNode(file)
	node.Right = t.deserializeNode(file)
	return node
}

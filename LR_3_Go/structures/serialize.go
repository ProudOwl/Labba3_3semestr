package structures

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"os"
	"strconv"
)

func (a *Array) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	writer.WriteString(strconv.Itoa(a.size) + "\n")
	
	for i := 0; i < a.size; i++ {
		writer.WriteString(a.data[i] + "\n")
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
	
	if !scanner.Scan() {
		return nil
	}
	
	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	a.Clear()
	a.ensureCapacity(size)
	
	for i := 0; i < size && scanner.Scan(); i++ {
		a.data[i] = scanner.Text()
		a.size++
	}
	
	return nil
}

func (l *List) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.Next
	}
	
	writer.WriteString(strconv.Itoa(count) + "\n")
	
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
	
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	l.head = nil
	
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

func (t *BinaryTree) SaveToBinaryFile(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    if err := t.serializeNode(writer, t.root); err != nil {
        return err
    }
    return writer.Flush()
}

func (t *BinaryTree) serializeNode(writer *bufio.Writer, node *TreeNode) error {
    if node == nil {
        marker := int32(-1)
        return binary.Write(writer, binary.LittleEndian, marker)
    }
    
    if err := binary.Write(writer, binary.LittleEndian, int32(node.Data)); err != nil {
        return err
    }
    
    if err := t.serializeNode(writer, node.Left); err != nil {
        return err
    }
    
    return t.serializeNode(writer, node.Right)
}

func (t *BinaryTree) LoadFromBinaryFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
	defer file.Close()

    reader := bufio.NewReader(file)
    t.root, err = t.deserializeNode(reader)
    return err
}

func (t *BinaryTree) deserializeNode(reader *bufio.Reader) (*TreeNode, error) {
    var value int32
    
    err := binary.Read(reader, binary.LittleEndian, &value)
    if err != nil {
        return nil, err
    }
    
    if value == -1 {
        return nil, nil
    }
    
    node := &TreeNode{Data: int(value)}
    
    node.Left, err = t.deserializeNode(reader)
    if err != nil {
        return nil, err
    }
    
    node.Right, err = t.deserializeNode(reader)
    if err != nil {
        return nil, err
    }
    
    return node, nil
}

func (d *DList) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := d.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	writer.WriteString(strconv.Itoa(len(values)) + "\n")
	
	for _, val := range values {
		writer.WriteString(val + "\n")
	}
	
	return writer.Flush()
}

func (d *DList) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	d.head, d.tail = nil, nil
	
	for i := 0; i < count && scanner.Scan(); i++ {
		d.AddTail(scanner.Text())
	}
	
	return nil
}

func (q *Queue) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := q.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	writer.WriteString(strconv.Itoa(len(values)) + "\n")
	
	for _, val := range values {
		writer.WriteString(val + "\n")
	}
	
	return writer.Flush()
}

func (q *Queue) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	q.head, q.tail = nil, nil
	
	for i := 0; i < count && scanner.Scan(); i++ {
		q.Enqueue(scanner.Text())
	}
	
	return nil
}

func (s *Stack) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := s.top
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	writer.WriteString(strconv.Itoa(len(values)) + "\n")
	
	for _, val := range values {
		writer.WriteString(val + "\n")
	}
	
	return writer.Flush()
}

func (s *Stack) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	s.top = nil
	
	var values []string
	for i := 0; i < count && scanner.Scan(); i++ {
		values = append(values, scanner.Text())
	}
	
	for i := len(values) - 1; i >= 0; i-- {
		s.Push(values[i])
	}
	
	return nil
}

type ChainHashEntry struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func (h *ChainHash) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	var entries []ChainHashEntry
	for i := 0; i < h.cap; i++ {
		current := h.table[i]
		for current != nil {
			entries = append(entries, ChainHashEntry{
				Key: current.Key,
				Val: current.Val,
			})
			current = current.Next
		}
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	data := struct {
		Capacity int               `json:"capacity"`
		Entries  []ChainHashEntry `json:"entries"`
	}{
		Capacity: h.cap,
		Entries:  entries,
	}
	
	return encoder.Encode(data)
}

func (h *ChainHash) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	
	var data struct {
		Capacity int               `json:"capacity"`
		Entries  []ChainHashEntry `json:"entries"`
	}
	
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	
	h.table = make([]*ChainHashNode, data.Capacity)
	h.cap = data.Capacity
	
	for _, entry := range data.Entries {
		h.Insert(entry.Key, entry.Val)
	}
	
	return nil
}

type OpenHashEntrySerializable struct {
	Key     string `json:"key"`
	Val     string `json:"val"`
	Used    bool   `json:"used"`
	Deleted bool   `json:"deleted"`
}

func (h *OpenHash) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	var entries []OpenHashEntrySerializable
	for i := 0; i < h.cap; i++ {
		entries = append(entries, OpenHashEntrySerializable{
			Key:     h.table[i].Key,
			Val:     h.table[i].Val,
			Used:    h.table[i].Used,
			Deleted: h.table[i].Deleted,
		})
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	
	data := struct {
		Capacity int                         `json:"capacity"`
		Entries  []OpenHashEntrySerializable `json:"entries"`
	}{
		Capacity: h.cap,
		Entries:  entries,
	}
	
	return encoder.Encode(data)
}

func (h *OpenHash) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	
	var data struct {
		Capacity int                         `json:"capacity"`
		Entries  []OpenHashEntrySerializable `json:"entries"`
	}
	
	if err := decoder.Decode(&data); err != nil {
		return err
	}
	
	h.table = make([]OpenHashEntry, data.Capacity)
	h.cap = data.Capacity
	
	for i, entry := range data.Entries {
		h.table[i] = OpenHashEntry{
			Key:     entry.Key,
			Val:     entry.Val,
			Used:    entry.Used,
			Deleted: entry.Deleted,
		}
	}
	
	return nil
}

func writeStringBinary(writer *bufio.Writer, s string) error {
	length := int32(len(s))
	if err := binary.Write(writer, binary.LittleEndian, length); err != nil {
		return err
	}
	
	if _, err := writer.WriteString(s); err != nil {
		return err
	}
	
	return nil
}

func readStringBinary(reader *bufio.Reader) (string, error) {
	var length int32
	if err := binary.Read(reader, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	
	buffer := make([]byte, length)
	if _, err := reader.Read(buffer); err != nil {
		return "", err
	}
	
	return string(buffer), nil
}

func (a *Array) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	count := int32(a.size)
	if err := binary.Write(writer, binary.LittleEndian, count); err != nil {
		return err
	}
	
	for i := 0; i < a.size; i++ {
		if err := writeStringBinary(writer, a.data[i]); err != nil {
			return err
		}
	}
	
	return writer.Flush()
}

func (a *Array) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var count int32
	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return err
	}
	
	a.Clear()
	a.ensureCapacity(int(count))
	
	for i := int32(0); i < count; i++ {
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		a.data[i] = val
		a.size++
	}
	
	return nil
}

func (l *List) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.Next
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(count)); err != nil {
		return err
	}
	
	current = l.head
	for current != nil {
		if err := writeStringBinary(writer, current.Value); err != nil {
			return err
		}
		current = current.Next
	}
	
	return writer.Flush()
}

func (l *List) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var count int32
	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return err
	}
	
	l.head = nil
	
	var tail *ListNode
	
	for i := int32(0); i < count; i++ {
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		
		node := &ListNode{Value: val}
		if l.head == nil {
			l.head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	
	return nil
}

func (d *DList) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := d.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(len(values))); err != nil {
		return err
	}
	
	for _, val := range values {
		if err := writeStringBinary(writer, val); err != nil {
			return err
		}
	}
	
	return writer.Flush()
}

func (d *DList) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var count int32
	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return err
	}
	
	d.head, d.tail = nil, nil
	
	for i := int32(0); i < count; i++ {
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		d.AddTail(val)
	}
	
	return nil
}

func (q *Queue) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := q.head
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(len(values))); err != nil {
		return err
	}
	
	for _, val := range values {
		if err := writeStringBinary(writer, val); err != nil {
			return err
		}
	}
	
	return writer.Flush()
}

func (q *Queue) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var count int32
	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return err
	}
	
	q.head, q.tail = nil, nil
	
	for i := int32(0); i < count; i++ {
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		q.Enqueue(val)
	}
	
	return nil
}

func (s *Stack) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var values []string
	current := s.top
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(len(values))); err != nil {
		return err
	}
	
	for _, val := range values {
		if err := writeStringBinary(writer, val); err != nil {
			return err
		}
	}
	
	return writer.Flush()
}

func (s *Stack) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var count int32
	if err := binary.Read(reader, binary.LittleEndian, &count); err != nil {
		return err
	}
	
	values := make([]string, count)
	for i := int32(0); i < count; i++ {
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		values[i] = val
	}
	
	s.top = nil
	
	for i := len(values) - 1; i >= 0; i-- {
		s.Push(values[i])
	}
	
	return nil
}

func (h *ChainHash) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	var pairs []struct {
		key string
		val string
	}
	for i := 0; i < h.cap; i++ {
		current := h.table[i]
		for current != nil {
			pairs = append(pairs, struct {
				key string
				val string
			}{
				key: current.Key,
				val: current.Val,
			})
			current = current.Next
		}
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(h.cap)); err != nil {
		return err
	}
	
	if err := binary.Write(writer, binary.LittleEndian, int32(len(pairs))); err != nil {
		return err
	}
	
	for _, pair := range pairs {
		if err := writeStringBinary(writer, pair.key); err != nil {
			return err
		}
		if err := writeStringBinary(writer, pair.val); err != nil {
			return err
		}
	}
	
	return writer.Flush()
}

func (h *ChainHash) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var capacity int32
	if err := binary.Read(reader, binary.LittleEndian, &capacity); err != nil {
		return err
	}
	
	var pairCount int32
	if err := binary.Read(reader, binary.LittleEndian, &pairCount); err != nil {
		return err
	}
	
	h.table = make([]*ChainHashNode, capacity)
	h.cap = int(capacity)
	
	for i := int32(0); i < pairCount; i++ {
		key, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		
		val, err := readStringBinary(reader)
		if err != nil {
			return err
		}
		
		h.Insert(key, val)
	}
	
	return nil
}

func (h *OpenHash) SaveToBinaryFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	if err := binary.Write(writer, binary.LittleEndian, int32(h.cap)); err != nil {
		return err
	}
	
	for i := 0; i < h.cap; i++ {
		flags := byte(0)
		if h.table[i].Used {
			flags |= 1 << 0
		}
		if h.table[i].Deleted {
			flags |= 1 << 1
		}
		
		if err := writer.WriteByte(flags); err != nil {
			return err
		}
		
		if h.table[i].Used {
			if err := writeStringBinary(writer, h.table[i].Key); err != nil {
				return err
			}
			if err := writeStringBinary(writer, h.table[i].Val); err != nil {
				return err
			}
		}
	}
	
	return writer.Flush()
}

func (h *OpenHash) LoadFromBinaryFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	
	var capacity int32
	if err := binary.Read(reader, binary.LittleEndian, &capacity); err != nil {
		return err
	}
	
	h.table = make([]OpenHashEntry, capacity)
	h.cap = int(capacity)
	
	for i := 0; i < h.cap; i++ {
		flags, err := reader.ReadByte()
		if err != nil {
			return err
		}
		
		used := (flags & (1 << 0)) != 0
		deleted := (flags & (1 << 1)) != 0
		
		var key, val string
		if used {
			key, err = readStringBinary(reader)
			if err != nil {
				return err
			}
			
			val, err = readStringBinary(reader)
			if err != nil {
				return err
			}
		}
		
		h.table[i] = OpenHashEntry{
			Key:     key,
			Val:     val,
			Used:    used,
			Deleted: deleted,
		}
	}
	
	return nil
}

func (t *BinaryTree) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	
	nodes := []*TreeNode{t.root}
	serialized := []string{}
	
	for len(nodes) > 0 {
		current := nodes[0]
		nodes = nodes[1:]
		
		if current == nil {
			serialized = append(serialized, "null")
		} else {
			serialized = append(serialized, strconv.Itoa(current.Data))
			nodes = append(nodes, current.Left, current.Right)
		}
	}
	
	for len(serialized) > 0 && serialized[len(serialized)-1] == "null" {
		serialized = serialized[:len(serialized)-1]
	}
	
	writer.WriteString(strconv.Itoa(len(serialized)) + "\n")
	
	for _, val := range serialized {
		writer.WriteString(val + "\n")
	}
	
	return writer.Flush()
}

func (t *BinaryTree) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if !scanner.Scan() {
		return nil
	}
	
	count, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}
	
	if count == 0 {
		t.root = nil
		t.size = 0
		return nil
	}
	
	values := make([]string, count)
	for i := 0; i < count && scanner.Scan(); i++ {
		values[i] = scanner.Text()
	}
	
	if values[0] == "null" {
		t.root = nil
		t.size = 0
		return nil
	}
	
	rootVal, err := strconv.Atoi(values[0])
	if err != nil {
		return err
	}
	
	t.root = &TreeNode{Data: rootVal}
	t.size = 1
	
	nodes := []*TreeNode{t.root}
	index := 1
	
	for len(nodes) > 0 && index < count {
		current := nodes[0]
		nodes = nodes[1:]
		
		if index < count && values[index] != "null" {
			leftVal, err := strconv.Atoi(values[index])
			if err != nil {
				return err
			}
			current.Left = &TreeNode{Data: leftVal}
			nodes = append(nodes, current.Left)
			t.size++
		}
		index++
		
		if index < count && values[index] != "null" {
			rightVal, err := strconv.Atoi(values[index])
			if err != nil {
				return err
			}
			current.Right = &TreeNode{Data: rightVal}
			nodes = append(nodes, current.Right)
			t.size++
		}
		index++
	}
	
	return nil
}

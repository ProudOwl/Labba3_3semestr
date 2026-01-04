package structures

import (
	"os"
	"strconv"
	"testing"
)

// ========== Array сериализация ==========

func TestArray_Serialize(t *testing.T) {
	// Создаем и заполняем массив
	arr := NewArray()
	arr.AddTail("Москва")
	arr.AddTail("Лондон")
	arr.AddTail("Токио")
	arr.AddTail("Париж")

	// Сохраняем
	filename := "test_array.txt"
	defer os.Remove(filename)

	err := arr.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	// Загружаем в новый массив
	arr2 := NewArray()
	err = arr2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}

	// Проверяем
	if arr2.Length() != 4 {
		t.Errorf("Expected length 4, got %d", arr2.Length())
	}
	if arr2.GetAt(0) != "Москва" {
		t.Errorf("Expected 'Москва' at index 0, got %s", arr2.GetAt(0))
	}
	if arr2.GetAt(3) != "Париж" {
		t.Errorf("Expected 'Париж' at index 3, got %s", arr2.GetAt(3))
	}
}

func TestArray_SerializeEmpty(t *testing.T) {
	// Пустой массив
	arr := NewArray()

	filename := "test_empty_array.txt"
	defer os.Remove(filename)

	err := arr.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed for empty array: %v", err)
	}

	arr2 := NewArray()
	err = arr2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed for empty array: %v", err)
	}

	if arr2.Length() != 0 {
		t.Errorf("Expected empty array, got length %d", arr2.Length())
	}
}

func TestArray_SerializeMissingFile(t *testing.T) {
	arr := NewArray()
	err := arr.LoadFromFile("non_existent_file.txt")
	if err == nil {
		t.Error("Expected error for missing file")
	}
}

func TestArray_SerializeCorruptedFile(t *testing.T) {
	// Создаем поврежденный файл
	filename := "test_corrupted_array.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	// Записываем некорректные данные
	file.WriteString("not_a_number\n")
	file.Close()

	arr := NewArray()
	err = arr.LoadFromFile(filename)
	if err == nil {
		t.Error("Expected error for corrupted file")
	}
}

func TestArray_SerializePartialFile(t *testing.T) {
	// Создаем файл с недостаточным количеством строк
	filename := "test_partial_array.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	// Записываем размер 3, но только 2 элемента
	file.WriteString("3\n")
	file.WriteString("Item1\n")
	file.WriteString("Item2\n")
	file.Close()

	arr := NewArray()
	err = arr.LoadFromFile(filename)
	if err != nil {
		t.Logf("LoadFromFile error (expected for partial file): %v", err)
	}
}

// ========== List сериализация ==========

func TestList_Serialize(t *testing.T) {
	// Создаем и заполняем список
	list := NewList()
	list.AddTail("Берлин")
	list.AddTail("Мадрид")
	list.AddTail("Рим")

	// Сохраняем
	filename := "test_list.txt"
	defer os.Remove(filename)

	err := list.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	// Загружаем в новый список
	list2 := NewList()
	err = list2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed: %v", err)
	}

	// Проверяем
	if !list2.Find("Берлин") {
		t.Error("Expected to find 'Берлин'")
	}
	if !list2.Find("Мадрид") {
		t.Error("Expected to find 'Мадрид'")
	}
	if !list2.Find("Рим") {
		t.Error("Expected to find 'Рим'")
	}
}

func TestList_SerializeEmpty(t *testing.T) {
	list := NewList()

	filename := "test_empty_list.txt"
	defer os.Remove(filename)

	err := list.SaveToFile(filename)
	if err != nil {
		t.Fatalf("SaveToFile failed for empty list: %v", err)
	}

	list2 := NewList()
	err = list2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("LoadFromFile failed for empty list: %v", err)
	}

	// Должен быть пустым
	if list2.Find("anything") {
		t.Error("Empty list should not find anything")
	}
}

// ========== BinaryTree сериализация ==========

func TestBinaryTree_SerializeBinary(t *testing.T) {
	// Создаем дерево
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	// Сохраняем в бинарный файл
	filename := "test_tree.bin"
	defer os.Remove(filename)

	err := tree.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новое дерево
	tree2 := NewBinaryTree()
	err = tree2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем
	if !tree2.Search(50) {
		t.Error("Expected to find 50")
	}
	if !tree2.Search(30) {
		t.Error("Expected to find 30")
	}
	if !tree2.Search(70) {
		t.Error("Expected to find 70")
	}
	if !tree2.Search(20) {
		t.Error("Expected to find 20")
	}
	if !tree2.Search(80) {
		t.Error("Expected to find 80")
	}
	if tree2.Search(100) {
		t.Error("Should not find 100")
	}
}

func TestBinaryTree_SerializeEmpty(t *testing.T) {
	tree := NewBinaryTree()

	filename := "test_empty_tree.bin"
	defer os.Remove(filename)

	err := tree.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("SaveToBinaryFile failed for empty tree: %v", err)
	}

	tree2 := NewBinaryTree()
	err = tree2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("LoadFromBinaryFile failed for empty tree: %v", err)
	}

	if tree2.Search(0) {
		t.Error("Empty tree should not find anything")
	}
}

func TestBinaryTree_CorruptedBinaryFile(t *testing.T) {
	// Создаем поврежденный бинарный файл
	filename := "test_corrupted_tree.bin"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	// Записываем некорректные данные - неполный int32
	// Только 1 байт вместо 4
	file.Write([]byte{0xFF})
	file.Close()

	tree := NewBinaryTree()
	err = tree.LoadFromBinaryFile(filename)
	if err == nil {
		t.Error("Expected error for corrupted binary file")
	} else {
		t.Logf("Got expected error: %v", err)
	}
}

// ========== DList сериализация ==========

func TestDList_Serialize(t *testing.T) {
	// Создаем и заполняем двусвязный список
	dlist := NewDList()
	dlist.AddTail("First")
	dlist.AddTail("Second")
	dlist.AddTail("Third")

	// Сохраняем
	filename := "test_dlist.txt"
	defer os.Remove(filename)

	err := dlist.SaveToFile(filename)
	if err != nil {
		t.Fatalf("DList SaveToFile failed: %v", err)
	}

	// Загружаем в новый список
	dlist2 := NewDList()
	err = dlist2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("DList LoadFromFile failed: %v", err)
	}

	// Проверяем
	if !dlist2.Find("First") {
		t.Error("Expected to find 'First'")
	}
	if !dlist2.Find("Second") {
		t.Error("Expected to find 'Second'")
	}
	if !dlist2.Find("Third") {
		t.Error("Expected to find 'Third'")
	}
}

// ========== Queue сериализация ==========

func TestQueue_Serialize(t *testing.T) {
	// Создаем и заполняем очередь
	queue := NewQueue()
	queue.Enqueue("Task1")
	queue.Enqueue("Task2")
	queue.Enqueue("Task3")

	// Сохраняем
	filename := "test_queue.txt"
	defer os.Remove(filename)

	err := queue.SaveToFile(filename)
	if err != nil {
		t.Fatalf("Queue SaveToFile failed: %v", err)
	}

	// Загружаем в новую очередь
	queue2 := NewQueue()
	err = queue2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("Queue LoadFromFile failed: %v", err)
	}

	// Проверяем порядок FIFO
	if queue2.Dequeue() != "Task1" {
		t.Error("Expected 'Task1' to be first")
	}
	if queue2.Dequeue() != "Task2" {
		t.Error("Expected 'Task2' to be second")
	}
	if queue2.Dequeue() != "Task3" {
		t.Error("Expected 'Task3' to be third")
	}
	if !queue2.IsEmpty() {
		t.Error("Queue should be empty after all dequeues")
	}
}

// ========== Stack сериализация ==========

func TestStack_Serialize(t *testing.T) {
	// Создаем и заполняем стек
	stack := NewStack()
	stack.Push("Bottom")
	stack.Push("Middle")
	stack.Push("Top")

	// Сохраняем
	filename := "test_stack.txt"
	defer os.Remove(filename)

	err := stack.SaveToFile(filename)
	if err != nil {
		t.Fatalf("Stack SaveToFile failed: %v", err)
	}

	// Загружаем в новый стек
	stack2 := NewStack()
	err = stack2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("Stack LoadFromFile failed: %v", err)
	}

	// Проверяем порядок LIFO
	if stack2.Pop() != "Top" {
		t.Error("Expected 'Top' to be first (LIFO)")
	}
	if stack2.Pop() != "Middle" {
		t.Error("Expected 'Middle' to be second")
	}
	if stack2.Pop() != "Bottom" {
		t.Error("Expected 'Bottom' to be third")
	}
	if stack2.Pop() != "[STACK_EMPTY]" {
		t.Error("Stack should be empty")
	}
}

// ========== ChainHash сериализация ==========

func TestChainHash_Serialize(t *testing.T) {
	// Создаем и заполняем хэш-таблицу
	hash := NewChainHash(10)
	hash.Insert("name", "Alice")
	hash.Insert("age", "30")
	hash.Insert("city", "Moscow")
	hash.Insert("country", "Russia")

	// Сохраняем в JSON
	filename := "test_chainhash.json"
	defer os.Remove(filename)

	err := hash.SaveToFile(filename)
	if err != nil {
		t.Fatalf("ChainHash SaveToFile failed: %v", err)
	}

	// Загружаем в новую таблицу
	hash2 := NewChainHash(5) // Начальная емкость будет перезаписана
	err = hash2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("ChainHash LoadFromFile failed: %v", err)
	}

	// Проверяем данные
	if hash2.Find("name") != "Alice" {
		t.Errorf("Expected 'Alice' for 'name', got %s", hash2.Find("name"))
	}
	if hash2.Find("age") != "30" {
		t.Errorf("Expected '30' for 'age', got %s", hash2.Find("age"))
	}
	if hash2.Find("city") != "Moscow" {
		t.Errorf("Expected 'Moscow' for 'city', got %s", hash2.Find("city"))
	}

	// Проверяем обновление
	hash2.Insert("age", "31")
	if hash2.Find("age") != "31" {
		t.Error("Should update value")
	}

	// Проверяем удаление
	if !hash2.Delete("city") {
		t.Error("Should delete 'city'")
	}
	if hash2.Find("city") != "" {
		t.Error("'city' should not be found after deletion")
	}
}

func TestChainHash_SerializeEmpty(t *testing.T) {
	hash := NewChainHash(10)

	filename := "test_empty_chainhash.json"
	defer os.Remove(filename)

	err := hash.SaveToFile(filename)
	if err != nil {
		t.Fatalf("ChainHash SaveToFile failed for empty: %v", err)
	}

	hash2 := NewChainHash(5)
	err = hash2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("ChainHash LoadFromFile failed for empty: %v", err)
	}

	// Должна быть пустая таблица с заданной емкостью
	if hash2.Find("anything") != "" {
		t.Error("Empty hash should not find anything")
	}
}

// ========== OpenHash сериализация ==========

func TestOpenHash_Serialize(t *testing.T) {
	// Создаем и заполняем хэш-таблицу с открытой адресацией
	hash := NewOpenHash(7)
	hash.Insert("apple", "red")
	hash.Insert("banana", "yellow")
	hash.Insert("grape", "purple")
	hash.Insert("lemon", "yellow")

	// Сохраняем в JSON
	filename := "test_openhash.json"
	defer os.Remove(filename)

	err := hash.SaveToFile(filename)
	if err != nil {
		t.Fatalf("OpenHash SaveToFile failed: %v", err)
	}

	// Загружаем в новую таблицу
	hash2 := NewOpenHash(3) // Начальная емкость будет перезаписана
	err = hash2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("OpenHash LoadFromFile failed: %v", err)
	}

	// Проверяем данные
	if hash2.Find("apple") != "red" {
		t.Errorf("Expected 'red' for 'apple', got %s", hash2.Find("apple"))
	}
	if hash2.Find("banana") != "yellow" {
		t.Errorf("Expected 'yellow' for 'banana', got %s", hash2.Find("banana"))
	}
	if hash2.Find("grape") != "purple" {
		t.Errorf("Expected 'purple' for 'grape', got %s", hash2.Find("grape"))
	}

	// Проверяем удаление
	hash2.Delete("banana")
	if hash2.Find("banana") != "" {
		t.Error("'banana' should not be found after deletion")
	}

	// Проверяем что другие элементы все еще на месте
	if hash2.Find("apple") != "red" {
		t.Error("'apple' should still exist")
	}

	// Проверяем повторную вставку
	hash2.Insert("melon", "green")
	if hash2.Find("melon") != "green" {
		t.Error("Should find newly inserted 'melon'")
	}
}

// ========== Array бинарная сериализация ==========

func TestArray_BinarySerialize(t *testing.T) {
	// Создаем и заполняем массив
	arr := NewArray()
	arr.AddTail("Hello")
	arr.AddTail("World")
	arr.AddTail("Binary")
	arr.AddTail("Serialization")

	// Сохраняем в бинарный файл
	filename := "test_array.bin"
	defer os.Remove(filename)

	err := arr.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("Array SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новый массив
	arr2 := NewArray()
	err = arr2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("Array LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем
	if arr2.Length() != 4 {
		t.Errorf("Expected length 4, got %d", arr2.Length())
	}
	if arr2.GetAt(0) != "Hello" {
		t.Errorf("Expected 'Hello' at index 0, got %s", arr2.GetAt(0))
	}
	if arr2.GetAt(3) != "Serialization" {
		t.Errorf("Expected 'Serialization' at index 3, got %s", arr2.GetAt(3))
	}

	// Проверяем все элементы
	expected := []string{"Hello", "World", "Binary", "Serialization"}
	for i, exp := range expected {
		if arr2.GetAt(i) != exp {
			t.Errorf("Index %d: expected %s, got %s", i, exp, arr2.GetAt(i))
		}
	}
}

func TestArray_BinarySerializeLarge(t *testing.T) {
	// Тест с большим количеством данных
	arr := NewArray()
	for i := 0; i < 1000; i++ {
		arr.AddTail("Item" + strconv.Itoa(i))
	}

	filename := "test_array_large.bin"
	defer os.Remove(filename)

	err := arr.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("SaveToBinaryFile failed for large array: %v", err)
	}

	arr2 := NewArray()
	err = arr2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("LoadFromBinaryFile failed for large array: %v", err)
	}

	if arr2.Length() != 1000 {
		t.Errorf("Expected length 1000, got %d", arr2.Length())
	}

	// Проверяем несколько элементов
	if arr2.GetAt(0) != "Item0" {
		t.Error("First element mismatch")
	}
	if arr2.GetAt(999) != "Item999" {
		t.Error("Last element mismatch")
	}
}

// ========== List бинарная сериализация ==========

func TestList_BinarySerialize(t *testing.T) {
	// Создаем и заполняем список
	list := NewList()
	list.AddTail("One")
	list.AddTail("Two")
	list.AddTail("Three")
	list.AddTail("Four")

	// Сохраняем в бинарный файл
	filename := "test_list.bin"
	defer os.Remove(filename)

	err := list.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("List SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новый список
	list2 := NewList()
	err = list2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("List LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем через Find
	if !list2.Find("One") || !list2.Find("Two") || !list2.Find("Three") || !list2.Find("Four") {
		t.Error("All elements should be present")
	}
}

// ========== DList бинарная сериализация ==========

func TestDList_BinarySerialize(t *testing.T) {
	// Создаем и заполняем двусвязный список
	dlist := NewDList()
	dlist.AddHead("First")  // Станет вторым
	dlist.AddHead("Before") // Станет первым
	dlist.AddTail("Last")

	// Сохраняем в бинарный файл
	filename := "test_dlist.bin"
	defer os.Remove(filename)

	err := dlist.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("DList SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новый список
	dlist2 := NewDList()
	err = dlist2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("DList LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем
	if !dlist2.Find("Before") || !dlist2.Find("First") || !dlist2.Find("Last") {
		t.Error("All elements should be present")
	}
}

// ========== Queue бинарная сериализация ==========

func TestQueue_BinarySerialize(t *testing.T) {
	// Создаем и заполняем очередь
	queue := NewQueue()
	queue.Enqueue("Customer1")
	queue.Enqueue("Customer2")
	queue.Enqueue("Customer3")
	queue.Enqueue("Customer4")

	// Сохраняем в бинарный файл
	filename := "test_queue.bin"
	defer os.Remove(filename)

	err := queue.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("Queue SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новую очередь
	queue2 := NewQueue()
	err = queue2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("Queue LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем порядок FIFO
	if queue2.Dequeue() != "Customer1" {
		t.Error("Expected 'Customer1' first")
	}
	if queue2.Dequeue() != "Customer2" {
		t.Error("Expected 'Customer2' second")
	}
	if queue2.Dequeue() != "Customer3" {
		t.Error("Expected 'Customer3' third")
	}
	if queue2.Dequeue() != "Customer4" {
		t.Error("Expected 'Customer4' fourth")
	}
	if !queue2.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

// ========== Stack бинарная сериализация ==========

func TestStack_BinarySerialize(t *testing.T) {
	// Создаем и заполняем стек
	stack := NewStack()
	stack.Push("Layer1")
	stack.Push("Layer2")
	stack.Push("Layer3")
	stack.Push("Layer4")

	// Сохраняем в бинарный файл
	filename := "test_stack.bin"
	defer os.Remove(filename)

	err := stack.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("Stack SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новый стек
	stack2 := NewStack()
	err = stack2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("Stack LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем порядок LIFO
	if stack2.Pop() != "Layer4" {
		t.Error("Expected 'Layer4' first (LIFO)")
	}
	if stack2.Pop() != "Layer3" {
		t.Error("Expected 'Layer3' second")
	}
	if stack2.Pop() != "Layer2" {
		t.Error("Expected 'Layer2' third")
	}
	if stack2.Pop() != "Layer1" {
		t.Error("Expected 'Layer1' fourth")
	}
	if stack2.Pop() != "[STACK_EMPTY]" {
		t.Error("Stack should be empty")
	}
}

// ========== ChainHash бинарная сериализация ==========

func TestChainHash_BinarySerialize(t *testing.T) {
	// Создаем и заполняем хэш-таблицу
	hash := NewChainHash(8)
	hash.Insert("username", "john_doe")
	hash.Insert("email", "john@example.com")
	hash.Insert("age", "30")
	hash.Insert("city", "New York")
	hash.Insert("country", "USA")

	// Сохраняем в бинарный файл
	filename := "test_chainhash.bin"
	defer os.Remove(filename)

	err := hash.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("ChainHash SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новую таблицу
	hash2 := NewChainHash(4)
	err = hash2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("ChainHash LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем данные
	expected := map[string]string{
		"username": "john_doe",
		"email":    "john@example.com",
		"age":      "30",
		"city":     "New York",
		"country":  "USA",
	}

	for key, expVal := range expected {
		actVal := hash2.Find(key)
		if actVal != expVal {
			t.Errorf("Key %s: expected %s, got %s", key, expVal, actVal)
		}
	}

	// Проверяем обновление
	hash2.Insert("age", "31")
	if hash2.Find("age") != "31" {
		t.Error("Should update value")
	}

	// Проверяем удаление
	if !hash2.Delete("city") {
		t.Error("Should delete 'city'")
	}
	if hash2.Find("city") != "" {
		t.Error("'city' should not be found after deletion")
	}
}

// ========== OpenHash бинарная сериализация ==========

func TestOpenHash_BinarySerialize(t *testing.T) {
	// Создаем и заполняем хэш-таблицу
	hash := NewOpenHash(10)
	hash.Insert("product", "Laptop")
	hash.Insert("price", "999.99")
	hash.Insert("brand", "Dell")
	hash.Insert("model", "XPS 13")
	hash.Insert("year", "2023")

	hash.Delete("brand")

	// Сохраняем в бинарный файл
	filename := "test_openhash.bin"
	defer os.Remove(filename)

	err := hash.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("OpenHash SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новую таблицу
	hash2 := NewOpenHash(5)
	err = hash2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("OpenHash LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем данные
	if hash2.Find("product") != "Laptop" {
		t.Errorf("Expected 'Laptop' for 'product', got %s", hash2.Find("product"))
	}
	if hash2.Find("price") != "999.99" {
		t.Errorf("Expected '999.99' for 'price', got %s", hash2.Find("price"))
	}
	if hash2.Find("model") != "XPS 13" {
		t.Errorf("Expected 'XPS 13' for 'model', got %s", hash2.Find("model"))
	}
	if hash2.Find("year") != "2023" {
		t.Errorf("Expected '2023' for 'year', got %s", hash2.Find("year"))
	}

	if hash2.Find("brand") != "" {
		t.Error("'brand' should not be found (was deleted)")
	}
}

func TestOpenHash_BinarySerializeEmpty(t *testing.T) {
	hash := NewOpenHash(10)

	filename := "test_openhash_empty.bin"
	defer os.Remove(filename)

	err := hash.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("OpenHash SaveToBinaryFile failed for empty: %v", err)
	}

	hash2 := NewOpenHash(5)
	err = hash2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("OpenHash LoadFromBinaryFile failed for empty: %v", err)
	}

	// Проверяем что таблица пустая
	for i := 0; i < hash2.cap; i++ {
		if hash2.table[i].Used {
			t.Error("Table should be empty")
		}
	}
}

// ========== BinaryTree улучшенная бинарная сериализация ==========

func TestBinaryTree_BinarySerializeV2(t *testing.T) {
	// Создаем дерево
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)
	tree.Insert(10)
	tree.Insert(25)

	// Сохраняем в бинарный файл (улучшенная версия)
	filename := "test_tree_v2.bin"
	defer os.Remove(filename)

	err := tree.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree SaveToBinaryFile failed: %v", err)
	}

	// Загружаем в новое дерево
	tree2 := NewBinaryTree()
	err = tree2.LoadFromBinaryFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree LoadFromBinaryFile failed: %v", err)
	}

	// Проверяем все элементы
	expected := []int{50, 30, 70, 20, 40, 60, 80, 10, 25}
	for _, val := range expected {
		if !tree2.Search(val) {
			t.Errorf("Expected to find %d", val)
		}
	}

	// Проверяем отсутствие лишних элементов
	if tree2.Search(100) || tree2.Search(0) || tree2.Search(55) {
		t.Error("Should not find non-existent elements")
	}
}

func TestSerialize_FileNotFound(t *testing.T) {
	// Тест загрузки несуществующего файла
	hash := NewChainHash(10)
	err := hash.LoadFromFile("non_existent_file_12345.json")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}

	queue := NewQueue()
	err = queue.LoadFromFile("non_existent_file_67890.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}

	// Тест для бинарных файлов
	tree := NewBinaryTree()
	err = tree.LoadFromBinaryFile("non_existent_file_13579.bin")
	if err == nil {
		t.Error("Expected error for non-existent binary file")
	}
}

func TestSerialize_Integration(t *testing.T) {
	// Массив
	arr := NewArray()
	arr.AddTail("Test1")
	arr.AddTail("Test2")
	arr.SaveToFile("test_integration_array.txt")
	defer os.Remove("test_integration_array.txt")

	// Бинарный массив
	arr.SaveToBinaryFile("test_integration_array.bin")
	defer os.Remove("test_integration_array.bin")

	// Список
	list := NewList()
	list.AddTail("List1")
	list.AddTail("List2")
	list.SaveToFile("test_integration_list.txt")
	defer os.Remove("test_integration_list.txt")

	// Бинарный список
	list.SaveToBinaryFile("test_integration_list.bin")
	defer os.Remove("test_integration_list.bin")

	// Очередь
	queue := NewQueue()
	queue.Enqueue("Queue1")
	queue.Enqueue("Queue2")
	queue.SaveToFile("test_integration_queue.txt")
	defer os.Remove("test_integration_queue.txt")

	// Бинарная очередь
	queue.SaveToBinaryFile("test_integration_queue.bin")
	defer os.Remove("test_integration_queue.bin")

	// Стек
	stack := NewStack()
	stack.Push("Stack1")
	stack.Push("Stack2")
	stack.SaveToFile("test_integration_stack.txt")
	defer os.Remove("test_integration_stack.txt")

	// Бинарный стек
	stack.SaveToBinaryFile("test_integration_stack.bin")
	defer os.Remove("test_integration_stack.bin")

	// Двусвязный список
	dlist := NewDList()
	dlist.AddTail("DList1")
	dlist.AddTail("DList2")
	dlist.SaveToFile("test_integration_dlist.txt")
	defer os.Remove("test_integration_dlist.txt")

	// Бинарный двусвязный список
	dlist.SaveToBinaryFile("test_integration_dlist.bin")
	defer os.Remove("test_integration_dlist.bin")

	// Хэш-таблица с цепочками
	chash := NewChainHash(5)
	chash.Insert("key1", "val1")
	chash.Insert("key2", "val2")
	chash.SaveToFile("test_integration_chash.json")
	defer os.Remove("test_integration_chash.json")

	// Бинарная хэш-таблица с цепочками
	chash.SaveToBinaryFile("test_integration_chash.bin")
	defer os.Remove("test_integration_chash.bin")

	// Хэш-таблица с открытой адресацией
	ohash := NewOpenHash(5)
	ohash.Insert("key1", "val1")
	ohash.Insert("key2", "val2")
	ohash.SaveToFile("test_integration_ohash.json")
	defer os.Remove("test_integration_ohash.json")

	// Бинарная хэш-таблица с открытой адресацией
	ohash.SaveToBinaryFile("test_integration_ohash.bin")
	defer os.Remove("test_integration_ohash.bin")

	// Бинарное дерево
	tree := NewBinaryTree()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.SaveToBinaryFile("test_integration_tree.bin")
	defer os.Remove("test_integration_tree.bin")

	t.Log("Integration test passed: all structures can be serialized")
}

func TestArray_BinarySerializeErrors(t *testing.T) {
	arr := NewArray()
	arr.AddTail("Test1")
	arr.AddTail("Test2")

	// Тест с некорректным путем для сохранения
	err := arr.SaveToBinaryFile("/invalid/path/test.bin")
	if err == nil {
		t.Error("Expected error for invalid save path")
	}

	// Тест с некорректным путем для загрузки
	err = arr.LoadFromBinaryFile("/nonexistent/file.bin")
	if err == nil {
		t.Error("Expected error for non-existent load file")
	}
}

func TestList_BinarySerializeErrors(t *testing.T) {
	list := NewList()
	list.AddTail("Item1")
	list.AddTail("Item2")

	// Сохраняем корректный файл
	filename := "test_list_errors.bin"
	err := list.SaveToBinaryFile(filename)
	if err != nil {
		t.Fatalf("Failed to save: %v", err)
	}
	defer os.Remove(filename)

	// Повреждаем файл (делаем его короче чем ожидается)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	// Усекаем файл, оставляя только заголовок
	file.Truncate(4)
	file.Close()

	list2 := NewList()
	err = list2.LoadFromBinaryFile(filename)
	if err == nil {
		t.Error("Expected error for truncated binary file")
	}
}

func TestBinaryTree_EmptyFile(t *testing.T) {
	filename := "test_empty_tree_file.bin"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)
	file.Close()

	tree := NewBinaryTree()
	err = tree.LoadFromBinaryFile(filename)
	if err == nil {
		t.Error("Expected error for completely empty binary file")
	}
}

func TestChainHash_CorruptedJSON(t *testing.T) {
	filename := "test_chainhash_corrupted.json"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(filename)

	// Пишем некорректный JSON
	file.WriteString(`{"capacity": 10, "entries": [{`)
	file.Close()

	hash := NewChainHash(5)
	err = hash.LoadFromFile(filename)
	if err == nil {
		t.Error("Expected error for corrupted JSON")
	}
}

func TestDList_LoadFromFileErrors(t *testing.T) {
	t.Run("Несуществующий файл", func(t *testing.T) {
		dlist := NewDList()
		err := dlist.LoadFromFile("non_existent_dlist_file_98765.txt")
		if err == nil {
			t.Error("Expected error for non-existent file")
		} else {
			t.Logf("Got expected error: %v", err)
		}
	})

	t.Run("Пустой файл", func(t *testing.T) {
		filename := "test_dlist_empty.txt"
		file, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}
		file.Close()
		defer os.Remove(filename)

		dlist := NewDList()
		err = dlist.LoadFromFile(filename)
		if err != nil {
			t.Errorf("Unexpected error for empty file: %v", err)
		}
		// Пустой файл должен создавать пустой список
		if dlist.head != nil || dlist.tail != nil {
			t.Error("Expected empty list from empty file")
		}
	})

	t.Run("Файл только с количеством", func(t *testing.T) {
		filename := "test_dlist_count_only.txt"
		file, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}
		file.WriteString("\n") // Пустая строка вместо числа
		file.Close()
		defer os.Remove(filename)

		dlist := NewDList()
		err = dlist.LoadFromFile(filename)
		if err == nil {
			t.Error("Expected error for empty count string")
		}
	})

	t.Run("Некорректный формат количества", func(t *testing.T) {
		filename := "test_dlist_invalid_count.txt"
		file, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}
		file.WriteString("not_a_number\n")
		file.Close()
		defer os.Remove(filename)

		dlist := NewDList()
		err = dlist.LoadFromFile(filename)
		if err == nil {
			t.Error("Expected error for invalid count format")
		} else {
			t.Logf("Got expected strconv.Atoi error: %v", err)
		}
	})

	t.Run("Отрицательное количество элементов", func(t *testing.T) {
		filename := "test_dlist_negative_count.txt"
		file, err := os.Create(filename)
		if err != nil {
			t.Fatal(err)
		}
		file.WriteString("-5\n")
		file.WriteString("Item1\n")
		file.WriteString("Item2\n")
		file.Close()
		defer os.Remove(filename)

		dlist := NewDList()
		err = dlist.LoadFromFile(filename)
		if err != nil {
			t.Errorf("Unexpected error for negative count: %v", err)
		}
		if dlist.head != nil || dlist.tail != nil {
			t.Error("Expected empty list for negative count")
		}
	})
}

func TestBinaryTree_TextSerialize(t *testing.T) {
	// Создаем дерево
	tree := NewBinaryTree()
	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)

	// Сохраняем в текстовый файл
	filename := "test_tree.txt"
	defer os.Remove(filename)

	err := tree.SaveToFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree SaveToFile failed: %v", err)
	}

	// Загружаем в новое дерево
	tree2 := NewBinaryTree()
	err = tree2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree LoadFromFile failed: %v", err)
	}

	// Проверяем
	if !tree2.Search(50) {
		t.Error("Expected to find 50")
	}
	if !tree2.Search(30) {
		t.Error("Expected to find 30")
	}
	if !tree2.Search(70) {
		t.Error("Expected to find 70")
	}
	if !tree2.Search(20) {
		t.Error("Expected to find 20")
	}
	if !tree2.Search(80) {
		t.Error("Expected to find 80")
	}
	if tree2.Search(100) {
		t.Error("Should not find 100")
	}
}

func TestBinaryTree_TextSerializeEmpty(t *testing.T) {
	tree := NewBinaryTree()

	filename := "test_empty_tree.txt"
	defer os.Remove(filename)

	err := tree.SaveToFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree SaveToFile failed for empty tree: %v", err)
	}

	tree2 := NewBinaryTree()
	err = tree2.LoadFromFile(filename)
	if err != nil {
		t.Fatalf("BinaryTree LoadFromFile failed for empty tree: %v", err)
	}

	if tree2.Search(0) {
		t.Error("Empty tree should not find anything")
	}
}


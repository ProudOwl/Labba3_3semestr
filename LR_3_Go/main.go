package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"go-data-structures/structures"
)

func showHelp() {
	fmt.Println("=== ОБЩИЕ КОМАНДЫ ===")
	fmt.Println("PRINT - Показать содержимое всех структур")
	fmt.Println("HELP - Показать это сообщение")
	fmt.Println("EXIT - Выйти из программы")
	fmt.Println()
	fmt.Println("=== СОХРАНЕНИЕ/ЗАГРУЗКА ===")
	fmt.Println("SAVE <basename> - Сохранить все структуры")
	fmt.Println("LOAD <basename> - Загрузить все структуры")
	fmt.Println()
	fmt.Println("=== МАССИВ (Array) ===")
	fmt.Println("M_PUSH_END <val> - Добавить в конец")
	fmt.Println("M_PUSH_HEAD <val> - Добавить в начало")
	fmt.Println("M_PUSH_AT <idx> <val> - Добавить по индексу")
	fmt.Println("M_DEL_AT <idx> - Удалить по индексу")
	fmt.Println("M_DEL_HEAD - Удалить из начала")
	fmt.Println("M_DEL_TAIL - Удалить из конца")
	fmt.Println("M_GET_AT <idx> - Получить по индексу")
	fmt.Println("M_REPLACE_AT <idx> <val> - Заменить по индексу")
	fmt.Println("M_LENGTH - Узнать длину массива")
	fmt.Println()
	fmt.Println("=== ОДНОСВЯЗНЫЙ СПИСОК (List) ===")
	fmt.Println("F_PUSH_HEAD <val> - Добавить в начало")
	fmt.Println("F_PUSH_TAIL <val> - Добавить в конец")
	fmt.Println("F_DEL_HEAD - Удалить первый элемент")
	fmt.Println("F_DEL_TAIL - Удалить последний элемент")
	fmt.Println("F_DEL_VAL <val> - Удалить по значению")
	fmt.Println("F_GET_VAL <val> - Найти по значению")
	fmt.Println()
	fmt.Println("=== ДВУСВЯЗНЫЙ СПИСОК (DList) ===")
	fmt.Println("L_PUSH_HEAD <val> - Добавить в начало")
	fmt.Println("L_PUSH_TAIL <val> - Добавить в конец")
	fmt.Println("L_DEL_HEAD - Удалить первый элемент")
	fmt.Println("L_DEL_TAIL - Удалить последний элемент")
	fmt.Println("L_DEL_VAL <val> - Удалить по значению")
	fmt.Println("L_GET_VAL <val> - Найти по значению")
	fmt.Println("L_PRINT_REV - Печать в обратном порядке")
	fmt.Println()
	fmt.Println("=== ОЧЕРЕДЬ (Queue) ===")
	fmt.Println("Q_PUSH <val> - Добавить в очередь")
	fmt.Println("Q_POP - Извлечь из очереди")
	fmt.Println("Q_GET - Прочитать первый элемент")
	fmt.Println()
	fmt.Println("=== СТЕК (Stack) ===")
	fmt.Println("S_PUSH <val> - Добавить в стек")
	fmt.Println("S_POP - Извлечь из стека")
	fmt.Println("S_GET - Прочитать верхний элемент")
	fmt.Println()
	fmt.Println("=== БИНАРНОЕ ДЕРЕВО (BinaryTree) ===")
	fmt.Println("CBT_INSERT <num> - Вставить число")
	fmt.Println("CBT_REMOVE <num> - Удалить число")
	fmt.Println("CBT_SEARCH <num> - Найти число")
	fmt.Println("CBT_PRINT - Вывести структуру дерева")
	fmt.Println("CBT_CLEAR - Очистить дерево")
	fmt.Println("CBT_SIZE - Узнать размер")
	fmt.Println()
	fmt.Println("=== ХЕШ-ТАБЛИЦА (Hash) ===")
	fmt.Println("HASH_MAN - Открыть меню хеш-таблиц")
	fmt.Println("-----------------------")
}

func parseCommand(line string) (string, string, string) {
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", "", ""
	}
	if len(parts) == 1 {
		return parts[0], "", ""
	}
	if len(parts) == 2 {
		return parts[0], parts[1], ""
	}
	return parts[0], parts[1], strings.Join(parts[2:], " ")
}

func hashManagement() {
	chainHash := structures.NewChainHash(10)
	openHash := structures.NewOpenHash(10)
	
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("УПРАВЛЕНИЕ ХЕШ-ТАБЛИЦАМИ")
	fmt.Println("Команды: INSERT k v | DELETE k | SEARCH k | SHOW | BACK")
	
	for {
		fmt.Print("Введите команду: ")
		if !scanner.Scan() {
			break
		}
		
		line := scanner.Text()
		cmd, arg1, arg2 := parseCommand(line)
		
		if cmd == "" {
			continue
		}
		if cmd == "BACK" {
			fmt.Println("Выход")
			break
		}
		if cmd == "SHOW" {
			fmt.Print("ChainHash: ")
			chainHash.Print()
			fmt.Print("OpenHash:  ")
			openHash.Print()
			continue
		}
		if cmd == "INSERT" {
			chainHash.Insert(arg1, arg2)
			openHash.Insert(arg1, arg2)
			continue
		}
		if cmd == "DELETE" {
			chainHash.Delete(arg1)
			openHash.Delete(arg1)
			continue
		}
		if cmd == "SEARCH" {
			v1 := chainHash.Find(arg1)
			v2 := openHash.Find(arg1)
			fmt.Printf("Chain: %s\n", v1)
			fmt.Printf("Open:  %s\n", v2)
			continue
		}
	}
}

func main() {
	arr := structures.NewArray()
	list := structures.NewList()
	dlist := structures.NewDList()
	queue := structures.NewQueue()
	stack := structures.NewStack()
	tree := structures.NewBinaryTree()
	
	showHelp()
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("\n>> ")
		if !scanner.Scan() {
			break
		}
		
		line := scanner.Text()
		cmd, arg1, arg2 := parseCommand(line)
		
		if cmd == "" {
			continue
		}
		if cmd == "EXIT" {
			fmt.Println("До свидания!")
			break
		}
		if cmd == "HELP" {
			showHelp()
			continue
		}
		if cmd == "PRINT" {
			fmt.Println("--- Содержимое структур ---")
			fmt.Print("Массив:     ")
			arr.Print()
			fmt.Print("Список:     ")
			list.PrintForward()
			fmt.Print("Дв.Список:  ")
			dlist.PrintForward()
			fmt.Print("Стек:       ")
			stack.Print()
			fmt.Print("Очередь:    ")
			queue.Print()
			fmt.Println("CBT:")
			tree.Print()
			fmt.Println("---------------------------")
			continue
		}
		
		if cmd == "SAVE" {
			if arg1 == "" {
				fmt.Println("Ошибка: нужно имя файла-основы.")
				continue
			}
			arr.SaveToFile(arg1 + ".array.txt")
			list.SaveToFile(arg1 + ".list.txt")
			dlist.SaveToFile(arg1 + ".dlist.txt")
			queue.SaveToFile(arg1 + ".queue.txt")
			stack.SaveToFile(arg1 + ".stack.txt")
			tree.SaveToFile(arg1 + ".tree.txt")
			fmt.Printf("Структуры сохранены в текстовом формате с базовым именем: %s\n", arg1)
			continue
		}
		
		if cmd == "SAVE_BIN" {
			if arg1 == "" {
				fmt.Println("Ошибка: нужно имя файла-основы.")
				continue
			}
			arr.SaveToBinaryFile(arg1 + ".array.bin")
			list.SaveToBinaryFile(arg1 + ".list.bin")
			dlist.SaveToBinaryFile(arg1 + ".dlist.bin")
			queue.SaveToBinaryFile(arg1 + ".queue.bin")
			stack.SaveToBinaryFile(arg1 + ".stack.bin")
			tree.SaveToBinaryFile(arg1 + ".tree.bin")
			fmt.Printf("Структуры сохранены в бинарном формате с базовым именем: %s\n", arg1)
			continue
		}
		
		if cmd == "LOAD" {
			if arg1 == "" {
				fmt.Println("Ошибка: нужно имя файла-основы.")
				continue
			}
			arr.LoadFromFile(arg1 + ".array.txt")
			list.LoadFromFile(arg1 + ".list.txt")
			dlist.LoadFromFile(arg1 + ".dlist.txt")
			queue.LoadFromFile(arg1 + ".queue.txt")
			stack.LoadFromFile(arg1 + ".stack.txt")
			tree.LoadFromFile(arg1 + ".tree.txt")
			fmt.Printf("Структуры загружены из текстового формата с базовым именем: %s\n", arg1)
			continue
		}
		
		if cmd == "LOAD_BIN" {
			if arg1 == "" {
				fmt.Println("Ошибка: нужно имя файла-основы.")
				continue
			}
			arr.LoadFromBinaryFile(arg1 + ".array.bin")
			list.LoadFromBinaryFile(arg1 + ".list.bin")
			dlist.LoadFromBinaryFile(arg1 + ".dlist.bin")
			queue.LoadFromBinaryFile(arg1 + ".queue.bin")
			stack.LoadFromBinaryFile(arg1 + ".stack.bin")
			tree.LoadFromBinaryFile(arg1 + ".tree.bin")
			fmt.Printf("Структуры загружены из бинарного формата с базовым именем: %s\n", arg1)
			continue
		}
		
		switch cmd {
		case "M_PUSH_END":
			if arg1 == "" {
				fmt.Println("Ошибка: команда требует значение")
				continue
			}
			arr.AddTail(arg1)
			arr.Print()
			
		case "M_PUSH_HEAD":
			if arg1 == "" {
				fmt.Println("Ошибка: команда требует значение")
				continue
			}
			arr.AddHead(arg1)
			arr.Print()
			
		case "M_PUSH_AT":
			if arg1 == "" || arg2 == "" {
				fmt.Println("Ошибка: команда требует индекс и значение")
				continue
			}
			idx, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: индекс должен быть числом")
				continue
			}
			arr.AddAt(idx, arg2)
			arr.Print()
			
		case "M_DEL_AT":
			if arg1 == "" {
				fmt.Println("Ошибка: команда требует индекс")
				continue
			}
			idx, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: индекс должен быть числом")
				continue
			}
			arr.DeleteAt(idx)
			arr.Print()
			
		case "M_DEL_HEAD":
			arr.DeleteHead()
			arr.Print()
			
		case "M_DEL_TAIL":
			arr.DeleteTail()
			arr.Print()
			
		case "M_GET_AT":
			if arg1 == "" {
				fmt.Println("Ошибка: команда требует индекс")
				continue
			}
			idx, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: индекс должен быть числом")
				continue
			}
			fmt.Printf("Элемент[%d]: %s\n", idx, arr.GetAt(idx))
			
		case "M_REPLACE_AT":
			if arg1 == "" || arg2 == "" {
				fmt.Println("Ошибка: команда требует индекс и значение")
				continue
			}
			idx, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: индекс должен быть числом")
				continue
			}
			arr.ReplaceAt(idx, arg2)
			arr.Print()
			
		case "M_LENGTH":
			fmt.Printf("Длина массива: %d\n", arr.Length())
			
		case "F_PUSH_HEAD":
			list.AddHead(arg1)
			list.PrintForward()
			
		case "F_PUSH_TAIL":
			list.AddTail(arg1)
			list.PrintForward()
			
		case "F_DEL_HEAD":
			list.DeleteHead()
			list.PrintForward()
			
		case "F_DEL_TAIL":
			list.DeleteTail()
			list.PrintForward()
			
		case "F_DEL_VAL":
			list.DeleteByValue(arg1)
			list.PrintForward()
			
		case "F_GET_VAL":
			found := list.Find(arg1)
			fmt.Printf("Элемент \"%s\" найден: %v\n", arg1, found)
			
		case "L_PUSH_HEAD":
			dlist.AddHead(arg1)
			dlist.PrintForward()
			
		case "L_PUSH_TAIL":
			dlist.AddTail(arg1)
			dlist.PrintForward()
			
		case "L_DEL_HEAD":
			dlist.DeleteHead()
			dlist.PrintForward()
			
		case "L_DEL_TAIL":
			dlist.DeleteTail()
			dlist.PrintForward()
			
		case "L_DEL_VAL":
			dlist.DeleteByValue(arg1)
			dlist.PrintForward()
			
		case "L_GET_VAL":
			found := dlist.Find(arg1)
			fmt.Printf("Элемент \"%s\" найден: %v\n", arg1, found)
			
		case "L_PRINT_REV":
			dlist.PrintBackward()
			
		case "Q_PUSH":
			queue.Enqueue(arg1)
			queue.Print()
			
		case "Q_POP":
			fmt.Printf("Извлечено: %s\n", queue.Dequeue())
			queue.Print()
			
		case "Q_GET":
			fmt.Printf("Первый: %s\n", queue.Peek())
			
		case "S_PUSH":
			stack.Push(arg1)
			stack.Print()
			
		case "S_POP":
			fmt.Printf("Извлечено: %s\n", stack.Pop())
			stack.Print()
			
		case "S_GET":
			fmt.Printf("Верхний: %s\n", stack.Peek())
			
		case "CBT_INSERT":
			val, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: команда требует число")
				continue
			}
			tree.Insert(val)
			
		case "CBT_REMOVE":
			val, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: команда требует число")
				continue
			}
			tree.Remove(val)
			
		case "CBT_SEARCH":
			val, err := strconv.Atoi(arg1)
			if err != nil {
				fmt.Println("Ошибка: команда требует число")
				continue
			}
			found := tree.Search(val)
			fmt.Printf("Элемент %d найден: %v\n", val, found)
			
		case "CBT_PRINT":
			tree.Print()
			
		case "CBT_CLEAR":
			tree.Clear()
			
		case "CBT_SIZE":
			fmt.Printf("Размер дерева: %d\n", tree.Size())
			
		case "HASH_MAN":
			hashManagement()
			
		default:
			fmt.Printf("Неизвестная команда: '%s'. Введите HELP для списка команд.\n", cmd)
		}
	}
}

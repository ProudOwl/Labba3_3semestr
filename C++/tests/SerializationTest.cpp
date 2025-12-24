#include "gtest/gtest.h"
#include "../Array.h"
#include "../SinglyList.h"
#include "../DoublyList.h"
#include "../Stack.h"
#include "../Queue.h"
#include "../CBT.h"
#include "../Hash.h"
#include "../HashSerialize.h"
#include "../ArraySerialize.h"
#include "../SinglyListSerialize.h"
#include "../DoublyListSerialize.h"
#include "../StackSerialize.h"
#include "../QueueSerialize.h"
#include "../CBTSerialize.h"
#include <sstream>
#include <iostream>
#include <cstdio>

using namespace std;

class OutputCapture {
    stringstream buffer;
    streambuf* old;
public:
    OutputCapture() : old(cout.rdbuf(buffer.rdbuf())) {}
    ~OutputCapture() { cout.rdbuf(old); }
    string str() const { return buffer.str(); }
};

void testArray() {
    cout << "=== Тестирование Array (Массив) ===" << endl;
    
    Array arr1;
    cout << "Добавляем элементы: Москва, Лондон, Токио, Париж" << endl;
    arr1.addEnd("Москва");
    arr1.addEnd("Лондон");
    arr1.addEnd("Токио");
    arr1.addEnd("Париж");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в Array.txt (Text)" << endl;
    ArraySerializer::saveToFile(arr1, "Array.txt"); 
    
    cout << "Загрузка из файла..." << endl;
    Array arrText;
    ArraySerializer::loadFromFile(arrText, "Array.txt");
    
    cout << "Проверка данных..." << endl;
    assert(arrText.lenArr() == 4);
    cout << "  - Индекс 0: " << arrText.getAt(0) << " (Ожидалось: Москва)" << endl;
    assert(arrText.getAt(0) == "Москва");
    cout << "  - Индекс 1: " << arrText.getAt(1) << " (Ожидалось: Лондон)" << endl;
    assert(arrText.getAt(1) == "Лондон");
    cout << "  - Индекс 3: " << arrText.getAt(3) << " (Ожидалось: Париж)" << endl;
    assert(arrText.getAt(3) == "Париж");
    cout << "Сериализация (Text): ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сериализация в файл: Array.bin (Binary)" << endl;
    ArraySerializer::saveToBinaryFile(arr1, "Array.bin"); 
    
    cout << "Загрузка из файла..." << endl;
    Array arrBin;
    ArraySerializer::loadFromBinaryFile(arrBin, "Array.bin");
    
    cout << "Проверка данных..." << endl;
    assert(arrBin.lenArr() == 4);
    cout << "  - Индекс 0: " << arrBin.getAt(0) << " (Ожидалось: Москва)" << endl;
    assert(arrBin.getAt(0) == "Москва");
    cout << "  - Индекс 1: " << arrBin.getAt(1) << " (Ожидалось: Лондон)" << endl;
    assert(arrBin.getAt(1) == "Лондон");
    cout << "  - Индекс 3: " << arrBin.getAt(3) << " (Ожидалось: Париж)" << endl;
    assert(arrBin.getAt(3) == "Париж");
    cout << "Сериализация (Binary): ПРОЙДЕНА" << endl;
    
    cout << "Array: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testSinglyList() {
    cout << "=== Тестирование SinglyList (Односвязный список) ===" << endl;
    
    SinglyList list1;
    cout << "Добавляем в хвост: Берлин, Мадрид, Рим" << endl;
    list1.addTail("Берлин");
    list1.addTail("Мадрид");
    list1.addTail("Рим");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в SinglyList.txt (Text)" << endl;
    SinglyListSerializer::saveToFile(list1, "SinglyList.txt");
    
    cout << "Загрузка из файла..." << endl;
    SinglyList listText;
    SinglyListSerializer::loadFromFile(listText, "SinglyList.txt");
    
    cout << "Проверка наличия значений..." << endl;
    assert(listText.findValue("Берлин"));
    cout << "  - Найдено: Берлин" << endl;
    assert(listText.findValue("Мадрид"));
    cout << "  - Найдено: Мадрид" << endl;
    assert(listText.findValue("Рим"));
    cout << "  - Найдено: Рим" << endl;
    cout << "Сериализация (Text): ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сериализация в файл: SinglyList_test.bin (Binary)" << endl;
    SinglyListSerializer::saveToBinaryFile(list1, "SinglyList_test.bin");
    
    cout << "Загрузка из файла..." << endl;
    SinglyList listBin;
    SinglyListSerializer::loadFromBinaryFile(listBin, "SinglyList_test.bin");
    
    cout << "Проверка наличия значений..." << endl;
    assert(listBin.findValue("Берлин"));
    cout << "  - Найдено: Берлин" << endl;
    assert(listBin.findValue("Мадрид"));
    cout << "  - Найдено: Мадрид" << endl;
    assert(listBin.findValue("Рим"));
    cout << "  - Найдено: Рим" << endl;
    cout << "Сериализация (Binary): ПРОЙДЕНА" << endl;
    
    cout << "SinglyList: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testDoublyList() {
    cout << "=== Тестирование DoublyList (Двусвязный список) ===" << endl;
    
    DoublyList list1;
    cout << "Добавляем в хвост: Пекин, Сеул, Бангкок" << endl;
    list1.addTail("Пекин");
    list1.addTail("Сеул");
    list1.addTail("Бангкок");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в DoublyList_test.txt (Text)" << endl;
    DoublyListSerializer::saveToFile(list1, "DoublyList_test.txt");
    
    cout << "Загрузка из файла..." << endl;
    DoublyList listText;
    DoublyListSerializer::loadFromFile(listText, "DoublyList_test.txt");
    
    cout << "Проверка наличия значений..." << endl;
    assert(listText.contains("Пекин"));
    cout << "  - Найдено: Пекин" << endl;
    assert(listText.contains("Сеул"));
    cout << "  - Найдено: Сеул" << endl;
    assert(listText.contains("Бангкок"));
    cout << "  - Найдено: Бангкок" << endl;
    cout << "Сериализация (Text): ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сериализация в файл: DoublyList_test.bin (Binary)" << endl;
    DoublyListSerializer::saveToBinaryFile(list1, "DoublyList_test.bin");
    
    cout << "Загрузка из файла..." << endl;
    DoublyList listBin;
    DoublyListSerializer::loadFromBinaryFile(listBin, "DoublyList_test.bin");
    
    cout << "Проверка наличия значений..." << endl;
    assert(listBin.contains("Пекин"));
    cout << "  - Найдено: Пекин" << endl;
    assert(listBin.contains("Сеул"));
    cout << "  - Найдено: Сеул" << endl;
    assert(listBin.contains("Бангкок"));
    cout << "  - Найдено: Бангкок" << endl;
    cout << "Сериализация (Binary): ПРОЙДЕНА" << endl;
    
    cout << "DoublyList: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testStack() {
    cout << "=== Тестирование Stack (Стек) ===" << endl;
    
    Stack stack1;
    cout << "Push: Осло -> Стокгольм -> Копенгаген (верх)" << endl;
    stack1.push("Осло");
    stack1.push("Стокгольм");
    stack1.push("Копенгаген");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в stack_test.txt (Text)" << endl;
    StackSerializer::saveToFile(stack1, "stack_test.txt");
    
    cout << "Загрузка из файла..." << endl;
    Stack stackText;
    StackSerializer::loadFromFile(stackText, "stack_test.txt");
    
    cout << "Проверка порядка (LIFO)..." << endl;
    string valT1 = stackText.pop();
    cout << "  - Pop 1: " << valT1 << " (Ожидалось: Копенгаген)" << endl;
    assert(valT1 == "Копенгаген");
    string valT2 = stackText.pop();
    cout << "  - Pop 2: " << valT2 << " (Ожидалось: Стокгольм)" << endl;
    assert(valT2 == "Стокгольм");
    string valT3 = stackText.pop();
    cout << "  - Pop 3: " << valT3 << " (Ожидалось: Осло)" << endl;
    assert(valT3 == "Осло");
    cout << "Сериализация (Text): ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сериализация в файл: stack_test.bin (Binary)" << endl;
    StackSerializer::saveToBinaryFile(stack1, "stack_test.bin");
    
    cout << "Загрузка из файла..." << endl;
    Stack stackBin;
    StackSerializer::loadFromBinaryFile(stackBin, "stack_test.bin");
    
    cout << "Проверка порядка (LIFO)..." << endl;
    string valB1 = stackBin.pop();
    cout << "  - Pop 1: " << valB1 << " (Ожидалось: Копенгаген)" << endl;
    assert(valB1 == "Копенгаген");
    string valB2 = stackBin.pop();
    cout << "  - Pop 2: " << valB2 << " (Ожидалось: Стокгольм)" << endl;
    assert(valB2 == "Стокгольм");
    string valB3 = stackBin.pop();
    cout << "  - Pop 3: " << valB3 << " (Ожидалось: Осло)" << endl;
    assert(valB3 == "Осло");
    cout << "Сериализация (Binary): ПРОЙДЕНА" << endl;
    
    cout << "Stack: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testQueue() {
    cout << "=== Тестирование Queue (Очередь) ===" << endl;
    
    Queue queue1;
    cout << "Push: Дели -> Мумбаи -> Калькутта" << endl;
    queue1.push("Дели");
    queue1.push("Мумбаи");
    queue1.push("Калькутта");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в queue_test.txt (Text)" << endl;
    QueueSerializer::saveToFile(queue1, "queue_test.txt");
    
    cout << "Загрузка из файла..." << endl;
    Queue queueText;
    QueueSerializer::loadFromFile(queueText, "queue_test.txt");
    
    cout << "Проверка порядка (FIFO)..." << endl;
    string valT1 = queueText.pop();
    cout << "  - Pop 1: " << valT1 << " (Ожидалось: Дели)" << endl;
    assert(valT1 == "Дели");
    string valT2 = queueText.pop();
    cout << "  - Pop 2: " << valT2 << " (Ожидалось: Мумбаи)" << endl;
    assert(valT2 == "Мумбаи");
    string valT3 = queueText.pop();
    cout << "  - Pop 3: " << valT3 << " (Ожидалось: Калькутта)" << endl;
    assert(valT3 == "Калькутта");
    cout << "Сериализация (Text): ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сериализация в файл: queue_test.bin (Binary)" << endl;
    QueueSerializer::saveToBinaryFile(queue1, "queue_test.bin");
    
    cout << "Загрузка из файла..." << endl;
    Queue queueBin;
    QueueSerializer::loadFromBinaryFile(queueBin, "queue_test.bin");
    
    cout << "Проверка порядка (FIFO)..." << endl;
    string valB1 = queueBin.pop();
    cout << "  - Pop 1: " << valB1 << " (Ожидалось: Дели)" << endl;
    assert(valB1 == "Дели");
    string valB2 = queueBin.pop();
    cout << "  - Pop 2: " << valB2 << " (Ожидалось: Мумбаи)" << endl;
    assert(valB2 == "Мумбаи");
    string valB3 = queueBin.pop();
    cout << "  - Pop 3: " << valB3 << " (Ожидалось: Калькутта)" << endl;
    assert(valB3 == "Калькутта");
    cout << "Сериализация (Binary): ПРОЙДЕНА" << endl;
    
    cout << "Queue: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testCBT() {
    cout << "=== Тестирование CompleteBinaryTree (CBT) ===" << endl;
    
    CompleteBinaryTree tree1;
    cout << "Вставка чисел: 50, 30, 70" << endl;
    tree1.insert(50);
    tree1.insert(30);
    tree1.insert(70);
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в cbt.txt (Text)" << endl;
    CBTSerializer::saveToFile(tree1, "cbt.txt"); 
    
    cout << "Загрузка из cbt.txt..." << endl;
    CompleteBinaryTree treeText;
    CBTSerializer::loadFromFile(treeText, "cbt.txt");
    
    cout << " Поиск узлов..." << endl;
    assert(treeText.search(50) == true);
    cout << "  - 50 найдено" << endl;
    assert(treeText.search(30) == true);
    cout << "  - 30 найдено" << endl;
    assert(treeText.search(70) == true);
    cout << "  - 70 найдено" << endl;
    cout << "Текстовая сериализация: ПРОЙДЕНА" << endl;

    cout << "\n Тест бинарного формата ---" << endl;
    cout << "Сохранение в cbt.bin (Binary)" << endl;
    CBTSerializer::saveToBinaryFile(tree1, "cbt.bin"); 
    
    cout << "Загрузка из cbt.bin..." << endl;
    CompleteBinaryTree treeBin;
    CBTSerializer::loadFromBinaryFile(treeBin, "cbt.bin");
    
    cout << "Поиск узлов..." << endl;
    assert(treeBin.search(50) == true);
    cout << "  - 50 найдено" << endl;
    assert(treeBin.search(30) == true);
    cout << "  - 30 найдено" << endl;
    assert(treeBin.search(70) == true);
    cout << "  - 70 найдено" << endl;
    cout << "Бинарная сериализация: ПРОЙДЕНА" << endl;
    
    cout << "CompleteBinaryTree: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

void testHashTable() {
    cout << "=== Тестирование HashTables (Хеш-таблицы) ===" << endl;
    
    cout << "Hash" << endl;
    Hash ch1(10);
    cout << "Вставка: key1->value1, key2->value2" << endl;
    ch1.insert("key1", "value1");
    ch1.insert("key2", "value2");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в chain.txt" << endl;
    HashSerializer::saveToFile(ch1, "chain.txt");
    
    Hash chText(10);
    HashSerializer::loadFromFile(chText, "chain.txt");
    
    cout << " Проверка значений..." << endl;
    assert(chText.find("key1") == "value1");
    cout << "  - key1 -> " << chText.find("key1") << endl;
    assert(chText.find("key2") == "value2");
    cout << "  - key2 -> " << chText.find("key2") << endl;
    cout << "Hash Текстовая Сериализация: ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сохранение в chain.bin" << endl;
    HashSerializer::saveToBinaryFile(ch1, "chain.bin");
    
    Hash chBin(10);
    HashSerializer::loadFromBinaryFile(chBin, "chain.bin");
    
    cout << " Проверка значений..." << endl;
    assert(chBin.find("key1") == "value1");
    cout << "  - key1 -> " << chBin.find("key1") << endl;
    assert(chBin.find("key2") == "value2");
    cout << "  - key2 -> " << chBin.find("key2") << endl;
    cout << "Hash Бинарная Сериализация: ПРОЙДЕНА" << endl;

    cout << "\n[ OpenHash ---" << endl;
    OpenHash oh1(10);
    cout << "Вставка: k1->v1, k2->v2" << endl;
    oh1.insert("k1", "v1");
    oh1.insert("k2", "v2");
    
    cout << "\nТест текстового формата ---" << endl;
    cout << "Сохранение в openhash_test.txt" << endl;
    HashSerializer::saveToFile(oh1, "openhash_test.txt");
    
    OpenHash ohText(10);
    HashSerializer::loadFromFile(ohText, "openhash_test.txt");
    
    cout << "Проверка значений..." << endl;
    assert(ohText.find("k1") == "v1");
    cout << "  - k1 -> " << ohText.find("k1") << endl;
    assert(ohText.find("k2") == "v2");
    cout << "  - k2 -> " << ohText.find("k2") << endl;
    cout << "OpenHash Текстовая Сериализация: ПРОЙДЕНА" << endl;

    cout << "\nТест бинарного формата ---" << endl;
    cout << "Сохранение в openhash_test.bin" << endl;
    HashSerializer::saveToBinaryFile(oh1, "openhash_test.bin");
    
    OpenHash ohBin(10);
    HashSerializer::loadFromBinaryFile(ohBin, "openhash_test.bin");
    
    cout << "Проверка значений..." << endl;
    assert(ohBin.find("k1") == "v1");
    cout << "  - k1 -> " << ohBin.find("k1") << endl;
    assert(ohBin.find("k2") == "v2");
    cout << "  - k2 -> " << ohBin.find("k2") << endl;
    cout << "OpenHash Бинарная Сериализация: ПРОЙДЕНА" << endl;
    
    cout << "HashTables: ВСЕ ТЕСТЫ ПРОЙДЕНЫ\n" << endl;
}

int main() {
    setlocale(LC_ALL, "");

    cout << "ЗАПУСК ТЕСТОВ СЕРИАЛИЗАЦИИ СТРУКТУР ДАННЫХ\n" << endl;
    
    try {
        testArray();
        testSinglyList();
        testDoublyList();
        testStack();
        testQueue();
        testCBT();
        testHashTable();
    } catch (const exception& e) {
        cout << "ОШИБКА: " << e.what() << endl;
        return 1;
    }
    
    return 0;
}

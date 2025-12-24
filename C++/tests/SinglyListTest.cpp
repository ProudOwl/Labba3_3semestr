#include "gtest/gtest.h"
#include "../SinglyList.h"
#include "../SinglyListSerialize.h"
#include "../Serialize.h"
#include <string>
#include <sstream>
#include <fstream>
#include <cstdio>

using namespace std;

class OutputCapture {
    stringstream buffer;
    streambuf* prev;
public:
    OutputCapture() : prev(cout.rdbuf(buffer.rdbuf())) {}
    ~OutputCapture() { cout.rdbuf(prev); }
    string str() { return buffer.str(); }
};

// проверка на различные сценарии удаления элементов
TEST(SinglyListTest, Branch_Deletions) {
    SinglyList list;
    OutputCapture cap;

    list.delTail();
    list.addHead("A");
    list.delTail();
    EXPECT_EQ(list.getHead_test(), nullptr);
    
    list.addTail("A");
    list.addTail("B");
    list.delTail();
    EXPECT_EQ(list.getHead_test()->value, "A");

    list.delByValue("Z");
    list.delByValue("A");
    EXPECT_EQ(list.getHead_test(), nullptr);

    list.addTail("X"); list.addTail("Y");
    list.delByValue("Y");
    EXPECT_FALSE(list.findValue("Y"));
}

// проверка на сложные операций удаления (до и после значений)
TEST(SinglyListTest, Branch_AdvancedOps) {
    SinglyList list;
    OutputCapture cap;
    
    list.delAfterValue("A"); 
    list.addTail("A");
    list.delAfterValue("A"); 
    list.addTail("B");
    list.delAfterValue("Z"); 
    list.delAfterValue("A"); 
    EXPECT_FALSE(list.findValue("B"));

    list.delBeforeValue("A");
    list.addHead("A");
    list.delBeforeValue("A"); 
    
    SinglyList l2;
    l2.addTail("1"); l2.addTail("2"); l2.addTail("3");
    l2.delBeforeValue("2"); 
    EXPECT_EQ(l2.getHead_test()->value, "2");

    SinglyList l3;
    l3.addTail("X"); l3.addTail("Y"); l3.addTail("Z"); l3.addTail("W");
    l3.delBeforeValue("W"); 
    EXPECT_FALSE(l3.findValue("Z"));
    
    l3.readBack();
}

// проверка на вставку относительно значений
TEST(SinglyListTest, AddBeforeHeadValue) {
    OutputCapture cap;
    SinglyList l;
    l.addBefore("A", "B");
    l.addHead("A");
    l.addBefore("A", "Z"); 
    EXPECT_EQ(l.getHead_test()->value, "Z");
}

// проверка на вставку после несуществующего значения
TEST(SinglyListTest, AddAfterNotFound) {
    OutputCapture cap;
    SinglyList l;
    l.addTail("A");
    l.addAfter("Z", "B"); 
}

// проверка на удаление по значению из пустого списка
TEST(SinglyListTest, DelByValueEmpty) {
    OutputCapture cap;
    SinglyList l;
    l.delByValue("A"); 
}

// проверка на многократное удаление хвоста
TEST(SinglyListTest, DelTailMultiple) {
    OutputCapture cap;
    SinglyList l;
    l.addTail("A"); l.addTail("B"); l.addTail("C");
    l.delTail(); 
    EXPECT_TRUE(l.findValue("B"));
    EXPECT_FALSE(l.findValue("C"));
}

// проверка на комбинации сложных условий удаления
TEST(ListTest, BranchBooster_ComplexConditions) {
    OutputCapture cap;
    SinglyList l;

    l.delAfterValue("A"); 
    l.addHead("A");
    l.delAfterValue("A"); 

    l.addTail("B");
    l.delAfterValue("Z");
    l.delAfterValue("B");
    
    SinglyList l2;
    l2.delBeforeValue("A"); 
    l2.addHead("A");
    l2.delBeforeValue("A");

    l2.addTail("B");
    l2.delBeforeValue("B"); 
    EXPECT_EQ(l2.getHead_test()->value, "B");

    SinglyList l3;
    l3.addTail("A"); l3.addTail("B"); l3.addTail("C");
    l3.delBeforeValue("Z");

    SinglyList l4;
    l4.addTail("A"); l4.addTail("B");
    l4.delBeforeValue("Z");
}

// проверка на корректность работы деструктора
TEST(DestructorTest, SinglyListCleanup) {
    OutputCapture cap;
    {
        SinglyList l;
        l.addTail("1");
        l.addTail("2");
    }
}

// проверка на вывод пустого списка
TEST(SinglyListTest, Coverage_EmptyPrints) {
    SinglyList l;
    OutputCapture cap;
    
    l.readForward();
    l.readBack();
    
    string out = cap.str();
    EXPECT_TRUE(out.find("пуст") != string::npos);
}

// проверка на удаление головы из пустого списка
TEST(SinglyListTest, Coverage_DelHeadEmpty) {
    SinglyList l;
    l.delHead(); 
    EXPECT_EQ(l.getHead_test(), nullptr);
}

// проверка на ошибки сериализации
TEST(SinglyListTest, Coverage_Serialization_Fail) {
    SinglyList l;
    
    SinglyListSerializer::saveToFile(l, ""); 

    SinglyListSerializer::loadFromFile(l, "non_existent_file_123.dat");
}

// проверка на загрузку поврежденных файлов
TEST(SinglyListTest, Coverage_Load_Corrupted) {
    SinglyList l;
    
    {
        ofstream f("corrupted_list.dat", ios::binary);
        int count = 5;
        f.write((char*)&count, sizeof(count));
        f.close();
    }
    
    SinglyListSerializer::loadFromFile(l, "corrupted_list.dat");
    
    remove("corrupted_list.dat");
}

// проверка на все виды удаления из пустого списка
TEST(SinglyListTest, Coverage_Delete_From_Empty) {
    SinglyList list;
    OutputCapture cap;
    
    // Пытаемся удалить из пустого
    list.delHead();
    list.delTail();
    list.delByValue("A");
    list.delAfterValue("A");
    list.delBeforeValue("A");
    
    string out = cap.str();
    EXPECT_NE(out.find("пуст"), string::npos); // Проверяем вывод ошибки
}

// проверка на граничные случаи добавления
TEST(SinglyListTest, Coverage_Add_EdgeCases) {
    SinglyList list;
    OutputCapture cap;
    
    // addBefore в пустой список
    list.addBefore("Target", "Val"); 
    EXPECT_EQ(list.getHead_test(), nullptr);

    list.addAfter("Target", "Val");
    EXPECT_NE(cap.str().find("пуст"), string::npos);
    
    list.addTail("A");
    list.addBefore("A", "NewHead");
    EXPECT_EQ(list.getHead_test()->value, "NewHead");
    
    list.addAfter("NonExistent", "X");
    
    list.addBefore("NonExistent", "X");
}

// проверка на удаление из списка с одним элементом
TEST(SinglyListTest, Coverage_Delete_Head_Tail_SingleElement) {
    SinglyList list;
    list.addHead("Single");
    

    list.delTail();
    EXPECT_EQ(list.getHead_test(), nullptr);
    
    list.addHead("Single2");
    list.delHead();
    EXPECT_EQ(list.getHead_test(), nullptr);
}

// проверка на различные сценарии удаления по значению
TEST(SinglyListTest, Coverage_DelByValue_Cases) {
    SinglyList list;
    OutputCapture cap;
    
    list.addTail("A");
    list.addTail("B");
    list.addTail("C");
    
    // Удаляем голову по значению
    list.delByValue("A");
    EXPECT_FALSE(list.findValue("A"));
    EXPECT_EQ(list.getHead_test()->value, "B");
    
    // Удаляем несуществующий
    list.delByValue("Z");
    EXPECT_NE(cap.str().find("не найден"), string::npos);
    
    // Удаляем хвост по значению
    list.delByValue("C");
    EXPECT_FALSE(list.findValue("C"));
}

// проверка на удаление после/до в разных позициях
TEST(SinglyListTest, Coverage_DelAfter_DelBefore) {
    SinglyList list;
    
    list.addTail("A");
    list.delAfterValue("A"); 
    
    list.addTail("B");
    list.delAfterValue("A");
    EXPECT_FALSE(list.findValue("B"));
    
    list.delHead();
    list.addTail("1");
    list.addTail("2");
    list.addTail("3");
    
    list.delBeforeValue("3");
    EXPECT_FALSE(list.findValue("2"));
    
    list.delBeforeValue("1");
    EXPECT_TRUE(list.findValue("1"));
    
    list.delBeforeValue("3"); 
    EXPECT_FALSE(list.findValue("1"));
}

// проверка на различные ошибки ввода-вывода
TEST(SinglyListTest, Coverage_IO_Failures) {
    SinglyList list;
    // Запись в недоступный файл
    SinglyListSerializer::saveToFile(list, "");
    
    // Чтение из несуществующего
    SinglyListSerializer::loadFromFile(list, "does_not_exist.txt");
    EXPECT_EQ(list.getHead_test(), nullptr);
    
    // Чтение битого файла
    {
        ofstream f("broken_list.txt");
        f << "NotANumber";
        f.close();
    }
    SinglyListSerializer::loadFromFile(list, "broken_list.txt");
    remove("broken_list.txt");
}

// проверка на загрузку с перезаписью существующих данных
TEST(SinglyListTest, Coverage_Load_Overwrite) {
    SinglyList l;
    l.addTail("Old1");
    l.addTail("Old2");
    
    remove("good_list.dat");

    {
        ofstream f("good_list.dat"); 
        if (f.is_open()) {
            f << "1\n";      
            f << "New\n";  
            f.close();
        }
    }
    
    SinglyListSerializer::loadFromFile(l, "good_list.dat");
    
   
    ASSERT_NE(l.getHead_test(), nullptr);
    EXPECT_EQ(l.getHead_test()->value, "New");
    EXPECT_EQ(l.getHead_test()->next, nullptr);
    
    remove("good_list.dat");
}

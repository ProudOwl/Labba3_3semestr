#include "gtest/gtest.h"
#include "../Array.h"
#include "../ArraySerialize.h"
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

// проверка на управление памятью при расширении и сжатии массива
TEST(ArrayTest, Branch_CapacityAndShrink) {
    Array arr;

    for(int i=0; i<65; ++i) {
        arr.addEnd("x");
    }
    EXPECT_GE(arr.lenArr(), 65);
    
    int startSize = arr.lenArr();
    for(int i=0; i < startSize; ++i) {
        arr.delEnd(); 
    }
    EXPECT_EQ(arr.lenArr(), 0);
    
    for(int i=0; i<20; ++i) arr.addEnd("y");
    for(int i=0; i<20; ++i) arr.delHead();
    EXPECT_EQ(arr.lenArr(), 0);

    for(int i=0; i<20; ++i) arr.addEnd("z");
    for(int i=0; i<20; ++i) arr.delAt(0);
    EXPECT_EQ(arr.lenArr(), 0);
}

// проверка на граничные случаи и обработку ошибок индексов
TEST(ArrayTest, Branch_EdgeCases) {
    Array arr;
    OutputCapture cap; 

    // Операции с пустым массивом
    arr.readArray();
    arr.delHead();
    arr.delEnd();
    arr.delAt(0);
    arr.repArr(0, "val");
    
    // Неверные индексы 
    arr.addEnd("A");
    arr.addAt(-1, "err");
    arr.addAt(5, "err");
    arr.delAt(-1);
    arr.delAt(5);
    arr.repArr(-1, "err");
    arr.repArr(5, "err");
    
    EXPECT_EQ(arr.getAt(-1), "[INVALID_INDEX]");
    EXPECT_EQ(arr.getAt(5), "[INVALID_INDEX]");

    arr.addHead("Head");
    EXPECT_EQ(arr.getAt(0), "Head");
    
    // Вставка в середину
    arr.addAt(1, "Mid"); 
    EXPECT_EQ(arr.getAt(1), "Mid");

    EXPECT_NE(arr.getData_test(), nullptr);
}

// проверка на обработку ошибок ввода-вывода
TEST(ArrayTest, IO_Errors_Coverage) {
    Array arr;
    
    ArraySerializer::loadFromFile(arr, "no_such_file.arr");
    ArraySerializer::loadFromBinaryFile(arr, "no_such_file.bin");

    {
        ofstream f("corrupt.arr");
        f << "NotANumber"; 
    }
    ArraySerializer::loadFromFile(arr, "corrupt.arr");
    remove("corrupt.arr");

    {
        ofstream f("corrupt.bin", ios::binary);
        int size = 100; 
        f.write((char*)&size, sizeof(size));
        // данные не пишем
    }
    ArraySerializer::loadFromBinaryFile(arr, "corrupt.bin");
    remove("corrupt.bin");
    

    ArraySerializer::saveToFile(arr, ""); 
    ArraySerializer::saveToBinaryFile(arr, "");
}

// проверка на корректную работу сериализации и десериализации
TEST(ArrayTest, SaveAndLoad_HappyPath) {
    Array arr;
    arr.addEnd("1");
    arr.addEnd("2");
    
    ArraySerializer::saveToFile(arr, "good.arr");
    Array arr2;
    ArraySerializer::loadFromFile(arr2, "good.arr");
    EXPECT_EQ(arr2.lenArr(), 2);
    EXPECT_EQ(arr2.getAt(0), "1");
    remove("good.arr");

    ArraySerializer::saveToBinaryFile(arr, "good.bin");
    Array arr3;
    ArraySerializer::loadFromBinaryFile(arr3, "good.bin");
    EXPECT_EQ(arr3.lenArr(), 2);
    remove("good.bin");
}

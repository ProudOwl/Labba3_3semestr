#ifndef ARRAY_H
#define ARRAY_H

#include <string>
#include <iostream>
#include <fstream>

using namespace std;

class ArraySerializer;

class Array {
public:
    string* data;
    int size;
    int capacity;

    void ensureCapacity(int newSize);

    Array();
    ~Array();

    Array(const Array&) = delete;
    Array& operator=(const Array&) = delete;
    Array(Array&&) = delete;
    Array& operator=(Array&&) = delete;

    void addHead(const string& val);
    void delHead();
    void delEnd();

    void addEnd(const string& val);
    string getAt(int idx) const;
    void delAt(int idx);
    void readArray() const;
    void addAt(int idx, const string& val);
    void repArr(int idx, const string& val);

    auto lenArr() const -> int;

    string* getData_test() const { return data; }

private:
    friend class ArrSerializer;
};

#endif

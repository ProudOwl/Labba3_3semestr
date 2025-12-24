#include "ArraySerialize.h"
#include "Serialize.h"
#include <iostream>

using namespace std;

void ArraySerializer::saveToFile(const Array& arr, const string& filename) {
    ofstream file(filename);
    if (!file) return;

    file << arr.size << "\n";
    for (int i = 0; i < arr.size; ++i) {
        writeStringText(file, arr.data[i]);
    }
}

void ArraySerializer::loadFromFile(Array& arr, const string& filename) {
    ifstream file(filename);
    if (!file) return;

    delete[] arr.data;
    arr.data = nullptr;
    arr.size = 0;
    arr.capacity = 0;

    int newSize;
    file >> newSize;
    string dummy; getline(file, dummy);

    if (file.fail()) return;

    arr.ensureCapacity(newSize);
    arr.size = newSize;

    for (int i = 0; i < arr.size; ++i) {
        arr.data[i] = readStringText(file);
    }
}

void ArraySerializer::saveToBinaryFile(const Array& arr, const string& filename) {
    ofstream file(filename, ios::binary | ios::trunc);
    if (!file) return;

    file.write(reinterpret_cast<const char*>(&arr.size), sizeof(arr.size));
    for (int i = 0; i < arr.size; ++i) {
        writeString(file, arr.data[i]);
    }
}

void ArraySerializer::loadFromBinaryFile(Array& arr, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) return;

    delete[] arr.data;
    arr.data = nullptr;
    arr.size = 0;
    arr.capacity = 0;

    int newSize;
    file.read(reinterpret_cast<char*>(&newSize), sizeof(newSize));
    if (file.fail()) return;

    arr.ensureCapacity(newSize);
    arr.size = newSize;

    for (int i = 0; i < arr.size; ++i) {
        arr.data[i] = readString(file);
        if (file.fail()) {
            arr.size = i;
            break;
        }
    }
}

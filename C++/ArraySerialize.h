#ifndef ARRAYSERIALIZE_H
#define ARRAYSERIALIZE_H

#include "Array.h"

class ArraySerializer {
public:
    static void saveToFile(const Array& arr, const std::string& filename);
    static void loadFromFile(Array& arr, const std::string& filename);
    static void saveToBinaryFile(const Array& arr, const std::string& filename);
    static void loadFromBinaryFile(Array& arr, const std::string& filename);
};

#endif

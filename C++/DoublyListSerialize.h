#ifndef DOUBLYLIST_SERIALIZE_H
#define DOUBLYLIST_SERIALIZE_H

#include "DoublyList.h"

class DoublyListSerializer {
public:
    static void saveToFile(const DoublyList& list, const std::string& filename);
    static void loadFromFile(DoublyList& list, const std::string& filename);
    static void saveToBinaryFile(const DoublyList& list, const std::string& filename);
    static void loadFromBinaryFile(DoublyList& list, const std::string& filename);
};

#endif

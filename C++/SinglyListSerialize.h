#ifndef LIST_SERIALIZE_H
#define LIST_SERIALIZE_H

#include "SinglyList.h"

class SinglyListSerializer {
public:
    static void saveToFile(const SinglyList& list, const std::string& filename);
    static void loadFromFile(SinglyList& list, const std::string& filename);
    static void saveToBinaryFile(const SinglyList& list, const std::string& filename);
    static void loadFromBinaryFile(SinglyList& list, const std::string& filename);
};

#endif

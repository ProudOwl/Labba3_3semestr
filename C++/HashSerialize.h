#ifndef HASHSERIALIZER_H
#define HASHSERIALIZER_H

#include "Hash.h"

class HashSerializer {
public:
    static void saveToFile(const Hash& table, const std::string& filename);
    static void loadFromFile(Hash& table, const std::string& filename);
    static void saveToBinaryFile(const Hash& table, const std::string& filename);
    static void loadFromBinaryFile(Hash& table, const std::string& filename);

    static void saveToFile(const OpenHash& table, const std::string& filename);
    static void loadFromFile(OpenHash& table, const std::string& filename);
    static void saveToBinaryFile(const OpenHash& table, const std::string& filename);
    static void loadFromBinaryFile(OpenHash& table, const std::string& filename);
};

#endif

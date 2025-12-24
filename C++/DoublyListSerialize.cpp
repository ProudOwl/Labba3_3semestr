#include "DoublyListSerialize.h"
#include "Serialize.h"
#include <iostream>

using namespace std;

void DoublyListSerializer::saveToFile(const DoublyList& list, const string& filename) {
    ofstream file(filename);
    if (!file) return;

    int count = 0;
    for (DoublyList::DNode* curr = list.head; curr != nullptr; curr = curr->next) {
        count++;
    }
    file << count << "\n";

    for (DoublyList::DNode* curr = list.head; curr != nullptr; curr = curr->next) {
        writeStringText(file, curr->value);
    }
}

void DoublyListSerializer::loadFromFile(DoublyList& list, const string& filename) {
    ifstream file(filename);
    if (!file) return;

    while (list.head != nullptr) {
        list.delHead();
    }

    int count;
    file >> count;
    string dummy; getline(file, dummy);

    if (file.fail()) return;

    for (int i = 0; i < count; ++i) {
        string val = readStringText(file);
        list.addTail(val);
    }
}

void DoublyListSerializer::saveToBinaryFile(const DoublyList& list, const string& filename) {
    ofstream file(filename, ios::binary | ios::trunc);
    if (!file) return;

    int count = 0;
    for (DoublyList::DNode* curr = list.head; curr != nullptr; curr = curr->next) {
        count++;
    }
    file.write(reinterpret_cast<const char*>(&count), sizeof(count));

    for (DoublyList::DNode* curr = list.head; curr != nullptr; curr = curr->next) {
        writeString(file, curr->value);
    }
}

void DoublyListSerializer::loadFromBinaryFile(DoublyList& list, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) return;

    while (list.head != nullptr) {
        list.delHead();
    }

    int count;
    file.read(reinterpret_cast<char*>(&count), sizeof(count));
    if (file.fail()) return;

    for (int i = 0; i < count; ++i) {
        string val = readString(file);
        if (file.fail()) break;
        list.addTail(val);
    }
}

#ifndef DOUBLYLIST_H
#define DOUBLYLIST_H

#include <iostream>
#include <string>
#include <fstream>
using namespace std;

class DoublyListSerializer;

class DoublyList {
private:
    class DNode {
    public:
        string value;
        DNode* next;
        DNode* prev;

        DNode(const string& v, DNode* n, DNode* p) : value(v), next(n), prev(p) {}
    };

    DNode* head;
    DNode* tail;

    auto findValue(const string& val) const -> DNode*;

public:
    DoublyList();
    ~DoublyList();

    DoublyList(const DoublyList&) = delete;
    DoublyList& operator=(const DoublyList&) = delete;
    DoublyList(DoublyList&&) = delete;
    DoublyList& operator=(DoublyList&&) = delete;

    void addHead(const string& val);
    void addTail(const string& val);
    void addBefore(const string& target, const string& val);
    void addAfter(const string& target, const string& val);
    void delHead();
    void delTail();
    void delByVal(const string& val);
    auto contains(const string& val) const -> bool;
    void readForward() const;
    void readBackward() const;
    void delAfterValue(const string& val);
    void delBeforeValue(const string& val);

    DNode* getHead_Test() const { return head; }
    DNode* getTail_Test() const { return tail; }

private:
    friend class DoublyListSerializer;
};

#endif

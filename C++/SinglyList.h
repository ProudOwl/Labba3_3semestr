#ifndef SINGLYLIST_H
#define SINGLYLIST_H

#include <string>
#include <iostream>
#include <fstream>

using namespace std;

class SinglyListSerializer;

class SinglyList {
public:
    class Node {
    public:
        string value;
        Node* next;

        Node(const string& v, Node* n) : value(v), next(n) {}
    };

    Node* head;

    void printBack(Node* node) const;

    SinglyList();
    ~SinglyList();

    SinglyList(const SinglyList&) = delete;
    SinglyList& operator=(const SinglyList&) = delete;
    SinglyList(SinglyList&&) = delete;
    SinglyList& operator=(SinglyList&&) = delete;

    void addHead(const string& val);
    void addTail(const string& val);
    void delHead();
    void addAfter(const string& target, const string& val);
    void addBefore(const string& target, const string& val);
    void delTail();
    void delByValue(const string& val);
    bool findValue(const string& val) const;
    void readForward() const;
    void readBack() const;
    void delAfterValue(const string& val);
    void delBeforeValue(const string& val);

    Node* getHead_test() const { return head; }

private:
    friend class SinhlyListSerializer;
};

#endif

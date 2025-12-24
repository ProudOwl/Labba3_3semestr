#include "SinglyList.h"
#include <iostream>
#include <fstream>

using namespace std;

SinglyList::SinglyList() : head(nullptr) {}

SinglyList::~SinglyList() {
    while (head) delHead();
}

void SinglyList::addHead(const string& val) {
    head = new Node(val, head);
}

void SinglyList::addTail(const string& val) {
    if (!head) {
        head = new Node(val, nullptr);
        return;
    }
    Node* curr = head;
    while (curr->next) curr = curr->next;
    curr->next = new Node(val, nullptr);
}

void SinglyList::delHead() {
    if (!head) {
        cout << "Список пуст. Удаление невозможно." << endl;
        return;
    }
    Node* temp = head;
    head = head->next;
    delete temp;
}

void SinglyList::addAfter(const string& target, const string& val) {
    Node* curr = head;
    while (curr && curr->value != target) curr = curr->next;
    if (!curr) {
        cout << "Элемент '" << target << "' не найден." << endl;
        return;
    }
    curr->next = new Node(val, curr->next);
}

void SinglyList::addBefore(const string& target, const string& val) {
    if (!head) {
        cout << "Список пуст. Добавление невозможно." << endl;
        return;
    }
    if (head->value == target) {
        addHead(val);
        return;
    }
    Node* curr = head;
    while (curr->next && curr->next->value != target) curr = curr->next;
    if (!curr->next) {
        cout << "Элемент '" << target << "' не найден." << endl;
        return;
    }
    curr->next = new Node(val, curr->next);
}

void SinglyList::delTail() {
    if (!head) {
        cout << "Список пуст. Удаление невозможно." << endl;
        return;
    }
    if (!head->next) {
        delete head;
        head = nullptr;
        return;
    }
    Node* curr = head;
    while (curr->next->next) curr = curr->next;
    delete curr->next;
    curr->next = nullptr;
}

void SinglyList::delByValue(const string& val) {
    if (!head) {
        cout << "Список пуст. Удаление невозможно." << endl;
        return;
    }
    if (head->value == val) {
        delHead();
        return;
    }
    Node* curr = head;
    while (curr->next && curr->next->value != val) curr = curr->next;
    if (!curr->next) {
        cout << "Элемент '" << val << "' не найден." << endl;
        return;
    }
    Node* temp = curr->next;
    curr->next = curr->next->next;
    delete temp;
}

bool SinglyList::findValue(const string& val) const {
    Node* curr = head;
    while (curr) {
        if (curr->value == val) return true;
        curr = curr->next;
    }
    return false;
}

void SinglyList::readForward() const {
    if (!head) {
        cout << "Список пуст." << endl;
        return;
    }
    cout << "Список: ";
    Node* curr = head;
    while (curr) {
        cout << curr->value;
        if (curr->next) cout << " -> ";
        curr = curr->next;
    }
    cout << endl;
}

void SinglyList::printBack(Node* node) const {
    if (!node) return;
    printBack(node->next);
    cout << node->value;
    if (node != head) cout << " <- ";
}

void SinglyList::readBack() const {
    if (!head) {
        cout << "Список пуст." << endl;
        return;
    }
    cout << "Список (обратный): ";
    printBack(head);
    cout << endl;
}

void SinglyList::delAfterValue(const string& val) {
    Node* curr = head;
    while (curr && curr->value != val) curr = curr->next;
    if (!curr || !curr->next) {
        cout << "Невозможно удалить элемент после '" << val << "'." << endl;
        return;
    }
    Node* temp = curr->next;
    curr->next = curr->next->next;
    delete temp;
}

void SinglyList::delBeforeValue(const string& val) {
    if (!head || !head->next || head->value == val) {
        cout << "Невозможно удалить элемент перед '" << val << "'." << endl;
        return;
    }
    if (head->next->value == val) {
        delHead();
        return;
    }
    Node* curr = head;
    while (curr->next->next && curr->next->next->value != val) curr = curr->next;
    if (!curr->next->next) {
        cout << "Элемент '" << val << "' не найден или перед ним нет элемента." << endl;
        return;
    }
    Node* temp = curr->next;
    curr->next = curr->next->next;
    delete temp;
}

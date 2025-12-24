#include "CBTSerialize.h"
#include "Serialize.h"
#include <iostream>
#include <queue>
#include <string>

using namespace std;

// Вспомогательная функция для сериализации дерева в текстовом формате
void CBTSerializer::serializeTree(ostream& os, CompleteBinaryTree::TreeNode* node) {
    if (!node) {
        // Для текстового формата используем обычный вывод
        os << "null\n";
        return;
    }
    os << node->data << "\n";
    serializeTree(os, node->left);
    serializeTree(os, node->right);
}

// Вспомогательная функция для десериализации дерева из текстового формата
CompleteBinaryTree::TreeNode* CBTSerializer::deserializeTree(istream& is) {
    string token;
    if (!getline(is, token) || token.empty()) return nullptr;
    
    if (token == "null") return nullptr;
    
    try {
        int value = stoi(token);
        CompleteBinaryTree::TreeNode* node = new CompleteBinaryTree::TreeNode(value);
        node->left = deserializeTree(is);
        node->right = deserializeTree(is);
        return node;
    } catch (...) {
        return nullptr;
    }
}

// Вспомогательная функция для сериализации дерева в бинарном формате
void CBTSerializer::serializeTreeBinary(ostream& os, CompleteBinaryTree::TreeNode* node) {
    if (!node) {
        int marker = -1;
        os.write(reinterpret_cast<const char*>(&marker), sizeof(marker));
        return;
    }
    os.write(reinterpret_cast<const char*>(&node->data), sizeof(node->data));
    serializeTreeBinary(os, node->left);
    serializeTreeBinary(os, node->right);
}

// Вспомогательная функция для десериализации дерева из бинарного формата
CompleteBinaryTree::TreeNode* CBTSerializer::deserializeTreeBinary(istream& is) {
    int value;
    is.read(reinterpret_cast<char*>(&value), sizeof(value));
    if (is.fail() || value == -1) return nullptr;
    
    CompleteBinaryTree::TreeNode* node = new CompleteBinaryTree::TreeNode(value);
    node->left = deserializeTreeBinary(is);
    node->right = deserializeTreeBinary(is);
    return node;
}

void CBTSerializer::saveToFile(const CompleteBinaryTree& tree, const string& filename) {
    ofstream file(filename);
    if (!file) return;
    serializeTree(file, tree.getRoot_test());
}

void CBTSerializer::loadFromFile(CompleteBinaryTree& tree, const string& filename) {
    ifstream file(filename);
    if (!file) return;
    
    // Очищаем существующее дерево
    tree.clear();
    
    // Десериализуем новое дерево
    CompleteBinaryTree::TreeNode* newRoot = deserializeTree(file);
    
    // Устанавливаем новый корень
    tree.setRoot(newRoot);
}

void CBTSerializer::saveToBinaryFile(const CompleteBinaryTree& tree, const string& filename) {
    ofstream file(filename, ios::binary);
    if (!file) return;
    serializeTreeBinary(file, tree.getRoot_test());
}

void CBTSerializer::loadFromBinaryFile(CompleteBinaryTree& tree, const string& filename) {
    ifstream file(filename, ios::binary);
    if (!file) return;
    
    // Очищаем существующее дерево
    tree.clear();
    
    // Десериализуем новое дерево
    CompleteBinaryTree::TreeNode* newRoot = deserializeTreeBinary(file);
    
    // Устанавливаем новый корень
    tree.setRoot(newRoot);
}

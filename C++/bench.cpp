#include <benchmark/benchmark.h>
#include "Array.h"
#include "SinglyList.h"
#include "DoublyList.h"
#include "Stack.h"
#include "Queue.h"
#include "Hash.h"
#include "CBT.h"

//для динамического массива - добавление в конец
static void BM_DynArr_PushEnd(benchmark::State& state) {
    for (auto _ : state) {
        Array arr;
        for (int i = 0; i < state.range(0); i++)
            arr.addEnd("X");
    }
}
BENCHMARK(BM_DynArr_PushEnd)->Range(1000, 100000);

//для односвязного списка - добавление в хвост
static void BM_List_AddTail(benchmark::State& state) {
    for (auto _ : state) {
        SinglyList l;
        for (int i = 0; i < state.range(0); i++)
            l.addTail("X");
    }
}
BENCHMARK(BM_List_AddTail)->Range(1000, 100000);

//для двусвязного списка - добавление в хвост
static void BM_DList_AddTail(benchmark::State& state) {
    for (auto _ : state) {
        DoublyList l;
        for (int i = 0; i < state.range(0); i++)
            l.addTail("Y");
    }
}
BENCHMARK(BM_DList_AddTail)->Range(1000, 100000);

//для стека - операции push и pop
static void BM_Stack_PushPop(benchmark::State& state) {
    for (auto _ : state) {
        Stack s;
        for (int i = 0; i < state.range(0); i++)
            s.push("A");
        for (int i = 0; i < state.range(0); i++)
            s.pop();
    }
}
BENCHMARK(BM_Stack_PushPop)->Range(1000, 100000);

//для очереди - операции push и pop
static void BM_Queue_PushPop(benchmark::State& state) {
    for (auto _ : state) {
        Queue q;
        for (int i = 0; i < state.range(0); i++)
            q.push("A");
        for (int i = 0; i < state.range(0); i++)
            q.pop();
    }
}
BENCHMARK(BM_Queue_PushPop)->Range(1000, 100000);

//для полного бинарного дерева - вставка элементов
static void BM_CBT_Insert(benchmark::State& state) {
    for (auto _ : state) {
        CompleteBinaryTree t;
        for (int i = 0; i < state.range(0); i++)
            t.insert(i);
    }
}
BENCHMARK(BM_CBT_Insert)->Range(1000, 50000);

//для полного бинарного дерева - поиск элемента
static void BM_CBT_Search(benchmark::State& state) {
    CompleteBinaryTree t;
    for (int i = 0; i < state.range(0); i++)
        t.insert(i);

    for (auto _ : state)
        t.search(state.range(0) / 2);
}
BENCHMARK(BM_CBT_Search)->Range(1000, 50000);

//для хеш-таблицы с цепочками - вставка
static void BM_ChainHash_Insert(benchmark::State& state) {
    for (auto _ : state) {
        Hash h(1000);
        for (int i = 0; i < state.range(0); i++)
            h.insert(std::to_string(i), "X");
    }
}
BENCHMARK(BM_ChainHash_Insert)->Range(1000, 50000);

//для хеш-таблицы с открытой адресацией - вставка
static void BM_OpenHash_Insert(benchmark::State& state) {
    for (auto _ : state) {
        OpenHash h(state.range(0) * 2);
        for (int i = 0; i < state.range(0); i++)
            h.insert(std::to_string(i), "X");
    }
}
BENCHMARK(BM_OpenHash_Insert)->Range(1000, 50000);


BENCHMARK_MAIN();

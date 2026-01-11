package main

import (
	"testing"
	"go-data-structures/structures"
)

func BenchmarkArrayPush(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := structures.NewArray()
		for j := 0; j < 1000; j++ {
			arr.AddTail("X")
		}
	}
}

func BenchmarkListAddTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := structures.NewList()
		for j := 0; j < 1000; j++ {
			list.AddTail("X")
		}
	}
}

func BenchmarkStackPushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack := structures.NewStack()
		for j := 0; j < 1000; j++ {
			stack.Push("X")
		}
		for j := 0; j < 1000; j++ {
			stack.Pop()
		}
	}
}

func BenchmarkQueueEnqueueDequeue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queue := structures.NewQueue()
		for j := 0; j < 1000; j++ {
			queue.Enqueue("X")
		}
		for j := 0; j < 1000; j++ {
			queue.Dequeue()
		}
	}
}

func BenchmarkBinaryTreeInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := structures.NewBinaryTree()
		for j := 0; j < 1000; j++ {
			tree.Insert(j)
		}
	}
}

func BenchmarkChainHashInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash := structures.NewChainHash(100)
		for j := 0; j < 1000; j++ {
			hash.Insert(string(rune(j)), "value")
		}
	}
}

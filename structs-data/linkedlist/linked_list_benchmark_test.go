package linkedlist

import (
	"testing"
)

func helperCreateList(b *testing.B, size int) *linkedList {
	b.Helper()

	l := New()
	for i := range size {
		l.Append(i)
	}
	return l
}

func BenchmarkLinkedList_Append(b *testing.B) {
	l := New()
	for b.Loop() {
		l.Append(1)
	}
}

func BenchmarkLinkedList_Prepend(b *testing.B) {
	l := New()
	for b.Loop() {
		l.Prepend(1)
	}
}

func BenchmarkLinkedList_Insert(b *testing.B) {
	l := helperCreateList(b, 1000)
	b.ResetTimer()

	index := 999
	for b.Loop() {
		l.Insert(index, 9999)
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	l := helperCreateList(b, 1000)
	b.ResetTimer()

	for b.Loop() {
		l.Get(l.size - 1)
	}
}

func BenchmarkLinkedList_Find(b *testing.B) {
	l := helperCreateList(b, 1000)
	b.ResetTimer()

	for b.Loop() {
		// Ищем -1, которого точно нет в списке
		l.Find(-1)
	}
}

func BenchmarkLinkedList_RemoveAt(b *testing.B) {
	l := helperCreateList(b, 1000)
	b.ResetTimer()

	for b.Loop() {
		l.RemoveAt(l.size - 1) // Замеряется только эта строка
		b.StopTimer()
		l.Append(1)
		b.StartTimer()
	}
}

func BenchmarkLinkedList_MoveToFront(b *testing.B) {
	// b.N позволяет задать достаточное количество элементов в саязаном саписке,
	// чтобы хватило на весь периаод работы теста
	l := helperCreateList(b, 1000)
	b.ResetTimer()

	for b.Loop() {
		l.MoveToFront(l.size - 1)
	}
}

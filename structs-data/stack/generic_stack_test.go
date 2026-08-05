package stack

import (
	"fmt"
	"testing"
)

// BenchmarkGenericStackOperations Комплексаная оценка работы стуктуры
func BenchmarkGenericStackOperations(b *testing.B) {
	// Разные размеры серий операций для проверки производительности под нагрузкой
	sizes := []int{10, 100, 1_000, 100_000, 1_000_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			s := NewGenericStack[int]()
			b.ResetTimer()

			// Одна итерация бенчмарка совершает `size` добавлений и извлечений.
			// Это позволяет измерить, как стек ведет себя при разном объеме данных.
			for b.Loop() {
				// Наполняем стек
				for i := range size {
					s.Push(i)
				}

				// Полностью очищаем стек
				for range size {
					_, _ = s.Pop()
				}
			}
		})
	}
}

func BenchmarkGenericStack_Push(b *testing.B) {
	b.ReportAllocs() // Включаем сбор метрик памяти

	s := NewGenericStack[int]()
	b.ResetTimer()

	// Измеряем только скорость добавления в конец среза
	for b.Loop() {
		s.Push(100)
	}
}

func BenchmarkGenericStack_Pop(b *testing.B) {
	b.ReportAllocs() // Включаем сбор метрик памяти

	s := NewGenericStack[int]()
	// Наполняем стек заранее, чтобы методу Pop всегда было что извлекать
	// Используем b.N, чтобы элементов точно хватило на все итерации
	for i := range b.N {
		s.Push(i)
	}

	b.ResetTimer() // Сбрасываем таймер, подготовка стека не учитывается

	for b.Loop() {
		_, _ = s.Pop()
	}
}

func BenchmarkGenericStack_Read(b *testing.B) {
	s := NewGenericStack[int]()
	s.Push(42) // Метод Read всегда будет читать этот элемент
	b.ResetTimer()

	for b.Loop() {
		_, _ = s.Read()
	}
}

func TestGenericStack(t *testing.T) {
	s := NewGenericStack[int]()

	// 1. Проверяем консистентность zero-value на пустом стеке
	readVal, readOk := s.Read()
	popVal, popOk := s.Pop()

	if readOk || popOk {
		t.Fatalf("Пустой стек вернул успешый флаг: Read() ok=%t, Pop() ok=%t", readOk, popOk)
	}

	if readVal != popVal {
		t.Errorf("Неконсистентные значения для пустого стека: Read()=%d, Pop()=%d", readVal, popVal)
	}

	if popVal != 0 {
		t.Errorf("Ожидалось zero-value (0) для пустого стека, получено %d", popVal)
	}

	// 2. Проверяем стандартный LIFO цикл
	s.Push(7)
	s.Push(14)

	// Проверяем Read
	if val, ok := s.Read(); !ok || val != 14 {
		t.Errorf("Read() ожидал (14, true), получил (%d, %t)", val, ok)
	}

	// Извлекаем 14
	if val, ok := s.Pop(); !ok || val != 14 {
		t.Errorf("Pop() ожидал (14, true), получил (%d, %t)", val, ok)
	}

	// Извлекаем 7
	if val, ok := s.Pop(); !ok || val != 7 {
		t.Errorf("Pop() ожидал (7, true), получил (%d, %t)", val, ok)
	}

	// Снова проверяем, что стек пуст и возвращает (0, false)
	if val, ok := s.Pop(); ok || val != 0 {
		t.Errorf("После очистки Pop() ожидал (0, false), получил (%d, %t)", val, ok)
	}
}

package sorting

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/alkosuv/go-playground/algorithmes"
)

func BenchmarkBruteForce(b *testing.B) {
	sizes := []int{10, 100, 1_000, 100_000, 1_000_000}

	for _, size := range sizes {
		data := algorithmes.GenerateSlice(b, size)

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for b.Loop() {
				BruteForce(data)
			}
		})
	}
}

func TestBruteForce(t *testing.T) {
	// Определяем тестовые сценарии
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Обычный массив",
			input:    []int{5, 3, 8, 2, 1, 4},
			expected: []int{1, 2, 3, 4, 5, 8},
		},
		{
			name:     "Уже отсортированный",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Отсортированный в обратном порядке",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Массив с дубликатами",
			input:    []int{3, 1, 3, 2, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
		{
			name:     "Пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Массив из одного элемента",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Отрицательные числа",
			input:    []int{-3, 10, -5, 0, 2},
			expected: []int{-5, -3, 0, 2, 10},
		},
	}

	// Запускаем каждый тест
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Передаем копию слайса, чтобы не портить исходные данные в структуре
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			result := BruteForce(inputCopy)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Провал в тесте '%s':\nОжидали: %v\nПолучили: %v", tt.name, tt.expected, result)
			}
		})
	}
}

package searching

import (
	"fmt"
	"testing"

	"github.com/alkosuv/go-playground/algorithmes"
)

func BenchmarkLine(b *testing.B) {
	sizes := []int{10, 100, 1_000, 100_000, 500_000, 1_000_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			data := algorithmes.GenerateSlice(b, size)
			target := data[len(data)-1]

			b.ResetTimer() // Сбрасываем таймер: время генерации массива не учитывается

			for b.Loop() {
				Line(data, target)
			}
		})
	}
}

func TestLine(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			name:   "Элемент находится в начале массива",
			array:  []int{10, 20, 30, 40, 50},
			target: 10,
			want:   0,
		},
		{
			name:   "Элемент находится в середине массива",
			array:  []int{10, 20, 30, 40, 50},
			target: 30,
			want:   2,
		},
		{
			name:   "Элемент находится в конце массива",
			array:  []int{10, 20, 30, 40, 50},
			target: 50,
			want:   4,
		},
		{
			name:   "Элемента нет в массиве",
			array:  []int{10, 20, 30, 40, 50},
			target: 99,
			want:   -1,
		},
		{
			name:   "Поиск в пустом массиве",
			array:  []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "Поиск в массиве из одного совпадающего элемента",
			array:  []int{7},
			target: 7,
			want:   0,
		},
		{
			name:   "Поиск первого совпадения при дубликатах",
			array:  []int{5, 2, 8, 2, 9},
			target: 2,
			want:   1, // Должен вернуть индекс первого вхождения
		},
		{
			name:   "Поиск отрицательного числа",
			array:  []int{-5, -3, 0, 3, 5},
			target: -3,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Line(tt.array, tt.target)
			if got != tt.want {
				t.Errorf("Line() = %d, want %d (массив: %v, цель: %d)", got, tt.want, tt.array, tt.target)
			}
		})
	}
}

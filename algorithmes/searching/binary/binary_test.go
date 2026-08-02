package searching

import (
	"fmt"
	"testing"

	"github.com/alkosuv/go-playground/algorithmes"
)

func BenchmarkBinary(b *testing.B) {
	sizes := []int{10, 100, 1_000, 100_000, 500_000, 1_000_000, 1_000_000_000}

	for _, size := range sizes {
		data := algorithmes.GenerateSortSlice(b, size)

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()

			for b.Loop() {
				Binary(data, data[len(data)-1])
			}
		})
	}
}

func TestBinary(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			name:   "Элемент в центре",
			array:  []int{10, 20, 30, 40, 50},
			target: 30,
			want:   2,
		},
		{
			name:   "Элемент на границах (начало и конец)",
			array:  []int{10, 20, 30},
			target: 10,
			want:   0,
		},
		{
			name:   "Элемента нет в массиве",
			array:  []int{10, 20, 30},
			target: 25,
			want:   -1,
		},
		{
			name:   "Пустой массив",
			array:  []int{},
			target: 5,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Binary(tt.array, tt.target); got != tt.want {
				t.Errorf("Binary() = %d, want %d", got, tt.want)
			}
		})
	}
}

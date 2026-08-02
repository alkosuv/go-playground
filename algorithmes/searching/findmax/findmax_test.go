package findmax

import (
	"fmt"
	"testing"

	"github.com/alkosuv/go-playground/algorithmes"
)

// BenchmarkFindMaxFast тестирует только линейный алгоритм O(N).
// Сюда включены все размеры, так как алгоритм справится с ними мгновенно.
func BenchmarkFindMaxFast(b *testing.B) {
	sizes := []int{10, 100, 1_000, 100_000, 1_000_000}

	for _, size := range sizes {
		data := algorithmes.GenerateSlice(b, size)

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				FindMax(data)
			}
		})
	}
}

// BenchmarkFindMaxSlow тестирует квадратичный алгоритм O(N^2).
// ВНИМАНИЕ: Большие размеры (100k и 1M) намеренно исключены,
// чтобы тест не завис намертво. Максимум для O(N^2) — 10 000 элементов.
func BenchmarkFindMaxSlow(b *testing.B) {
	sizes := []int{10, 100, 1_000, 100_000, 1_000_000}

	for _, size := range sizes {
		data := algorithmes.GenerateSlice(b, size)

		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				FindMaxSlow(data)
			}
		})
	}
}

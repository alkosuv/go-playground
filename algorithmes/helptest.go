package algorithmes

import (
	"math/rand"
	"testing"
)

// GenerateSlice – вспомогательная функция для генерации случайного слайса указаной длины
func GenerateSlice(b *testing.B, size int) []int {
	b.Helper()

	slice := make([]int, size)
	for i := range size {
		slice[i] = rand.Intn(1_000_000)
	}
	return slice
}

// GenerateSortSlice – вспомогательная функция для генерации отсортированного слайса указаной длины
func GenerateSortSlice(b *testing.B, size int) []int {
	b.Helper()

	slice := make([]int, size)
	for i := range size {
		slice[i] = i
	}
	return slice
}

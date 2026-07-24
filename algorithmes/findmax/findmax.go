package findmax

import "errors"

//go:generate go test -bench=. -benchmem

var ErrEmptySlice = errors.New("cannot find maximum of an empty slice")

// FindMax находит наибольшее число в массиве за линейное время O(N).
func FindMax(array []int) int {
	if len(array) == 0 {
		panic(ErrEmptySlice)
	}

	result := array[0]

	for index := range array {
		if result < array[index] {
			result = array[index]
		}
	}

	return result
}

// FindMaxSlow Квадратичный алгоритм O(N^2)
func FindMaxSlow(array []int) int {
	if len(array) == 0 {
		return 0
	}

	// Ищем элемент, который больше или равен всем остальным в массиве
	for index := range array {
		isMax := true
		for j := range array {
			if array[j] > array[index] {
				isMax = false
				break // этот элемент точно не максимум, переходим к следующему
			}
		}
		if isMax {
			return array[index]
		}
	}

	return array[0]
}

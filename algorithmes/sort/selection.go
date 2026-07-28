package sort

//go:generate go test -run=^BenchmarkSelection$ -bench="^BenchmarkSelection$"

// Selection – алгоритм сортировки выбором O(N^2)
func Selection(array []int) []int {
	for i := range array {
		indexSwap := i
		for j := i + 1; j < len(array); j++ {
			if array[indexSwap] > array[j] {
				indexSwap = j
			}
		}

		array[i], array[indexSwap] = array[indexSwap], array[i]
	}

	return array
}

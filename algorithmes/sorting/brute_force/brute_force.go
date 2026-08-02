package sorting

//go:generate go test -v -run=.
//go:generate go test -bench=. -benchmem

// BruteForce – алгоритм линейной сортировки O(N^2)
func BruteForce(array []int) []int {

	for i := range array {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}

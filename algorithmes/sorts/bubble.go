package sorts

//go:generate go test -bench=BenchmarkBubble -benchmem

// Bubble – алгоритм пузырькой сортировки O(N^2)
func Bubble(array []int) []int {
	var (
		sorted bool
		l      int = len(array) - 1
	)

	for !sorted {
		sorted = true
		for i := range l {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sorted = false
			}
		}
		l--
	}

	return array
}

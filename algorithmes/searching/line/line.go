package searching

//go:generate go test -v -run=.
//go:generate go test -bench=. -benchmem

// Line – алгоритм линейного поиска (O(N))
// Возращет индекс массива в случае нахождения target в массиве или -1
// В случаее если target в массиве несколько, то возращаетс индекс первого найденного target в массиве
func Line(array []int, target int) int {
	for index := range array {
		if array[index] == target {
			return index
		}
	}

	return -1
}

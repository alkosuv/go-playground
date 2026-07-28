package search

// Binary – алгоритм линейного поиска (O(logN))
// Возращет индекс массива в случае нахождения target в массиве или -1
// В случаее если target в массиве несколько, то возращаетс индекс первого найденного target в массиве
func Binary(array []int, target int) int {
	leftIndex := 0
	rightIndex := len(array) - 1

	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex-leftIndex)/2
		// middleIndex := leftIndex + ((rightIndex - leftIndex) >> 1)
		if array[middleIndex] == target {
			return middleIndex
		}

		if array[middleIndex] > target {
			rightIndex = middleIndex - 1
		}

		if array[middleIndex] < target {
			leftIndex = middleIndex + 1
		}
	}

	return -1
}

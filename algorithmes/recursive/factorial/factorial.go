package factorial

import "math"

//go:generate go test -v -run=TestFactorialRecursive
//go:generate go test -bench=BenchmarkFactorialRecursive -benchmem

func FactorialRecursive(num int64) (int64, error) {
	if num < 0 {
		return 0, ErrNegativNumber
	}

	if num == 0 || num == 1 {
		return 1, nil
	}

	result, err := FactorialRecursive(num - 1)
	if err != nil {
		return 0, err
	}

	if result > math.MaxInt64/num {
		return 0, ErrOverflow
	}

	return num * result, nil
}

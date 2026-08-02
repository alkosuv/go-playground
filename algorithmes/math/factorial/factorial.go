package factorial

import (
	"math"
)

//go:generate go test -v -run=TestFactorialLoop
//go:generate go test -bench=BenchmarkFactorialLoop -benchmem

// FactorialLoop – фунция расчёта факториала через цыкл
// возврщает ошибку, если не возможно произвести вычисления
func FactorialLoop(num int64) (int64, error) {
	if num < 0 {
		return 0, ErrNegativNumber
	}

	if num == 0 || num == 1 {
		return 1, nil
	}

	var result int64 = 1
	var i int64
	for i = 2; i <= num; i++ {
		if result > math.MaxInt64/i {
			return 0, ErrOverflow
		}

		result *= i
	}
	return result, nil
}

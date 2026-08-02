package stack

//go:generate go test -v -run TestStackLockFree
//go:generate go test -bench=BenchmarkStackLockFreeOperations -benchmem

// stack – простоя раелизация стуктуры Стек через слайс
type stackLockFree struct {
	data []int
}

func NewStackLockFree() *stack {
	return new(stack)
}

func (s *stackLockFree) Push(value int) {}

func (s *stackLockFree) Pop() (int, bool) {
	return 0, false
}

func (s *stackLockFree) Read() (int, bool) {
	return 0, false
}

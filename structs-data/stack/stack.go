package stack

//go:generate go test -v -run TestStack
//go:generate go test -bench=BenchmarkStackOperations -benchmem

// stack – простоя раелизация стуктуры Стек через слайс
type stack struct {
	data []int
}

func NewStack() *stack {
	return new(stack)
}

// Push – операция добавления нового элемента в стек (слайс) (O(N))
func (s *stack) Push(value int) {
	s.data = append(s.data, value)
}

// Pop – операция выталкивания элемента из стек (слайс) (O(N))
func (s *stack) Pop() (int, bool) {
	value, ok := s.Read()
	if ok {
		s.data = s.data[:len(s.data)-1]
		return value, ok
	}

	return value, ok
}

// Read – операция чтения верхниго элемента стека (слайс) (O(N))
func (s *stack) Read() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.data[len(s.data)-1], true
}

// IsEmpty – метод проверяет stack.data на пустоту
func (s *stack) IsEmpty() bool {
	if s.data == nil {
		return true
	}

	if len(s.data) == 0 {
		return true
	}

	return false
}

package stack

import "sync"

//go:generate go test -v -run TestStackWithMutex
//go:generate go test -bench=BenchmarkStackWithMutexOperations -benchmem

// stack – простоя раелизация стуктуры Стек через слайс
type stackWithMutex struct {
	mu   sync.Mutex
	data []int
}

func NewStackWithMutex() *stackWithMutex {
	return new(stackWithMutex)
}

// Push – операция добавления нового элемента в стек (слайс) (O(N))
func (s *stackWithMutex) Push(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, value)
}

// Pop – операция выталкивания элемента из стек (слайс) (O(N))
func (s *stackWithMutex) Pop() (int, bool) {
	value, ok := s.Read()
	if ok {
		s.mu.Lock()
		s.data = s.data[:len(s.data)-1]
		s.mu.Unlock()
		return value, ok
	}

	return value, ok
}

// Read – операция чтения верхниго элемента стека (слайс) (O(N))
func (s *stackWithMutex) Read() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[len(s.data)-1], true
}

// IsEmpty – метод проверяет stack.data на пустоту
func (s *stackWithMutex) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.data == nil {
		return true
	}

	if len(s.data) == 0 {
		return true
	}

	return false
}

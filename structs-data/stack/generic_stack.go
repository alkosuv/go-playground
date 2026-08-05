package stack

//go:generate go test -v -run TestGenericStack
//go:generate go test -bench=BenchmarkGenericStackOperations -benchmem

type genericStack[T any] struct {
	data []T
}

func NewGenericStack[T any]() *genericStack[T] {
	return new(genericStack[T])
}

func (s *genericStack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *genericStack[T]) Pop() (T, bool) {
	value, ok := s.Read()
	if ok {
		s.data = s.data[:len(s.data)-1]
		return value, ok
	}

	var zero T
	return zero, false
}

func (s *genericStack[T]) Read() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	return s.data[len(s.data)-1], true
}

func (s *genericStack[T]) IsEmpty() bool {
	if s.data == nil {
		return true
	}

	if len(s.data) == 0 {
		return true
	}

	return false
}

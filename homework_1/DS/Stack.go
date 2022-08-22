package DS

type stackType interface {
	int | string
}

type Stack[T stackType] struct {
	value []T
}

func (s *Stack[T]) Get() (v []T) {
	return s.value
}

func (s *Stack[T]) Push(v T) {
	s.value = append(s.value, v)
	println(v)
}

func (s *Stack[T]) Pop() (v T) {
	popped := s.value[len(s.value)-1]
	s.value = append([]T{}, s.value[:len(s.value)-1]...)
	return popped
}

func NewStack[T stackType]() Stack[T] {
	s := Stack[T]{}
	s.value = []T{}

	return s
}

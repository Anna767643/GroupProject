package stack

type Stack struct {
	data []int
}

func New() *Stack {
	return &Stack{}
}

// добавить на вершину
func (s *Stack) Push(v int) {
	s.data = append(s.data, v)
}

// извлечь и удалить
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	lenS := s.size() - 1
	val := s.data[lenS]
	s.data = s.data[:lenS]
	return val, true
}

/* Получить элемент на вершине стека */
func (s *Stack) Top() (int, bool) {

	if s.IsEmpty() {
		return 0, false
	}

	lenS := s.size()
	return s.data[lenS-1], true
}

func (s *Stack) size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return s.size() == 0
}

package main

type Stack struct {
	items []int
}

func (s *Stack) Push(valor int) {
	s.items = append(s.items, valor)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	valor := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return valor, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

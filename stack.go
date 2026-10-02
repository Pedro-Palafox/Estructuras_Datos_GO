package main

type Stack struct {
	items []int
}

// Agrega un elemento al Stack
func (s *Stack) Push(valor int) {
	s.items = append(s.items, valor)
}

// Comprueba si está vacío
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Elimina y devuelve el último elemento del Stack
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	valor := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return valor, true
}

// Devuelve el último elemento sin eliminarlo
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	return s.items[len(s.items)-1], true
}

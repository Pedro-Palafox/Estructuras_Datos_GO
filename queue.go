package main

type Queue struct {
	items []int
}

// Agrega un elemento al final de la Queue
func (q *Queue) Enqueue(valor int) {
	q.items = append(q.items, valor)
}

// Comprueba si la Queue está vacía
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Elimina y devuelve el primer elemento de la Queue
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	valor := q.items[0]
	q.items = q.items[1:]
	return valor, true

}

// Devuelve el primer elemento sin eliminarlo
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[0], true
}

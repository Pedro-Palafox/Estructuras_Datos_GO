package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(valor int) {
	q.items = append(q.items, valor)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	valor := q.items[0]
	q.items = q.items[1:]
	return valor, true

}

func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[0], true
}

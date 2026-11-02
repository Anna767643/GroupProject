package queue

type Queue struct {
	data []int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(v int) {
	q.data = append(q.data, v)
}

// извлечь
func (q *Queue) Pop() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *Queue) size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return q.size() == 0
}

/* Получить элемент в голове очереди */
func (q *Queue) Front() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.data[0], true
}

func (q *Queue) Back() any {
	if q.IsEmpty() {
		return nil
	}

	lenS := q.size()
	return q.data[lenS-1]
}

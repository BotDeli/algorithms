package queue

type Queue struct {
	list []any
}

func CreateQueue() *Queue {
	return &Queue{[]any{}}
}

func (q *Queue) Enqueue(val any) {
	q.list = append(q.list, val)
}

func (q *Queue) Dequeue() any {
	val := q.list[0]
	q.list = q.list[1:]
	return val
}

func (q *Queue) NotEmpty() bool {
	return len(q.list) > 0
}

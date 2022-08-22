package DS

type queueType interface {
	int | string
}

type Queue[T queueType] struct {
	value []T
}

func (q *Queue[T]) Get() (v []T) {
	return q.value
}

func (q *Queue[T]) Enqueue(v T) {
	q.value = append(q.value, v)
	println(v)
}

func (q *Queue[T]) Dequeue() (v T) {
	popped := q.value[0]
	q.value = append([]T{}, q.value[1:len(q.value)]...)
	return popped
}

func NewQueue[T queueType]() Queue[T] {
	q := Queue[T]{}
	q.value = []T{}

	return q
}

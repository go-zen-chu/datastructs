package queue

import "fmt"

// Queue implementation
type queue struct {
	data []int
	size int
}

type Queue interface {
	Enqueue(int)
	Dequeue() (bool, int)
	IsEmpty() bool
}

// NewQueue instantiates a new queue
func NewQueue(cap int) Queue {
	return &queue{data: make([]int, 0, cap), size: 0}
}

// Enqueue adds a new element at the end of the queue
func (q *queue) Enqueue(n int) {
	q.data = append(q.data, n)
	q.size++
}

// Dequeue removes the first element from queue
func (q *queue) Dequeue() (bool, int) {
	if q.IsEmpty() {
		return false, -1
	}
	q.size--
	val := q.data[0]
	q.data = q.data[1:]
	return true, val
}

// IsEmpty checks if the queue is empty
func (q *queue) IsEmpty() bool {
	return q.size == 0
}

// String implements Stringer interface
func (q *queue) String() string {
	return fmt.Sprint(q.data)
}

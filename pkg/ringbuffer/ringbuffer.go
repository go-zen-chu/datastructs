package ringbuffer

import "fmt"

type ringBuffer struct {
	buf      []int
	len      int
	startIdx int
	endIdx   int
}

type RingBuffer interface {
	Len() int
	IsEmpty() bool
	IsFull() bool
	Enqueue(data int) error
	Dequeue() (data int, err error)
}

func NewRingBuffer(len int) (RingBuffer, error) {
	if len <= 1 {
		return nil, fmt.Errorf("length of ring buffer can't be equal or less than 1: %d\n", len)
	}
	return &ringBuffer{
		// buf is fixed size
		buf: make([]int, len),
		len: len,
	}, nil
}

func (r *ringBuffer) Len() int {
	return len(r.buf)
}

func (r *ringBuffer) IsEmpty() bool {
	return r.startIdx == r.endIdx
}

func (r *ringBuffer) IsFull() bool {
	return r.startIdx-1 == r.endIdx
}

func (r *ringBuffer) ShowQueue() []int {
	q := make([]int, 0, r.len)
	for idx := r.startIdx; idx != r.endIdx; idx = (idx + 1) % r.len {
		q = append(q, r.buf[idx])
		// startIdx can be +1 larger than endIdx, so you need to break before endIdx
		if (idx+1)%r.len == r.endIdx {
			q = append(q, r.buf[r.endIdx])
			break
		}
	}
	return q
}

func (r *ringBuffer) Enqueue(data int) error {
	if r.IsFull() {
		return fmt.Errorf("ring buffer is full: start %d, end %d\n", r.startIdx, r.endIdx)
	} else {
		insertIdx := (r.endIdx + 1) % r.len
		r.buf[insertIdx] = data
		r.endIdx = insertIdx
	}
	return nil
}

func (r *ringBuffer) Dequeue() (data int, err error) {
	if r.IsEmpty() {
		return 0, fmt.Errorf("ring buffer is empty: start %d, end %d\n", r.startIdx, r.endIdx)
	} else {
		data = r.buf[r.startIdx]
		r.startIdx = (r.startIdx + 1) % r.len
		return data, nil
	}
}

package stack

import "fmt"

type stack struct {
	data []int
	size int
}

type Stack interface {
	Push(int)
	Pop() (bool, int)
	Top() int
	IsEmpty() bool
}

// NewStack instantiates a new stack
func NewStack(cap int) Stack {
	return &stack{data: make([]int, 0, cap), size: 0}
}

// Push adds a new element at the end of the stack
func (s *stack) Push(n int) {
	s.data = append(s.data, n)
	s.size++
}

// Pop removes the last element from stack
func (s *stack) Pop() (bool, int) {
	if s.IsEmpty() {
		return false, -1
	}
	s.size--
	val := s.data[s.size]
	s.data = s.data[:s.size]
	return true, val
}

// Top returns the last element of stack
func (s *stack) Top() int {
	return s.data[s.size-1]
}

// IsEmpty checks if the stack is empty
func (s *stack) IsEmpty() bool {
	return s.size == 0
}

// String implements Stringer interface
func (s *stack) String() string {
	return fmt.Sprint(s.data)
}

package main

type Stacker interface {
	Push(v int)
	Pop() int
}

type stack struct {
	slice []int
}

func (s *stack) Push(v int) {
	s.slice = append(s.slice, v)
}

func (s *stack) Pop() int {
	r := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return r
}

func New() *stack {
	return &stack{}
}

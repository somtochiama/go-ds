package stack

import ()

type Stack struct {
	Elem []interface{}
	Size int
}

func New() *Stack{
	return &Stack{}
}

func (s *Stack) Pop() interface{}{
	if s.Size == 0 {
		return nil
	}

	ret := s.Elem[s.Size-1]
	s.Elem = s.Elem[:s.Size-1]
	s.Size--
	return ret
}

func (s *Stack) Push(item interface{}) int{
	s.Elem = append(s.Elem, item)
	s.Size++

	return s.Size
}
package sets 

import (
	"errors"
)

type Set struct {
	arr []interface{}
}

func (s *Set) add(e interface{}) error{
	for _, item := range(s.arr) {
		if item == e {
			return errors.New("Item already in set.")
		}
	}

	s.arr = append(s.arr, e)
	return nil
}

func (s *Set) List() []interface{}{
	return s.arr
}

func (s *Set) Delete(e interface{}) {
	for i, item := range(s.arr) {
		if item == e {
			s.arr = append(s.arr[:i], s.arr[i+1:]...)
		}
	}
}
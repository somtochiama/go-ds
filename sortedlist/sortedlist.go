package hashtable

import "errors"

type SortedList struct {
	arr []int
}

func (s *SortedList) Add(e int) {
	if len(s.arr) == 0 {
		s.arr = append(s.arr, e )
		return
	}

	i := 0
	for e > s.arr[0] {
		i++
	}

	insert(s.arr, e, i)
}

func (s *SortedList) Remove(n int) error{
	if n > len(s.arr) - 1 {
		return errors.New("Index out of ranger")
	}

	if  n == len(s.arr) - 1 {
		s.arr = s.arr[:n]
		return nil
	}

	s.arr = append(s.arr[:n], s.arr[n+1:]...)
	return nil
}

func insert(arr []int, elem int, index int) []int{
	if index == len(arr) {
		arr = append(arr, elem)
	} else if index == 0 {
		arr = append([]int{elem}, arr...)
	} else {
		arr = append(arr[:index+1], arr[index:]...)
		arr[index] = elem
	}

	return arr
}

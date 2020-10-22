package queue 

type Queue struct {
	Elem []interface{}
	Size int
}


func New() *Queue{
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) int{
	q.Elem = append(q.Elem, item)
	q.Size++
	return q.Size
} 

func (q *Queue) Dequeue() interface{}{
	if q.Size == 0 {
		return nil
	}

	ret := q.Elem[0]
	q.Elem = q.Elem[1:]
	q.Size--

	return ret
}
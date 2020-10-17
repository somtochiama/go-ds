package priorityq

type item struct {
	e interface{}
	p int
}

type PriorQ struct {
	Elem []item
	Size int
}

func New() *PriorQ{
	return &PriorQ{}
}

func (pq *PriorQ) Enqueue(i interface{}, p int) int {
	value := item{
		e: i,
		p: p,
	}

	j := 0
	for j < pq.Size && p <= pq.Elem[j].p {
		j++
	}
	if j == pq.Size {
		pq.Elem = append(pq.Elem, value)
	} else if j == 0 {
		pq.Elem = append([]item{value}, pq.Elem...)
	} else {
		pq.Elem = append(pq.Elem[:j+1], pq.Elem[j:]...)
		pq.Elem[j] = value
	}

	pq.Size++
	return pq.Size
}


func (pq *PriorQ) Dequeue() interface{}{
	if pq.Size == 0 {
		return nil
	}

	ret := pq.Elem[0]
	pq.Elem = pq.Elem[1:]
	pq.Size--

	return ret
}
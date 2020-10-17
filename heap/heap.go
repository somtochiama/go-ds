package heap

// Max heap
type Heap struct {
	arr []int
	size int
}

func New() *Heap{
	return &Heap{}
}

func (h *Heap) Add(val int) {
	h.arr = append(h.arr, val)
	h.size++
	h.upHeapify(h.size - 1)
}

func (h *Heap) Remove() int{
	top := h.arr[0]
	h.arr[0] = h.arr[h.size-1]
	h.arr = h.arr[:h.size-1]
	h.size--
	h.downHeapify(0)

	return top
}

func (h *Heap) Peep() int{
	return h.arr[0]
}


func (h *Heap) upHeapify(index int) {
	parentIndex := (index-1)/2
	for parentIndex >= 0 && h.arr[parentIndex] < h.arr[index] {
		h.swap(parentIndex, index)
		index = parentIndex
		parentIndex = (index-1)/2
	}
}

func (h *Heap) downHeapify(index int) {
	leftChild := index * 2 + 1
	rightChild := index * 2 + 2
	biggest := index
	
	for biggest <= (h.size + 1)/2 {
		if h.arr[rightChild] > h.arr[biggest] {
			biggest = rightChild
		}

		if h.arr[leftChild] > h.arr[biggest] {
			biggest = leftChild
		}

		if biggest != index {
			h.swap(index, biggest)
		}
	}
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) List() []int{
	return h.arr
}

// func (h *Heap) downHeapify(index int) {
	
// 	if h.arr[parentIndex] < h.arr[index] {
// 		h.swap(parentIndex, index)
// 	}
// }
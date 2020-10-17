package doublylinkedlist

// import "fmt"

type DLL struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	val interface{}
	Next *Node
	Prev *Node
}

func New() *DLL {
	return &DLL{}
}

func(d *DLL) Add(val interface{}) {
	newnode := &Node{
		val: val,
	}
	if d.Head == nil {
		d.Head = newnode
		d.Tail = newnode
	}

	d.Tail.Next = newnode
	newnode.Prev = d.Tail
	d.Tail = newnode
}

func(d *DLL) Remove(val interface{}) {
	if d.Head == nil {
		return
	}
	
	if d.Head.val == val {
		if d.Head == d.Tail {
			d.Head = nil
			d.Tail = nil
		} else {
			d.Head = d.Head.Next
			d.Head.Prev = nil
		}

		return
	}

	current := d.Head.Next
	for current != nil && current.val != val {
		current = current.Next
	}

	if current == nil {
		return
	}

	if current == d.Tail {
		d.Tail = current.Prev
		current.Prev.Next = nil
	} else {
		current.Prev.Next = current.Next
		current.Next.Prev = current.Prev
	}
}

func (l *DLL) Print() []interface{}{
	var ret []interface{}
	current := l.Head
	for current != nil {
		// fmt.Println(current.val)
		ret = append(ret, current.val)
		current = current.Next
	}

	return ret
} 
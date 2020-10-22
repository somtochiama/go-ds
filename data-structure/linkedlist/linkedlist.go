package linkedlist

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	val interface{}
	Next *Node
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(val interface{}) int{
	newnode := &Node{
		val: val,
	}

	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
		l.Size++
		return l.Size
	} else {
		l.Tail.Next = newnode
		l.Tail = newnode
	}
	
	l.Size++
	return l.Size
}

func (l *LinkedList) Remove(val interface{}) {
	if l.Head == nil {
		return
	}

	current := l.Head
	if current.val == val {
		if l.Head == l.Tail {
			l.Head = nil
			l.Tail = nil
		} else {
			l.Head = current.Next
		}
		return
	}

	for current.Next != nil && current.Next.val != val {
		current = current.Next
	}

	fmt.Println(current, l.Tail)
	if current.Next != nil {
		if current.Next == l.Tail {
			current.Next = nil
			l.Tail = current
			return
		}

		current.Next = current.Next.Next 
	}
}

func (l *LinkedList) Print() []interface{}{
	var ret []interface{}
	current := l.Head
	for current != nil {
		// fmt.Println(current.val)
		ret = append(ret, current.val)
		current = current.Next
	}

	return ret
} 
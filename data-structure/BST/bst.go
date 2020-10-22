package BST

type BST struct {
	Root *Node
}

type Node struct {
	val int
	Left *Node
	Right *Node
}

func New(val int) *BST {
	return &BST{
		Root: &Node{
			val: val,
		},
	}
}

func(bst *BST) Insert(val int) {
	newnode := &Node{
		val: val,
	}


	current := bst.Root
	var lastNode *Node
	for current != nil {
		lastNode = current
		if val > current.val {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	if val > lastNode.val {
		current.Right = newnode
	} else {
		current.Left = newnode
	}
}
package linkedlist

type intNode struct {
	data int
	next *intNode
}

// LinkedList - struct for a singly-linked list.
type LinkedList struct {
	Length int
	Head   *intNode
}

// Push adds a node holding input data to end of the list.
func (l *LinkedList) Push(x int) {
	if l.Head == nil {
		node := new(intNode)
		node.data = x
		l.Head = node
		l.Length++
	} else {
		curr := l.Head
		for curr.next != nil {
			curr = curr.next
		}
		node := new(intNode)
		node.data = x
		curr.next = node
		l.Length++
	}
}

// Pop removes last node from the list.
func (l *LinkedList) Pop() {
	if l.Head != nil {
		var prev *intNode
		curr := l.Head
		for curr.next != nil {
			prev = curr
			curr = curr.next
		}
		prev.next = nil // GC collects now unreachable curr node
		l.Length--
	} else {
		panic("No node to pop; list is empty.")
	}
}

// Inserts a node holding input data into given index of the list.
func (l *LinkedList) insert(input, index int) {
	if index > l.Length {
		panic("Invalid index; given index larger than Length of the list.")
	}
	// TODO: insert
}

// Removes node at given index from the list.
func (l *LinkedList) remove(index int) {
	if index > l.Length {
		panic("Invalid index; given index larger than Length of the list.")
	}
}

// Returns data of node at given index.
func (l LinkedList) get(index int) int {
	if index > l.Length {
		panic("Invalid index; given index larger than Length of the list.")
	}
	return 0
}

// Searches and returns index of target's first occurrence.
func (l LinkedList) search(target int) int {
	return 0
}

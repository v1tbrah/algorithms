package design_linked_list

type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

type MyLinkedList struct {
	Head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if l.Head == nil {
		return -1
	}
	if index < 0 {
		return -1
	}

	temp := l.Head
	currIdx := 0
	for temp.Next != nil && currIdx < index {
		temp = temp.Next
		currIdx++
	}

	if currIdx < index {
		return -1
	}

	if index == 0 {
		return l.Head.Val
	}

	return temp.Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	if l.Head == nil {
		l.Head = &Node{Val: val}
		return
	}

	l.Head.Prev = &Node{Val: val, Next: l.Head}

	l.Head = l.Head.Prev

	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.Head == nil {
		l.Head = &Node{Val: val}
		return
	}

	temp := l.Head
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = &Node{Val: val, Prev: temp}

	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	if l.Head == nil {
		if index == 0 {
			l.Head = &Node{Val: val}
		}
		return
	}

	temp := l.Head
	currIdx := 0
	for temp.Next != nil && currIdx < index {
		temp = temp.Next
		currIdx++
	}

	if index == 0 {
		l.Head.Prev = &Node{Val: val, Next: l.Head}
		l.Head = l.Head.Prev
		return
	}

	idxNextAfterLast := currIdx
	nextAfterLast := temp
	for nextAfterLast != nil {
		nextAfterLast = nextAfterLast.Next
		idxNextAfterLast++
	}

	if idxNextAfterLast < index {
		return
	}

	if index == idxNextAfterLast {
		newNode := &Node{Val: val, Prev: temp}
		temp.Next = newNode

		return
	}

	prev := temp.Prev
	next := temp

	newNode := &Node{Val: val, Prev: temp.Prev, Next: temp}

	prev.Next = newNode
	next.Prev = newNode

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l.Head == nil {
		return
	}
	if index < 0 {
		return
	}

	temp := l.Head
	currIdx := 0
	for temp.Next != nil && currIdx < index {
		temp = temp.Next
		currIdx++
	}

	if index == 0 {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		}
		return
	}

	lastIndex := currIdx
	last := temp
	for last.Next != nil {
		last = last.Next
		lastIndex++
	}

	if lastIndex < index {
		return
	}

	if currIdx == lastIndex {
		temp.Prev.Next = nil
		return
	}

	prev := temp.Prev
	next := temp.Next

	prev.Next = next
	next.Prev = prev

	return
}

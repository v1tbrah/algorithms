package design_linked_list

type Node struct {
	val  int
	prev *Node
	next *Node
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if l.head == nil {
		return -1
	}
	if index < 0 {
		return -1
	}

	temp := l.head
	currIdx := 0
	for temp.next != nil && currIdx < index {
		temp = temp.next
		currIdx++
	}

	if currIdx < index {
		return -1
	}

	if index == 0 {
		return l.head.val
	}

	return temp.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	if l.head == nil {
		l.head = &Node{val: val}
		return
	}

	l.head.prev = &Node{val: val, next: l.head}

	l.head = l.head.prev

	return
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.head == nil {
		l.head = &Node{val: val}
		return
	}

	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = &Node{val: val, prev: temp}

	return
}

func (l *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	if l.head == nil {
		if index == 0 {
			l.head = &Node{val: val}
		}
		return
	}

	temp := l.head
	currIdx := 0
	for temp.next != nil && currIdx < index {
		temp = temp.next
		currIdx++
	}

	if index == 0 {
		l.head.prev = &Node{val: val, next: l.head}
		l.head = l.head.prev
		return
	}

	idxnextAfterLast := currIdx
	nextAfterLast := temp
	for nextAfterLast != nil {
		nextAfterLast = nextAfterLast.next
		idxnextAfterLast++
	}

	if idxnextAfterLast < index {
		return
	}

	if index == idxnextAfterLast {
		newNode := &Node{val: val, prev: temp}
		temp.next = newNode

		return
	}

	prev := temp.prev
	next := temp

	newNode := &Node{val: val, prev: temp.prev, next: temp}

	prev.next = newNode
	next.prev = newNode

	return
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if l.head == nil {
		return
	}
	if index < 0 {
		return
	}

	temp := l.head
	currIdx := 0
	for temp.next != nil && currIdx < index {
		temp = temp.next
		currIdx++
	}

	if index == 0 {
		l.head = l.head.next
		if l.head != nil {
			l.head.prev = nil
		}
		return
	}

	lastIndex := currIdx
	last := temp
	for last.next != nil {
		last = last.next
		lastIndex++
	}

	if lastIndex < index {
		return
	}

	if currIdx == lastIndex {
		temp.prev.next = nil
		return
	}

	prev := temp.prev
	next := temp.next

	prev.next = next
	next.prev = prev

	return
}

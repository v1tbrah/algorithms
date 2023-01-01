package design_linked_list

type Node struct {
	next *Node
	val  int
}

type MyLinkedList struct {
	head *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}

	head := l.head
	if head == nil {
		return -1
	}

	currNode := head
	currIdx := 0
	for currIdx < index {
		if currNode.next == nil {
			break
		}
		currNode = currNode.next
		currIdx++
	}

	if currIdx != index {
		return -1
	}

	return currNode.val
}

func (l *MyLinkedList) AddAtHead(val int) {
	newHead := &Node{val: val}
	currHead := l.head
	if currHead != nil {
		newHead.next = currHead
	}
	l.head = newHead
}

func (l *MyLinkedList) AddAtTail(val int) {
	newTail := &Node{val: val}
	last := l.head
	if last == nil {
		l.head = newTail
		return
	}
	for last.next != nil {
		last = last.next
	}
	last.next = newTail
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

	length := 1
	curr, prev, last := l.head, l.head, l.head
	for last.next != nil {
		if length <= index {
			prev = curr
			curr = curr.next
		}
		last = last.next
		length++
	}
	if index > length {
		return
	}
	if index == length {
		last.next = &Node{val: val}
		return
	}
	if index == 0 {
		newNode := Node{val: val}
		newNode.next = l.head
		l.head = &newNode
		return
	}

	newNode := Node{val: val}
	prev.next = &newNode
	newNode.next = curr
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	length := 1
	prev, curr, last := m.head, m.head, m.head
	for last.next != nil {
		if length <= index {
			prev = curr
			curr = curr.next
		}
		last = last.next
		length++
	}
	if index >= length {
		return
	}
	if index == 0 {
		if length == 1 {
			*m = Constructor()
		} else {
			m.head = m.head.next
		}
		return
	}
	prev.next = curr.next
}

package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k <= 0 {
		return head
	}

	length := 0
	last := head
	for last != nil && last.Next != nil {
		last = last.Next
		length++
	}
	length++

	k %= length
	if k == length {
		return head
	}

	last.Next = head

	newLast := head
	for i := 0; i != length-k-1; i++ {
		newLast = newLast.Next
	}

	newHead := newLast.Next
	newLast.Next = nil
	return newHead
}

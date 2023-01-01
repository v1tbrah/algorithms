package add_two_numbers

type ListNode struct {
	Next *ListNode
	Val  int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var head *ListNode
	var result *ListNode = new(ListNode)

	head = result

	var haveRemainder bool
	for isFirst := true; l1 != nil && l2 != nil; l1, l2, isFirst = l1.Next, l2.Next, false {
		if !isFirst {
			newNode := new(ListNode)
			result.Next = newNode
			result = result.Next
		}

		c := l1.Val + l2.Val
		if haveRemainder {
			c += 1
		}
		s := c % 10
		result.Val = s

		haveRemainder = s != c
	}

	for ; l1 != nil; l1 = l1.Next {
		newNode := new(ListNode)
		result.Next = newNode
		result = result.Next

		c := l1.Val
		if haveRemainder {
			c += 1
		}
		s := c % 10
		result.Val = s

		haveRemainder = s != c
	}

	for ; l2 != nil; l2 = l2.Next {
		newNode := new(ListNode)
		result.Next = newNode
		result = result.Next

		c := l2.Val
		if haveRemainder {
			c += 1
		}
		s := c % 10
		result.Val = s

		haveRemainder = s != c
	}

	if haveRemainder {
		newNode := new(ListNode)
		result.Next = newNode
		result = result.Next
		result.Val = 1
	}

	return head
}

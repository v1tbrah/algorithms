package problem2

// Task: https://leetcode.com/problems/add-two-numbers/

// Solution: https://leetcode.com/problems/add-two-numbers/submissions/1085506121/

// tags:
// #Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(N)
// Space: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy

	reminder := 0
	for l1 != nil || l2 != nil || reminder != 0 {
		curr.Next = new(ListNode)
		curr = curr.Next

		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}
		curr.Val += reminder

		reminder = curr.Val / 10
		curr.Val %= 10
	}

	return dummy.Next
}

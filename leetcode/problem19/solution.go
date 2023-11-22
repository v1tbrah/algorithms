package problem19

// Task: https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Solution: https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/1104443049/

// tags:
// #Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(N)
// Space: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	n = lenList(head) + 1 - n

	if n == 1 {
		head = head.Next
		return head
	}

	var prev *ListNode
	curr := head
	currN := 1
	for curr != nil && currN < n {
		currN++
		prev, curr = curr, curr.Next
	}

	prev.Next = curr.Next

	return head
}

func lenList(head *ListNode) int {
	lenL := 0
	for head != nil {
		lenL++
		head = head.Next
	}
	return lenL
}

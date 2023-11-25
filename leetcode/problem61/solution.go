package problem61

// Task: https://leetcode.com/problems/rotate-list/

// Solution: https://leetcode.com/problems/rotate-list/submissions/1106359695/

// tags:
// #Linked List
// #Two Pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(N)
// Space: O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	lenList := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		lenList++
	}

	if k > lenList {
		k %= lenList
	}
	if k == 0 || k == lenList {
		return head
	}

	currNumber := 1
	curr = head

	for currNumber < lenList-k {
		curr = curr.Next
		currNumber++
	}

	startListK := curr.Next

	curr.Next = nil

	endListK := startListK
	for endListK.Next != nil {
		endListK = endListK.Next
	}

	endListK.Next = head

	return startListK
}

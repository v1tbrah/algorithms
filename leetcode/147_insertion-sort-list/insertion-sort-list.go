package insertion_sort_list

// Task: https://leetcode.com/problems/insertion-sort-list/

// Solution: https://leetcode.com/problems/insertion-sort-list/submissions/1048564135/

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy := new(ListNode)

	curr := head
	for curr != nil {
		prev := dummy

		for prev.Next != nil && prev.Next.Val <= curr.Val {
			prev = prev.Next
		}

		next := curr.Next

		curr.Next = prev.Next
		prev.Next = curr

		curr = next
	}

	return dummy.Next
}

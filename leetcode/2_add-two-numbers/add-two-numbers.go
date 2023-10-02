package add_two_numbers

// Task: https://leetcode.com/problems/add-two-numbers/

// Solution: https://leetcode.com/problems/add-two-numbers/submissions/1065137383/

// tags:
// #Linked List

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(n)
// Space: O(1)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	result := new(ListNode)
	dummy.Next = result

	var prev *ListNode

	hasRemainder := false
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		if prev != nil {
			result = new(ListNode)
			prev.Next = result
		}

		result.Val = l1.Val + l2.Val
		if hasRemainder {
			result.Val++
		}

		if result.Val >= 10 {
			result.Val = result.Val % 10
			hasRemainder = true
		} else {
			hasRemainder = false
		}
		prev = result
	}

	for ; l1 != nil; l1 = l1.Next {
		if prev != nil {
			result = new(ListNode)
			prev.Next = result
		}

		result.Val = l1.Val
		if hasRemainder {
			result.Val++
		}
		if result.Val >= 10 {
			result.Val = result.Val % 10
			hasRemainder = true
		} else {
			hasRemainder = false
		}
		prev = result
	}

	for ; l2 != nil; l2 = l2.Next {
		if prev != nil {
			result = new(ListNode)
			prev.Next = result
		}

		result.Val = l2.Val
		if hasRemainder {
			result.Val++
		}
		if result.Val >= 10 {
			result.Val = result.Val % 10
			hasRemainder = true
		} else {
			hasRemainder = false
		}
		prev = result
	}

	if hasRemainder {
		if prev != nil {
			result = new(ListNode)
			prev.Next = result
		}

		result.Val = 1
	}

	return dummy.Next
}

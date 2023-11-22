package problem23

// Task: https://leetcode.com/problems/merge-k-sorted-lists/

// Solution: https://leetcode.com/problems/merge-k-sorted-lists/submissions/1104448354/

// tags:
// #Linked List
// #Heap (Priority Queue)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time: O(N*LogK)
// Space: O(N)
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := new(ListNode)
	result := new(ListNode)
	dummy = result

	newLists := make([]*ListNode, 0)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			newLists = append(newLists, lists[i])
		}
	}

	lastNonLeafIdx := (len(newLists) - 1) / 2
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(newLists, i)
	}

	if len(newLists) == 0 {
		return nil
	}

	for len(newLists) > 0 {
		minList := newLists[0]

		newNode := new(ListNode)
		newNode.Val = minList.Val

		result.Next = newNode
		result = result.Next

		minList = minList.Next
		newLists[0] = minList

		if minList == nil {
			newLists[0] = newLists[len(newLists)-1]
			newLists = newLists[:len(newLists)-1]
		}
		heapify(newLists, 0)
	}

	return dummy.Next
}

func heapify(arr []*ListNode, idx int) {
	smallestIdx := idx
	leftChildIdx := idx*2 + 1
	rightChildIdx := idx*2 + 2

	if leftChildIdx < len(arr) && !less(arr[smallestIdx], arr[leftChildIdx]) {
		smallestIdx = leftChildIdx
	}
	if rightChildIdx < len(arr) && !less(arr[smallestIdx], arr[rightChildIdx]) {
		smallestIdx = rightChildIdx
	}

	if idx != smallestIdx {
		arr[idx], arr[smallestIdx] = arr[smallestIdx], arr[idx]
		heapify(arr, smallestIdx)
	}
}

func less(list1, list2 *ListNode) bool {
	return list1.Val < list2.Val
}

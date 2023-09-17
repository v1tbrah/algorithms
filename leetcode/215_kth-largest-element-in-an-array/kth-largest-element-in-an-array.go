package kth_largest_element_in_an_array

// Task: https://leetcode.com/problems/kth-largest-element-in-an-array/

// Solution: https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/1051696277/

// tags:
// #Array
// #Sorting
// #Heap (Priority Queue)

// Time: O(n+k*logN)
// n - for heapify input array; k*logN for get k elements from Heap and heapify after all get operations
// Space: O(logN)
// n - max recursion stack in heapify operation
func findKthLargest(nums []int, k int) int {
	// 1. Heapify array to MaxHeap
	// 2. Get and remove max element from MaxHeap k times

	// in Heap lastNonLeafIdx is: (length array / 2) - 1
	lastNonLeafIdx := len(nums)/2 - 1

	// heapify all non leaf nodes
	for idx := lastNonLeafIdx; idx >= 0; idx-- {
		heapify(nums, len(nums), idx)
	}

	for i := 0; i < k; i++ {
		// current heap size
		heapSize := len(nums) - i
		// swap top element and last element
		nums[0], nums[heapSize-1] = nums[heapSize-1], nums[0]
		// heapify current array after swap
		heapify(nums, heapSize-1, 0)
	}

	// return k-th element from end of array
	return nums[len(nums)-k]
}

func heapify(nums []int, numsSize, idx int) {
	largestElIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < numsSize && nums[largestElIdx] < nums[leftChildNodeIdx] {
		largestElIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < numsSize && nums[largestElIdx] < nums[rightChildNodeIdx] {
		largestElIdx = rightChildNodeIdx
	}

	if idx != largestElIdx {
		nums[idx], nums[largestElIdx] = nums[largestElIdx], nums[idx]
		heapify(nums, numsSize, largestElIdx)
	}
}

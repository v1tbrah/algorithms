package kth_largest_element_in_an_array

// Task: https://leetcode.com/problems/kth-largest-element-in-an-array/

// Solution: https://leetcode.com/problems/kth-largest-element-in-an-array/submissions/1051794276/

// tags:
// #Array
// #Sorting
// #Heap (Priority Queue)

// Time: O(n*logK)
// need k for heapify inputArray[:k]; logK for each loop iterate after k-th element
// Space: O(logK)
// k - max recursion stack in heapify operation
func findKthLargest(nums []int, k int) int {
	// 1. Heapify array[:k] to MinHeap
	// 2. For each element after k, check if element > minElement in heap, then swap and heapify

	numsHeap := nums[:k]

	lastNonLeafIdx := len(numsHeap)/2 - 1

	for idx := lastNonLeafIdx; idx >= 0; idx-- {
		heapify(numsHeap, idx)
	}

	for i := k; i < len(nums); i++ {
		if nums[i] > numsHeap[0] {
			numsHeap[0] = nums[i]
			heapify(numsHeap, 0)
		}
	}

	return numsHeap[0]
}

func heapify(nums []int, idx int) {
	smallestElIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(nums) && nums[smallestElIdx] > nums[leftChildNodeIdx] {
		smallestElIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(nums) && nums[smallestElIdx] > nums[rightChildNodeIdx] {
		smallestElIdx = rightChildNodeIdx
	}

	if idx != smallestElIdx {
		nums[idx], nums[smallestElIdx] = nums[smallestElIdx], nums[idx]
		heapify(nums, smallestElIdx)
	}
}

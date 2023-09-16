package sort_an_array

// Task: https://leetcode.com/problems/sort-an-array/

// Solution: https://leetcode.com/problems/sort-an-array/submissions/1051240022/

// tags:
// #Array
// #Sorting
// #Heap (Priority Queue)

func sortArray(nums []int) []int {
	lastNonLeafIdx := len(nums)/2 - 1

	for idx := lastNonLeafIdx; idx >= 0; idx-- {
		heapify(nums, idx)
	}

	result := make([]int, 0, len(nums))
	for len(nums) != 0 {
		var minEl int
		nums, minEl = pop(nums)
		result = append(result, minEl)
	}

	return result
}

func heapify(nums []int, idx int) {
	smallestValueIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(nums) && nums[smallestValueIdx] > nums[leftChildNodeIdx] {
		smallestValueIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(nums) && nums[smallestValueIdx] > nums[rightChildNodeIdx] {
		smallestValueIdx = rightChildNodeIdx
	}

	if idx != smallestValueIdx {
		nums[idx], nums[smallestValueIdx] = nums[smallestValueIdx], nums[idx]
		heapify(nums, smallestValueIdx)
	}
}

func pop(nums []int) ([]int, int) {
	if len(nums) == 0 {
		return nums, 0
	}

	// remember top element
	result := nums[0]

	// move last element to top position
	nums[0] = nums[len(nums)-1]
	// cut last element
	nums = nums[:len(nums)-1]

	heapify(nums, 0)

	return nums, result
}

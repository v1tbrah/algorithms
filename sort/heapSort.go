package sort

// Сортировка кучей.
// Time: O(n*logN)
// Space: O(n+logN) - где, n - новый массив для мутаций, logN - максимальный размер стека при рекурсии
func heapSort(nums []int) {
	lastNonLeafIdx := len(nums)/2 - 1

	for idx := lastNonLeafIdx; idx >= 0; idx-- {
		heapify(nums, idx)
	}

	copyForMutate := make([]int, len(nums))
	copy(copyForMutate, nums)

	for i := 0; i < len(nums); i++ {
		var minEl int
		copyForMutate, minEl = pop(copyForMutate)
		nums[i] = minEl
	}
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

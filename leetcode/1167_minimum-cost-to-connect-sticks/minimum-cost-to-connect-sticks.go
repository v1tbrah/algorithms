package minimum_cost_to_connect_sticks

// Task: https://leetcode.com/problems/minimum-cost-to-connect-sticks/

// Solution: https://leetcode.com/problems/minimum-cost-to-connect-sticks/submissions/1059804199/

// tags:
// #Array
// #Heap (Priority Queue)

// Time: O(n*logN)
// Space: O(logN) - where logN for maximum length recursive stack heapify
func connectSticks(sticks []int) int {
	lastNonLeafIdx := (len(sticks) - 1) / 2
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(sticks, i)
	}

	var sum int
	for len(sticks) > 1 {
		first := sticks[0]
		sticks[0] = sticks[len(sticks)-1]
		sticks = sticks[:len(sticks)-1]
		heapify(sticks, 0)

		second := sticks[0]
		sumFirstSecond := first + second
		sum += sumFirstSecond
		sticks[0] = sumFirstSecond
		heapify(sticks, 0)
	}

	return sum
}

func heapify(arr []int, idx int) {
	smallestValIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[smallestValIdx] > arr[leftChildNodeIdx] {
		smallestValIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[smallestValIdx] > arr[rightChildNodeIdx] {
		smallestValIdx = rightChildNodeIdx
	}

	if idx != smallestValIdx {
		arr[idx], arr[smallestValIdx] = arr[smallestValIdx], arr[idx]
		heapify(arr, smallestValIdx)
	}
}

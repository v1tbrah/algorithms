package minheap

// Task: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

// Solution: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/submissions/1056666163/

// tags:
// #Array
// #Sorting
// #Matrix
// #Heap (Priority Queue)

// Let's imagine, that n=rows*cols
// Idea:
// 1. Make array with all elements - O(n)
// 2. Make MinHeap - O(n)
// 3. Get topElement from MinHeap k times O(klogn)

// Time: O(n+klogn)
// Space: O(n)
func kthSmallest(matrix [][]int, k int) int {
	all := make([]int, 0, len(matrix)*len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			all = append(all, matrix[i][j])
		}
	}

	lastNonLeafIdx := len(all)/2 - 1
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(all, i)
	}

	var result int
	for i := 0; i < k; i++ {
		result = all[0]
		all[0] = all[len(all)-1]
		all = all[:len(all)-1]
		heapify(all, 0)
	}

	return result
}

func heapify(arr []int, idx int) {
	smallestIdx := idx
	leftChildIdx := idx*2 + 1
	rightChildIdx := idx*2 + 2

	if leftChildIdx < len(arr) && arr[smallestIdx] > arr[leftChildIdx] {
		smallestIdx = leftChildIdx
	}
	if rightChildIdx < len(arr) && arr[smallestIdx] > arr[rightChildIdx] {
		smallestIdx = rightChildIdx
	}

	if idx != smallestIdx {
		arr[idx], arr[smallestIdx] = arr[smallestIdx], arr[idx]
		heapify(arr, smallestIdx)
	}
}

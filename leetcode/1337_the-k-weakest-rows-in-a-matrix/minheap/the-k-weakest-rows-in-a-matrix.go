package minheap

// Task: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/

// Solution: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/submissions/1055743676/

// tags:
// #Array
// #Sorting
// #Matrix
// #Heap (Priority Queue)
// #Binary Search

// Idea:
// 1. Make new matrix, where matrix[i] - pair[count, idx]
// 2. Make MinHeap
// 3. Get topElement from MinHeap k times

// Time: O(m*logn*k)
// Space: O(k)
func kWeakestRows(mat [][]int, k int) []int {
	countAndIdx := make([][2]int, 0, len(mat))
	for i, row := range mat {
		countAndIdx = append(countAndIdx, [2]int{binarySearch(row), i})
	}

	lastNonLeafIdx := len(countAndIdx)/2 - 1
	for i := lastNonLeafIdx; i >= 0; i-- {
		heapify(countAndIdx, i)
	}

	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		minIdx := countAndIdx[0][1]
		result = append(result, minIdx)
		countAndIdx[0] = countAndIdx[len(countAndIdx)-1]
		countAndIdx = countAndIdx[:len(countAndIdx)-1]
		heapify(countAndIdx, 0)
	}

	return result
}

func binarySearch(arr []int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] == 1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func heapify(arr [][2]int, idx int) {
	smallestIdx := idx
	leftChildIdx := smallestIdx*2 + 1
	rightChildIdx := smallestIdx*2 + 2

	if leftChildIdx < len(arr) && less(arr[leftChildIdx], arr[smallestIdx]) {
		smallestIdx = leftChildIdx
	}
	if rightChildIdx < len(arr) && less(arr[rightChildIdx], arr[smallestIdx]) {
		smallestIdx = rightChildIdx
	}

	if idx != smallestIdx {
		arr[idx], arr[smallestIdx] = arr[smallestIdx], arr[idx]
		heapify(arr, smallestIdx)
	}
}

func less(a, b [2]int) bool {
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return a[0] < b[0]
}

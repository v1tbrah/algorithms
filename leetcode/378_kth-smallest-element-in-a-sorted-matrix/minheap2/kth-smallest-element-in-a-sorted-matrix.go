package minheap2

// Task: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

// Solution: https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/submissions/1056944379/

// tags:
// #Array
// #Sorting
// #Matrix
// #Heap (Priority Queue)

// Time: O(m+klogm)
// Space: O(m)
func kthSmallest(matrix [][]int, k int) int {
	minHeap := make([][3]int, 0, len(matrix)) // [3]value,row,col
	for i := 0; i < len(matrix); i++ {
		minHeap = append(minHeap, [3]int{matrix[i][0], i, 0})
	}

	var kThSmallest int
	for i := 0; i < k; i++ {
		kThSmallestObj := minHeap[0]

		// get top, remove from minHeap, hepify
		minHeap[0] = minHeap[len(minHeap)-1]
		minHeap = minHeap[:len(minHeap)-1]
		heapify(minHeap, 0)

		currRow, currCol := kThSmallestObj[1], kThSmallestObj[2]
		if currCol < len(matrix)-1 { // if there are elements in this row

			// add element
			minHeap = append(minHeap, [3]int{
				matrix[currRow][currCol+1],
				currRow,
				currCol + 1,
			},
			)

			// hepify
			idxNewEl := len(minHeap) - 1
			parentIdx := (idxNewEl - 1) / 2
			for parentIdx >= 0 && minHeap[idxNewEl][0] < minHeap[parentIdx][0] {
				minHeap[idxNewEl], minHeap[parentIdx] = minHeap[parentIdx], minHeap[idxNewEl]

				idxNewEl = parentIdx
				parentIdx = (idxNewEl - 1) / 2
			}
		}

		kThSmallest = kThSmallestObj[0]
	}

	return kThSmallest
}

func heapify(arr [][3]int, idx int) {
	smallestValIdx := idx
	leftChildIdx := idx*2 + 1
	rightChildIdx := idx*2 + 2

	if leftChildIdx < len(arr) && arr[smallestValIdx][0] > arr[leftChildIdx][0] {
		smallestValIdx = leftChildIdx
	}
	if rightChildIdx < len(arr) && arr[smallestValIdx][0] > arr[rightChildIdx][0] {
		smallestValIdx = rightChildIdx
	}

	if idx != smallestValIdx {
		arr[idx], arr[smallestValIdx] = arr[smallestValIdx], arr[idx]
		heapify(arr, smallestValIdx)
	}
}

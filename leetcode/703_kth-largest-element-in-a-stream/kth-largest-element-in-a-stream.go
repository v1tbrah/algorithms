package kth_largest_element_in_a_stream

// Task: https://leetcode.com/problems/kth-largest-element-in-a-stream/

// Solution: https://leetcode.com/problems/kth-largest-element-in-a-stream/submissions/1052859389/

// tags:
// #Tree
// #Design
// #Binary Search Tree
// #Heap (Priority Queue)
// #Binary Tree
// #Data Stream

type KthLargest struct {
	arr []int // MaxHeap

	topK []int // MinHeap with the K largest elements

	k int
}

func Constructor(k int, nums []int) KthLargest {
	// copy nums to array, because we don't want, that somebody change base-array of nums array after end of function
	arr := make([]int, len(nums))
	copy(arr, nums)

	// make our arr MaxHeap
	lastNonLeafIdx := len(arr)/2 - 1

	for i := lastNonLeafIdx; i >= 0; i-- {
		maxHeapify(arr, i)
	}

	// make topK array - MinHeap with the K largest elements
	topK := make([]int, 0, k)
	copyArr := make([]int, len(arr))
	copy(copyArr, arr)

	for i := 0; i < len(arr) && i < k; i++ {
		topK = append(topK, copyArr[0])
		copyArr[0] = copyArr[len(copyArr)-1]
		copyArr = copyArr[:len(copyArr)-1]
		maxHeapify(copyArr, 0)
	}
	// now topK is MaxHeap the largest elements

	// reverse topK, because we need MinHeap
	for i, j := 0, len(topK)-1; i <= j; i, j = i+1, j-1 {
		topK[i], topK[j] = topK[j], topK[i]
	}

	return KthLargest{
		arr:  arr,
		k:    k,
		topK: topK,
	}
}

func (this *KthLargest) Add(val int) int {
	// add new value
	this.arr = append(this.arr, val)

	// heapify it from bottom to up
	idx := len(this.arr) - 1
	parentIdx := (idx - 1) / 2

	for parentIdx >= 0 && this.arr[idx] > this.arr[parentIdx] {
		this.arr[idx], this.arr[parentIdx] = this.arr[parentIdx], this.arr[idx]

		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}

	// if val grater then minimum element of topK largest elements,
	// we need swap it with minimum element and heapify new topK
	if len(this.topK) < this.k {
		this.topK = append(this.topK, val)
	} else if val > this.topK[0] {
		this.topK[0] = val
	}

	minHeapify(this.topK, 0)

	// return minimum of topK largest element -> it's Kth element
	return this.topK[0]
}

func maxHeapify(arr []int, idx int) {
	largestElIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[largestElIdx] < arr[leftChildNodeIdx] {
		largestElIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[largestElIdx] < arr[rightChildNodeIdx] {
		largestElIdx = rightChildNodeIdx
	}

	if idx != largestElIdx {
		arr[idx], arr[largestElIdx] = arr[largestElIdx], arr[idx]
		maxHeapify(arr, largestElIdx)
	}
}

func minHeapify(arr []int, idx int) {
	smallestElIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(arr) && arr[smallestElIdx] > arr[leftChildNodeIdx] {
		smallestElIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(arr) && arr[smallestElIdx] > arr[rightChildNodeIdx] {
		smallestElIdx = rightChildNodeIdx
	}

	if idx != smallestElIdx {
		arr[idx], arr[smallestElIdx] = arr[smallestElIdx], arr[idx]
		minHeapify(arr, smallestElIdx)
	}
}

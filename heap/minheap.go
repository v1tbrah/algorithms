package heap

import (
	"fmt"
	"strings"
	"sync"
)

// The main idea:
// 1. index of the parent node of any node is [index of the node / 2]
// 2. index of the left child node is [index of the node * 2]
// 3. index of the right child node is [index of the node * 2 + 1]
// 4. node is a leaf node, when idx > h.realSize/2

type MinHeap[V Numeric] struct {
	// Create a complete binary tree using an array
	// Then use the binary tree to construct a Heap
	arr []V

	rwmu *sync.RWMutex
}

func NewMinHeap[V Numeric](initData []V) *MinHeap[V] {
	h := &MinHeap[V]{
		arr:  make([]V, len(initData)),
		rwmu: &sync.RWMutex{},
	}

	copy(h.arr, initData)

	lastNonLeafIdx := len(h.arr)/2 - 1

	for idx := lastNonLeafIdx; idx >= 0; idx-- {
		h.heapify(idx)
	}

	return h
}

func (h *MinHeap[V]) heapify(idx int) {
	smallestValueIdx := idx
	leftChildNodeIdx := idx*2 + 1
	rightChildNodeIdx := idx*2 + 2

	if leftChildNodeIdx < len(h.arr) && h.arr[smallestValueIdx] > h.arr[leftChildNodeIdx] {
		smallestValueIdx = leftChildNodeIdx
	}
	if rightChildNodeIdx < len(h.arr) && h.arr[smallestValueIdx] > h.arr[rightChildNodeIdx] {
		smallestValueIdx = rightChildNodeIdx
	}

	if idx != smallestValueIdx {
		h.arr[idx], h.arr[smallestValueIdx] = h.arr[smallestValueIdx], h.arr[idx]
		h.heapify(smallestValueIdx)
	}
}

// Push an element to Heap
func (h *MinHeap[V]) Push(el V) {
	h.rwmu.Lock()
	defer h.rwmu.Unlock()

	// add new element
	h.arr = append(h.arr, el)

	// index of the newly added element
	idx := len(h.arr) - 1
	// index of parent element
	parentIdx := (idx - 1) / 2

	// swap node and parent until node < parent
	for idx >= 0 && h.arr[idx] < h.arr[parentIdx] {
		h.arr[idx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[idx]

		idx = parentIdx
		parentIdx = (idx - 1) / 2
	}
}

// Pop removes the top element of the Heap
func (h *MinHeap[V]) Pop() {
	h.rwmu.Lock()
	defer h.rwmu.Unlock()

	if len(h.arr) == 0 {
		return
	}

	// insert last element to top
	h.arr[0] = h.arr[len(h.arr)-1]

	// remove last element
	h.arr = h.arr[:len(h.arr)-1]

	idx := 0

	// swap child and node until node > child
	for idx <= len(h.arr)-1 {
		smallestValueIdx := idx
		leftChildIdx := idx*2 + 1
		rightChildIdx := idx*2 + 2

		if leftChildIdx <= len(h.arr)-1 && h.arr[smallestValueIdx] > h.arr[leftChildIdx] {
			smallestValueIdx = leftChildIdx
		}
		if rightChildIdx <= len(h.arr)-1 && h.arr[smallestValueIdx] > h.arr[rightChildIdx] {
			smallestValueIdx = rightChildIdx
		}

		if idx != smallestValueIdx {
			h.arr[idx], h.arr[smallestValueIdx] = h.arr[smallestValueIdx], h.arr[idx]
		} else {
			break
		}
	}
}

// GetSize returns size of the Heap
func (h *MinHeap[V]) GetSize() int {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	return len(h.arr)
}

// GetPeek returns the top element of the Heap
func (h *MinHeap[V]) GetPeek() V {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	if len(h.arr) == 0 {
		var res V
		return res
	}

	return h.arr[0]
}

func (h *MinHeap[V]) String() string {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	var b strings.Builder
	b.WriteString("[")
	for i := 0; i < len(h.arr); i++ {
		b.WriteString(fmt.Sprintf("%v", h.arr[i]))
		if i < len(h.arr)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("]")

	return b.String()
}

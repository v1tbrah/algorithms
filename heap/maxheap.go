package heap

import (
	"fmt"
	"strings"
	"sync"
)

type MaxHeap[V Numeric] struct {
	arr []V

	rwmu *sync.RWMutex
}

func NewMaxHeap[V Numeric](inputData []V) *MaxHeap[V] {
	h := &MaxHeap[V]{
		arr:  make([]V, len(inputData)),
		rwmu: &sync.RWMutex{},
	}

	copy(h.arr, inputData)

	lastNonLeafNodeIdx := (len(h.arr) / 2) - 1

	for idx := lastNonLeafNodeIdx; idx >= 0; idx-- {
		h.heapify(idx)
	}

	return h
}

func (h *MaxHeap[V]) heapify(idx int) {
	largestValueIdx := idx
	leftChildNode := 2*idx + 1
	rightChildNode := 2*idx + 2

	if leftChildNode < len(h.arr) && h.arr[leftChildNode] > h.arr[largestValueIdx] {
		largestValueIdx = leftChildNode
	}

	if rightChildNode < len(h.arr) && h.arr[rightChildNode] > h.arr[largestValueIdx] {
		largestValueIdx = rightChildNode
	}

	if largestValueIdx != idx {
		h.arr[idx], h.arr[largestValueIdx] = h.arr[largestValueIdx], h.arr[idx]
		h.heapify(largestValueIdx)
	}
}

// Push an element to Heap
func (h *MaxHeap[V]) Push(el V) {
	h.rwmu.Lock()
	defer h.rwmu.Unlock()

	h.arr = append(h.arr, el)

	// index of the newly added element
	idx := len(h.arr) - 1
	// index of parent element
	parentIdx := (idx - 1) / 2

	// swap node and parent until node > parent
	for parentIdx >= 0 && h.arr[idx] > h.arr[parentIdx] {
		h.arr[idx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[idx]

		idx = parentIdx

		parentIdx = (idx - 1) / 2
	}
}

// Pop removes the top element of the Heap
func (h *MaxHeap[V]) Pop() {
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

	// swap child and node until node < child
	for idx < len(h.arr) {
		biggestValueIdx := idx
		leftChildIdx := idx*2 + 1
		rightChildIdx := idx*2 + 2

		if leftChildIdx < len(h.arr) && h.arr[biggestValueIdx] < h.arr[leftChildIdx] {
			biggestValueIdx = leftChildIdx
		}
		if rightChildIdx < len(h.arr) && h.arr[biggestValueIdx] < h.arr[rightChildIdx] {
			biggestValueIdx = rightChildIdx
		}

		if idx != biggestValueIdx {
			h.arr[idx], h.arr[biggestValueIdx] = h.arr[biggestValueIdx], h.arr[idx]
		} else {
			break
		}
	}
}

// GetSize returns size of the Heap
func (h *MaxHeap[V]) GetSize() int {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	return len(h.arr)
}

// GetPeek returns the top element of the Heap
func (h *MaxHeap[V]) GetPeek() V {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	if len(h.arr) == 0 {
		var res V
		return res
	}

	return h.arr[0]
}

func (h *MaxHeap[V]) String() string {
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

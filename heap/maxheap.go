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
		arr:  make([]V, 0, len(inputData)),
		rwmu: &sync.RWMutex{},
	}

	for i := 0; i < len(inputData); i++ {
		h.Push(inputData[i])
	}

	return h
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
		leftChildIdx := idx*2 + 1
		rightChildIdx := idx*2 + 2

		var isLeftChildMustSwap, isRightChildMustSwap bool
		bigger := h.arr[idx]
		if leftChildIdx < len(h.arr) && bigger < h.arr[leftChildIdx] {
			bigger = h.arr[leftChildIdx]
			isLeftChildMustSwap = true
		}
		if rightChildIdx < len(h.arr) && bigger < h.arr[rightChildIdx] {
			bigger = h.arr[rightChildIdx]
			isRightChildMustSwap = true
		}

		if isLeftChildMustSwap && !isRightChildMustSwap {
			h.arr[idx], h.arr[leftChildIdx] = h.arr[leftChildIdx], h.arr[idx]

			idx = leftChildIdx
		} else if isRightChildMustSwap {
			h.arr[idx], h.arr[rightChildIdx] = h.arr[rightChildIdx], h.arr[idx]

			idx = rightChildIdx
		}

		if !isLeftChildMustSwap && !isRightChildMustSwap {
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
	for i := 0; i <= len(h.arr); i++ {
		b.WriteString(fmt.Sprintf("%v", h.arr[i]))
		if i < len(h.arr)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("]")

	return b.String()
}

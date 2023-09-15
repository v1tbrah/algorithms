package heap

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type MaxHeap[V Numeric] struct {
	arr []V

	size int

	realSize int

	rwmu *sync.RWMutex
}

func NewMaxHeap[V Numeric](size int) (*MaxHeap[V], error) {
	if size <= 0 {
		return nil, errors.New("size must be grater then 0")
	}

	return &MaxHeap[V]{
		arr:      make([]V, size+1),
		size:     size,
		realSize: 0,
		rwmu:     &sync.RWMutex{},
	}, nil
}

func (h *MaxHeap[V]) Add(el V) error {
	h.rwmu.Lock()
	defer h.rwmu.Unlock()

	if h.realSize+1 > h.size {
		return errors.New("to many elements")
	}
	h.realSize++

	h.arr[h.realSize] = el

	idx := h.realSize
	parentIdx := idx / 2

	for parentIdx >= 1 && h.arr[idx] > h.arr[parentIdx] {
		tmp := h.arr[idx]
		h.arr[idx] = h.arr[parentIdx]
		h.arr[parentIdx] = tmp

		idx = parentIdx
		parentIdx = idx / 2
	}

	return nil
}

func (h *MaxHeap[V]) Pop() error {
	h.rwmu.Lock()
	defer h.rwmu.Unlock()

	if h.realSize < 1 {
		return errors.New("nothing elements")
	}

	h.arr[1] = h.arr[h.realSize]
	h.realSize--

	idx := 1
	leftChildIdx := idx * 2
	rightChildIdx := idx*2 + 1

	for idx <= h.realSize/2 { // before idx is not leaf node idx
		if h.arr[idx] < h.arr[leftChildIdx] {
			tmp := h.arr[idx]
			h.arr[idx] = h.arr[leftChildIdx]
			h.arr[leftChildIdx] = tmp

			idx = leftChildIdx
			leftChildIdx = idx * 2
			rightChildIdx = idx*2 + 1

			continue
		} else if h.arr[idx] < h.arr[rightChildIdx] {
			tmp := h.arr[idx]
			h.arr[idx] = h.arr[rightChildIdx]
			h.arr[rightChildIdx] = tmp

			idx = rightChildIdx
			leftChildIdx = idx * 2
			rightChildIdx = idx*2 + 1

			continue
		}

		break
	}

	return nil
}

func (h *MaxHeap[V]) GetSize() int {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	return h.realSize
}

func (h *MaxHeap[V]) GetPeek() (V, error) {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	if h.realSize < 1 {
		return 0, errors.New("nothing elements")
	}

	return h.arr[1], nil
}

func (h *MaxHeap[V]) String() string {
	h.rwmu.RLock()
	defer h.rwmu.RUnlock()

	if h.realSize < 1 {
		return "nothing elements"
	}

	var b strings.Builder
	b.WriteString("[")
	for i := 1; i <= h.realSize; i++ {
		b.WriteString(fmt.Sprintf("%v", h.arr[i]))
		if i < h.realSize {
			b.WriteString(",")
		}
	}
	b.WriteString("]")

	return b.String()
}

package heap

import (
	"errors"
	"strconv"
	"strings"
)

// The main idea:
// 1. index of the parent node of any node is [index of the node / 2]
// 2. index of the left child node is [index of the node * 2]
// 3. index of the right child node is [index of the node * 2 + 1]
// 4. node is a leaf node, when idx > h.realSize/2

type MaxHeap struct {
	arr []int

	size int

	realSize int
}

func NewMaxHeap(size int) (*MaxHeap, error) {
	if size <= 0 {
		return nil, errors.New("size must be grater then 0")
	}

	return &MaxHeap{
		arr:      make([]int, size+1),
		size:     size,
		realSize: 0,
	}, nil
}

func (h *MaxHeap) Add(el int) error {
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

func (h *MaxHeap) Pop() error {
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

func (h *MaxHeap) Size() int {
	return h.realSize
}

func (h *MaxHeap) GetPeek() (int, error) {
	if h.realSize < 1 {
		return 0, errors.New("nothing elements")
	}

	return h.arr[1], nil
}

func (h *MaxHeap) String() string {
	if h.realSize < 1 {
		return "nothing elements"
	}

	var b strings.Builder
	b.WriteString("[")
	for i := 1; i <= h.realSize; i++ {
		b.WriteString(strconv.Itoa(h.arr[i]))
		if i < h.realSize {
			b.WriteString(",")
		}
	}
	b.WriteString("]")

	return b.String()
}

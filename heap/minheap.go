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

type MinHeap struct {
	// Create a complete binary tree using an array
	// Then use the binary tree to construct a Heap
	arr []int

	// the number of elements is needed when instantiating an array
	// heapSize records the size of the array
	heapSize int

	// realSize records the number of elements in the Heap
	realSize int
}

func NewMinHeap(size int) (*MinHeap, error) {
	if size <= 0 {
		return nil, errors.New("size must be grater then 0")
	}

	return &MinHeap{
		arr:      make([]int, size+1),
		heapSize: size,
		realSize: 0,
	}, nil
}

// Add an element to Heap
func (h *MinHeap) Add(el int) error {
	if h.realSize+1 > h.heapSize {
		return errors.New("added to many elements")
	}
	h.realSize++

	// Add the element into the array
	h.arr[h.realSize] = el

	// Index of the newly added element
	idx := h.realSize

	// Parent node of the newly added element
	// Note if we use an array to represent the complete binary tree
	// and store the root node at index 1
	// index of the parent node of any node is [index of the node / 2]
	// index of the left child node is [index of the node * 2]
	// index of the right child node is [index of the node * 2 + 1]
	parentIdx := idx / 2

	// If the newly added element is smaller than its parent node,
	// its value will be exchanged with that of the parent node
	for idx >= 1 && h.arr[idx] < h.arr[parentIdx] {
		tmp := h.arr[idx]
		h.arr[idx] = h.arr[parentIdx]
		h.arr[parentIdx] = tmp
		// in Go language we can swap it like this:
		// h.arr[idx], h.arr[parenIdx] = h.arr[parentIdx], h.arr[idx]

		idx = parentIdx
		parentIdx = idx / 2
	}

	return nil
}

// GetPeek returns the top element of the Heap
func (h *MinHeap) GetPeek() (int, error) {
	if h.realSize < 1 {
		return 0, errors.New("nothing elements")
	}
	return h.arr[1], nil
}

// Pop removes the top element of the Heap
func (h *MinHeap) Pop() error {
	if h.realSize < 1 {
		return errors.New("nothing elements")
	}

	// Put the last element in the Heap to the top of Heap
	h.arr[1] = h.arr[h.realSize]
	// Decrease Heap size
	h.realSize--

	// index of the parent node of any node is [index of the node / 2]
	// index of the left child node is [index of the node * 2]
	// index of the right child node is [index of the node * 2 + 1]
	idx := 1
	leftChildIdx := idx * 2
	rightChildIdx := idx*2 + 1

	// When the deleted element is not a leaf node
	for idx <= h.realSize/2 {
		// If the deleted element is smaller than the left or right child
		// its value needs to be exchanged with the larger value
		// of the left and right child
		if leftChildIdx <= h.realSize && h.arr[idx] > h.arr[leftChildIdx] {
			tmp := h.arr[idx]
			h.arr[idx] = h.arr[leftChildIdx]
			h.arr[leftChildIdx] = tmp

			idx = leftChildIdx
			leftChildIdx = idx * 2
			rightChildIdx = idx*2 + 1

			continue
		} else if rightChildIdx <= h.realSize && h.arr[idx] > h.arr[rightChildIdx] {
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

// GetSize returns size of the Heap
func (h *MinHeap) GetSize() int {
	return h.realSize
}

func (h *MinHeap) String() string {
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

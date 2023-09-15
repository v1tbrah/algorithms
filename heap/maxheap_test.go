package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	mh := NewMaxHeap[int]([]int{1, 2, 3})

	mh.Push(4)
	assert.Equal(t, 4, mh.GetPeek())
	assert.Equal(t, 4, mh.GetSize())

	mh.Push(-1)
	assert.Equal(t, 4, mh.GetPeek())
	assert.Equal(t, 5, mh.GetSize())

	mh.Pop()
	assert.Equal(t, 3, mh.GetPeek())
	assert.Equal(t, 4, mh.GetSize())
}

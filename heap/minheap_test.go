package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinHeap(t *testing.T) {
	mh := NewMinHeap[int]([]int{3, 2, 1})
	assert.Equal(t, 1, mh.GetPeek())
	assert.Equal(t, 3, mh.GetSize())

	mh.Push(-1)
	assert.Equal(t, -1, mh.GetPeek())
	assert.Equal(t, 4, mh.GetSize())

	mh.Push(-2)
	assert.Equal(t, -2, mh.GetPeek())
	assert.Equal(t, 5, mh.GetSize())

	mh.Push(-3)
	assert.Equal(t, -3, mh.GetPeek())
	assert.Equal(t, 6, mh.GetSize())

	mh.Push(4)
	assert.Equal(t, -3, mh.GetPeek())
	assert.Equal(t, 7, mh.GetSize())

	mh.Pop()
	assert.Equal(t, -2, mh.GetPeek())
	assert.Equal(t, 6, mh.GetSize())

	mh.Pop()
	assert.Equal(t, -1, mh.GetPeek())
	assert.Equal(t, 5, mh.GetSize())
}

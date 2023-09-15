package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaxHeap(t *testing.T) {
	mh, err := NewMaxHeap[int](3)
	assert.NoError(t, err)

	assert.NoError(t, mh.Add(1))
	// 1. Add(1)

	// arr:
	// [1]

	// tree:
	//		1

	assert.NoError(t, mh.Add(2))
	// 2. Add(2)

	// arr:
	// [2,1]

	// tree:
	//		2
	//	1

	assert.NoError(t, mh.Add(3))
	// 3. Add(3)

	// arr:
	// [3,1,2]

	// tree:
	//		3
	//	1		2

	require.Error(t, mh.Add(4))

	assert.NoError(t, mh.Pop())
	currPeak, err := mh.GetPeek()
	assert.NoError(t, err)
	assert.Equal(t, 2, currPeak)

	assert.NoError(t, mh.Pop())
	currPeak, err = mh.GetPeek()
	assert.NoError(t, err)
	assert.Equal(t, 1, currPeak)

	assert.NoError(t, mh.Pop())
	require.Error(t, mh.Pop())
	currPeak, err = mh.GetPeek()
	require.Error(t, err)
}

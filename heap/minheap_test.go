package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinHeap(t *testing.T) {
	mh, err := NewMinHeap(3)
	assert.NoError(t, err)

	assert.NoError(t, mh.Add(3))
	// 1. Add(3)

	// arr:
	// [3]

	// tree:
	//		3

	assert.NoError(t, mh.Add(2))
	// 2. Add(2)

	// arr:
	// [2,3]

	// tree:
	//		2
	//	3

	assert.NoError(t, mh.Add(1))
	// 3. Add(1)

	// arr:
	// [1,3,2]

	// tree:
	//		1
	//	3		2

	require.Error(t, mh.Add(4))

	assert.NoError(t, mh.Pop())
	currPeak, err := mh.GetPeek()
	assert.NoError(t, err)
	assert.Equal(t, 2, currPeak)

	assert.NoError(t, mh.Pop())
	currPeak, err = mh.GetPeek()
	assert.NoError(t, err)
	assert.Equal(t, 3, currPeak)

	assert.NoError(t, mh.Pop())
	require.Error(t, mh.Pop())
	currPeak, err = mh.GetPeek()
	require.Error(t, err)
}

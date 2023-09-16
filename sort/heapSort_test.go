package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_heapSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  []int
	}{
		{
			name: "simple",
			in:   []int{7, 3, 2, 5, 6, 10, 9, 8, 1},
			out:  []int{1, 2, 3, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "nil",
			in:   nil,
			out:  nil,
		},
		{
			name: "empty",
			in:   []int{},
			out:  []int{},
		},
		{
			name: "equal elements",
			in:   []int{1, 1, 1},
			out:  []int{1, 1, 1},
		},
		{
			name: "simple #2",
			in:   []int{3, 2, 1},
			out:  []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heapSort(tt.in)
			assert.Equal(t, tt.out, tt.in)
		})
	}
}

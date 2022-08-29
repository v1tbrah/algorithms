package sort

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Zero elements",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "One element",
			input: []int{0},
			want:  []int{0},
		},
		{
			name:  "Two elements",
			input: []int{2, 1},
			want:  []int{1, 2},
		},
		{
			name:  "Three elements",
			input: []int{2, 1, 0},
			want:  []int{0, 1, 2},
		},
		{
			name:  "Many elements #1",
			input: []int{2, 1, 0, 88, 22, 14, 654, 7, 2, 2, 2, 9},
			want:  []int{0, 1, 2, 2, 2, 2, 7, 9, 14, 22, 88, 654},
		},
		{
			name:  "Many elements #2",
			input: []int{0, 0, 0, 0, 0, 1, 0, 0, 0},
			want:  []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			name:  "Many elements #3",
			input: []int{5, 1, 2, 3, 4},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.input)
			if got := tt.input; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

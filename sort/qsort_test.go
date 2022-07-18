package sort

import (
	"reflect"
	"testing"
)

func Test_Qsort(t *testing.T) {
	type args struct {
		input []int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Zero elements",
			args: args{input: []int{}},
			want: []int{},
		},
		{
			name: "One element",
			args: args{input: []int{0}},
			want: []int{0},
		},
		{
			name: "Two elements",
			args: args{input: []int{2, 1}},
			want: []int{1, 2},
		},
		{
			name: "Three elements",
			args: args{input: []int{2, 1, 0}},
			want: []int{0, 1, 2},
		},
		{
			name: "Many elements #1",
			args: args{input: []int{2, 1, 0, 88, 22, 14, 654, 7, 2, 2, 2, 9}},
			want: []int{0, 1, 2, 2, 2, 2, 7, 9, 14, 22, 88, 654},
		},
		{
			name: "Many elements #2",
			args: args{input: []int{0, 0, 0, 0, 0, 1, 0, 0, 0}},
			want: []int{0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Qsort(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("qsort() = %v, want %v", got, tt.want)
			}
		})
	}
}

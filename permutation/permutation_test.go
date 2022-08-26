package permutation

import (
	"reflect"
	"testing"
)

func Test_fillPermutations(t *testing.T) {
	type args struct {
		s            []byte
		i            int
		n            int
		permutations *[][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test 1",
			args: args{
				s:            []byte{1, 2, 3},
				i:            0,
				n:            3,
				permutations: &[][]byte{},
			},
			want: [][]byte{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fillPermutations(tt.args.s, tt.args.i, tt.args.n, tt.args.permutations)
			if !reflect.DeepEqual(*tt.args.permutations, tt.want) {
				t.Errorf("permutations = %v, want %v", *tt.args.permutations, tt.want)
			}
		})
	}
}

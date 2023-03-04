package perfect_squares

import "testing"

func Test_numSquares(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "13",
			n:    13,
			want: 2,
		},
		{
			name: "12",
			n:    12,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}

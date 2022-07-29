package sqrtx

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name   string
		target int
		result int
	}{
		{
			name:   "Test 0",
			target: 0,
			result: 0,
		},
		{
			name:   "Test 1",
			target: 4,
			result: 2,
		},
		{
			name:   "Test 2",
			target: 2147483647,
			result: 46340,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.target); got != tt.result {
				t.Errorf("mySqrt() = %v, want %v", got, tt.result)
			}
		})
	}
}

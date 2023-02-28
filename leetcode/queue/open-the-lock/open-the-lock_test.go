package open_the_lock

import "testing"

func Test_openLock(t *testing.T) {
	tests := []struct {
		name     string
		deadends []string
		target   string
		want     int
	}{
		{
			name:     "test1",
			deadends: []string{"0201", "0101", "0102", "1212", "2002"},
			target:   "0202",
			want:     6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.deadends, tt.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

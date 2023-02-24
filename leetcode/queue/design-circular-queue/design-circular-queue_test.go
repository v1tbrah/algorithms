package design_circular_queue

import (
	"fmt"
	"testing"
)

// It's not working like test. It's just for self-checking.
func TestLengthOfLastWordV1(t *testing.T) {
	tests := []struct {
		name     string
		behavior func()
	}{
		{
			name: "Test1",
			behavior: func() {
				q := Constructor(3)
				fmt.Println(q.EnQueue(1))
				fmt.Println(q.EnQueue(2))
				fmt.Println(q.EnQueue(3))
				fmt.Println(q.EnQueue(4))
				fmt.Println(q.Rear())
				fmt.Println(q.IsFull())
				fmt.Println(q.DeQueue())
				fmt.Println(q.EnQueue(4))
				fmt.Println(q.Rear())

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.behavior()
		})
	}
}

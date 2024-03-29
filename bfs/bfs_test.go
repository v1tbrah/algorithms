package bfs

import "testing"

func TestBFS(t *testing.T) {
	type args struct {
		graph     map[string][]string
		head      string
		isDesired func(tested string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				graph: func() map[string][]string {
					graph := map[string][]string{}
					graph["s"] = []string{"a", "b"}
					graph["a"] = []string{"c"}
					graph["b"] = []string{"c", "d", "e"}
					graph["e"] = []string{"d"}
					graph["d"] = []string{"e"}
					return graph
				}(),
				head: "s",
				isDesired: func(verifiable string) bool {
					return verifiable == "t"
				},
			},
			want: false,
		},
		{
			name: "Test 2",
			args: args{
				graph: func() map[string][]string {
					graph := map[string][]string{}
					graph["s"] = []string{"a", "b"}
					graph["a"] = []string{"c"}
					graph["b"] = []string{"c", "d", "e"}
					graph["e"] = []string{"d"}
					graph["d"] = []string{"e", "t"}
					return graph
				}(),
				head: "s",
				isDesired: func(verifiable string) bool {
					return verifiable == "t"
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.graph, tt.args.head, tt.args.isDesired); got != tt.want {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

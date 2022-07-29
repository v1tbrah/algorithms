package bfs

import "testing"

func TestBFS(t *testing.T) {
	type args struct {
		friends   graph
		desired   string
		isDesired func(verifiable string) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Have friend",
			args: args{
				friends: graph{
					"Vitya": []string{"Vanya", "Dasha", "Kirill"},
				},
				isDesired: func(verifiable string) bool {
					return verifiable == "Kirill"
				},
			},
			want: true,
		},
		{
			name: "Have not friend",
			args: args{
				friends: graph{
					"Vitya":  []string{"Vanya", "Dasha", "Kirill"},
					"Vanya":  []string{"Oleg", "Dasha", "Olga"},
					"Kirill": []string{"Valera", "Dima"},
				},
				isDesired: func(verifiable string) bool {
					return verifiable == "Inokentiy"
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.friends, tt.args.isDesired); got != tt.want {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

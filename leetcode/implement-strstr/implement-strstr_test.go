package implement_strstr

import "testing"

func TestStrStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty haystack",
			args: args{},
			want: 0,
		},
		{
			name: "Empty needle",
			args: args{},
			want: 0,
		},
		{
			name: "needle > haystack",
			args: args{haystack: "aaa", needle: "aaaa"},
			want: -1,
		},
		{
			name: "aaaaa -> bba",
			args: args{haystack: "aaaaa", needle: "bba"},
			want: -1,
		},
		{
			name: "hello -> ll",
			args: args{haystack: "hello", needle: "ll"},
			want: 2,
		},
		{
			name: "mississippi -> issip",
			args: args{haystack: "mississippi", needle: "issip"},
			want: 4,
		},
		{
			name: "lll -> lll",
			args: args{haystack: "lll", needle: "lll"},
			want: 0,
		},
		{
			name: "aaaBB -> BB",
			args: args{haystack: "aaaBB", needle: "BB"},
			want: 3,
		},
		{
			name: "bbaaa -> aa",
			args: args{haystack: "bbaaa", needle: "aa"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

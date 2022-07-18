package length_of_last_word

import "testing"

func TestLengthOfLastWordV1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Big space string",
			args: args{s: "               "},
			want: 0,
		},
		{
			name: "Hi            ",
			args: args{s: "Hi            "},
			want: 2,
		},
		{
			name: "Hello World",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "   fly me   to   the moon  ",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "luffy is still joyboy",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: "     end",
			args: args{s: "     end"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWordV1(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWordV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLastWordV2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "Big space string",
			args: args{s: "               "},
			want: 0,
		},
		{
			name: "Hi            ",
			args: args{s: "Hi            "},
			want: 2,
		},
		{
			name: "Hello World",
			args: args{s: "Hello World"},
			want: 5,
		},
		{
			name: "   fly me   to   the moon  ",
			args: args{s: "   fly me   to   the moon  "},
			want: 4,
		},
		{
			name: "luffy is still joyboy",
			args: args{s: "luffy is still joyboy"},
			want: 6,
		},
		{
			name: "     end",
			args: args{s: "     end"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLastWordV2(tt.args.s); got != tt.want {
				t.Errorf("LengthOfLastWordV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

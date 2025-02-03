package api

import "testing"

func TestMd5Pattern(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "valid md5",
			in:   "ab5bac7aa256e7d3a59685a5cd92b3e5",
			want: true,
		},
		{
			name: "invalid md5",
			in:   "0123456789abcdef0123456789abcdeg",
			want: false,
		},
		{
			name: "too short",
			in:   "0123456789abcdef0123456789abcde",
			want: false,
		},
		{
			name: "too long",
			in:   "0123456789abcdef0123456789abcdef0",
			want: false,
		},
		{
			name: "empty",
			in:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := md5Pattern.MatchString(tt.in)
			if got != tt.want {
				t.Errorf("md5Pattern.MatchString(%q) = %v; want %v", tt.in, got, tt.want)
			}
		})
	}

}

package leetcode

import "testing"

func TestSmallestNumber(t *testing.T) {
	tests := []struct {
		pattern string
		want    string
	}{
		{
			pattern: "I",
			want:    "12",
		},
		{
			pattern: "D",
			want:    "21",
		},
		{
			pattern: "IIDD",
			want:    "12543",
		},
		{
			pattern: "DDD",
			want:    "4321",
		},
		{
			pattern: "III",
			want:    "1234",
		},
		{
			pattern: "IIIDIDDD",
			want:    "123549876",
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := smallestNumber(tt.pattern); got != tt.want {
				t.Errorf("smallestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

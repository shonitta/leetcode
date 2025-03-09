package leetcode

import "testing"

func Test_findKthBit(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want byte
	}{
		{
			name: "Example 1: n=3, k=1",
			n:    3,
			k:    1,
			want: '0',
		},
		{
			name: "Example 2: n=4, k=11",
			n:    4,
			k:    11,
			want: '1',
		},
		{
			name: "Example 3: n=1, k=1",
			n:    1,
			k:    1,
			want: '0',
		},
		{
			name: "Example 4: n=2, k=3",
			n:    2,
			k:    3,
			want: '1',
		},
		{
			name: "Example 5: n=3, k=7",
			n:    3,
			k:    7,
			want: '1',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthBit(tt.n, tt.k); got != tt.want {
				t.Errorf("findKthBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

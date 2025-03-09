package leetcode

import "testing"

func Test_numberOfAlternatingGroups(t *testing.T) {
	tests := []struct {
		name   string
		colors []int
		k      int
		want   int
	}{
		{
			name:   "Example 1",
			colors: []int{0, 1, 0, 1, 0},
			k:      3,
			want:   3,
		},
		{
			name:   "Example 2",
			colors: []int{0, 1, 0, 0, 1, 0, 1},
			k:      6,
			want:   2,
		},
		{
			name:   "Example 3",
			colors: []int{1, 1, 0, 1},
			k:      4,
			want:   0,
		},
		{
			name:   "Example 4",
			colors: []int{1, 1, 1, 1, 1, 1, 1},
			k:      3,
			want:   0,
		},
		{
			name:   "Example 5",
			colors: []int{0, 0, 1},
			k:      3,
			want:   1,
		},
		{
			name:   "Example 6",
			colors: []int{1, 0, 1, 0, 1},
			k:      5,
			want:   1,
		},
		{
			name:   "Example 7",
			colors: []int{0, 1, 1},
			k:      3,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfAlternatingGroups(tt.colors, tt.k); got != tt.want {
				t.Errorf("numberOfAlternatingGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}

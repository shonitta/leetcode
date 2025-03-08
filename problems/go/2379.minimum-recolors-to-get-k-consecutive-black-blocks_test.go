package leetcode

import "testing"

func Test_minimumRecolors(t *testing.T) {
	tests := []struct {
		name   string
		blocks string
		k      int
		want   int
	}{
		{
			name:   "Example 1",
			blocks: "WBBWWBBWBW",
			k:      7,
			want:   3,
		},
		{
			name:   "Example 2",
			blocks: "WBWBBBW",
			k:      2,
			want:   0,
		},
		{
			name:   "All white",
			blocks: "WWWWW",
			k:      3,
			want:   3,
		},
		{
			name:   "All black",
			blocks: "BBBBB",
			k:      3,
			want:   0,
		},
		{
			name:   "Single block required",
			blocks: "WBW",
			k:      1,
			want:   0,
		},
		{
			name:   "K equals length",
			blocks: "WBWBW",
			k:      5,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRecolors(tt.blocks, tt.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}

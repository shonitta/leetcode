package leetcode

import "testing"

func TestCountDays(t *testing.T) {
	tests := []struct {
		name     string
		days     int
		meetings [][]int
		want     int
	}{
		{
			name:     "One meeting",
			days:     5,
			meetings: [][]int{{2, 3}},
			want:     3,
		},
		{
			name:     "Multiple meetings",
			days:     10,
			meetings: [][]int{{1, 2}, {3, 4}, {7, 8}},
			want:     4,
		},
		{
			name:     "Overlapping meetings",
			days:     10,
			meetings: [][]int{{1, 3}, {2, 4}, {5, 7}},
			want:     3,
		},
		{
			name:     "LeetCode Test Case 1",
			days:     10,
			meetings: [][]int{{5, 7}, {1, 3}, {9, 10}},
			want:     2,
		},
		{
			name:     "LeetCode Test Case 2",
			days:     5,
			meetings: [][]int{{2, 4}, {1, 3}},
			want:     1,
		},
		{
			name:     "LeetCode Test Case 3",
			days:     6,
			meetings: [][]int{{1, 6}},
			want:     0,
		},
		{
			name: "LeetCode Test Case 4",
			days: 57,
			meetings: [][]int{
				{3, 49},
				{23, 44},
				{21, 56},
				{26, 55},
				{23, 52},
				{2, 9},
				{1, 48},
				{3, 31},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countDays(tt.days, tt.meetings)
			if got != tt.want {
				t.Errorf("countDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

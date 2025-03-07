package leetcode

import (
	"reflect"
	"testing"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	tests := []struct {
		grid     [][]int
		expected []int
	}{
		{
			grid: [][]int{
				{1, 2, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expected: []int{2, 9},
		},
		{
			grid: [][]int{
				{1, 1},
				{2, 3},
			},
			expected: []int{1, 4},
		},
	}

	for _, test := range tests {
		result := findMissingAndRepeatedValues(test.grid)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For grid %v, expected %v, but got %v", test.grid, test.expected, result)
		}
	}
}

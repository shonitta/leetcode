package leetcode

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		weights []int
		wants   []int
	}{
		{[]int{1}, []int{0}},
		{[]int{1, 3}, []int{0, 1}},
		{[]int{1, 3, 2, 4}, []int{0, 1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			solution := Constructor(tt.weights)
			for i := 0; i < 1000; i++ {
				index := solution.PickIndex()
				if !slices.Contains(tt.wants, index) {
					t.Fatalf("got: %v, wants: %v", index, tt.wants)
				}
			}
		})
	}
}

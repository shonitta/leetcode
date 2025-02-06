package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleSameProduct(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 4, 6}, 8},
		{[]int{1, 2, 4, 5, 10}, 16},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 72},
	}

	for _, tt := range tests {
		got := tupleSameProduct(tt.input)
		assert.Equal(t, tt.expected, got, "input=%v", tt.input)
	}
}

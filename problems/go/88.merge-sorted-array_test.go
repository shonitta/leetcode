package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{-10, -9, -8, 0, 0, 0}, 3, []int{-8, -7, 4}, 3, []int{-10, -9, -8, -8, -7, 4}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("merge() = %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}

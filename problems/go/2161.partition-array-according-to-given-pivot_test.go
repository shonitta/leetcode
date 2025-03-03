package leetcode

import (
	"reflect"
	"testing"
)

func Test_pivotArray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		pivot int
		want  []int
	}{
		{
			name:  "example 1",
			nums:  []int{9, 12, 5, 10, 14, 3, 10},
			pivot: 10,
			want:  []int{9, 5, 3, 10, 10, 12, 14},
		},
		{
			name:  "example 2",
			nums:  []int{-3, 4, 3, 2},
			pivot: 2,
			want:  []int{-3, 2, 4, 3},
		},
		{
			name:  "all equal to pivot",
			nums:  []int{5, 5, 5, 5},
			pivot: 5,
			want:  []int{5, 5, 5, 5},
		},
		{
			name:  "single element",
			nums:  []int{1},
			pivot: 1,
			want:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pivotArray(tt.nums, tt.pivot)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

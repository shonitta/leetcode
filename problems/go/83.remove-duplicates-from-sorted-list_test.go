package leetcode

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{1, 1, 2, 3, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{1, 1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			head := sliceToList(tt.input)
			got := deleteDuplicates(head)
			if !reflect.DeepEqual(listToSlice(got), tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", listToSlice(got), tt.want)
			}
		})
	}
}

func sliceToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	var nums []int
	for current := head; current != nil; current = current.Next {
		nums = append(nums, current.Val)
	}
	return nums
}

/*
 * @lc app=leetcode id=2206 lang=golang
 *
 * [2206] Divide Array Into Equal Pairs
 */

package leetcode

import "sort"

// @lc code=start
func divideArray(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}
	return true
}

// @lc code=end

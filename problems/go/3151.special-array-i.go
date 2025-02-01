/*
 * @lc app=leetcode id=3151 lang=golang
 *
 * [3151] Special Array I
 */

package leetcode

// @lc code=start
func isArraySpecial(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]&1 == nums[i+1]&1 {
			return false
		}
	}

	return true
}

// @lc code=end

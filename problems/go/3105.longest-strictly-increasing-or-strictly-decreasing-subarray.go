/*
 * @lc app=leetcode id=3105 lang=golang
 *
 * [3105] Longest Strictly Increasing or Strictly Decreasing Subarray
 */

package leetcode

// @lc code=start
func longestMonotonicSubarray(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	maxLen, inc, dec := 1, 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			inc++
		} else {
			inc = 1
		}
		if nums[i] < nums[i-1] {
			dec++
		} else {
			dec = 1
		}
		maxLen = max(maxLen, inc, dec)

	}
	return maxLen
}

// @lc code=end

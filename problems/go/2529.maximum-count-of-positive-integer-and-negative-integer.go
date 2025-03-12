/*
 * @lc app=leetcode id=2529 lang=golang
 *
 * [2529] Maximum Count of Positive Integer and Negative Integer
 */

package leetcode

// @lc code=start
// findFirstPositive returns the index of the first positive number
// If no positive number exists, returns length of array
func findFirstPositive(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// findLastNegative returns the index of the last negative number
// If no negative number exists, returns -1
func findLastNegative(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

func maximumCount(nums []int) int {
	// Edge cases
	if len(nums) == 0 {
		return 0
	}
	if nums[0] > 0 {
		return len(nums) // all positive
	}
	if nums[len(nums)-1] < 0 {
		return len(nums) // all negative
	}

	// Find counts using binary search
	lastNegIdx := findLastNegative(nums)
	firstPosIdx := findFirstPositive(nums)

	negCount := lastNegIdx + 1
	posCount := len(nums) - firstPosIdx

	if negCount > posCount {
		return negCount
	}
	return posCount
}

// @lc code=end

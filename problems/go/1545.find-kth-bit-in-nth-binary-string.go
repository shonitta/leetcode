/*
 * @lc app=leetcode id=1545 lang=golang
 *
 * [1545] Find Kth Bit in Nth Binary String
 */

package leetcode

// @lc code=start
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}

	length := (1 << n) - 1 // 2^n - 1

	// If k is in the middle of the string
	if k == (length+1)/2 {
		return '1' // Middle bit is always '1' for n > 1
	}

	// If k is in the first half
	if k < (length+1)/2 {
		return findKthBit(n-1, k)
	}

	// If k is in the second half (after the middle bit)
	// We need to invert and reverse the bit from the corresponding position
	// in the first half
	invertedBit := findKthBit(n-1, length+1-k)
	if invertedBit == '0' {
		return '1'
	}
	return '0'
}

// @lc code=end

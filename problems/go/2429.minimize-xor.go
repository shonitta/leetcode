/*
 * @lc app=leetcode id=2429 lang=golang
 *
 * [2429] Minimize XOR
 */

package leetcode

// @lc code=start
func minimizeXor(num1 int, num2 int) int {
	targetBits := countBits(num2)
	res := 0

	// Set bits from num1 to res
	for i := 30; i >= 0 && targetBits > 0; i-- {
		if (num1 & (1 << i)) != 0 {
			res |= 1 << i
			targetBits--
		}
	}

	// Set remaining bits to res
	for i := 0; i <= 30 && targetBits > 0; i++ {
		if (res & (1 << i)) == 0 {
			res |= 1 << i
			targetBits--
		}
	}

	return res
}

func countBits(num int) int {
	count := 0
	for num != 0 {
		count += num & 1 // Add the last bit to count
		num >>= 1        // Remove the last bit
	}
	return count
}

// @lc code=end

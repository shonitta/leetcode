/*
 * @lc app=leetcode id=2116 lang=golang
 *
 * [2116] Check if a Parentheses String Can Be Valid
 */

package leetcode

// @lc code=start
func canBeValid(s string, locked string) bool {
	l := len(s)
	if l%2 != 0 {
		return false
	}

	// Check balance of left parentheses from left to right
	balance, countCanModify := 0, 0
	for i := 0; i < l; i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				balance++
			} else {
				balance--
			}
		} else {
			countCanModify++
		}
		if balance+countCanModify < 0 {
			return false
		}
	}
	// Check balance of right parentheses from right to left
	balance, countCanModify = 0, 0
	for i := l - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				balance++
			} else {
				balance--
			}
		} else {
			countCanModify++
		}
		if balance+countCanModify < 0 {
			return false
		}
	}

	return true
}

// @lc code=end

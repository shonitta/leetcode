/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

package leetcode

// @lc code=start
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		cl := toLower(s[l])
		cr := toLower(s[r])
		if !isAlphanumeric(cl) {
			l++
		} else if !isAlphanumeric(cr) {
			r--
		} else if cl != cr {
			return false
		} else {
			l++
			r--
		}
	}

	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// @lc code=end

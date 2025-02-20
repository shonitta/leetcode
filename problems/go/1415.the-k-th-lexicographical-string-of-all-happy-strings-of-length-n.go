/*
 * @lc app=leetcode id=1415 lang=golang
 *
 * [1415] The k-th Lexicographical String of All Happy Strings of Length n
 */

package leetcode

// @lc code=start
func getHappyString(n int, k int) string {
	abcThreshold := (1 << (n - 1))
	if k > 3*abcThreshold { // 3: {a, b, c}
		return ""
	}

	k-- // 0-based index
	happyString := ""

	firstChar := k / abcThreshold
	if firstChar == 0 {
		happyString = "a"
	} else if firstChar == 1 {
		happyString = "b"
	} else {
		happyString = "c"
	}
	k %= abcThreshold

	for i := 1; i < n; i++ {
		abcThreshold /= 2 // 2: {b, c} or {a, c} or {a, b}
		prev := happyString[len(happyString)-1]
		var options []byte

		if prev == 'a' {
			options = []byte{'b', 'c'}
		} else if prev == 'b' {
			options = []byte{'a', 'c'}
		} else { // prev == 'c'
			options = []byte{'a', 'b'}
		}

		nextChar := options[k/abcThreshold]
		happyString += string(nextChar)
		k %= abcThreshold
	}

	return happyString
}

// @lc code=end

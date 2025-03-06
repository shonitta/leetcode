/*
 * @lc app=leetcode id=2965 lang=golang
 *
 * [2965] Find Missing and Repeated Values
 */

package leetcode

// @lc code=start
func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	count := make([]int, n*n+1)
	var repeated, missing int

	for _, row := range grid {
		for _, num := range row {
			count[num]++
		}
	}

	for i := 1; i <= n*n; i++ {
		if count[i] == 2 {
			repeated = i
		}
		if count[i] == 0 {
			missing = i
		}
	}

	return []int{repeated, missing}
}

// @lc code=end

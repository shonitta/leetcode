/*
 * @lc app=leetcode id=1267 lang=golang
 *
 * [1267] Count Servers that Communicate
 */

package leetcode

// @lc code=start
func countServers(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	rowCounts, colCounts := make([]int, m), make([]int, n)
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				rowCounts[i]++
				colCounts[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (rowCounts[i] > 1 || colCounts[j] > 1) {
				res++
			}
		}
	}

	return res
}

// @lc code=end

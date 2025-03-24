/*
 * @lc app=leetcode id=3169 lang=golang
 *
 * [3169] Count Days Without Meetings
 */

package leetcode

import "sort"

// @lc code=start
func countDays(days int, meetings [][]int) int {
	// Sort meetings by start time and end time
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] < meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	// Merge overlapping meetings
	mergedMeetings := [][]int{meetings[0]}
	for _, meeting := range meetings[1:] {
		currentEnd := mergedMeetings[len(mergedMeetings)-1][1]
		if meeting[0] <= currentEnd {
			mergedMeetings[len(mergedMeetings)-1][1] = max(currentEnd, meeting[1])
		} else {
			mergedMeetings = append(mergedMeetings, meeting)
		}
	}

	// Count days without meetings
	countMeetingDays := 0
	for _, meeting := range mergedMeetings {
		countMeetingDays += meeting[1] - meeting[0] + 1
	}

	return days - countMeetingDays
}

// @lc code=end

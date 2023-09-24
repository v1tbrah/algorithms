package intuition

// Task: https://leetcode.com/problems/meeting-rooms-ii/

// Solution: https://leetcode.com/problems/meeting-rooms-ii/submissions/1058042566/

// tags:
// #Array

// Time: O(1000000+2*n)
// Space: O(2n)
func minMeetingRooms(intervals [][]int) int {
	start := make(map[int]int)
	end := make(map[int]int)

	for i := 0; i < len(intervals); i++ {
		start[intervals[i][0]]++
		end[intervals[i][1]]++
	}

	var currStarts, maxStarts int
	for i := 0; i <= 1000000; i++ {
		currStarts += start[i]
		currStarts -= end[i]
		if currStarts > maxStarts {
			maxStarts = currStarts
		}
	}

	return maxStarts
}

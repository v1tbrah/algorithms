package maximize_distance_to_closest_person

// Task: https://leetcode.com/problems/maximize-distance-to-closest-person/

// Solution: https://leetcode.com/problems/maximize-distance-to-closest-person/submissions/1065904695/

// tags:
// #Array

// Time: O(n)
// Space: O(1)
func maxDistToClosest(seats []int) int {
	leftSeatIdx := -1
	var maxDistance int
	var currDistance int
	for i := 0; i < len(seats); i++ {
		if seats[i] == 0 {
			continue
		}

		if leftSeatIdx == -1 {
			currDistance = i
		} else {
			currDistance = (i - leftSeatIdx) / 2
		}

		if currDistance > maxDistance {
			maxDistance = currDistance
		}
		leftSeatIdx = i
	}

	if seats[len(seats)-1] == 0 {
		currDistance = len(seats) - 1 - leftSeatIdx
		if currDistance > maxDistance {
			maxDistance = currDistance
		}
	}

	return maxDistance
}

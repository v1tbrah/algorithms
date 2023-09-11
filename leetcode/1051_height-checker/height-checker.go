package height_checker

// Task: https://leetcode.com/problems/height-checker/

// Solution: https://leetcode.com/problems/height-checker/submissions/1046626300/

// tags:
// #Array
// #Sorting
// #Counting Sort

func heightChecker(heights []int) int {
	if len(heights) < 2 {
		return 0
	}

	maxEl := heights[0]
	for i := 1; i < len(heights); i++ {
		if heights[i] > maxEl {
			maxEl = heights[i]
		}
	}

	idxCount := make([]int, maxEl+1)
	for i := range heights {
		idxCount[heights[i]]++
	}

	sorted := make([]int, len(heights))
	for i, k := 0, 0; i < len(idxCount); i++ {
		for j := 1; j <= idxCount[i]; j, k = j+1, k+1 {
			sorted[k] = i
		}
	}

	var result int
	for i := range heights {
		if sorted[i] != heights[i] {
			result++
		}
	}

	return result
}

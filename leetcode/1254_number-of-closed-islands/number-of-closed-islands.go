package number_of_closed_islands

// Task: https://leetcode.com/problems/number-of-closed-islands/

// Solution: https://leetcode.com/problems/number-of-closed-islands/submissions/1068061310/

// tags:
// #Matrix
// #Depth-First Search

// Time: O(m*n)
// Space: O(1)
func closedIsland(grid [][]int) int {
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				isClosedIsland := true
				dfs(grid, i, j, &isClosedIsland)
				if isClosedIsland {
					count++
				}
			}
		}
	}

	return count
}

func dfs(grid [][]int, i, j int, isClosedIsland *bool) {
	if i < 0 || j < 0 || i > len(grid)-1 || j > len(grid[i])-1 {
		*isClosedIsland = false
		return
	}

	if grid[i][j] != 0 {
		return
	}

	grid[i][j] = -1

	// r
	dfs(grid, i, j+1, isClosedIsland)
	// d
	dfs(grid, i+1, j, isClosedIsland)
	// l
	dfs(grid, i, j-1, isClosedIsland)
	// u
	dfs(grid, i-1, j, isClosedIsland)
}

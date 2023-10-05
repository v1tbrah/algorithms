package number_of_islands

// Task: https://leetcode.com/problems/number-of-islands/

// Solution: https://leetcode.com/problems/number-of-islands/submissions/1067953419/

// tags:
// #Matrix
// #Breadth-First Search

const water = '0'

// Time: O(m*n)
// Space: O(1)
func numIslands(grid [][]byte) int {
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != water {
				dfs(grid, i, j)
				count++
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 ||
		col < 0 ||
		row > len(grid)-1 ||
		col > len(grid[row])-1 ||
		grid[row][col] == water {
		return
	}

	grid[row][col] = water

	// r
	dfs(grid, row, col+1)
	// d
	dfs(grid, row+1, col)
	// l
	dfs(grid, row, col-1)
	// u
	dfs(grid, row-1, col)
}

package number_of_islands_dfs

func numIslands(grid [][]byte) int {
	var result int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 49 {
				dfs(i, j, grid)
				result++
			}
		}
	}
	return result
}

func dfs(i, j int, grid[][]byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 48 {
		return
	}
	grid[i][j] = 48

	dfs(i-1, j, grid) // up
	dfs(i+1, j, grid) // down
	dfs(i, j-1, grid) // left
	dfs(i, j+1, grid) // right
}
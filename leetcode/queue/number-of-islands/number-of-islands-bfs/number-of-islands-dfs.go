package number_of_islands_dfs

func numIslands(grid [][]byte) int {
	var result int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 49 {
				bfs(grid, i, j)
				result++
			}
		}
	}

	return result
}

func bfs(grid [][]byte, r, c int) {
	queue := make([][2]int, 0, 1)

	queue = append(queue, [2]int{r, c})

	grid[r][c] = 48

	offsetsToNeighbours := [4][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		curr := queue[0]

		queue = queue[1:]

		for _, offset := range offsetsToNeighbours {
			x := curr[0] + offset[0]
			y := curr[1] + offset[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == 49 {
				queue = append(queue, [2]int{x, y})
				grid[x][y] = 48
			}
		}
	}
}

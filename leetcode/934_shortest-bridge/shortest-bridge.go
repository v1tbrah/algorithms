package shortest_bridge

// Task: https://leetcode.com/problems/shortest-bridge/

// Solution: https://leetcode.com/problems/shortest-bridge/submissions/1068037723/

// tags:
// #Matrix
// #Breadth-First Search
// #Depth-First Search

// Time: O(m*n)
// Space: O(m*n)
func shortestBridge(grid [][]int) int {
	var firstFilled bool
	var bfsQueue [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				dfs(grid, i, j, &bfsQueue)
				firstFilled = true
				break
			}
		}
		if firstFilled {
			break
		}
	}

	var distance int
	for len(bfsQueue) > 0 {
		var newBfsQueue [][2]int
		for _, coord := range bfsQueue {
			x, y := coord[0], coord[1]
			neighbours := [4][2]int{
				{x, y + 1},
				{x + 1, y},
				{x, y - 1},
				{x - 1, y},
			}
			for _, neighbour := range neighbours {
				nX, nY := neighbour[0], neighbour[1]
				if nX < 0 || nY < 0 || nX > len(grid)-1 || nY > len(grid[nX])-1 {
					continue
				}
				if grid[nX][nY] == 1 {
					return distance
				}
				if grid[nX][nY] == 0 {
					newBfsQueue = append(newBfsQueue, neighbour)
					grid[nX][nY] = -1
				}
			}
		}

		bfsQueue = newBfsQueue
		distance++
	}

	return distance
}

func dfs(grid [][]int, i, j int, bfsQueuePtr *[][2]int) {
	if i < 0 ||
		j < 0 ||
		i > len(grid)-1 ||
		j > len(grid[i])-1 ||
		grid[i][j] != 1 {
		return
	}

	*bfsQueuePtr = append(*bfsQueuePtr, [2]int{i, j})

	grid[i][j] = 2

	// r
	dfs(grid, i, j+1, bfsQueuePtr)
	// d
	dfs(grid, i+1, j, bfsQueuePtr)
	// l
	dfs(grid, i, j-1, bfsQueuePtr)
	// u
	dfs(grid, i-1, j, bfsQueuePtr)
}

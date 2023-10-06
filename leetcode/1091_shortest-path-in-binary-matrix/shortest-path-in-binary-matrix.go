package task1091

// Task: https://leetcode.com/problems/shortest-path-in-binary-matrix/

// Solution: https://leetcode.com/problems/shortest-path-in-binary-matrix/submissions/1068838531/

// tags:
// #Array
// #Breadth-First Search
// #Matrix

// Time: O(m*n)
// Space: O(m*n)
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
		return -1
	}
	if len(grid) == 1 && grid[0][0] == 0 {
		return 1
	}

	pathLen := 1
	neighbours := getNeighbours(0, 0)

	// [
	// [[0,0,0]
	// [[1,1,0]
	// [[1,1,1]
	// ]
	used := make(map[[2]int]bool)
	for len(neighbours) > 0 {
		newNeighbours := make([][]int, 0)
		for _, pair := range neighbours {
			nR, nC := pair[0], pair[1]
			if nR == len(grid)-1 && nC == len(grid)-1 {
				return pathLen + 1
			}

			toCheck := [2]int{nR, nC}
			if used[toCheck] {
				continue
			} else {
				used[toCheck] = true
			}

			if nR < 0 || nC < 0 || nR > len(grid)-1 || nC > len(grid[nR])-1 || grid[nR][nC] == 1 {
				continue
			}
			newNeighbours = append(newNeighbours, getNeighbours(nR, nC)...)
		}

		pathLen++
		neighbours = newNeighbours
	}

	return -1
}

func getNeighbours(i, j int) [][]int {
	return [][]int{
		{i, j + 1},     // r
		{i + 1, j},     // d
		{i, j - 1},     // l
		{i - 1, j},     // u
		{i - 1, j + 1}, // ru
		{i - 1, j - 1}, // lu
		{i + 1, j + 1}, // rd
		{i + 1, j - 1}, // ld
	}
}

package task542

// Task: https://leetcode.com/problems/01-matrix/

// Solution: https://leetcode.com/problems/01-matrix/submissions/1068730201/

// tags:
// #Array
// #Breadth-First Search
// #Matrix

// Time: O(m*n)
// Space: O(m*n)
func updateMatrix(mat [][]int) [][]int {
	result := make([][]int, len(mat))

	neighbours := make([][]int, 0)
	checked := make(map[[2]int]bool, len(mat)*len(mat[0]))
	for i := 0; i < len(mat); i++ {
		result[i] = make([]int, len(mat[i]))
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 0 {
				checked[[2]int{i, j}] = true
				neighbours = append(neighbours, getNeighbours(i, j)...)
			}
		}
	}

	currDist := 1

	for len(neighbours) > 0 {
		newNeighbours := make([][]int, 0)

		for _, pair := range neighbours {
			r, c := pair[0], pair[1]
			if r < 0 || c < 0 || r > len(mat)-1 || c > len(mat[r])-1 || checked[[2]int{r, c}] {
				continue
			}
			checked[[2]int{r, c}] = true
			result[r][c] = currDist
			newNeighbours = append(newNeighbours, getNeighbours(r, c)...)
		}

		currDist++
		neighbours = newNeighbours
	}

	return result
}

func getNeighbours(i, j int) [][]int {
	return [][]int{
		{i, j + 1},
		{i + 1, j},
		{i, j - 1},
		{i - 1, j},
	}
}

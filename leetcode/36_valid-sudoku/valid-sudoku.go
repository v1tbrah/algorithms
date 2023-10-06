package valid_sudoku

// Task: https://leetcode.com/problems/valid-sudoku/

// Solution: https://leetcode.com/problems/valid-sudoku/submissions/1000397052/

// tags:
// #Array
// #Hash Table
// #Matrix

// Time: O(m*n)
// Space: O(m*n)
func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		row := make(map[byte]struct{})
		for j := 0; j < len(board); j++ {
			if board[i][j] != '.' {
				if _, ok := row[board[i][j]]; !ok {
					row[board[i][j]] = struct{}{}
				} else {
					return false
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		col := make(map[byte]struct{})
		for j := 0; j < len(board); j++ {
			if board[j][i] != '.' {
				if _, ok := col[board[j][i]]; !ok {
					col[board[j][i]] = struct{}{}
				} else {
					return false
				}
			}
		}
	}

	r, c := 0, 0
	for r != 9 {
		unique := make(map[byte]struct{})
		for i := r; i < 3+r; i++ {
			for j := c; j < 3+c; j++ {
				if board[i][j] != '.' {
					if _, ok := unique[board[i][j]]; !ok {
						unique[board[i][j]] = struct{}{}
					} else {
						return false
					}
				}
			}
		}

		if c == 0 || c == 3 {
			c += 3
			continue
		}
		if c == 6 {
			r += 3
			c = 0
			continue
		}
	}

	return true
}

package valid_sudoku

// Task: https://leetcode.com/problems/valid-sudoku/

// Solution: https://leetcode.com/problems/valid-sudoku/submissions/1000397052/

// tags:
// #Array
// #Hash Table
// #Matrix

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		currCheckMap := make(map[byte]struct{}, 9)
		for j := 0; j < len(board[i]); j++ {
			if _, ok := currCheckMap[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' { // not point
				currCheckMap[board[i][j]] = struct{}{}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		currCheckMap := make(map[byte]struct{}, 9)
		for j := 0; j < len(board[i]); j++ {
			if _, ok := currCheckMap[board[j][i]]; ok {
				return false
			} else if board[j][i] != '.' { // not point
				currCheckMap[board[j][i]] = struct{}{}
			}
		}
	}

	vShift, hShift := 0, 0

	for {
		currCheckMap := make(map[byte]struct{}, 9)
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if _, ok := currCheckMap[board[i+vShift][j+hShift]]; ok {
					return false
				} else if board[i+vShift][j+hShift] != '.' { // not point
					currCheckMap[board[i+vShift][j+hShift]] = struct{}{}
				}
			}
		}

		if vShift == 6 && hShift == 6 {
			break
		}

		if hShift == 0 || hShift == 3 {
			hShift += 3
		} else {
			vShift += 3
			hShift = 0
		}
	}

	return true
}

package flood_fill

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	rows, cols := len(image), len(image[0])
	flooded := map[string]struct{}{}
	target := image[sr][sc]
	floodAndCheckNeighbors(image, flooded, sr, sc, target, color, rows, cols)
	return image
}

func checkNeighbors(image [][]int, flooded map[string]struct{}, rowIdx, colIdx, target, color, rows, cols int) {
	if haveRight := colIdx+1 <= cols-1; haveRight {
		currRowIdx := rowIdx
		currColIdx := colIdx + 1
		floodAndCheckNeighbors(image, flooded, currRowIdx, currColIdx, target, color, rows, cols)
	}
	if haveLeft := colIdx != 0; haveLeft {
		currRowIdx := rowIdx
		currColIdx := colIdx - 1
		floodAndCheckNeighbors(image, flooded, currRowIdx, currColIdx, target, color, rows, cols)

	}
	if haveUp := rowIdx != 0; haveUp {
		currRowIdx := rowIdx - 1
		currColIdx := colIdx
		floodAndCheckNeighbors(image, flooded, currRowIdx, currColIdx, target, color, rows, cols)
	}
	if haveDown := rowIdx+1 <= rows-1; haveDown {
		currRowIdx := rowIdx + 1
		currColIdx := colIdx
		floodAndCheckNeighbors(image, flooded, currRowIdx, currColIdx, target, color, rows, cols)
	}
}

func floodAndCheckNeighbors(image [][]int, flooded map[string]struct{}, rowIdx, colIdx, target, color, rows, cols int) {
	curr := image[rowIdx][colIdx]
	if curr == target {
		if _, isFlooded := flooded[fmt.Sprintf("%d:%d", rowIdx, colIdx)]; !isFlooded {
			image[rowIdx][colIdx] = color
			flooded[fmt.Sprintf("%d:%d", rowIdx, colIdx)] = struct{}{}
			checkNeighbors(image, flooded, rowIdx, colIdx, target, color, rows, cols)
		}
	}
}

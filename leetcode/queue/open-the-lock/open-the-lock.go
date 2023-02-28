package open_the_lock

import (
	"strings"
)

var builder strings.Builder

func openLock(deadends []string, target string) int {
	start := "0000"
	if target == start {
		return 0
	}

	deadendsMap := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsMap[deadend] = true
	}
	if deadendsMap[start] {
		return -1
	}

	var counter int
	queue := []string{start}
	visited := map[string]bool{start: true}

	for len(queue) != 0 {
		counter++
		for _, code := range queue {
			queue = queue[1:]
			for _, neighbour := range neighbours(code) {
				if neighbour == target {
					return counter
				}

				if !deadendsMap[neighbour] && !visited[neighbour] {
					visited[neighbour] = true
					queue = append(queue, neighbour)
				}
			}
		}
	}

	return -1
}

func neighbours(code string) (result [8]string) {
	if len(code) != 4 {
		panic("unexpected string")
	}

	var iter int
	for i := 0; i < 4; i++ {
		builder.Grow(4)
		for j := 0; j < 4; j++ {
			if j == i {
				builder.WriteByte(next(code[i]))
			} else {
				builder.WriteByte(code[j])
			}
		}
		result[iter] = builder.String()

		iter++
		builder.Reset()

		builder.Grow(4)
		for j := 0; j < 4; j++ {
			if j == i {
				builder.WriteByte(prev(code[i]))
			} else {
				builder.WriteByte(code[j])
			}
		}
		result[iter] = builder.String()

		iter++
		builder.Reset()
	}

	return result
}

func next(curr byte) byte {
	switch curr {
	case '9':
		curr = '0'
	default:
		curr++
	}

	return curr
}

func prev(curr byte) byte {
	switch curr {
	case '0':
		curr = '9'
	default:
		curr--
	}

	return curr
}

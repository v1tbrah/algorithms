package bfs

func BFS(graph map[string][]string, head string, isDesired func(tested string) bool) bool {
	if isDesired(head) {
		return true
	}
	searched := map[string]struct{}{head: {}}

	searchQueue := make([]string, 0)
	searchQueue = append(searchQueue, graph[head]...)

	for len(searchQueue) > 0 {
		currNode := searchQueue[0]
		if isDesired(currNode) {
			return true
		}

		// shift first element in searchQueue
		searchQueue = searchQueue[1:]

		searched[currNode] = struct{}{}

		for _, node := range graph[currNode] {
			if _, isSearched := searched[node]; isSearched {
				continue
			}
			searchQueue = append(searchQueue, node)
		}
	}
	return false
}

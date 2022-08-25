package dfs

func DFS(graph map[string][]string, searched map[string]struct{}, thisNode string, isDesired func(tested string) bool) bool {
	if isDesired(thisNode) {
		return true
	}
	searched[thisNode] = struct{}{}

	for _, currNode := range graph[thisNode] {
		if _, isSearched := searched[currNode]; isSearched {
			continue
		}
		if isTarget := DFS(graph, searched, currNode, isDesired); isTarget {
			return true
		}
	}

	return false
}

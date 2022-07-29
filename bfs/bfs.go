package bfs

type graph map[string][]string

func BFS(friends graph, isDesired func(verifiable string) bool) bool {
	searchQueue := []string{}
	searchQueue = append(searchQueue, friends["Vitya"]...)
	searched := []string{}
	for len(searchQueue) > 0 {
		if isDesired(searchQueue[0]) {
			return true
		} else {
			searched = append(searched, searchQueue[0])
			searchQueue = append(searchQueue[1:], friends[searchQueue[0]]...)
		}
	}
	return false
}

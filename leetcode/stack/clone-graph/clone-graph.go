package clone_graph

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	srcToCopy := make(map[*Node]*Node)

	dfs(node, srcToCopy)

	return srcToCopy[node]
}

func dfs(node *Node, srcToCopy map[*Node]*Node) {
	if node == nil {
		return
	}

	copyNode := &Node{Val: node.Val}

	srcToCopy[node] = copyNode

	for _, neighbor := range node.Neighbors {
		if _, copyIsExists := srcToCopy[neighbor]; !copyIsExists {
			dfs(neighbor, srcToCopy)
		}

		copyNode.Neighbors = append(copyNode.Neighbors, srcToCopy[neighbor])
	}
}

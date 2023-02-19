package lowest_common_ancestor_of_a_binary_tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var pPath, qPath []*TreeNode
	findPathToTarget(root, p, &pPath)
	findPathToTarget(root, q, &qPath)

	var idx int
	for ; idx < len(pPath) && idx < len(qPath); idx++ {
		if pPath[idx] != qPath[idx] {
			return pPath[idx-1]
		}
	}

	return pPath[idx-1]
}

func findPathToTarget(node, target *TreeNode, path *[]*TreeNode) bool {
	if node == nil {
		return false
	}

	*path = append(*path, node)

	if node.Val == target.Val ||
		findPathToTarget(node.Left, target, path) ||
		findPathToTarget(node.Right, target, path) {
		return true
	}
	temp := *path
	temp = temp[:len(temp)-1]
	*path = temp

	return false
}

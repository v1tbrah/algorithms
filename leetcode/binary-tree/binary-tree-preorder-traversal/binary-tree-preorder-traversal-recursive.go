package binary_tree_preorder_traversal

// recursive solution
func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 1)
	result[0] = root.Val

	dfs(root.Left, &result)
	dfs(root.Right, &result)

	return result
}

func dfs(node *TreeNode, resultPtr *[]int) {
	if node == nil {
		return
	}

	*resultPtr = append(*resultPtr, node.Val)

	dfs(node.Left, resultPtr)
	dfs(node.Right, resultPtr)
}

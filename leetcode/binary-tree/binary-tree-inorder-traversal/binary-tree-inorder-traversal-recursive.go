package binary_tree_inorder_traversal

// recursive solution
func inorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0, 1)
	dfs(root, &result)

	return result
}

func dfs(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, result)
	*result = append(*result, root.Val)
	dfs(root.Right, result)
}

package binary_tree_postorder_traversal

// recursive solution
func postorderTraversalRecursive(root *TreeNode) []int {
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

	if root.Left != nil {
		dfs(root.Left, result)
	}

	if root.Right != nil {
		dfs(root.Right, result)
	}

	*result = append(*result, root.Val)
}

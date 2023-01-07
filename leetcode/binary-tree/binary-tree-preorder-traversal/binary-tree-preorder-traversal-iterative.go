package binary_tree_preorder_traversal

// iterative solution
func preorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := make([]*TreeNode, 1)
	stack[0] = root

	result := make([]int, 0, 1)

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root != nil {
			result = append(result, root.Val)
			stack = append(stack, root.Right)
			stack = append(stack, root.Left)
		}
	}

	return result
}

package binary_tree_inorder_traversal

// iterative solution
func inorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	wasThere := make(map[*TreeNode]bool)

	result := make([]int, 0, 1)

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		curr := stack[len(stack)-1]

		if curr.Left != nil && !wasThere[curr.Left] {
			stack = append(stack, curr.Left)
			wasThere[curr.Left] = true
			continue
		}

		result = append(result, curr.Val)
		stack = stack[:len(stack)-1]

		if curr.Right != nil {
			stack = append(stack, curr.Right)
			continue
		}
	}

	return result
}

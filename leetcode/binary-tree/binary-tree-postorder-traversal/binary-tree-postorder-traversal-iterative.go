package binary_tree_postorder_traversal

// iterative solution
func postorderTraversalIterative(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0, 1)

	wasThere := make(map[*TreeNode]bool)
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]

		if curr.Left != nil && !wasThere[curr.Left] {
			stack = append(stack, curr.Left)
			wasThere[curr.Left] = true
			continue
		}

		if curr.Right != nil && !wasThere[curr.Right] {
			stack = append(stack, curr.Right)
			wasThere[curr.Right] = true
			continue
		}

		result = append(result, curr.Val)
		stack = stack[:len(stack)-1]
	}

	return result
}

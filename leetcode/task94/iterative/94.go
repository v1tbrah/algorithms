package task94

// Task: https://leetcode.com/problems/binary-tree-inorder-traversal/

// Solution: https://leetcode.com/problems/binary-tree-inorder-traversal/submissions/1074966680/

// tags:
// #Binary Tree
// #Stack
// #Tree

// Desc: Binary tree inorder traversal - is than we add 1) left, 2) root, 3) right

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(N), where N - is count of nodes
// Space: O(N), cause we need place to save stack
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// left -> root -> right
	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	used := make(map[*TreeNode]bool)

	stack = append(stack, root)
	for len(stack) != 0 {
		top := stack[len(stack)-1]

		if top.Left != nil && !used[top.Left] {
			stack = append(stack, top.Left)
			continue
		}

		result = append(result, top.Val)
		stack = stack[:len(stack)-1]
		used[top] = true

		if top.Right != nil && !used[top.Right] {
			stack = append(stack, top.Right)
			continue
		}
	}

	return result
}

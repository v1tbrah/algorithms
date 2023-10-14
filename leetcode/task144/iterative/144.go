package task144

// Task: https://leetcode.com/problems/binary-tree-preorder-traversal/

// Solution: https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/1074910668/

// tags:
// #Binary Tree
// #Stack
// #Tree

// Desc: Binary tree preorder traversal - is than we add 1) root, 2) left, 3) right

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(N), where N - is count of nodes
// Space: O(N), cause we need place to save stack
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, top.Val)

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
		if top.Left != nil {
			stack = append(stack, top.Left)
		}
	}

	return result
}

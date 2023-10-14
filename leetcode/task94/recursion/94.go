package task94

// Task: https://leetcode.com/problems/binary-tree-inorder-traversal/

// Solution: https://leetcode.com/problems/binary-tree-inorder-traversal/submissions/1074961216/

// tags:
// #Binary Tree
// #Depth-First Search
// #Tree

// Desc: Binary tree inorder traversal - is than we add 1) left, 2) root, 3) right

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(N), where N - is count of nodes
// Space: O(N), cause we need place to save recursion stack
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 0)
	dfs(root.Left, &result)
	result = append(result, root.Val)
	dfs(root.Right, &result)

	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, result)
	*result = append(*result, node.Val)
	dfs(node.Right, result)
}

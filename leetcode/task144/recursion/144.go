package task144

// Task: https://leetcode.com/problems/binary-tree-preorder-traversal/

// Solution: https://leetcode.com/problems/binary-tree-preorder-traversal/submissions/1074906175/

// tags:
// #Binary Tree
// #Depth-First Search
// #Tree

// Desc: Binary tree preorder traversal - is than we add 1) root, 2) left, 3) right

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(N), where N - is count of nodes
// Space: O(N), cause we need place to save recursion stack
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}
	dfs(root.Left, &result)
	dfs(root.Right, &result)

	return result
}

func dfs(node *TreeNode, toSave *[]int) {
	if node == nil {
		return
	}

	*toSave = append(*toSave, node.Val)
	dfs(node.Left, toSave)
	dfs(node.Right, toSave)
}

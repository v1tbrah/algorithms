package maximum_depth_of_binary_tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1
	right := 1
	if root.Left != nil {
		left = 1 + maxDepth(root.Left)

	}
	if root.Right != nil {
		right = 1 + maxDepth(root.Right)
	}
	if left > right {
		return left
	}
	return right
}

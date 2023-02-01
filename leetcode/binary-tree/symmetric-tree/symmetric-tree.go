package symmetric_tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSym(root.Left, root.Right)
}

func isSym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	isSymRes := isSym(left.Left, right.Right)
	if !isSymRes {
		return false
	}
	isSymRes = isSym(left.Right, right.Left)
	if !isSymRes {
		return false
	}
	return true
}

package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := make([]int, 1)
	result[0] = root.Val

	preorderTraversalRecursive(root.Left, &result)
	preorderTraversalRecursive(root.Right, &result)

	return result
}

func preorderTraversalRecursive(node *TreeNode, resultPtr *[]int) {
	if node == nil {
		return
	}

	*resultPtr = append(*resultPtr, node.Val)

	preorderTraversalRecursive(node.Left, resultPtr)
	preorderTraversalRecursive(node.Right, resultPtr)
}

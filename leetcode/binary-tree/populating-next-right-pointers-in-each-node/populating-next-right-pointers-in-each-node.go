package populating_next_right_pointers_in_each_node

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	recConnect(root.Left, root.Right, nil)

	return root
}

func recConnect(left, right, nextL *Node) {
	if left == nil || right == nil {
		return
	}

	left.Next = right
	right.Next = nextL

	var nextR *Node
	if nextL != nil {
		nextR = nextL.Left
	}

	recConnect(left.Left, left.Right, right.Left)
	recConnect(right.Left, right.Right, nextR)
}

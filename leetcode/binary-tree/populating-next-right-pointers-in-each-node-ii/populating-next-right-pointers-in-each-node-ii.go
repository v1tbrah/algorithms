package populating_next_right_pointers_in_each_node_ii

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		thisLvlLen := len(queue)
		for i := 0; i < thisLvlLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if isLastOnThisLvl := i == thisLvlLen-1; !isLastOnThisLvl {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

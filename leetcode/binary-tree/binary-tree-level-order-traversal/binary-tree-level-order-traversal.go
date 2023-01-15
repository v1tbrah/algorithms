package binary_tree_level_order_traversal

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0, 1)
	result = append(result, []int{root.Val})

	resMap := make(map[int]*[]int)

	levelOrderRecursive(root.Left, root.Right, 1, resMap)

	for i := 1; i <= len(resMap); i++ {
		values := resMap[i]
		result = append(result, *values)
	}

	return result
}

func levelOrderRecursive(left, right *TreeNode, lvl int, lvlToNodes map[int]*[]int) {
	if left != nil {
		nodesOfLvl := lvlToNodes[lvl]
		if nodesOfLvl == nil {
			nodesOfLvl = &[]int{}
		}
		*nodesOfLvl = append(*nodesOfLvl, left.Val)
		lvlToNodes[lvl] = nodesOfLvl
	}

	if right != nil {
		nodesOfLvl := lvlToNodes[lvl]
		if nodesOfLvl == nil {
			nodesOfLvl = &[]int{}
		}
		*nodesOfLvl = append(*nodesOfLvl, right.Val)
		lvlToNodes[lvl] = nodesOfLvl
	}

	if left == nil && right == nil {
		return
	}

	newLvl := lvl + 1
	if left != nil {
		levelOrderRecursive(left.Left, left.Right, newLvl, lvlToNodes)
	}
	if right != nil {
		levelOrderRecursive(right.Left, right.Right, newLvl, lvlToNodes)
	}
}

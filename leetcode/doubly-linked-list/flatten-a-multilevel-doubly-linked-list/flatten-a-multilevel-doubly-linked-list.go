package flatten_a_multilevel_doubly_linked_list

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	head := root

	flattenRecursive(root)

	return head
}

func flattenRecursive(root *Node) *Node {
	if root == nil {
		return root
	}

	next := root.Next
	if root.Child != nil {
		root.Next = root.Child
		root.Child = nil
		if root.Next != nil {
			root.Next.Prev = root
		}
		root = flattenRecursive(root.Next)
	}

	if next != nil {
		root.Next = next // 2
		if root.Next != nil {
			root.Next.Prev = root
		}
		root = flattenRecursive(root.Next)
	}

	return root
}

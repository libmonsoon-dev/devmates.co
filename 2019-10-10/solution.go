package solution

func Solution(root *Node) *Node {
	start := root

	for root != nil && root.Left != nil {
		next := root.Left
		for root != nil {
			root.Left.Next = root.Right
			if root.Next != nil {
				root.Right.Next = root.Next.Left
			} else {
				root.Right.Next = nil
			}
			root = root.Next
		}
		root = next
	}

	return start
}

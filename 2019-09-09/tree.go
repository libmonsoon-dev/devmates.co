package binary_search_tree

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func (n *Node) AddLeft(child *Node) {
	n.Left = child
	child.Parent = n
}

func (n *Node) AddRight(child *Node) {
	n.Right = child
	child.Parent = n
}

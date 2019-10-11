package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func CreateNode(value int) *Node {
	return &Node{
		Val: value,
	}

}

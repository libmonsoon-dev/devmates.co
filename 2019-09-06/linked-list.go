package linked_list

import (
	"fmt"
	"strconv"
)

func New(rootValue int, values ...int) *Node {
	root := &Node{rootValue, nil}
	current := root

	for _, value := range values {
		current = current.AppendInt(value)
	}

	return root
}

type Node struct {
	Value int
	Next  *Node
}

func (node *Node) GoString() string {
	if node.Next == nil {
		return strconv.Itoa(node.Value)
	}
	return fmt.Sprintf("%v->%v", node.Value, node.Next.GoString())
}

func (node *Node) Append(next *Node) *Node {
	node.Next = next
	return next
}

func (node *Node) AppendInt(val int) *Node {
	return node.Append(&Node{val, nil})
}

func Merge(input []*Node) *Node {
	rootIndex, _ := getSmallerNodeIndex(input)
	root := input[rootIndex]
	current := root

	for {
		i, ok := getSmallerNodeIndex(input)

		if !ok {
			break
		}
		node := input[i]
		input[i] = node.Next
		current = current.Append(node)
	}

	return root
}

func getSmallerNodeIndex(input []*Node) (int, bool) {
	gotValue := false
	var smallest int
	var index int

	for i, node := range input {
		if node == nil {
			continue
		}
		if !gotValue {
			index = i
			smallest = node.Value
			gotValue = true
			continue
		}
		if node.Value < smallest {
			smallest = node.Value
			index = i
		}

	}

	return index, gotValue
}

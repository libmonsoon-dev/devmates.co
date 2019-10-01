package solution

type listNode struct {
	Value int
	Next  *listNode
}

func (ln *listNode) Empty() bool {
	return ln == nil
}

func listFromSlice(first int, slice ...int) *listNode {
	root := &listNode{Value: first}
	current := root

	for _, val := range slice {
		next := &listNode{Value: val}
		current.Next, current = next, next

	}
	return root
}

func sliceFromList(node *listNode) []int {
	result := make([]int, 0)
	for node != nil {
		result = append(result, node.Value)
		node = node.Next
	}
	return result
}

func Solution(_a, _b []int) []int {
	a, b := listFromSlice(_a[0], _a[1:]...), listFromSlice(_b[0], _b[1:]...)

	var prev *listNode
	var start *listNode
	carry := 0

	for !a.Empty() || !b.Empty() {
		sum := 0

		if !a.Empty() {
			sum = sum + a.Value
			a = a.Next
		}

		if !b.Empty() {
			sum = sum + b.Value
			b = b.Next
		}

		current := sum + carry
		newNode := &listNode{Value: current % 10}
		carry = current / 10

		if start == nil {
			start = newNode
		}

		if prev != nil {
			prev.Next = newNode
		}

		prev = newNode

	}

	if carry != 0 {
		newNode := &listNode{Value: carry}
		prev.Next = newNode
	}

	return sliceFromList(start)
}

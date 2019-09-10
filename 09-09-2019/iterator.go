package binary_search_tree

type Iterator struct {
	next *Node
}

func (i *Iterator) HasNext() bool {
	return i.next != nil
}

func (i *Iterator) Next() int {
	r := i.next
	// If you can walk right, walk right, then fully left.
	// otherwise, walk up until you come from left.
	if i.next.Right != nil {
		i.next = i.next.Right
		for i.next.Left != nil {
			i.next = i.next.Left
		}
		return r.Value
	}

	for {
		if i.next.Parent == nil {
			i.next = nil
			return r.Value
		}
		if i.next.Parent.Left == i.next {
			i.next = i.next.Parent
			return r.Value
		}
		i.next = i.next.Parent
	}
}

func NewIterator(root *Node) *Iterator {
	next := root
	for next.Left != nil {
		next = next.Left
	}
	return &Iterator{next: next}
}

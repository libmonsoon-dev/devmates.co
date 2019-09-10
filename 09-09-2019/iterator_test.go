package binary_search_tree

import (
	"testing"

	"gotest.tools/assert"
)

func TestIterator(t *testing.T) {
	root := &Node{Value: 7}
	_3 := &Node{Value: 3}
	_15 := &Node{Value: 15}
	_9 := &Node{Value: 9}
	_20 := &Node{Value: 20}

	_15.AddLeft(_9)
	_15.AddRight(_20)

	root.AddLeft(_3)
	root.AddRight(_15)

	i := NewIterator(root)

	assertEqual := func(a, b interface{}) {
		assert.Equal(t, a, b)
	}

	assertEqual(i.Next(), 3)
	assertEqual(i.Next(), 7)
	assertEqual(i.HasNext(), true)
	assertEqual(i.Next(), 9)
	assertEqual(i.HasNext(), true)
	assertEqual(i.Next(), 15)
	assertEqual(i.HasNext(), true)
	assertEqual(i.Next(), 20)
	assertEqual(i.HasNext(), false)
}

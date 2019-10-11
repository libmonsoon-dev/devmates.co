package solution

import "testing"

func treeEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil {
		return a.Val == b.Val &&
			treeEqual(a.Next, b.Next) &&
			treeEqual(a.Left, b.Left) &&
			treeEqual(a.Right, b.Right)
	}

	return false
}

func TestSolution(t *testing.T) {
	test := func(input, expect *Node) {
		if actual := Solution(input); !treeEqual(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	input := CreateNode(1)

	input.Left = CreateNode(2)
	input.Right = CreateNode(3)

	input.Left.Left = CreateNode(4)
	input.Left.Right = CreateNode(5)

	input.Right.Left = CreateNode(6)
	input.Right.Right = CreateNode(7)

	expect := CreateNode(1)

	expect.Left = CreateNode(2)
	expect.Right = CreateNode(3)

	expect.Left.Left = CreateNode(4)
	expect.Left.Right = CreateNode(5)

	expect.Right.Left = CreateNode(6)
	expect.Right.Right = CreateNode(7)

	expect.Left.Next = expect.Right             // 2 -> 3
	expect.Left.Left.Next = expect.Left.Right   // 4 -> 5
	expect.Left.Right.Next = expect.Right.Left  // 5 -> 6
	expect.Right.Left.Next = expect.Right.Right // 6 -> 7

	test(input, expect)
}

package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input []*int, expect int) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]*int{Ptr(1), Ptr(2), Ptr(3)}, 6)
	test([]*int{Ptr(-10), Ptr(9), Ptr(20), nil, nil, Ptr(15), Ptr(7)}, 42)
}

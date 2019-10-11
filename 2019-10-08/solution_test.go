package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input [][]int, expect int) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(
		[][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
		7,
	)

}

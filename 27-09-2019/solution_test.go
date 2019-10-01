package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(matrix [][]int, expect int) {
		if actual := Solution(matrix); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(
		[][]int{
			{1, 0, 1, 0, 0},
			{1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1},
			{1, 0, 0, 1, 0},
		},
		4,
	)
}

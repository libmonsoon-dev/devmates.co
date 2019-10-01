package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input []int, target int, expect [2]int) {
		if actual := Solution(input, target); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]int{3, 4, 7, 7, 9, 9, 9, 10}, 9, [2]int{4, 6})
	test([]int{3, 4, 7, 7, 9, 9, 9, 10}, 8, [2]int{-1, -1})
}

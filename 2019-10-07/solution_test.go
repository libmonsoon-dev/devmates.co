package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	equal := func(a, b [][]*int) bool {
		if len(a) != len(b) {
			return false
		}

		for i := range a {
			if len(a[i]) != len(b[i]) {
				return false
			}

			for j := range a[i] {
				if a[i][j] == nil && b[i][j] == nil {
					continue
				}

				if a[i][j] == nil || b[i][j] == nil {
					return false
				}

				if *a[i][j] != *b[i][j] {
					return false
				}
			}
		}
		return true
	}

	test := func(input, toDelete []int, expect [][]*int) {
		if actual := Solution(input, toDelete); !equal(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(
		[]int{1, 2, 3, 4, 5, 6, 7},
		[]int{3, 5},
		[][]*int{
			{ptrOf(1), ptrOf(2), nil, ptrOf(4)},
			{ptrOf(6)},
			{ptrOf(7)},
		},
	)
}

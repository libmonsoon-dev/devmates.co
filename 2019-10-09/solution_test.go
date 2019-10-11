package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(m, n, expect int) {
		if actual := Solution(m, n); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(3, 2, 3)
}

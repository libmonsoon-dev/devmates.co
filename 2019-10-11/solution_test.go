package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input []string, expect int) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]string{"23:29", "00:00", "23:11"}, 18)

}

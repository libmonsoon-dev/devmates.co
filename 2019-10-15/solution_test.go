package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input int, expect string) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(3, "III")
	test(4, "IV")
	test(9, "IX")
	test(58, "LVIII")
	test(1999, "MCMXCIX")
}

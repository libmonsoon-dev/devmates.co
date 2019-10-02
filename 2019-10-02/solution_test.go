package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input string, expect int) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test("III", 3)
	test("IV", 4)
	test("IX", 9)
	test("LVIII", 58)
	test("MCMXCIV", 1994)

}

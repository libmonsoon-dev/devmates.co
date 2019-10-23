package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(input, expect string) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test("aaabbbc", "c")
	test("aabbbacd", "cd")
	test("aabbccddeeedcba", "")
	test("aabbccddeeedcba", "")
}

package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(input, expect string) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test("hello", "holle")
	test("leetcode", "leotcede")

}

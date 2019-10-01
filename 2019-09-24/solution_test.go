package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(str, expect string) {
		if actual := Solution(str); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test("hello this is amadeo", "olleh siht si oedama")
}

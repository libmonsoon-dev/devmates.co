package solution

import (
	"testing"

	"go.uber.org/goleak"
)

func TestSolution(t *testing.T) {
	test := func(list []int, target, expect int) {
		defer goleak.VerifyNoLeaks(t) //VerifyNone in master

		if actual := Solution(list, target); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]int{1, 1, 1, 1, 1}, 3, 5)
}

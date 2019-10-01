package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(n, k int, expect [][]int) {
		if actual := Solution(n, k); !reflect.DeepEqual(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}
	test(
		4,
		2,
		[][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		},
	)
}

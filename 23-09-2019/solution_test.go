package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(a, b, expect []int) {
		if actual := Solution(a, b); !reflect.DeepEqual(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]int{6, 5, 4}, []int{8, 7, 3}, []int{4, 3, 8})

}

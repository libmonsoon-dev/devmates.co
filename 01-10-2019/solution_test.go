package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(input, expect []int) {
		if actual := Solution(input); !reflect.DeepEqual(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test([]int{1, 2, 3, 4}, []int{24, 12, 8, 6})

}

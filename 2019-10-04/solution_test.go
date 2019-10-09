package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	test := func(input []string, expect [][]string) {
		if actual := Solution(input); !reflect.DeepEqual(expect, actual) {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}

	test(
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		},
	)
}

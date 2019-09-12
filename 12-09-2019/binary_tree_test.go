package binary_tree

import (
	"reflect"
	"testing"
)

func TestGetRightValues(t *testing.T) {
	toPtr := func(i int) *int {
		return &i
	}

	input := []*int{toPtr(1), toPtr(2), toPtr(3), nil, toPtr(5), nil, toPtr(4)}
	expect := []int{1, 3, 4}

	actual := GetRightValues(input...)

	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%v != %v", expect, actual)
	}

}

package linked_list

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		args []*Node
		want *Node
	}{
		{
			[]*Node{
				New(1, 10, 20),
				New(4, 11, 13),
				New(3, 8, 9),
			},
			New(1, 3, 4, 8, 9, 10, 11, 13, 20),
		},
		{
			[]*Node{
				New(1, 4, 5),
				New(2, 6),
			},
			New(1, 2, 4, 5, 6),
		},
	}

	for _, test := range tests {
		if got := Merge(test.args); !reflect.DeepEqual(got, test.want) {
			t.Errorf("Merge() = %#v, want %#v", got, test.want)
		}
	}
}

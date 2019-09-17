package substring_counter

import (
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		Input  string
		Output int
	}{
		{"adddw", 2},
		{"abcdabcbb", 4},
		{"ccccccc", 1},
	}

	for _, test := range tests {
		actual := Count(test.Input)

		if test.Output != actual {
			t.Errorf("assertion failed: %v (test.Output %[1]T) != %v (actual %[2]T)", test.Output, actual)
		}
	}

}

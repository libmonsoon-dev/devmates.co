package counter

import "testing"

var tests = [...]struct {
	Input  [][]int
	Output int
}{
	{
		Input: [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		},
		Output: 6,
	},
	{
		Input: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
		Output: 0,
	},
}

func TestCount(t *testing.T) {
	for _, test := range tests {
		got := Count(test.Input)

		if got != test.Output {
			t.Errorf("Count(%#v) got = %#v, want %#v", test.Input, got, test.Output)
		}

	}

}

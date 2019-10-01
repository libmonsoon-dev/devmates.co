package solution

import "testing"

func TestSolution(t *testing.T) {
	tests := [...]struct {
		Tasks  []string
		N      int
		Output int
	}{
		{[]string{"A", "A", "A", "B", "B", "B"}, 2, 8},
	}

	for _, test := range tests {
		actual := Solution(test.Tasks, test.N)

		if test.Output != actual {
			t.Errorf("assertion failed: %v (test.Output %[1]T) != %v (actual %[2]T)", test.Output, actual)
		}
	}
}

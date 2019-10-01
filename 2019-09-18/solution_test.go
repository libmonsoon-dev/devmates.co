package solution

import (
	"testing"

	"gotest.tools/assert"
)

func TestSolution(t *testing.T) {
	want := permutations{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	}

	actual := Solution([]int{1, 1, 2})

	assert.DeepEqual(t, want, actual)

}

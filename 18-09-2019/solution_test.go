package solution

import (
	"testing"

	"gotest.tools/assert"
)

func TestSolution(t *testing.T) {
	want := permutations{
		permutation{1, 1, 2}: struct{}{},
		permutation{1, 2, 1}: struct{}{},
		permutation{2, 1, 1}: struct{}{},
	}

	actual := Solution(permutation{1, 1, 2})

	assert.DeepEqual(t, want, actual)

}

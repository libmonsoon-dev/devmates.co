package solution

import (
	"testing"

	"gotest.tools/assert"
)

var want = []string{
	"((()))",
	"(()())",
	"(())()",
	"()(())",
	"()()()",
}

var n = 3

func TestRecursiveSolution(t *testing.T) {

	actual := RecursiveSolution(n)

	assert.DeepEqual(t, want, actual)
}

func TestSolution(t *testing.T) {

	actual := Solution(n)

	assert.DeepEqual(t, want, actual)
}

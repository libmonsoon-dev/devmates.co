package fill

import (
	"testing"

	"gotest.tools/assert"
)

func TestFill(t *testing.T) {
	img := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}

	want := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}

	Fill(&img, 1, 1, 2)
	assert.DeepEqual(t, img, want)
}

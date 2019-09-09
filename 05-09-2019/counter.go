package counter

import "math"

func count(grid [][]int, i, j int) int {
	if i < 0 ||
		i >= len(grid) ||
		j < 0 ||
		j >= len(grid[i]) ||
		grid[i][j] != 1 {
		return 0
	}
	grid[i][j] = 0

	return 1 +
		count(grid, i+1, j) +
		count(grid, i-1, j) +
		count(grid, i, j+1) +
		count(grid, i, j-1)
}

func Count(grid [][]int) int {
	result := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			result = int(math.Max(
				float64(count(grid, i, j)),
				float64(result),
			))
		}

	}

	return result
}

package solution

import "math"

const placeholder = -1
const maxInt = int(^uint(0) >> 1)

func Solution(input [][]int) int {
	if len(input) == 0 || len(input[0]) == 0 {
		return 0
	}

	visited := make([][]int, 0, len(input))

	for i := range input {
		visited = append(visited, make([]int, len(input[i])))
		for j := range visited[i] {
			visited[i][j] = placeholder
		}
	}

	visited[0][0] = input[0][0]

	var dfs func(x, y int) int

	dfs = func(x, y int) int {
		if visited[x][y] != placeholder {
			return visited[x][y]
		}

		min := maxInt

		if x-1 >= 0 {
			min = int(math.Min(
				float64(min),
				float64(dfs(x-1, y)),
			))
		}

		if y-1 >= 0 {
			min = int(math.Min(
				float64(min),
				float64(dfs(x, y-1)),
			))
		}

		visited[x][y] = min + input[x][y]
		return visited[x][y]
	}

	return dfs(len(input)-1, len(input[0])-1)
}

package solution

func Solution(input []int) []int {
	helper := func(items []int) int {
		result := items[0]
		for _, item := range items[1:] {
			result *= item
		}
		return result
	}
	result := make([]int, len(input))

	for i := range result {
		before := input[0:i]
		after := input[i+1:]
		result[i] = helper(append(after, before...))
	}

	return result
}

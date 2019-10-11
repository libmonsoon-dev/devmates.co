package solution

func CreateSlice(from, to int) []int {
	lenght := to - from

	result := make([]int, lenght)

	for i := range result {
		result[i] = from + i
	}

	return result

}

func Solution(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	current := make([]int, n)

	for i := range current {
		current[i] = 1
	}

	for range CreateSlice(1, m) {
		for _, j := range CreateSlice(1, n) {
			current[j] += current[j-1]
		}

	}

	return current[len(current)-1]
}

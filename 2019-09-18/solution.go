package solution

type permutation []int
type permutations = []permutation

func Solution(input []int) permutations {
	arrLen := len(input)
	result := make(permutations, 0, arrLen)

	for i := range input {
		p := make(permutation, arrLen)
		for j := 0; j < arrLen; j++ {
			index := (i + j) % arrLen
			p[j] = input[index]
		}

		result = append(result, p)
	}

	return result
}

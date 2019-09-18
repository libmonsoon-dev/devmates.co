package solution

const arrLen = 3

type permutation [arrLen]int
type permutations = map[permutation]struct{}

func Solution(input permutation) permutations {
	result := make(permutations)

	for i := range input {
		p := permutation{}
		for j := 0; j < arrLen; j++ {
			index := (i + j) % arrLen
			p[j] = input[index]
		}

		result[p] = struct{}{}
	}

	return result
}

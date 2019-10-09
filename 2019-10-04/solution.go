package solution

func Solution(input []string) [][]string {
	lookup := make(map[anagram][]string)

	for i := 0; i < len(input); i++ {
		a := NewAnagram(input[i])

		lookup[a] = append(lookup[a], input[i])

	}

	result := make([][]string, 0, len(lookup))

	for key := range lookup {
		result = append(result, lookup[key])
	}

	return result
}

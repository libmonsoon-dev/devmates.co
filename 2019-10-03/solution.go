package solution

func Solution(input string) string {
	chars := []rune(input)
	vowelsIndexes := make([]int, 0)

	for i := range chars {
		if Vowels.Has(chars[i]) {
			vowelsIndexes = append(vowelsIndexes, i)
		}
	}

	for i, a := range vowelsIndexes[:len(vowelsIndexes)/2] {
		b := vowelsIndexes[len(vowelsIndexes)-1-i]

		chars[a], chars[b] = chars[b], chars[a]
	}

	return string(chars)
}

package solution

import (
	"math"
	"sort"
)

const maxChars = 3

type anagram [maxChars]rune

func NewAnagram(input string) anagram {
	chars := []rune(input)
	result := anagram{}

	sort.Slice(chars, func(i, j int) bool {
		return string(chars[i]) < string(chars[j])
	})

	minLen := int(math.Min(float64(len(chars)), maxChars))

	for i := 0; i < minLen; i++ {
		result[i] = chars[i]
	}

	return result
}

package solution

import "fmt"

var lookup = romanToInt{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

type romanToInt map[rune]int

func (lookup romanToInt) Get(key rune) int {
	val, ok := lookup[key]

	if !ok {
		panic(fmt.Sprintf("romanToInt.Get(%v) item not found", key))
	}

	return val

}

func (lookup romanToInt) GetList(keys []rune) []int {
	result := make([]int, len(keys))

	for i := 0; i < len(keys); i++ {
		result[i] = lookup.Get(keys[i])
	}

	return result
}

func Solution(input string) int {
	runes := []rune(input)
	result := 0
	i := 0

	for i < len(runes)-1 {
		val := lookup.Get(runes[i])
		next := lookup.Get(runes[i+1])

		if val >= next {
			result = result + val
			i = i + 1
		} else {
			result = result + next - val
			i = i + 2
		}
	}

	if i < len(runes) {
		result += lookup.Get(runes[i])
	}

	return result
}

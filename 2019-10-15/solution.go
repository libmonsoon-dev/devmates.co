package solution

import "strings"

func Solution(input int) string {
	builder := &strings.Builder{}

	for _, value := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		char := lookup.Get(value)
		for j := 0; j < input/value; j++ {
			builder.Write(char)
		}
		input %= value

	}

	return builder.String()
}

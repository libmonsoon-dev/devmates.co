package solution

import "strings"

func Solution(str string) string {
	slice := []rune(str)
	result := &strings.Builder{}

	for i := 0; i < len(slice); i++ {
		tmp := &strings.Builder{}
		if slice[i] != ' ' {
			for i < len(slice) && slice[i] != ' ' {
				builder := &strings.Builder{}
				builder.WriteRune(slice[i])
				builder.WriteString(tmp.String())
				tmp = builder
				i++
			}
			result.WriteString(tmp.String())
			if i < len(slice)-1 {
				result.WriteRune(slice[i])
			}
		} else {
			result.WriteRune(slice[i])
		}
	}

	return result.String()
}

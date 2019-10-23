package solution

func Solution(input string) string {
	data := &CharsList{[]rune(input), 3}

	for {
		if changed := data.RemoveDuplicates(); !changed {
			break
		}
	}

	return data.String()
}

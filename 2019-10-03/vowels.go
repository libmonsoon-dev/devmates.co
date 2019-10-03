package solution

type vowels struct {
	set map[rune]struct{}
}

func (v *vowels) Has(char rune) bool {
	_, ok := v.set[char]

	return ok
}

var Vowels = vowels{
	map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'y': {},
	},
}

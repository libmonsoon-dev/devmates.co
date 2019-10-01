package decoder

import (
	"strconv"
	"strings"
	"unicode"
)

type Parser struct {
	Before     string
	After      string
	Pattern    string
	Multiplier int

	openIndex  int
	closeIndex int
}

func (parser *Parser) String() string {
	return strings.Join(
		[]string{parser.Before, strings.Repeat(parser.Pattern, parser.Multiplier), parser.After},
		"",
	)
}

func parse(input []rune) string {
	parser := &Parser{}
	opened := 0
	for i, char := range input {
		if char == '[' {
			if opened == 0 {
				parser.openIndex = i
			}
			opened++
			continue
		}

		if char == ']' {
			opened--
			if opened == 0 {
				parser.closeIndex = i
				break
			}
		}
	}

	if parser.openIndex == 0 || parser.closeIndex == 0 {
		parser.Before = string(input)
		return parser.String()
	}

	multiplierIndex := parser.openIndex - 1

	for i := multiplierIndex; i >= 0; i-- {
		if !unicode.IsDigit(input[i]) {
			break
		}
		multiplierIndex = i
	}
	var err error

	strMultiplier := input[multiplierIndex:parser.openIndex]
	parser.Multiplier, err = strconv.Atoi(string(strMultiplier))

	if err != nil {
		panic(err)
	}

	parser.Before = string(input[:parser.openIndex-len(strMultiplier)])
	parser.After = parse(input[parser.closeIndex+1:])
	parser.Pattern = parse(input[parser.openIndex+1: parser.closeIndex])

	return parser.String()
}

func Recursive(input string) string {
	return parse([]rune(input))
}

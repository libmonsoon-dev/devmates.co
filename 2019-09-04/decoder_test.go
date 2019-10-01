package decoder

import "testing"

var tests = []struct {
	Input  string
	Output string
}{
	{Input: "3[a]2[bc]", Output: "aaabcbc"},
	{Input: "2[abc]3[cd]ef", Output: "abcabccdcdcdef"},
	{Input: "3[a2[c]]", Output: "accaccacc"},
}

func TestRecursive(t *testing.T) {
	for _, test := range tests {
		got := Recursive(test.Input)
		
		if got != test.Output {
			t.Errorf("Recursive(%#v) got = %#v, want %#v", test.Input, got, test.Output)
		}
	}
}

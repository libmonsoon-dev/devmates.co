package solution

import "testing"

func TestSolution(t *testing.T) {
	test := func(beginWord, endWord string, wordList []string, expect int) {
		if actual := Solution(beginWord, endWord, wordList); expect != actual {
			t.Errorf("expect != actual (%v != %v)", expect, actual)
		}
	}

	test("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}, 5)
	test("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}, 0)
}

func TestInclude(t *testing.T) {
	type args struct {
		items []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"include",
			args{[]string{"hot", "dot", "dog", "lot", "log", "cog"}, "cog"},
			true,
		},
		{
			"not include",
			args{[]string{"hot", "dot", "dog", "lot", "log"}, "cog"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := include(tt.args.items, tt.args.item); got != tt.want {
				t.Errorf("include() = %v, want %v", got, tt.want)
			}
		})
	}
}

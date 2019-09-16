package mutations

import (
	"testing"
)

func TestCountMutations(t *testing.T) {
	tests := []struct {
		start  str
		end    str
		bank   []str
		expect int
	}{
		{
			toStr("AACCGGTT"),
			toStr("AAACGGTA"),
			[]str{
				toStr("AACCGGTA"),
				toStr("AACCGCTA"),
				toStr("AAACGGTA"),
			},
			2,
		},
		{
			toStr("AAAAACCC"),
			toStr("AACCCCCC"),
			[]str{
				toStr("AAAACCCC"),
				toStr("AAACCCCC"),
				toStr("AACCCCCC"),
			},
			3,
		},
	}

	for _, test := range tests {
		actual := CountMutations(test.start, test.end, test.bank)

		if test.expect != actual {
			t.Errorf("expect != actual (%v != %v)", test.expect, actual)
		}
	}

}

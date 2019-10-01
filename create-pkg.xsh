#!/usr/bin/env xonsh

$RAISE_SUBPROC_ERROR = True

from datetime import datetime

date_str = datetime.now().strftime("%Y-%m-%d")
header = "package solution\n"

mkdir -p @(date_str)
touch @(date_str)/README.md
echo @(header + '''
func Solution(input interface{}) interface{} {
	return nil
}''') > @(date_str)/solution.go

echo @(header + '''
import "testing"

func TestSolution(t *testing.T) {
	test := func(input, expect interface{}) {
		if actual := Solution(input); expect != actual {
			t.Errorf("expect != actual (%#v != %#v)", expect, actual)
		}
	}
}''') > @(date_str)/solution_test.go
